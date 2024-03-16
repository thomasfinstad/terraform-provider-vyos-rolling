package tests

import (
	"context"
	"math/big"
	"sort"
	"strings"
	"testing"

	"github.com/go-test/deep"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/policy/access-list/resourcemodel"
)

// TODO figure out print logging vs tflog for testing

// resource "vyos_policy_access_list" "name" {
// 	access_list_id = 2

//   }

func TestPolicyAccessListEmptyMarshal(t *testing.T) {
	model := &resourcemodel.PolicyAccessList{
		SelfIdentifier: basetypes.NewNumberValue(big.NewFloat(42)),
	}

	want := make(map[string]any)

	ctx := context.Background()
	lvl := tflog.WithLevel(hclog.Error)
	ctx = tflog.NewSubsystem(ctx, "my-subsystem", lvl)

	got, err := helpers.MarshalVyos(ctx, model)
	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("compare failed: %v, %v", diff, err)
	}
}

// TestFirewallIPvfourNameMarshalVyos does some simple marshalling tests
func TestPolicyAccessListEmptyGenerateVyosOps(t *testing.T) {
	hasPath := []string{"policy", "access-list", "42"}
	hasData := make(map[string]any)

	want := [][]string{
		{"policy", "access-list", "42"},
	}

	ctx := context.Background()
	lvl := tflog.WithLevel(hclog.Error)
	ctx = tflog.NewSubsystem(ctx, "my-subsystem", lvl)

	got := helpers.GenerateVyosOps(ctx, hasPath, hasData)

	// Sort results
	sort.Slice(want, func(i, j int) bool {
		return strings.Join(want[i], "") < strings.Join(want[j], "")
	})
	sort.Slice(got, func(i, j int) bool {
		return strings.Join(got[i], "") < strings.Join(got[j], "")
	})

	// Diff results
	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("Want: %v", want)
		t.Errorf("Got: %v", got)
		t.Errorf("compare failed: %v", diff)
	}
}
