package tests

import (
	"context"
	"os"
	"testing"

	"github.com/go-test/deep"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflogtest"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/firewall/zone/resourcemodel"
)

// TestFirewallZoneMarshalVyos does some simple marshalling tests to prevent a bug with tags handling from reappearing
func TestFirewallZoneMarshalVyos(t *testing.T) {
	has := &resourcemodel.FirewallZone{
		// SelfIdentifier: basetypes.NewStringValue("Testing"),
		SelfIdentifier: basetypes.NewObjectNull(
			map[string]attr.Type{}),
		NodeFirewallZoneIntraZoneFiltering: &resourcemodel.FirewallZoneIntraZoneFiltering{
			LeafFirewallZoneIntraZoneFilteringAction: basetypes.NewStringValue("accept"),
			NodeFirewallZoneIntraZoneFilteringFirewall: &resourcemodel.FirewallZoneIntraZoneFilteringFirewall{
				LeafFirewallZoneIntraZoneFilteringFirewallName: basetypes.NewStringValue("test"),
			},
		},
	}

	want := map[string]interface{}{
		// "id":                    "firewall__name__rule-one__rule__42",
		"intra-zone-filtering": map[string]interface{}{
			"action": "accept",
			"firewall": map[string]interface{}{
				"name": "test",
			},
		},
	}

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	got, err := helpers.MarshalVyos(ctx, has)
	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("Want: %v", want)
		t.Errorf("Got: %v", got)
		t.Errorf("compare failed: %v, %v", diff, err)
	}
}

// TestFirewallNameRuleMarshalVyos does some simple marshalling tests
func TestFirewallZoneUnmarshalVyos(t *testing.T) {
	has := map[string]interface{}{
		// "id":                    "firewall__name__rule-one__rule__42",
		"intra-zone-filtering": map[string]interface{}{
			"action": "accept",
			"firewall": map[string]interface{}{
				"name": "test",
			},
		},
	}

	want := &resourcemodel.FirewallZone{
		ID: basetypes.NewStringValue("Testing"),
		NodeFirewallZoneIntraZoneFiltering: &resourcemodel.FirewallZoneIntraZoneFiltering{
			LeafFirewallZoneIntraZoneFilteringAction: basetypes.NewStringValue("accept"),
			NodeFirewallZoneIntraZoneFilteringFirewall: &resourcemodel.FirewallZoneIntraZoneFilteringFirewall{
				LeafFirewallZoneIntraZoneFilteringFirewallName: basetypes.NewStringValue("test"),
			},
		},
	}

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	got := &resourcemodel.FirewallZone{}

	err := helpers.UnmarshalVyos(ctx, has, got)
	if err != nil {
		t.Fatalf(`desired value can not be unmarshaled: %v`, err)
	}

	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("compare failed: %v", diff)
	}
}
