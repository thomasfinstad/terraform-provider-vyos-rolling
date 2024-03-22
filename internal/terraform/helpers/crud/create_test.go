package crud

import (
	"context"
	"errors"
	"math/big"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflogtest"
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
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudCreateSuccess"]}`,
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
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudCreateSuccess","rule","42"]}`,
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
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateSuccess","rule","42","action","accept"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateSuccess","rule","42","description","Allow http outgoing traffic"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateSuccess","rule","42","protocol","tcp"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateSuccess","rule","42","destination","group","port-group","Web"]}`+
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
		ParentIDFirewallIPvfourName: basetypes.NewStringValue("TestCrudCreateSuccess"),
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
	apiAddress := "localhost:50001"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)

	// Execute test
	err := create(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Create failed: %v", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
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
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudCreateResourceAlreadyExsitsFailure"]}`,
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
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudCreateResourceAlreadyExsitsFailure","rule","42"]}`,
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
		ParentIDFirewallIPvfourName: basetypes.NewStringValue("TestCrudCreateResourceAlreadyExsitsFailure"),
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
	apiAddress := "localhost:50002"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)

	// Execute test
	err := create(ctx, providerData, client, model)
	if err == nil {
		t.Errorf("Should have failed to create existing resource")
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
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
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudCreateResourceAlreadyExsitsIgnore"]}`,
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
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateResourceAlreadyExsitsIgnore","rule","42","action","accept"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateResourceAlreadyExsitsIgnore","rule","42","description","Allow http outgoing traffic"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateResourceAlreadyExsitsIgnore","rule","42","protocol","tcp"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateResourceAlreadyExsitsIgnore","rule","42","destination","group","port-group","Web"]}`+
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
		ParentIDFirewallIPvfourName: basetypes.NewStringValue("TestCrudCreateResourceAlreadyExsitsIgnore"),
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
	apiAddress := "localhost:50003"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)
	providerData.Config.CrudSkipExistingResourceCheck = true

	// Execute test
	err := create(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Create failed: %v", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
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
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudCreateResourceParentMissingFailure"]}`,
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
		ParentIDFirewallIPvfourName: basetypes.NewStringValue("TestCrudCreateResourceParentMissingFailure"),
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
	apiAddress := "localhost:50004"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)

	// Execute test
	err := create(ctx, providerData, client, model)
	if err == nil {
		t.Errorf("Should have failed to create resource without parent")
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
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
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudCreateResourceParentMissingIgnore","rule","42"]}`,
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
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateResourceParentMissingIgnore","rule","42","action","accept"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateResourceParentMissingIgnore","rule","42","description","Allow http outgoing traffic"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateResourceParentMissingIgnore","rule","42","protocol","tcp"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateResourceParentMissingIgnore","rule","42","destination","group","port-group","Web"]}`+
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
		ParentIDFirewallIPvfourName: basetypes.NewStringValue("TestCrudCreateResourceParentMissingIgnore"),
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
	apiAddress := "localhost:50014"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)
	providerData.Config.CrudSkipCheckParentBeforeCreate = true

	// Execute test
	err := create(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Create failed: %v", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
	}
}

// TestCrudCreateTimeoutSuccess test CRUD helper: Create
//
//	Verify that resource will wait for configured timeout
func TestCrudCreateTimeoutSuccess(t *testing.T) {
	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// Parent check API call
	exchangeParentExistsCheck := eList.Add()
	exchangeParentExistsCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudCreateTimeoutSuccess"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": true,
			"error": null
		}`,
	).Delay(10 * time.Millisecond)

	// Self check API call
	exchangeExistingResourceCheck := eList.Add()
	exchangeExistingResourceCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudCreateTimeoutSuccess","rule","42"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": false,
			"error": null
		}`,
	).Delay(10 * time.Millisecond)

	// Create resource API call
	exchangeCreateResource := eList.Add()
	exchangeCreateResource.Expect(
		"/configure",
		apiKey,
		`[`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateTimeoutSuccess","rule","42","action","accept"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateTimeoutSuccess","rule","42","description","Allow http outgoing traffic"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateTimeoutSuccess","rule","42","protocol","tcp"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateTimeoutSuccess","rule","42","destination","group","port-group","Web"]}`+
			`]`,
	).Response(
		200,
		`{
			"success": true,
			"data": null,
			"error": null
		}`,
	).Delay(10 * time.Millisecond)

	// From resource model
	model := &resourcemodel.FirewallIPvfourNameRule{

		Timeouts: timeouts.Value{
			Object: types.ObjectValueMust(
				map[string]attr.Type{
					"create": types.StringType,
					"read":   types.StringType,
					"update": types.StringType,
					"delete": types.StringType,
				},
				map[string]attr.Value{
					"create": types.StringValue("1s"),
					"read":   types.StringValue("1s"),
					"update": types.StringValue("1s"),
					"delete": types.StringValue("1s"),
				},
			),
		},

		ParentIDFirewallIPvfourName: basetypes.NewStringValue("TestCrudCreateTimeoutSuccess"),
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
	apiAddress := "localhost:50005"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	start := time.Now()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)

	// Execute test
	err := create(ctx, providerData, client, model)
	if err != nil {
		d, ok := ctx.Deadline()
		t.Error("Start        :", start)
		t.Error("Now          :", time.Now())
		t.Error("DeadLine Time:", d, ok)
		t.Error("DeadLine Left:", time.Until(d))
		t.Error("Since        :", time.Since(start))
		t.Error("Create failed:", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
	}
}

// TestCrudCreateTimeoutFailure test CRUD helper: Create
//
//	Verify that resource will timeout after configured time
func TestCrudCreateTimeoutFailure(t *testing.T) {
	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// Parent check API call
	exchangeParentExistsCheck := eList.Add()
	exchangeParentExistsCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudCreateTimeoutFailure"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": true,
			"error": null
		}`,
	).Delay(10 * time.Millisecond)

	// Self check API call
	exchangeExistingResourceCheck := eList.Add()
	exchangeExistingResourceCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudCreateTimeoutFailure","rule","42"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": false,
			"error": null
		}`,
	).Delay(10 * time.Millisecond)

	// Create resource API call
	exchangeCreateResource := eList.Add()
	exchangeCreateResource.Expect(
		"/configure",
		apiKey,
		`[`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateTimeoutFailure","rule","42","action","accept"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateTimeoutFailure","rule","42","description","Allow http outgoing traffic"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateTimeoutFailure","rule","42","protocol","tcp"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateTimeoutFailure","rule","42","destination","group","port-group","Web"]}`+
			`]`,
	).Response(
		200,
		`{
			"success": true,
			"data": null,
			"error": null
		}`,
	).Delay(1 * time.Second)

	// From resource model
	model := &resourcemodel.FirewallIPvfourNameRule{

		// This logic is handled in the public function,
		// we are testing the private function, and will
		// just send a context.Context with timeout
		//
		// Timeouts: timeouts.Value{
		// 	Object: types.ObjectValueMust(
		// 		map[string]attr.Type{
		// 			"create": types.StringType,
		// 			"read":   types.StringType,
		// 			"update": types.StringType,
		// 			"delete": types.StringType,
		// 		},
		// 		map[string]attr.Value{
		// 			"create": types.StringValue("2s"),
		// 			"read":   types.StringValue("2s"),
		// 			"update": types.StringValue("2s"),
		// 			"delete": types.StringValue("2s"),
		// 		},
		// 	),
		// },

		ParentIDFirewallIPvfourName: basetypes.NewStringValue("TestCrudCreateTimeoutFailure"),
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
	apiAddress := "localhost:50006"
	srv := &http.Server{
		Addr: apiAddress,
	}
	api.Server(srv, eList)

	// Client
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	client := client.NewClient(ctx, "http://"+apiAddress, apiKey, "test-agent", true)
	providerData := data.NewProviderData(client)

	// Execute test
	err := create(ctx, providerData, client, model)

	var netErr net.Error
	if !errors.As(err, &netErr) && netErr.Timeout() {
		t.Errorf("Create timeout failed: %v", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) != 1 {
		for _, e := range eList.Unmatched() {
			t.Errorf("Unmatched exchange:\n%s", e.Sexpect())
		}
		t.Errorf("Expected 1 unmatched exchange. Total unmatched exchanges: %d", len(eList.Unmatched()))
	}
}

// TestCrudCreateRetrySuccess test CRUD helper: Create
//
// Default situation where we attempt to create a resource where the parent exists,
// but the resource is in the process of being destroyed so we must wait and retry
// until the resource has been removed so we can re-create it
func TestCrudCreateRetrySuccess(t *testing.T) {
	// API mocking
	eList := api.NewExchangeList()
	apiKey := "test-key"

	// Parent check API call
	exchangeParentExistsCheck := eList.Add()
	exchangeParentExistsCheck.Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudCreateRetrySuccess"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": true,
			"error": null
		}`,
	)

	// Self check API call: before remove x3
	for range 3 {
		exchangeExistingResourceCheckBefore := eList.Add()
		exchangeExistingResourceCheckBefore.Expect(
			"/retrieve",
			apiKey,
			`{"op":"exists","path":["firewall","ipv4","name","TestCrudCreateRetrySuccess","rule","42"]}`,
		).Response(
			200,
			`{
			"success": true,
			"data": true,
			"error": null
		}`,
		)
	}

	// Self check API call: after remove
	exchangeExistingResourceCheckAfter := eList.Add()
	exchangeExistingResourceCheckAfter.Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudCreateRetrySuccess","rule","42"]}`,
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
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateRetrySuccess","rule","42","action","accept"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateRetrySuccess","rule","42","description","Allow http outgoing traffic"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateRetrySuccess","rule","42","protocol","tcp"]},`+
			`{"op":"set","path":["firewall","ipv4","name","TestCrudCreateRetrySuccess","rule","42","destination","group","port-group","Web"]}`+
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
		ParentIDFirewallIPvfourName: basetypes.NewStringValue("TestCrudCreateRetrySuccess"),
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
	apiAddress := "localhost:50001"
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
	err := create(ctx, providerData, client, model)
	if err != nil {
		t.Errorf("Create failed: %v", err)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
	}
}
