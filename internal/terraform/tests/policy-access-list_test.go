package tests

import (
	"context"
	"math/big"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/go-test/deep"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflogtest"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/policy/access-list/resourcemodel"
)

// TestPolicyAccessListEmptyMarshal tests that that an empty resource marshals to the correct empty representation
//
//	Simulated resource:
//		resource "vyos_policy_access_list" "name" {
//			access_list_id = 42
//		}
func TestPolicyAccessListEmptyMarshal(t *testing.T) {
	model := &resourcemodel.PolicyAccessList{
		SelfIdentifier: basetypes.NewNumberValue(big.NewFloat(42)),
	}

	want := make(map[string]any)

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	got, err := helpers.MarshalVyos(ctx, model)
	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("compare failed: %v, %v", diff, err)
	}
}

// TestPolicyAccessListEmptyGenerateVyosOps tests that an empty resource generates the correct vyos operations
//
//	Simulated resource:
//		resource "vyos_policy_access_list" "name" {
//			access_list_id = 42
//		}
func TestPolicyAccessListEmptyGenerateVyosOps(t *testing.T) {
	hasPath := []string{"policy", "access-list", "42"}
	hasData := make(map[string]any)

	want := [][]string{
		{"policy", "access-list", "42"},
	}

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

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
