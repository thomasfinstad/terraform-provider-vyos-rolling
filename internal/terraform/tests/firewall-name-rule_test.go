package tests

import (
	"context"
	"encoding/json"
	"math/big"
	"reflect"
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/name-rule/resourcemodel"
)

// MarshalVyos calls greetings.Hello with a name, checking
// for a valid return value.
func TestFirewallNameRuleMarshalVyos(t *testing.T) {
	lst, _ := basetypes.NewListValue(basetypes.StringType{}, []attr.Value{basetypes.NewStringValue("420"), basetypes.NewStringValue("13-37")})

	model := &resourcemodel.FirewallNameRule{
		ID:                                      basetypes.NewNumberValue(big.NewFloat(42)),
		ParentIDFirewallName:                    basetypes.NewStringValue("rule-one"),
		LeafFirewallNameRuleAction:              basetypes.NewStringValue("accept"),
		LeafFirewallNameRuleDisable:             basetypes.NewBoolValue(true),
		LeafFirewallNameRuleQueue:               basetypes.NewNumberValue(big.NewFloat(28)),
		LeafFirewallNameRulePacketLengthExclude: lst,
		NodeFirewallNameRuleDestination: &resourcemodel.FirewallNameRuleDestination{
			LeafFirewallNameRuleDestinationAddress: basetypes.NewStringValue("127.0.0.2"),
		},
	}

	want, err := json.Marshal(
		map[string]interface{}{
			"action":                "accept",
			"disable":               map[string]interface{}{},
			"queue":                 28,
			"packet-length-exclude": []string{"420", "13-37"},
			"destination": map[string]interface{}{
				"address": "127.0.0.2",
			},
		},
	)
	if err != nil {
		t.Fatalf(`desired value can not be jsonified: %v`, err)
	}

	ctx := context.Background()
	lvl := tflog.WithLevel(hclog.Error)
	ctx = tflog.NewSubsystem(ctx, "my-subsystem", lvl)

	result, err := helpers.MarshalVyos(ctx, model)
	if !reflect.DeepEqual(result, want) {
		t.Fatalf(`model.MarshalVyos() = %#q, %v, want match for %#q, nil`, result, err, want)
	}
}
