package helpers_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	serviceDnsForwarding "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/service/dns-forwarding/resourcemodel"
	firewallBridgeInputFilterRule "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/firewall/bridge-input-filter-rule/resourcemodel"
)

func TestModifiersUnknownToNull(t *testing.T) {

	m := firewallBridgeInputFilterRule.FirewallBrIDgeInputFilterRule{
		LeafFirewallBrIDgeInputFilterRuleQueue: basetypes.NewNumberUnknown(),
		NodeFirewallBrIDgeInputFilterRuleLimit: &firewallBridgeInputFilterRule.FirewallBrIDgeInputFilterRuleLimit{
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

func TestModifiersUnknownToNullListInValue(t *testing.T) {
	var diags diag.Diagnostics
	_LeafServiceDNSForwardingAllowFrom, diag := basetypes.NewListValue(basetypes.StringType{}, []attr.Value{basetypes.NewStringValue("192.168.0.0/16")})
	diags = append(diags, diag...)
	_LeafServiceDNSForwardingListenAddress, diag := basetypes.NewListValue(basetypes.StringType{}, []attr.Value{basetypes.NewStringValue("192.168.2.253")})
	diags = append(diags, diag...)

	for _, d := range diags {
		t.Logf("TF diagnostic issue: %s", d)
		t.Fail()
	}

	m := serviceDnsForwarding.ServiceDNSForwarding{
		LeafServiceDNSForwardingAllowFrom:       _LeafServiceDNSForwardingAllowFrom,
		LeafServiceDNSForwardingIgnoreHostsFile: basetypes.NewBoolValue(false),
		LeafServiceDNSForwardingListenAddress:   _LeafServiceDNSForwardingListenAddress,
		LeafServiceDNSForwardingSourceAddress:   basetypes.NewListUnknown(basetypes.StringType{}),
	}

	helpers.UnknownToNull(context.Background(), &m)
}
