package helpers_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/firewall/bridge-input-filter-rule/resourcemodel"
)

func TestModifiersUnknownToNull(t *testing.T) {

	m := resourcemodel.FirewallBrIDgeInputFilterRule{
		LeafFirewallBrIDgeInputFilterRuleQueue: basetypes.NewNumberUnknown(),
		NodeFirewallBrIDgeInputFilterRuleLimit: &resourcemodel.FirewallBrIDgeInputFilterRuleLimit{
			LeafFirewallBrIDgeInputFilterRuleLimitBurst: basetypes.NewNumberUnknown(),
		},
	}

	helpers.UnknownToNull(context.Background(), &m)

	if m.LeafFirewallBrIDgeInputFilterRuleQueue.IsUnknown() {
		t.Errorf("Top level attribute is still unknown")
	}
	if !m.LeafFirewallBrIDgeInputFilterRuleQueue.IsNull() {
		t.Errorf("Top level attribute did not get converted to null")
	}

	if m.NodeFirewallBrIDgeInputFilterRuleLimit.LeafFirewallBrIDgeInputFilterRuleLimitBurst.IsUnknown() {
		t.Errorf("Nested attribute is still unknown")
	}
	if !m.NodeFirewallBrIDgeInputFilterRuleLimit.LeafFirewallBrIDgeInputFilterRuleLimitBurst.IsNull() {
		t.Errorf("Nested attribute did not get converted to null")
	}
}
