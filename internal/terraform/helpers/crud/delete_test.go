package crud

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflogtest"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	cruderrors "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers/crud/cruderror"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"
	ipv4FwFilterResModel "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/firewall/ipv4-forward-filter/resourcemodel"
	ipv4ResModel "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/firewall/ipv4-name/resourcemodel"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/tests/api"
)

// TestCrudDeleteSuccess test CRUD helper: Delete
//
//	Default situation where we attempt to delete a resource after
//	making sure it has no children.
func TestCrudDeleteSuccess(t *testing.T) {
	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// Resource exists check API call
	eList.Add().Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudDeleteSuccess"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": true,
			"error": null
		}`,
	)

	// Child check API call
	exchangeChildExistsCheck := eList.Add()
	exchangeChildExistsCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"showConfig","path":["firewall","ipv4","name","TestCrudDeleteSuccess"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": {
				"default-action": "reject",
				"default-log": {},
				"description": "Managed by terraform"
			},
			"error": null
		}`,
	)

	// Delete resource API call
	exchangeCreateResource := eList.Add()
	exchangeCreateResource.Expect(
		"/configure",
		apiKey,
		`[`+
			`{"op":"delete","path":["firewall","ipv4","name","TestCrudDeleteSuccess"]}`+
			`]`,
	).Response(
		200,
		`{
				"success": true,
				"data": null,
				"error": null
			}`,
	)

	// From resource model
	model := &ipv4ResModel.FirewallIPvfourName{
		// SelfIdentifier:                       basetypes.NewStringValue("TestCrudDeleteSuccess"),
		SelfIdentifier: basetypes.NewObjectValueMust(
			map[string]attr.Type{"name": basetypes.StringType{}},
			map[string]attr.Value{"name": basetypes.NewStringValue("TestCrudDeleteSuccess")}),

		LeafFirewallIPvfourNameDefaultAction: basetypes.NewStringValue("reject"),
		LeafFirewallIPvfourNameDefaultLog:    basetypes.NewBoolValue(true),
		LeafFirewallIPvfourNameDescrIPtion:   basetypes.NewStringValue("Managed by terraform"),
	}

	// Server
	apiAddress := "localhost:50007"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)

	// Execute test
	err := delete(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
	}
}

// TestCrudDeleteResourceHasChildFailure test CRUD helper: Delete
//
//	Default situation where the resource has a child causing
//	delete to fail
func TestCrudDeleteResourceHasChildFailure(t *testing.T) {
	{
		// API mocking
		eList := api.NewExchangeList()
		apiKey := "test-key"

		// x4, due to retry
		for range 4 {
			// Resource exists check API call
			eList.Add().Expect(
				"/retrieve",
				apiKey,
				`{"op":"exists","path":["firewall","ipv4","name","TestCrudDeleteResourceHasChildFailure"]}`,
			).Response(
				200,
				`{
					"success": true,
					"data": true,
					"error": null
				}`,
			)

			// Child check API call
			eList.Add().Expect(
				"/retrieve",
				apiKey,
				`{"op":"showConfig","path":["firewall","ipv4","name","TestCrudDeleteResourceHasChildFailure"]}`,
			).Response(
				200,
				`{
					"success": true,
					"data": {
						"default-action": "reject",
						"default-log": {},
						"description": "Managed by terraform",
						"rule": {
							"1": {"action": "accept", "description": "Allow established connections", "protocol": "all", "state": "established"},
							"2": {"action": "accept", "description": "Allow related connections", "protocol": "all", "state": "related"},
							"3": {"action": "drop", "description": "Disallow invalid packets", "log": {}, "protocol": "all", "state": "invalid"}
						}
					},
					"error": null
				}`,
			)
		}

		// From resource model
		model := &ipv4ResModel.FirewallIPvfourName{
			// SelfIdentifier:                       basetypes.NewStringValue("TestCrudDeleteResourceHasChildFailure"),
			SelfIdentifier: basetypes.NewObjectValueMust(
				map[string]attr.Type{"name": basetypes.StringType{}},
				map[string]attr.Value{"name": basetypes.NewStringValue("TestCrudDeleteResourceHasChildFailure")}),

			LeafFirewallIPvfourNameDefaultAction: basetypes.NewStringValue("reject"),
			LeafFirewallIPvfourNameDefaultLog:    basetypes.NewBoolValue(true),
			LeafFirewallIPvfourNameDescrIPtion:   basetypes.NewStringValue("Managed by terraform"),
		}

		// Server
		apiAddress := "localhost:50008"
		srv := &http.Server{
			Addr: apiAddress,
		}
		api.Server(srv, eList)

		// Client
		ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
		ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()
		client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
		providerData := data.NewProviderData(client)

		// Execute test
		err := delete(ctx, providerData, client, model)

		// Verify results
		if err == nil {
			t.Errorf("expected an error, got nil")
		} else {
			if !strings.Contains(err.Error(), "child resource detected") {
				t.Errorf("delete returned the wrong error: %s", err.Error())
			}
		}

		// Validate API calls
		if len(eList.Unmatched()) > 0 {
			t.Logf("Total matched exchanges: %d", len(eList.Matched()))
			t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
			t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
			t.Errorf("Received request:\n%s", eList.Failed())
		}
	}
}

// TestCrudDeleteResourceHasChildIgnore test CRUD helper: Delete
//
//	Situation where the resource has a child but provider
//	is configured to ignore it
func TestCrudDeleteResourceHasChildIgnore(t *testing.T) {
	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// Delete resource API call
	exchangeCreateResource := eList.Add()
	exchangeCreateResource.Expect(
		"/configure",
		apiKey,
		`[`+
			`{"op":"delete","path":["firewall","ipv4","name","TestCrudDeleteResourceHasChildIgnore"]}`+
			`]`,
	).Response(
		200,
		`{
				"success": true,
				"data": null,
				"error": null
			}`,
	)

	// From resource model
	model := &ipv4ResModel.FirewallIPvfourName{
		// SelfIdentifier:                       basetypes.NewStringValue("TestCrudDeleteResourceHasChildIgnore"),
		SelfIdentifier: basetypes.NewObjectValueMust(
			map[string]attr.Type{"name": basetypes.StringType{}},
			map[string]attr.Value{"name": basetypes.NewStringValue("TestCrudDeleteResourceHasChildIgnore")}),

		LeafFirewallIPvfourNameDefaultAction: basetypes.NewStringValue("reject"),
		LeafFirewallIPvfourNameDefaultLog:    basetypes.NewBoolValue(true),
		LeafFirewallIPvfourNameDescrIPtion:   basetypes.NewStringValue("Managed by terraform"),
	}

	// Server
	apiAddress := "localhost:50009"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)
	providerData.Config.CrudSkipCheckChildBeforeDelete = true

	// Execute test
	err := delete(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
	}
}

// TestCrudDeleteGlobalResourceWithChild test CRUD helper: Delete
//
//	Default situation where the global resource has a child
//	and the resource should only delete its own attributes
func TestCrudDeleteGlobalResourceWithChild(t *testing.T) {
	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// Resource exists check API call
	eList.Add().Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","forward","filter"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": true,
			"error": null
		}`,
	)

	// Child check API call
	exchangeChildExistsCheck := eList.Add()
	exchangeChildExistsCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"showConfig","path":["firewall","ipv4","forward","filter"]}`,
	).Response(
		200,
		`{
				"success": true,
				"data": {
					"default-action": "reject",
					"default-log": {},
					"rule": {
						"1": {
							"action": "drop",
							"log": {},
							"description": "child resource"
						}
					}
				},
				"error": null
			}`,
	)

	// Delete resource API call
	exchangeCreateResource := eList.Add()
	exchangeCreateResource.Expect(
		"/configure",
		apiKey,
		`[`+
			`{"op":"delete","path":["firewall","ipv4","forward","filter","default-action","reject"]},`+
			`{"op":"delete","path":["firewall","ipv4","forward","filter","default-log"]}`+
			`]`,
	).Response(
		200,
		`{
				"success": true,
				"data": null,
				"error": null
			}`,
	)

	// From resource model
	model := &ipv4FwFilterResModel.FirewallIPvfourForwardFilter{
		ID: basetypes.NewStringValue("GLOBAL"),
		LeafFirewallIPvfourForwardFilterDefaultAction: basetypes.NewStringValue("reject"),
		LeafFirewallIPvfourForwardFilterDefaultLog:    basetypes.NewBoolValue(true),
	}

	// Server
	apiAddress := "localhost:50010"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)

	// Execute test
	err := delete(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
	}
}

// TestCrudDeleteGlobalResourceWithoutChild test CRUD helper: Delete
//
//	Default situation where the global resource has no children
//	and the resource should delete itself completely, not just
//	the attributes
func TestCrudDeleteGlobalResourceWithoutChild(t *testing.T) {
	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// Resource exists check API call
	eList.Add().Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","forward","filter"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": true,
			"error": null
		}`,
	)

	// Child check API call
	exchangeChildExistsCheck := eList.Add()
	exchangeChildExistsCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"showConfig","path":["firewall","ipv4","forward","filter"]}`,
	).Response(
		200,
		`{
				"success": true,
				"data": {
					"default-action": "reject",
					"default-log": {}
				},
				"error": null
			}`,
	)

	// Delete resource API call
	exchangeCreateResource := eList.Add()
	exchangeCreateResource.Expect(
		"/configure",
		apiKey,
		`[`+
			`{"op":"delete","path":["firewall","ipv4","forward","filter"]}`+
			`]`,
	).Response(
		200,
		`{
				"success": true,
				"data": null,
				"error": null
			}`,
	)

	// From resource model
	model := &ipv4FwFilterResModel.FirewallIPvfourForwardFilter{
		ID: basetypes.NewStringValue("GLOBAL"),
		LeafFirewallIPvfourForwardFilterDefaultAction: basetypes.NewStringValue("reject"),
		LeafFirewallIPvfourForwardFilterDefaultLog:    basetypes.NewBoolValue(true),
	}

	// Server
	apiAddress := "localhost:50011"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)

	// Execute test
	err := delete(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
	}
}

// TestCrudDeleteRetrySuccess test CRUD helper: Delete
//
//	Default situation where we attempt to delete a resource, but
//	have to wait for the child resource to be deleted.
func TestCrudDeleteRetrySuccess(t *testing.T) {
	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// x3, due to retry
	for range 3 {
		// Resource exists check API call
		eList.Add().Expect(
			"/retrieve",
			apiKey,
			`{"op":"exists","path":["firewall","ipv4","name","TestCrudDeleteRetrySuccess"]}`,
		).Response(
			200,
			`{
				"success": true,
				"data": true,
				"error": null
			}`,
		)

		// Child check API call: before delete
		exchangeChildExistsCheck := eList.Add()
		exchangeChildExistsCheck.Expect(
			"/retrieve",
			apiKey,
			`{"op":"showConfig","path":["firewall","ipv4","name","TestCrudDeleteRetrySuccess"]}`,
		).Response(
			200,
			`{
				"success": true,
				"data": {
					"default-action": "reject",
					"default-log": {},
					"description": "Managed by terraform",
					"rule": {
						"1": {"action": "accept", "description": "Allow established connections", "protocol": "all", "state": "established"},
						"2": {"action": "accept", "description": "Allow related connections", "protocol": "all", "state": "related"},
						"3": {"action": "drop", "description": "Disallow invalid packets", "log": {}, "protocol": "all", "state": "invalid"}
					}
				},
				"error": null
			}`,
		)
	}

	// Resource exists check API call
	eList.Add().Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudDeleteRetrySuccess"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": true,
			"error": null
		}`,
	)

	// Child check API call: after delete
	exchangeChildExistsCheck := eList.Add()
	exchangeChildExistsCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"showConfig","path":["firewall","ipv4","name","TestCrudDeleteRetrySuccess"]}`,
	).Response(
		200,
		`{
				"success": true,
				"data": {
					"default-action": "reject",
					"default-log": {},
					"description": "Managed by terraform"
				},
				"error": null
			}`,
	)

	// Delete resource API call
	exchangeCreateResource := eList.Add()
	exchangeCreateResource.Expect(
		"/configure",
		apiKey,
		`[`+
			`{"op":"delete","path":["firewall","ipv4","name","TestCrudDeleteRetrySuccess"]}`+
			`]`,
	).Response(
		200,
		`{
				"success": true,
				"data": null,
				"error": null
			}`,
	)

	// From resource model
	model := &ipv4ResModel.FirewallIPvfourName{
		// SelfIdentifier:                       basetypes.NewStringValue("TestCrudDeleteRetrySuccess"),
		SelfIdentifier: basetypes.NewObjectValueMust(
			map[string]attr.Type{"name": basetypes.StringType{}},
			map[string]attr.Value{"name": basetypes.NewStringValue("TestCrudDeleteRetrySuccess")}),

		LeafFirewallIPvfourNameDefaultAction: basetypes.NewStringValue("reject"),
		LeafFirewallIPvfourNameDefaultLog:    basetypes.NewBoolValue(true),
		LeafFirewallIPvfourNameDescrIPtion:   basetypes.NewStringValue("Managed by terraform"),
	}

	// Server
	apiAddress := "localhost:50007"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)

	// Execute test
	err := delete(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
	}
}

// TestCrudDeleteEmptyGlobalResource test CRUD helper: Delete
//
//	Default situation where a global resource exists but is empty
func TestCrudDeleteEmptyGlobalResource(t *testing.T) {
	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// Resource exists check API call
	eList.Add().Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","forward","filter"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": true,
			"error": null
		}`,
	)

	// Child check API call
	exchangeChildExistsCheck := eList.Add()
	exchangeChildExistsCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"showConfig","path":["firewall","ipv4","forward","filter"]}`,
	).Response(
		400,
		`{"success": false, "error": "Configuration under specified path is empty\n", "data": null}`,
	)

	// Delete resource API call
	exchangeCreateResource := eList.Add()
	exchangeCreateResource.Expect(
		"/configure",
		apiKey,
		`[`+
			`{"op":"delete","path":["firewall","ipv4","forward","filter"]}`+
			`]`,
	).Response(
		200,
		`{
				"success": true,
				"data": null,
				"error": null
			}`,
	)

	// From resource model
	model := &ipv4FwFilterResModel.FirewallIPvfourForwardFilter{
		ID: basetypes.NewStringValue("GLOBAL"),
		LeafFirewallIPvfourForwardFilterDefaultAction: basetypes.NewStringValue("reject"),
		LeafFirewallIPvfourForwardFilterDefaultLog:    basetypes.NewBoolValue(true),
	}

	// Server
	apiAddress := "localhost:50011"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)

	// Execute test
	err := delete(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
	}
}

// TestCrudDeleteEmptyResource test CRUD helper: Delete
//
//	Default situation where the a named resource exists but is empty
func TestCrudDeleteEmptyResource(t *testing.T) {
	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// Resource exists check API call
	eList.Add().Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudDeleteEmptyResource"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": true,
			"error": null
		}`,
	)

	// Child check API call
	exchangeChildExistsCheck := eList.Add()
	exchangeChildExistsCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"showConfig","path":["firewall","ipv4","name","TestCrudDeleteEmptyResource"]}`,
	).Response(
		400,
		`{"success": false, "error": "Configuration under specified path is empty\n", "data": null}`,
	)

	// Delete resource API call
	exchangeCreateResource := eList.Add()
	exchangeCreateResource.Expect(
		"/configure",
		apiKey,
		`[`+
			`{"op":"delete","path":["firewall","ipv4","name","TestCrudDeleteEmptyResource"]}`+
			`]`,
	).Response(
		200,
		`{
				"success": true,
				"data": null,
				"error": null
			}`,
	)

	// From resource model
	model := &ipv4ResModel.FirewallIPvfourName{
		// SelfIdentifier:                       basetypes.NewStringValue("TestCrudDeleteEmptyResource"),
		SelfIdentifier: basetypes.NewObjectValueMust(
			map[string]attr.Type{"name": basetypes.StringType{}},
			map[string]attr.Value{"name": basetypes.NewStringValue("TestCrudDeleteEmptyResource")}),

		LeafFirewallIPvfourNameDefaultAction: basetypes.NewStringValue("reject"),
		LeafFirewallIPvfourNameDefaultLog:    basetypes.NewBoolValue(true),
		LeafFirewallIPvfourNameDescrIPtion:   basetypes.NewStringValue("Managed by terraform"),
	}

	// Server
	apiAddress := "localhost:50011"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)

	// Execute test
	err := delete(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
	}
}

// TestCrudDeleteDeletedGlobalResource test CRUD helper: Delete
//
//	Default situation where a global resource is already deleted
func TestCrudDeleteDeletedGlobalResource(t *testing.T) {
	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// Resource exists check API call
	eList.Add().Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","forward","filter"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": false,
			"error": null
		}`,
	)

	// From resource model
	model := &ipv4FwFilterResModel.FirewallIPvfourForwardFilter{
		ID: basetypes.NewStringValue("GLOBAL"),
		LeafFirewallIPvfourForwardFilterDefaultAction: basetypes.NewStringValue("reject"),
		LeafFirewallIPvfourForwardFilterDefaultLog:    basetypes.NewBoolValue(true),
	}

	// Server
	apiAddress := "localhost:50011"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)

	// Execute test
	err := delete(ctx, providerData, client, model)
	var wantErr cruderrors.ResourceNotFoundError
	if !errors.As(err, &wantErr) {
		t.Errorf("Delete failed: %v", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
	}
}

// TestCrudDeleteDeletedResource test CRUD helper: Delete
//
//	Default situation where the a named resource is already deleted
func TestCrudDeleteDeletedResource(t *testing.T) {
	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// Resource exists check API call
	eList.Add().Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudDeleteDeletedResource"]}`,
	).Response(
		400,
		`{"success": true, "error": null, "data": false}`,
	)

	// From resource model
	model := &ipv4ResModel.FirewallIPvfourName{
		// SelfIdentifier:                       basetypes.NewStringValue("TestCrudDeleteDeletedResource"),
		SelfIdentifier: basetypes.NewObjectValueMust(
			map[string]attr.Type{"name": basetypes.StringType{}},
			map[string]attr.Value{"name": basetypes.NewStringValue("TestCrudDeleteDeletedResource")}),

		LeafFirewallIPvfourNameDefaultAction: basetypes.NewStringValue("reject"),
		LeafFirewallIPvfourNameDefaultLog:    basetypes.NewBoolValue(true),
		LeafFirewallIPvfourNameDescrIPtion:   basetypes.NewStringValue("Managed by terraform"),
	}

	// Server
	apiAddress := "localhost:50011"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)

	// Execute test
	err := delete(ctx, providerData, client, model)
	var wantErr cruderrors.ResourceNotFoundError
	if !errors.As(err, &wantErr) {
		t.Errorf("Delete failed: %v", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
	}
}
