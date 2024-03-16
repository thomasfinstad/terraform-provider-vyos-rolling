package client

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

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
		tflog.Warn(ctx, "Disabling TLS Certificate Verification")
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
	tflog.Trace(ctx, "stageing set ops", map[string]interface{}{"client:httpClient": fmt.Sprintf("%p:%p", c, &c.httpClient), "paths": values, "current set ops": c.opsSet})
	c.opsSet = append(c.opsSet, values...)
}

// StageDelete saves vyos paths to delete during commit
func (c *Client) StageDelete(ctx context.Context, values [][]string) {
	tflog.Trace(ctx, "stageing delete ops", map[string]interface{}{"client:httpClient": fmt.Sprintf("%p:%p", c, &c.httpClient), "paths": values, "current del ops": c.opsDelete})
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

	tflog.Info(ctx, "Creating configure request", map[string]interface{}{"endpoint": endpoint, "payload": payload})

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(payload.Encode()))
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
		return nil, fmt.Errorf("failed to unmarshal http response body as json: %w, body: %s", err, body)
	}

	if ret["success"] == true {
		return ret["data"], nil
	}

	return nil, fmt.Errorf("API ERROR [%s]: %v", resp.Status, ret["error"])
}

func (c *Client) Read(ctx context.Context, path []string) (any, error) {
	endpoint := c.endpoint + "/retrieve"
	operation, err := json.Marshal(
		map[string]interface{}{
			"op":   "showConfig",
			"path": path,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("fail json marshal read operation: %w", err)
	}

	payload := url.Values{
		"key":  []string{c.apiKey},
		"data": []string{string(operation)},
	}

	tflog.Info(ctx, "Creating read request for endpoint", map[string]interface{}{"endpoint": endpoint, "payload": payload})
	log.Println("Creating read request for endpoint", map[string]interface{}{"endpoint": endpoint, "payload": payload})

	payloadEnc := payload.Encode()
	tflog.Debug(ctx, "Request payload encoded", map[string]interface{}{"payload": payloadEnc})
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(payloadEnc))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request object: %w", err)
	}

	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	log.Println("Sending Request")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to complete http request: %w", err)
	}
	defer resp.Body.Close()

	log.Println("Reading response")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read http response: %w", err)
	}

	var ret map[string]interface{}

	log.Println("Unmarshaling from json")
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal http response body as json: %w, body: %s", err, body)
	}

	if ret["success"] == true {
		log.Println("API read success")
		return ret["data"], nil
	}

	log.Println("API read failure")
	return nil, fmt.Errorf("API ERROR: %s", ret["error"])
}
