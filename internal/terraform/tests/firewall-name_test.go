package tests

import (
	"context"
	"testing"

	"github.com/go-test/deep"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/ipv4-name/resourcemodel"
)

// TODO figure out print logging vs tflog for testing

// TestFirewallIPvfourNameMarshalVyos does some simple marshalling tests
func TestFirewallIPvfourNameMarshalVyos(t *testing.T) {
	model := &resourcemodel.FirewallIPvfourName{
		SelfIdentifier:                       basetypes.NewStringValue("test-id"),
		LeafFirewallIPvfourNameDefaultAction: basetypes.NewStringValue("drop"),
		LeafFirewallIPvfourNameDefaultLog:    basetypes.NewBoolValue(false),
	}

	want := map[string]interface{}{
		"default-action": "drop",
	}

	got, err := helpers.MarshalVyos(context.Background(), model)
	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("compare failed: %v, %v", diff, err)
	}
}

// TestFirewallIPvfourNameUnmarshalVyos does some simple unmarshalling tests
func TestFirewallIPvfourNameUnmarshalVyos(t *testing.T) {
	has := map[string]interface{}{
		"default-action": "drop",
	}

	want := &resourcemodel.FirewallIPvfourName{
		LeafFirewallIPvfourNameDefaultAction: basetypes.NewStringValue("drop"),
		LeafFirewallIPvfourNameDefaultLog:    basetypes.NewBoolValue(false),
	}

	got := &resourcemodel.FirewallIPvfourName{}

	err := helpers.UnmarshalVyos(context.Background(), has, got)
	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("compare failed: %v, %v", diff, err)
	}
}
