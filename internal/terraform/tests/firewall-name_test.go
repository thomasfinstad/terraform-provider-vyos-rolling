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
	if err != nil {
		t.Fatalf(`desired value can not be marshalled: %v`, err)
	}

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
	if err != nil {
		t.Fatalf(`desired value can not be unmarshaled: %v`, err)
	}

	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("compare failed: %v, %v", diff, err)
	}
}

// TestFirewallIPvfourNameUnmarshalVyosHasChild does some simple unmarshalling tests
func TestFirewallIPvfourNameUnmarshalVyosHasChild(t *testing.T) {
	has := map[string]interface{}{
		"default-action": "drop",
		"rule": map[string]interface{}{
			"test-rule": map[string]interface{}{
				"action":                "accept",
				"disable":               map[string]interface{}{},
				"queue":                 float64(28),
				"packet-length-exclude": []string{"420", "13-37"},
				"destination": map[string]interface{}{
					"address": "127.0.0.2",
				},
			},
		},
	}

	want := true

	resource := &resourcemodel.FirewallIPvfourName{}
	err := helpers.UnmarshalVyos(context.Background(), has, resource)
	if err != nil {
		t.Fatalf(`desired value can not be unmarshaled: %v`, err)
	}

	_, got := helpers.GetChild(context.Background(), resource)

	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("compare failed: %v, %v", diff, err)
	}
}

// TestFirewallIPvfourNameUnmarshalVyosNoChild does some simple unmarshalling tests
func TestFirewallIPvfourNameUnmarshalVyosNoChild(t *testing.T) {
	has := map[string]interface{}{
		"default-action": "drop",
	}

	want := false

	resource := &resourcemodel.FirewallIPvfourName{}
	err := helpers.UnmarshalVyos(context.Background(), has, resource)
	if err != nil {
		t.Fatalf(`desired value can not be unmarshaled: %v`, err)
	}

	_, got := helpers.GetChild(context.Background(), resource)

	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("compare failed: %v, %v", diff, err)
	}
}
