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

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NewClient creates a new client object to use with VyOS CRUD functions
func NewClient(
	ctx context.Context,
	endpoint string,
	apiKey string,
	userAgent string,
	disableVerify bool,
) *Client {

	c := &Client{
		httpClient: &http.Client{},

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
	httpClient *http.Client

	userAgent string
	endpoint  string
	apiKey    string

	opsSet    [][]string
	opsDelete [][]string
}

// StageSet saves vyos paths to configure during commit
func (c *Client) StageSet(ctx context.Context, values [][]string) {
	c.opsSet = append(c.opsSet, values...)
}

// StageDelete saves vyos paths to delete during commit
func (c *Client) StageDelete(ctx context.Context, values [][]string) {
	c.opsDelete = append(c.opsDelete, values...)
}

// CommitChanges executes staged vyos paths.
// Order of operations as they are sent to VyOS:
//  1. delete
//  2. set
func (c *Client) CommitChanges(ctx context.Context) (any, error) {
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
		tflog.Error(ctx, "Fail json marshal delete ops", map[string]interface{}{"operations": operations, "error": err})
		return nil, err
	}

	payload := url.Values{
		"key":  []string{c.apiKey},
		"data": []string{string(jsonOperations)},
	}

	tflog.Info(ctx, "Creating request for <endpoint>/configure", map[string]interface{}{"payload": payload})

	req, err := http.NewRequest("POST", c.endpoint+"/configure", strings.NewReader(payload.Encode()))
	if err != nil {
		tflog.Error(ctx, "Failed to create http request object", map[string]interface{}{"error": err})
		return nil, err
	}

	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		tflog.Error(ctx, "Failed to complete http request", map[string]interface{}{"error": err})
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, "Failed to read http response", map[string]interface{}{"error": err})
		return nil, err
	}

	c.opsSet = [][]string{}
	c.opsDelete = [][]string{}

	var ret map[string]interface{}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		tflog.Error(ctx, "Failed to unmarshal http response body as json", map[string]interface{}{"response": body, "error": err})
		return nil, err
	}

	if ret["success"] == true {
		return ret["data"], nil
	}

	return nil, fmt.Errorf("API ERROR: %s", ret["error"])
}
