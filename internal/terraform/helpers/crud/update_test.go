package crud

import (
	"context"
	"net/http"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	fwgroup "github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/group-port-group/resourcemodel"
	fwname "github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/ipv4-name/resourcemodel"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/tests/api"
)

// TestCrudUpdate test CRUD helper: Update
// func TestCrudUpdate(t *testing.T) {
// 	ctx := context.Background()

// 	// When Mock API Server
// 	address := "localhost:8081"
// 	uri := "/retrieve"
// 	key := "test-key"
// 	srv := &http.Server{
// 		Addr: address,
// 	}

// 	eList := api.NewExchangeList()

// 	e1 := eList.Add()
// 	e1.Expect(uri, key, `{"op":"showConfig","path":["firewall","ipv4","name","SrvMain-WanTelia"]}`)
// 	e1.Response(
// 		200,
// 		`{
// 				"success": true, "data": {
// 					"default-action": "reject",
// 					"default-log": {},
// 					"description": "Managed by terraform",
// 					"rule": {
// 						"1": {"action": "accept", "description": "Allow established connections", "protocol": "all", "state": "established"},
// 						"2": {"action": "accept", "description": "Allow related connections", "protocol": "all", "state": "related"},
// 						"3": {"action": "drop", "description": "Disallow invalid packets", "log": {}, "protocol": "all", "state": "invalid"},
// 						"1000": {"action": "accept", "description": "Allow http outgoing traffic", "destination": {"group": {"port-group": "Web"}}, "protocol": "tcp"}
// 					}
// 				},
// 				"error": null
// 			}`,
// 	)

// 	api.ApiServer(srv, eList)

// 	// When API Client
// 	client := client.NewClient(ctx, "http://"+address, key, "test-agent", true)

// 	// With NewFirewallIPvfourName
// 	model := &fwname.FirewallIPvfourName{
// 		SelfIdentifier: basetypes.NewStringValue("SrvMain-WanTelia"),
// 	}

// 	// Should
// 	modelShould := &fwname.FirewallIPvfourName{
// 		SelfIdentifier:                       basetypes.NewStringValue("SrvMain-WanTelia"),
// 		LeafFirewallIPvfourNameDefaultAction: basetypes.NewStringValue("reject"),
// 		LeafFirewallIPvfourNameDefaultLog:    basetypes.NewBoolValue(true),
// 		LeafFirewallIPvfourNameDescrIPtion:   basetypes.NewStringValue("Managed by terraform"),
// 		ExistsTagFirewallIPvfourNameRule:     true,
// 	}

// 	err := read(ctx, &client, model)
// 	if err != nil {
// 		t.Errorf("read failed: %v", err)
// 	}

// 	deep.CompareUnexportedFields = true
// 	deep.LogErrors = true
// 	diff := deep.Equal(model, modelShould)
// 	if diff != nil {
// 		t.Errorf("compare failed: %v", diff)
// 	}
// }

// TestCrudUpdateCrossResourceContamination test CRUD helper: Update
// Had some issues where the client is shared across resources causing a clusterfuck during api commit
func TestCrudUpdateCrossResourceContamination(t *testing.T) {
	ctx := context.Background()

	// When Mock API Server
	address := "localhost:8081"
	uri := "/configure"
	key := "test-key"
	srv := &http.Server{
		Addr: address,
	}

	// 	curl -k --location --request POST "https://$VYOS_HOST/configure" --form key="$VYOS_KEY" --form data='[
	//     {"op":"delete","path":["firewall","ipv4","name","TF-Example","description","Another tf test"]},
	//     {"op":"set","path":["firewall","ipv4","name","TF-Example","description","Another tf test"]}
	// ]'
	// {"success": true, "data": null, "error": null}

	// $ curl -k --location --request POST "https://$VYOS_HOST/configure" --form key="$VYOS_KEY" --form data='[
	//     {"op":"delete","path":["firewall","group","port-group","TF-Examples", "description"]},
	//     {"op":"set","path":["firewall","group","port-group","TF-Examples", "description", "blabla"]}
	// ]'
	// {"success": true, "data": null, "error": null}

	// to minify json:
	// echo '...' | jq -c

	eList := api.NewExchangeList()

	e1 := eList.Add()
	e1.Expect(uri, key, `[{"op":"delete","path":["firewall","group","port-group","TF-grp","description","pre-grp"]},{"op":"set","path":["firewall","group","port-group","TF-grp","description","post-grp"]}]`)
	e1.Response(200, `{"success": true, "data": null, "error": null}`)

	e2 := eList.Add()
	e2.Expect(uri, key, `[{"op":"delete","path":["firewall","ipv4","name","TF-fw","description","pre-fw"]},{"op":"set","path":["firewall","ipv4","name","TF-fw","description","post-fw"]}]`)
	e2.Response(200, `{"success": true, "data": null, "error": null}`)

	api.Server(srv, eList)

	// When API Client
	providerDataClient := client.NewClient(ctx, "http://"+address, key, "test-agent", true)

	// With
	groupState := &fwgroup.FirewallGroupPortGroup{
		SelfIdentifier:                        basetypes.NewStringValue("TF-grp"),
		LeafFirewallGroupPortGroupDescrIPtion: basetypes.NewStringValue("pre-grp"),
	}
	fwState := &fwname.FirewallIPvfourName{
		SelfIdentifier:                     basetypes.NewStringValue("TF-fw"),
		LeafFirewallIPvfourNameDescrIPtion: basetypes.NewStringValue("pre-fw"),
	}

	// Should
	groupPlan := &fwgroup.FirewallGroupPortGroup{
		SelfIdentifier:                        basetypes.NewStringValue("TF-grp"),
		LeafFirewallGroupPortGroupDescrIPtion: basetypes.NewStringValue("post-grp"),
	}
	fwPlan := &fwname.FirewallIPvfourName{
		SelfIdentifier:                     basetypes.NewStringValue("TF-fw"),
		LeafFirewallIPvfourNameDescrIPtion: basetypes.NewStringValue("post-fw"),
	}

	// Run
	var wg sync.WaitGroup

	client1 := providerDataClient
	client2 := providerDataClient

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := update(ctx, &client1, groupState, groupPlan)
		if err != nil {
			t.Logf("[grp] update error: %v", err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := update(ctx, &client2, fwState, fwPlan)
		if err != nil {
			t.Logf("[fw] compare failed: %v", err.Error())
		}
	}()

	wg.Wait()

	if len(eList.Unmatched()) > 0 {
		for _, e := range eList.Unmatched() {
			t.Errorf("Unmatched exchange: %#v", e)
		}
		t.Errorf("Have %d unmatched exchanges", len(eList.Unmatched()))
	}
}
