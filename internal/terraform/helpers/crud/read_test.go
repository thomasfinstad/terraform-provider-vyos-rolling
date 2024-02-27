package crud

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-test/deep"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/ipv4-name/resourcemodel"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/tests/api"
)

// TestCrudRead test CRUD helper: Read
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
	model := &resourcemodel.FirewallIPvfourName{
		SelfIdentifier: basetypes.NewStringValue("SrvMain-WanTelia"),
	}

	// Should
	modelShould := &resourcemodel.FirewallIPvfourName{
		SelfIdentifier:                       basetypes.NewStringValue("SrvMain-WanTelia"),
		LeafFirewallIPvfourNameDefaultAction: basetypes.NewStringValue("reject"),
		LeafFirewallIPvfourNameDefaultLog:    basetypes.NewBoolValue(true),
		LeafFirewallIPvfourNameDescrIPtion:   basetypes.NewStringValue("Managed by terraform"),
		ExistsTagFirewallIPvfourNameRule:     true,
	}

	err := read(ctx, &client, model)
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
