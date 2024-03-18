package crud

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"
	ipv4FwFilterResModel "github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/global/firewall/ipv4-forward-filter/resourcemodel"
	ipv4ResModel "github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/ipv4-name/resourcemodel"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/tests/api"
)

// TestCrudDeleteSuccess test CRUD helper: Delete
//
//	Default situation where we attempt to delete a resource after
//	making sure it has no children.
func TestCrudDeleteSuccess(t *testing.T) {
	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// Child check API call
	exchangeChildExistsCheck := eList.Add()
	exchangeChildExistsCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"showConfig","path":["firewall","ipv4","name","Test-Fw-Name"]}`,
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
			`{"op":"delete","path":["firewall","ipv4","name","Test-Fw-Name"]}`+
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
		SelfIdentifier:                       basetypes.NewStringValue("Test-Fw-Name"),
		LeafFirewallIPvfourNameDefaultAction: basetypes.NewStringValue("reject"),
		LeafFirewallIPvfourNameDefaultLog:    basetypes.NewBoolValue(true),
		LeafFirewallIPvfourNameDescrIPtion:   basetypes.NewStringValue("Managed by terraform"),
	}

	// Server
	apiAddress := "localhost:8080"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := context.Background()
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)

	// Execute test
	err := delete(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		for _, e := range eList.Unmatched() {
			t.Errorf("Unmatched exchange:\n%s", e.Sexpect())
		}
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
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

		// Child check API call
		exchangeChildExistsCheck := eList.Add()
		exchangeChildExistsCheck.Expect(
			"/retrieve",
			apiKey,
			`{"op":"showConfig","path":["firewall","ipv4","name","Test-Fw-Name"]}`,
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

		// From resource model
		model := &ipv4ResModel.FirewallIPvfourName{
			SelfIdentifier:                       basetypes.NewStringValue("Test-Fw-Name"),
			LeafFirewallIPvfourNameDefaultAction: basetypes.NewStringValue("reject"),
			LeafFirewallIPvfourNameDefaultLog:    basetypes.NewBoolValue(true),
			LeafFirewallIPvfourNameDescrIPtion:   basetypes.NewStringValue("Managed by terraform"),
		}

		// Server
		apiAddress := "localhost:8080"
		srv := &http.Server{
			Addr: apiAddress,
		}
		api.Server(srv, eList)

		// Client
		ctx := context.Background()
		client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
		providerData := data.NewProviderData(client)

		// Execute test
		err := delete(ctx, providerData, client, model)

		// Verify results
		if err == nil {
			t.Errorf("expected an error, got nil")
			return
		}

		if !strings.Contains(err.Error(), "child resource detected") {
			t.Errorf("delete returned the wrong error: %s", err.Error())
			return
		}

		t.Logf("Found child!\n%s", err.Error())

		// Validate API calls
		if len(eList.Unmatched()) > 0 {
			for _, e := range eList.Unmatched() {
				t.Errorf("Unmatched exchange:\n%s", e.Sexpect())
			}
			t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
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
			`{"op":"delete","path":["firewall","ipv4","name","Test-Fw-Name"]}`+
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
		SelfIdentifier:                       basetypes.NewStringValue("Test-Fw-Name"),
		LeafFirewallIPvfourNameDefaultAction: basetypes.NewStringValue("reject"),
		LeafFirewallIPvfourNameDefaultLog:    basetypes.NewBoolValue(true),
		LeafFirewallIPvfourNameDescrIPtion:   basetypes.NewStringValue("Managed by terraform"),
	}

	// Server
	apiAddress := "localhost:8080"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := context.Background()
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)
	providerData.Config.CrudSkipCheckChildBeforeDelete = true

	// Execute test
	err := delete(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
		return
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		for _, e := range eList.Unmatched() {
			t.Errorf("Unmatched exchange:\n%s", e.Sexpect())
		}
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		return
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
	apiAddress := "localhost:8080"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := context.Background()
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)

	// Execute test
	err := delete(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		for _, e := range eList.Unmatched() {
			t.Errorf("Unmatched exchange:\n%s", e.Sexpect())
		}
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
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
	apiAddress := "localhost:8080"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := context.Background()
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)

	// Execute test
	err := delete(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		for _, e := range eList.Unmatched() {
			t.Errorf("Unmatched exchange:\n%s", e.Sexpect())
		}
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
	}
}
