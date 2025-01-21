package tests

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/go-test/deep"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflogtest"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/firewall/ipv4-name-rule/resourcemodel"
)

// TestFirewallIPvfourNameRuleMarshalVyos does some simple marshalling tests
func TestFirewallIPvfourNameRuleMarshalVyos(t *testing.T) {
	lst, _ := basetypes.NewListValue(basetypes.StringType{}, []attr.Value{basetypes.NewStringValue("420"), basetypes.NewStringValue("13-37")})
	model := &resourcemodel.FirewallIPvfourNameRule{
		SelfIdentifier: &resourcemodel.FirewallIPvfourNameRuleSelfIdentifier{
			FirewallIPvfourNameRule: basetypes.NewNumberValue(big.NewFloat(42)),
			FirewallIPvfourName:     basetypes.NewStringValue("rule-one"),
		},
		LeafFirewallIPvfourNameRuleAction:              basetypes.NewStringValue("accept"),
		LeafFirewallIPvfourNameRuleDisable:             basetypes.NewBoolValue(true),
		LeafFirewallIPvfourNameRuleQueue:               basetypes.NewNumberValue(big.NewFloat(28)),
		LeafFirewallIPvfourNameRulePacketLengthExclude: lst,
		NodeFirewallIPvfourNameRuleDestination: &resourcemodel.FirewallIPvfourNameRuleDestination{
			LeafFirewallIPvfourNameRuleDestinationAddress: basetypes.NewStringValue("127.0.0.2"),
		},
	}

	want := map[string]interface{}{
		// "id":                    "firewall__name__rule-one__rule__42",
		"action":                "accept",
		"disable":               map[string]interface{}{},
		"queue":                 float64(28),
		"packet-length-exclude": []string{"420", "13-37"},
		"destination": map[string]interface{}{
			"address": "127.0.0.2",
		},
	}

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	got, err := helpers.MarshalVyos(ctx, model)
	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("compare failed: %v, %v", diff, err)
	}
}

// TestFirewallIPvfourNameRuleMarshalVyos does some simple marshalling tests
func TestFirewallIPvfourNameRuleUnmarshalVyos(t *testing.T) {
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
	want := &resourcemodel.FirewallIPvfourNameRule{
		LeafFirewallIPvfourNameRuleAction:              basetypes.NewStringValue("accept"),
		LeafFirewallIPvfourNameRuleDisable:             basetypes.NewBoolValue(true),
		LeafFirewallIPvfourNameRuleQueue:               basetypes.NewNumberValue(big.NewFloat(28)),
		LeafFirewallIPvfourNameRulePacketLengthExclude: lst,
		NodeFirewallIPvfourNameRuleDestination: &resourcemodel.FirewallIPvfourNameRuleDestination{
			LeafFirewallIPvfourNameRuleDestinationAddress: basetypes.NewStringValue("127.0.0.2"),
		},
	}

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	got := &resourcemodel.FirewallIPvfourNameRule{}

	err := helpers.UnmarshalVyos(ctx, has, got)
	if err != nil {
		t.Fatalf(`desired value can not be unmarshaled: %v`, err)
	}

	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("compare failed: %v", diff)
	}
}
