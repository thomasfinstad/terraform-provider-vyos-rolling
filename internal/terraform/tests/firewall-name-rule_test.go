package tests

import (
	"context"
	"math/big"
	"testing"

	"github.com/go-test/deep"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/name-rule/resourcemodel"
)

// TestFirewallNameRuleMarshalVyos does some simple marshalling tests
func TestFirewallNameRuleMarshalVyos(t *testing.T) {
	lst, _ := basetypes.NewListValue(basetypes.StringType{}, []attr.Value{basetypes.NewStringValue("420"), basetypes.NewStringValue("13-37")})
	model := &resourcemodel.FirewallNameRule{
		SelfIdentifier:                          basetypes.NewNumberValue(big.NewFloat(42)),
		ParentIDFirewallName:                    basetypes.NewStringValue("rule-one"),
		LeafFirewallNameRuleAction:              basetypes.NewStringValue("accept"),
		LeafFirewallNameRuleDisable:             basetypes.NewBoolValue(true),
		LeafFirewallNameRuleQueue:               basetypes.NewNumberValue(big.NewFloat(28)),
		LeafFirewallNameRulePacketLengthExclude: lst,
		NodeFirewallNameRuleDestination: &resourcemodel.FirewallNameRuleDestination{
			LeafFirewallNameRuleDestinationAddress: basetypes.NewStringValue("127.0.0.2"),
		},
	}

	want := map[string]interface{}{
		"action":                "accept",
		"disable":               map[string]interface{}{},
		"queue":                 float64(28),
		"packet-length-exclude": []string{"420", "13-37"},
		"destination": map[string]interface{}{
			"address": "127.0.0.2",
		},
	}

	ctx := context.Background()
	lvl := tflog.WithLevel(hclog.Error)
	ctx = tflog.NewSubsystem(ctx, "my-subsystem", lvl)

	got, err := helpers.MarshalVyos(ctx, model)
	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("compare failed: %v, %v", diff, err)
	}
}

// TestFirewallNameRuleMarshalVyos does some simple marshalling tests
func TestFirewallNameRuleUnmarshalVyos(t *testing.T) {
	has := map[string]interface{}{
		"action":                "accept",
		"disable":               map[string]interface{}{},
		"queue":                 float64(28),
		"packet-length-exclude": []string{"420", "13-37"},
		"destination": map[string]interface{}{
			"address": "127.0.0.2",
		},
	}

	lst, _ := basetypes.NewListValue(basetypes.StringType{}, []attr.Value{basetypes.NewStringValue("420"), basetypes.NewStringValue("13-37")})
	want := &resourcemodel.FirewallNameRule{
		LeafFirewallNameRuleAction:              basetypes.NewStringValue("accept"),
		LeafFirewallNameRuleDisable:             basetypes.NewBoolValue(true),
		LeafFirewallNameRuleQueue:               basetypes.NewNumberValue(big.NewFloat(28)),
		LeafFirewallNameRulePacketLengthExclude: lst,
		NodeFirewallNameRuleDestination: &resourcemodel.FirewallNameRuleDestination{
			LeafFirewallNameRuleDestinationAddress: basetypes.NewStringValue("127.0.0.2"),
		},
	}

	got := &resourcemodel.FirewallNameRule{}

	err := helpers.UnmarshalVyos(context.Background(), has, got)
	if err != nil {
		t.Fatalf(`desired value can not be unmarshaled: %v`, err)
	}

	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("compare failed: %v", diff)
	}
}
