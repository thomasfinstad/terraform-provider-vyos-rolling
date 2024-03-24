package client

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client/clienterrors"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers/tools"
)

// TODO create a client request internal function
//  deduplicate the work being done in the public functions by
//  having a private function to send the request and receive the response
//  milestone:6

// NewClient creates a new client object to use with VyOS CRUD functions
func NewClient(
	ctx context.Context,
	endpoint string,
	apiKey string,
	userAgent string,
	disableVerify bool,
) Client {

	c := Client{
		httpClient: http.Client{},

		userAgent: userAgent,
		endpoint:  endpoint,
		apiKey:    apiKey,
	}

	if disableVerify {
		tools.Warn(ctx, "Disabling TLS Certificate Verification")
		c.httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	return c
}

// Client wrapper around http client with convinience functions
// Use NewClient() to generate a new client
type Client struct {
	httpClient http.Client

	userAgent string
	endpoint  string
	apiKey    string

	opsSet    [][]string
	opsDelete [][]string
}

// StageSet saves vyos paths to configure during commit
func (c *Client) StageSet(ctx context.Context, values [][]string) {
	tools.Trace(ctx, "stageing set ops", map[string]interface{}{"client:httpClient": fmt.Sprintf("%p:%p", c, &c.httpClient), "paths": values, "current set ops": c.opsSet})
	c.opsSet = append(c.opsSet, values...)
}

// StageDelete saves vyos paths to delete during commit
func (c *Client) StageDelete(ctx context.Context, values [][]string) {
	tools.Trace(ctx, "stageing delete ops", map[string]interface{}{"client:httpClient": fmt.Sprintf("%p:%p", c, &c.httpClient), "paths": values, "current del ops": c.opsDelete})
	c.opsDelete = append(c.opsDelete, values...)
}

// CommitChanges executes staged vyos paths.
// Order of operations as they are sent to VyOS:
//  1. delete
//  2. set
func (c *Client) CommitChanges(ctx context.Context) (any, error) {

	// TODO investigate options to speed up multiple resource config
	//	Suggestes pesudo solution:
	//  1. Make client, or some part, a pointer shared across all resources
	//	2. wait 500ms to allow multiple resources to gather up changes to be commited
	//  3. on failure remove some resources changes from the commit and try again
	//  4. on success return so to the resources that succeeded
	//  milestone:6

	endpoint := c.endpoint + "/configure"
	operations := []map[string]interface{}{}

	for _, path := range c.opsDelete {
		operations = append(
			operations,
			map[string]interface{}{
				"op":   "delete",
				"path": path,
			},
		)
	}

	for _, path := range c.opsSet {
		operations = append(
			operations,
			map[string]interface{}{
				"op":   "set",
				"path": path,
			},
		)
	}

	jsonOperations, err := json.Marshal(
		operations,
	)
	if err != nil {
		return nil, fmt.Errorf("fail json marshal delete ops: %w", err)
	}

	payload := url.Values{
		"key":  []string{c.apiKey},
		"data": []string{string(jsonOperations)},
	}

	tools.Info(ctx, "Creating configure request for endpoint", map[string]interface{}{"endpoint": endpoint, "payload": payload})

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request object: %w", err)
	}

	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to complete http request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read http response: %w", err)
	}

	if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("http error [%s]: %s", resp.Status, string(body))
	}

	c.opsSet = [][]string{}
	c.opsDelete = [][]string{}

	var ret map[string]interface{}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal http response body: '%s' as json: %w", body, err)
	}

	if ret["success"] == true {
		return ret["data"], nil
	}

	return nil, fmt.Errorf("API ERROR [%s]: %v", resp.Status, ret["error"])
}

// Has checks the provided path for a configuration and returns
// true if found, false otherwise.
// Also returns true for empty config blocks by
// using the `exists` API operation.
func (c *Client) Has(ctx context.Context, path []string) (bool, error) {
	endpoint := c.endpoint + "/retrieve"
	operation, err := json.Marshal(
		map[string]interface{}{
			"op":   "exists",
			"path": path,
		},
	)
	if err != nil {
		return false, &MarshalError{message: "read operation", marshalErr: err}
	}

	payload := url.Values{
		"key":  []string{c.apiKey},
		"data": []string{string(operation)},
	}

	tools.Info(ctx, "Creating exists request for endpoint", map[string]interface{}{"endpoint": endpoint, "payload": payload})

	payloadEnc := payload.Encode()
	tools.Debug(ctx, "Request payload encoded", map[string]interface{}{"payload": payloadEnc})
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, strings.NewReader(payloadEnc))
	if err != nil {
		return false, fmt.Errorf("failed to create http request object: %w", err)
	}

	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to complete http request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("failed to read http response: %w", err)
	}

	var ret map[string]interface{}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal http response body: '%s' as json: %w", body, err)
	}

	if ret["success"] == true {

		if retB, ok := ret["data"].(bool); ok {
			return retB, nil
		}
		return false, fmt.Errorf("[api error]: could not convert returned 'data' field to bool: %v", ret)
	}

	if errmsg, ok := ret["error"]; ok {
		if errmsg, ok := errmsg.(string); ok {
			return false, clienterrors.NewNotFoundError("[api error]: %s", errmsg)
		}
	}

	return false, clienterrors.NewNotFoundError("[api error]: %v", ret)
}

// Get returns the config found under path if it exists
//
// Returns:
//
//	error: if the resource was not found a clienterror.NotFoundError is returned, otherwise a generic error
func (c *Client) Get(ctx context.Context, path []string) (any, error) {
	endpoint := c.endpoint + "/retrieve"
	operation, err := json.Marshal(
		map[string]interface{}{
			"op":   "showConfig",
			"path": path,
		},
	)
	if err != nil {
		return nil, &MarshalError{message: "showConfig operation", marshalErr: err}
	}

	payload := url.Values{
		"key":  []string{c.apiKey},
		"data": []string{string(operation)},
	}

	tools.Info(ctx, "Creating showConfig request for endpoint", map[string]interface{}{"endpoint": endpoint, "payload": payload})

	payloadEnc := payload.Encode()
	tools.Debug(ctx, "Request payload encoded", map[string]interface{}{"payload": payloadEnc})
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, strings.NewReader(payloadEnc))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request object: %w", err)
	}

	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to complete http request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read http response: %w", err)
	}

	var ret map[string]interface{}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal http response body: '%s' as json: %w", body, err)
	}

	if ret["success"] == true {

		return ret["data"], nil
	}

	if errmsg, ok := ret["error"]; ok {
		if errmsg, ok := errmsg.(string); ok && errmsg == "Configuration under specified path is empty\n" {
			return nil, clienterrors.NewNotFoundError(
				"[%s]: %s",
				strings.Join(path, " "),
				strings.TrimSuffix(errmsg, "\n"),
			)
		}

		return nil, fmt.Errorf("[api error]: %s", errmsg)
	}

	return nil, fmt.Errorf("[api error]: %v", ret)
}
