package crud

import (
	"context"
	"math/big"
	"net/http"
	"testing"

	"github.com/go-test/deep"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	fw4res "github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/ipv4-name/resourcemodel"
	polalres "github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/policy/access-list/resourcemodel"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/tests/api"
)

// TestCrudRead test CRUD helper: Read
//
// curl -k --location --request POST "https://$VYOS_HOST/retrieve" --form key="$VYOS_KEY" --form data='{"op":"showConfig","path":["firewall","ipv4","name","SrvMain-WanTelia"]}'
//
//	Pseudo simulated Resources:
//	resource "vyos_firewall_ipv4_name" "name" {
//		name_id = "SrvMain-WanTelia"
//		default_action = "reject"
//		description    = "Managed by terraform"
//		default_log    = true
//	}
//
//	resource "vyos_firewall_ipv4_name_rule" "..." {
//		name_id = "SrvMain-WanTelia"
//		...
//		rule_id = 1...
//		rule_id = 2...
//		rule_id = 3...
//		rule_id = 1000...
//		...
//	}

func TestCrudRead(t *testing.T) {
	ctx := context.Background()

	// When Mock API Server
	address := "localhost:8080"
	uri := "/retrieve"
	key := "test-key"
	srv := &http.Server{
		Addr: address,
	}

	eList := api.NewExchangeList()

	e1 := eList.Add()
	e1.Expect(
		uri,
		key,
		`{"op":"showConfig","path":["firewall","ipv4","name","SrvMain-WanTelia"]}`,
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
	client := client.NewClient(ctx, "http://"+address, key, "test-agent", true)

	// With NewFirewallIPvfourName
	model := &fw4res.FirewallIPvfourName{
		SelfIdentifier: basetypes.NewStringValue("SrvMain-WanTelia"),
	}

	// Should
	modelShould := &fw4res.FirewallIPvfourName{
		SelfIdentifier:                       basetypes.NewStringValue("SrvMain-WanTelia"),
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
}

// TestPolicyEmptyResponseFromApi tests that an empty resource response from API is checked correctly
//
// curl -k --location --request POST "https://$VYOS_HOST/retrieve" --form key="$VYOS_KEY" --form data='{"op":"showConfig","path":["policy","access-list","2"]}'
//
//	Simulated resource:
//	resource "vyos_policy_access_list" "name" {
//		access_list_id = 42
//	}
func TestPolicyEmptyResponseFromApi(t *testing.T) {
	ctx := context.Background()

	// When Mock API Server
	address := "localhost:8080"
	uri := "/retrieve"
	key := "test-key"
	srv := &http.Server{
		Addr: address,
	}

	eList := api.NewExchangeList()

	// Initial (normal) request
	e1 := eList.Add()
	e1.Expect(
		uri,
		key,
		`{"op":"showConfig","path":["policy","access-list","42"]}`,
	).Response(
		400,
		`{"success": false, "error": "Configuration under specified path is empty\n", "data": null}`,
	)

	// Dedicated (special case) request to check for empty resources
	e2 := eList.Add()
	e2.Expect(
		uri,
		key,
		`{"op":"showConfig","path":["policy"]}`,
	).Response(
		200,
		`{"success": true, "data": {"access-list": {"42": {}}}, "error": null}`,
	)

	api.Server(srv, eList)

	// When API Client
	client := client.NewClient(ctx, "http://"+address, key, "test-agent", true)

	// With NewFirewallIPvfourName
	model := &polalres.PolicyAccessList{
		SelfIdentifier: basetypes.NewNumberValue(big.NewFloat(42)),
	}

	// Should
	modelShould := &polalres.PolicyAccessList{
		SelfIdentifier: basetypes.NewNumberValue(big.NewFloat(42)),
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
}
