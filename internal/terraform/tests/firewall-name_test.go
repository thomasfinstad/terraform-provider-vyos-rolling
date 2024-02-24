package tests

import (
	"context"
	"testing"

	"github.com/go-test/deep"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/firewall/name/resourcemodel"
)

// TODO figure out print logging vs tflog for testing

// TestFirewallNameMarshalVyos does some simple marshalling tests
func TestFirewallNameMarshalVyos(t *testing.T) {
	model := &resourcemodel.FirewallName{
		SelfIdentifier:                   basetypes.NewStringValue("test-id"),
		LeafFirewallNameDefaultAction:    basetypes.NewStringValue("drop"),
		LeafFirewallNameEnableDefaultLog: basetypes.NewBoolValue(false),
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

// TestFirewallNameUnmarshalVyos does some simple unmarshalling tests
func TestFirewallNameUnmarshalVyos(t *testing.T) {
	has := map[string]interface{}{
		"default-action": "drop",
	}

	want := &resourcemodel.FirewallName{
		LeafFirewallNameDefaultAction:    basetypes.NewStringValue("drop"),
		LeafFirewallNameEnableDefaultLog: basetypes.NewBoolValue(false),
	}

	got := &resourcemodel.FirewallName{}

	err := helpers.UnmarshalVyos(context.Background(), has, got)
	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("compare failed: %v, %v", diff, err)
	}
}
