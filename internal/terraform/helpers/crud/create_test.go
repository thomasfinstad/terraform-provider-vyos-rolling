package crud

import (
	"context"
	"math/big"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/ipv4-name-rule/resourcemodel"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/tests/api"
)

// TestCrudCreateSuccess test CRUD helper: Create
//
//	Default situation where we attempt to create a resource after
//	making sure the parent exists and the resource does not exists.
func TestCrudCreateSuccess(t *testing.T) {
	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// Parent check API call
	exchangeParentExistsCheck := eList.Add()
	exchangeParentExistsCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","Test-Fw-Name"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": true,
			"error": null
		}`,
	)

	// Self check API call
	exchangeExistingResourceCheck := eList.Add()
	exchangeExistingResourceCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","Test-Fw-Name","rule","42"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": false,
			"error": null
		}`,
	)

	// Create resource API call
	exchangeCreateResource := eList.Add()
	exchangeCreateResource.Expect(
		"/configure",
		apiKey,
		`[`+
			`{"op":"set","path":["firewall","ipv4","name","Test-Fw-Name","rule","42","action","accept"]},`+
			`{"op":"set","path":["firewall","ipv4","name","Test-Fw-Name","rule","42","description","Allow http outgoing traffic"]},`+
			`{"op":"set","path":["firewall","ipv4","name","Test-Fw-Name","rule","42","protocol","tcp"]},`+
			`{"op":"set","path":["firewall","ipv4","name","Test-Fw-Name","rule","42","destination","group","port-group","Web"]}`+
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
	model := &resourcemodel.FirewallIPvfourNameRule{
		ParentIDFirewallIPvfourName: basetypes.NewStringValue("Test-Fw-Name"),
		SelfIdentifier:              basetypes.NewNumberValue(big.NewFloat(42)),

		LeafFirewallIPvfourNameRuleAction:      basetypes.NewStringValue("accept"),
		LeafFirewallIPvfourNameRuleDescrIPtion: basetypes.NewStringValue("Allow http outgoing traffic"),
		NodeFirewallIPvfourNameRuleDestination: &resourcemodel.FirewallIPvfourNameRuleDestination{
			NodeFirewallIPvfourNameRuleDestinationGroup: &resourcemodel.FirewallIPvfourNameRuleDestinationGroup{
				LeafFirewallIPvfourNameRuleDestinationGroupPortGroup: basetypes.NewStringValue("Web"),
			},
		},
		LeafFirewallIPvfourNameRuleProtocol: basetypes.NewStringValue("tcp"),
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
	err := create(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Create failed: %v", err)
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

// TestCrudCreateResourceAlreadyExsitsFailure test CRUD helper: Create
//
//	Default situation where the resource already exists and we fail because of it
func TestCrudCreateResourceAlreadyExsitsFailure(t *testing.T) {
	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// Parent check API call
	exchangeParentExistsCheck := eList.Add()
	exchangeParentExistsCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","Test-Fw-Name"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": true,
			"error": null
		}`,
	)

	// Self check API call
	exchangeExistingResourceCheck := eList.Add()
	exchangeExistingResourceCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","Test-Fw-Name","rule","42"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": true,
			"error": null
		}`,
	)

	// From resource model
	model := &resourcemodel.FirewallIPvfourNameRule{
		ParentIDFirewallIPvfourName: basetypes.NewStringValue("Test-Fw-Name"),
		SelfIdentifier:              basetypes.NewNumberValue(big.NewFloat(42)),

		LeafFirewallIPvfourNameRuleAction:      basetypes.NewStringValue("accept"),
		LeafFirewallIPvfourNameRuleDescrIPtion: basetypes.NewStringValue("Allow http outgoing traffic"),
		NodeFirewallIPvfourNameRuleDestination: &resourcemodel.FirewallIPvfourNameRuleDestination{
			NodeFirewallIPvfourNameRuleDestinationGroup: &resourcemodel.FirewallIPvfourNameRuleDestinationGroup{
				LeafFirewallIPvfourNameRuleDestinationGroupPortGroup: basetypes.NewStringValue("Web"),
			},
		},
		LeafFirewallIPvfourNameRuleProtocol: basetypes.NewStringValue("tcp"),
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
	err := create(ctx, providerData, client, model)
	if err == nil {
		t.Errorf("Should have failed to create existing resource")
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

// TestCrudCreateResourceAlreadyExsitsIgnore test CRUD helper: Create
//
//	Configure provider to ignore an existing resource and overwrite it
func TestCrudCreateResourceAlreadyExsitsIgnore(t *testing.T) {
	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// Parent check API call
	exchangeParentExistsCheck := eList.Add()
	exchangeParentExistsCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","Test-Fw-Name"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": true,
			"error": null
		}`,
	)

	// Create resource API call
	exchangeCreateResource := eList.Add()
	exchangeCreateResource.Expect(
		"/configure",
		apiKey,
		`[`+
			`{"op":"set","path":["firewall","ipv4","name","Test-Fw-Name","rule","42","action","accept"]},`+
			`{"op":"set","path":["firewall","ipv4","name","Test-Fw-Name","rule","42","description","Allow http outgoing traffic"]},`+
			`{"op":"set","path":["firewall","ipv4","name","Test-Fw-Name","rule","42","protocol","tcp"]},`+
			`{"op":"set","path":["firewall","ipv4","name","Test-Fw-Name","rule","42","destination","group","port-group","Web"]}`+
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
	model := &resourcemodel.FirewallIPvfourNameRule{
		ParentIDFirewallIPvfourName: basetypes.NewStringValue("Test-Fw-Name"),
		SelfIdentifier:              basetypes.NewNumberValue(big.NewFloat(42)),

		LeafFirewallIPvfourNameRuleAction:      basetypes.NewStringValue("accept"),
		LeafFirewallIPvfourNameRuleDescrIPtion: basetypes.NewStringValue("Allow http outgoing traffic"),
		NodeFirewallIPvfourNameRuleDestination: &resourcemodel.FirewallIPvfourNameRuleDestination{
			NodeFirewallIPvfourNameRuleDestinationGroup: &resourcemodel.FirewallIPvfourNameRuleDestinationGroup{
				LeafFirewallIPvfourNameRuleDestinationGroupPortGroup: basetypes.NewStringValue("Web"),
			},
		},
		LeafFirewallIPvfourNameRuleProtocol: basetypes.NewStringValue("tcp"),
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
	providerData.Config.CrudSkipExistingResourceCheck = true

	// Execute test
	err := create(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Create failed: %v", err)
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

// TestCrudCreateResourceParentMissingFailure test CRUD helper: Create
//
//	Default situation where we attempt to create a resource when
//	the parent resource is missing and we fail because of it
func TestCrudCreateResourceParentMissingFailure(t *testing.T) {

	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// Parent check API call
	exchangeParentExistsCheck := eList.Add()
	exchangeParentExistsCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","Test-Fw-Name"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": false,
			"error": null
		}`,
	)

	// From resource model
	model := &resourcemodel.FirewallIPvfourNameRule{
		ParentIDFirewallIPvfourName: basetypes.NewStringValue("Test-Fw-Name"),
		SelfIdentifier:              basetypes.NewNumberValue(big.NewFloat(42)),

		LeafFirewallIPvfourNameRuleAction:      basetypes.NewStringValue("accept"),
		LeafFirewallIPvfourNameRuleDescrIPtion: basetypes.NewStringValue("Allow http outgoing traffic"),
		NodeFirewallIPvfourNameRuleDestination: &resourcemodel.FirewallIPvfourNameRuleDestination{
			NodeFirewallIPvfourNameRuleDestinationGroup: &resourcemodel.FirewallIPvfourNameRuleDestinationGroup{
				LeafFirewallIPvfourNameRuleDestinationGroupPortGroup: basetypes.NewStringValue("Web"),
			},
		},
		LeafFirewallIPvfourNameRuleProtocol: basetypes.NewStringValue("tcp"),
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
	err := create(ctx, providerData, client, model)
	if err == nil {
		t.Errorf("Should have failed to create resource without parent")
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

// TestCrudCreateResourceParentMissingIgnore test CRUD helper: Create
//
//	Configure provider to ignore the missing parent resource and go ahead
//	and create the resource anyway
func TestCrudCreateResourceParentMissingIgnore(t *testing.T) {

	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// Self check API call
	exchangeExistingResourceCheck := eList.Add()
	exchangeExistingResourceCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","Test-Fw-Name","rule","42"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": false,
			"error": null
		}`,
	)

	// Create resource API call
	exchangeCreateResource := eList.Add()
	exchangeCreateResource.Expect(
		"/configure",
		apiKey,
		`[`+
			`{"op":"set","path":["firewall","ipv4","name","Test-Fw-Name","rule","42","action","accept"]},`+
			`{"op":"set","path":["firewall","ipv4","name","Test-Fw-Name","rule","42","description","Allow http outgoing traffic"]},`+
			`{"op":"set","path":["firewall","ipv4","name","Test-Fw-Name","rule","42","protocol","tcp"]},`+
			`{"op":"set","path":["firewall","ipv4","name","Test-Fw-Name","rule","42","destination","group","port-group","Web"]}`+
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
	model := &resourcemodel.FirewallIPvfourNameRule{
		ParentIDFirewallIPvfourName: basetypes.NewStringValue("Test-Fw-Name"),
		SelfIdentifier:              basetypes.NewNumberValue(big.NewFloat(42)),

		LeafFirewallIPvfourNameRuleAction:      basetypes.NewStringValue("accept"),
		LeafFirewallIPvfourNameRuleDescrIPtion: basetypes.NewStringValue("Allow http outgoing traffic"),
		NodeFirewallIPvfourNameRuleDestination: &resourcemodel.FirewallIPvfourNameRuleDestination{
			NodeFirewallIPvfourNameRuleDestinationGroup: &resourcemodel.FirewallIPvfourNameRuleDestinationGroup{
				LeafFirewallIPvfourNameRuleDestinationGroupPortGroup: basetypes.NewStringValue("Web"),
			},
		},
		LeafFirewallIPvfourNameRuleProtocol: basetypes.NewStringValue("tcp"),
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
	providerData.Config.CrudSkipCheckParentBeforeCreate = true

	// Execute test
	err := create(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Create failed: %v", err)
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
