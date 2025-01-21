package crud

import (
	"context"
	"math/big"
	"net/http"
	"os"
	"testing"

	"github.com/go-test/deep"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflogtest"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	fw4res "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/firewall/ipv4-name/resourcemodel"
	polalres "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/policy/access-list/resourcemodel"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/tests/api"
)

// TestCrudReadSuccess test CRUD helper: Read
func TestCrudReadSuccess(t *testing.T) {
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	// When Mock API Server
	address := "localhost:50012"
	uri := "/retrieve"
	apiKey := "test-key"
	srv := &http.Server{
		Addr: address,
	}

	eList := api.NewExchangeList()

	// Resource exists check API call
	eList.Add().Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["firewall","ipv4","name","TestCrudReadSuccess"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": true,
			"error": null
		}`,
	)

	e1 := eList.Add()
	e1.Expect(
		uri,
		apiKey,
		`{"op":"showConfig","path":["firewall","ipv4","name","TestCrudReadSuccess"]}`,
	).Response(
		200,
		`{
				"success": true, "data": {
					"default-action": "reject",
					"default-log": {},
					"description": "Managed by terraform",
					"rule": {
						"1": {"action": "accept", "description": "Allow established connections", "protocol": "all", "state": "established"},
						"2": {"action": "accept", "description": "Allow related connections", "protocol": "all", "state": "related"},
						"3": {"action": "drop", "description": "Disallow invalid packets", "log": {}, "protocol": "all", "state": "invalid"},
						"1000": {"action": "accept", "description": "Allow http outgoing traffic", "destination": {"group": {"port-group": "Web"}}, "protocol": "tcp"}
					}
				},
				"error": null
			}`,
	)

	api.Server(srv, eList)

	// When API Client
	client := client.NewClient(ctx, "http://"+address, apiKey, "test-agent", true)

	// With Policy Access List
	model := &fw4res.FirewallIPvfourName{
		SelfIdentifier: &fw4res.FirewallIPvfourNameSelfIdentifier{FirewallIPvfourName: basetypes.NewStringValue("TestCrudReadSuccess")},
	}

	// Should
	modelShould := &fw4res.FirewallIPvfourName{
		SelfIdentifier: &fw4res.FirewallIPvfourNameSelfIdentifier{FirewallIPvfourName: basetypes.NewStringValue("TestCrudReadSuccess")},

		LeafFirewallIPvfourNameDefaultAction: basetypes.NewStringValue("reject"),
		LeafFirewallIPvfourNameDefaultLog:    basetypes.NewBoolValue(true),
		LeafFirewallIPvfourNameDescrIPtion:   basetypes.NewStringValue("Managed by terraform"),
		ExistsTagFirewallIPvfourNameRule:     true,
	}

	err := read(ctx, client, model)
	if err != nil {
		t.Errorf("read failed: %v", err)
	}

	deep.CompareUnexportedFields = true
	deep.LogErrors = true
	diff := deep.Equal(model, modelShould)
	if diff != nil {
		t.Errorf("compare failed: %v", diff)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
	}
}

// TestCrudReadEmptyGlobalResource tests that an empty resource response from API is checked correctly
//
// curl -k --location --request POST "https://$VYOS_HOST/retrieve" --form key="$VYOS_KEY" --form data='{"op":"showConfig","path":["policy","access-list","2"]}'
//
//	Simulated resource:
//	resource "vyos_policy_access_list" "name" {
//		access_list_id = 42
//	}
func TestCrudReadEmptyGlobalResource(t *testing.T) {
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	// When Mock API Server
	address := "localhost:50013"
	apiKey := "test-key"
	srv := &http.Server{
		Addr: address,
	}

	eList := api.NewExchangeList()

	// Resource exists check API call
	eList.Add().Expect(
		"/retrieve",
		apiKey,
		`{"op":"exists","path":["policy","access-list","42"]}`,
	).Response(
		200,
		`{
			"success": true,
			"data": true,
			"error": null
		}`,
	)

	// Initial retrieve request
	e1 := eList.Add()
	e1.Expect(
		"/retrieve",
		apiKey,
		`{"op":"showConfig","path":["policy","access-list","42"]}`,
	).Response(
		400,
		`{"success": false, "error": "Configuration under specified path is empty\n", "data": null}`,
	)

	api.Server(srv, eList)

	// When API Client
	client := client.NewClient(ctx, "http://"+address, apiKey, "test-agent", true)

	// With Policy Access List
	model := &polalres.PolicyAccessList{
		SelfIdentifier: &polalres.PolicyAccessListSelfIdentifier{PolicyAccessList: basetypes.NewNumberValue(big.NewFloat(42))},
	}

	// Should
	modelShould := &polalres.PolicyAccessList{
		SelfIdentifier: &polalres.PolicyAccessListSelfIdentifier{PolicyAccessList: basetypes.NewNumberValue(big.NewFloat(42))},
	}

	// Execute test
	err := read(ctx, client, model)
	if err != nil {
		t.Errorf("read failed: %v", err)
	}

	deep.CompareUnexportedFields = true
	deep.LogErrors = true
	diff := deep.Equal(model, modelShould)
	if diff != nil {
		t.Errorf("compare failed: %v", diff)
	}

	// Validate API calls
	if len(eList.Unmatched()) > 0 {
		t.Logf("Total matched exchanges: %d", len(eList.Matched()))
		t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
		t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
		t.Errorf("Received request:\n%s", eList.Failed())
	}
}
