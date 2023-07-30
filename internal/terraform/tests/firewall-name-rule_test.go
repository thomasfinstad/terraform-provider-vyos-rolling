package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/name-rule/resourcemodel"
)

// MarshalVyos calls greetings.Hello with a name, checking
// for a valid return value.
func TestFirewallNameRuleMarshalVyos(t *testing.T) {
	model := &resourcemodel.FirewallNameRule{
		ID:                         basetypes.NewStringValue("test-id"),
		ParentIDFirewallName:       basetypes.NewStringValue("rule-one"),
		LeafFirewallNameRuleAction: basetypes.NewStringValue("accept"),
		NodeFirewallNameRuleDestination: &resourcemodel.FirewallNameRuleDestination{
			LeafFirewallNameRuleDestinationAddress: basetypes.NewStringValue("127.0.0.2"),
		},
	}

	want, err := json.Marshal(
		map[string]interface{}{
			"action": "accept",
			"destination": map[string]interface{}{
				"address": "127.0.0.2",
			},
		},
	)
	if err != nil {
		t.Fatalf(`desired value can not be jsonified: %v`, err)
	}

	result, err := helpers.MarshalVyos(model)
	if !reflect.DeepEqual(result, want) {
		t.Fatalf(`model.MarshalVyos() = %q, %v, want match for %#q, nil`, result, err, want)
	}
}
