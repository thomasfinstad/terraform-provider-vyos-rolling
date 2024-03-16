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
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/global/high-availability/vrrp-global-parameters/resourcemodel"
)

// TestHighAvailabilityVrrpGlobalParametersMarshalVyos test global resource marshaling
func TestHighAvailabilityVrrpGlobalParametersMarshalVyos(t *testing.T) {
	has := &resourcemodel.HighAvailabilityVrrpGlobalParameters{
		ID: basetypes.NewStringValue("high-availability__vrrp__global-parameters"),
		LeafHighAvailabilityVrrpGlobalParametersStartupDelay: basetypes.NewNumberValue(big.NewFloat(12)),
		ExistsNodeHighAvailabilityVrrpGlobalParametersGarp:   true,
	}

	want := map[string]interface{}{
		"startup-delay": float64(12),
	}

	ctx := context.Background()
	lvl := tflog.WithLevel(hclog.Error)
	ctx = tflog.NewSubsystem(ctx, "my-subsystem", lvl)

	got, err := helpers.MarshalVyos(ctx, has)
	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("Want: %v", want)
		t.Errorf("Got: %v", got)
		t.Errorf("compare failed: %v, %v", diff, err)
	}
}

// TestHighAvailabilityVrrpGlobalParametersUnmarshalVyos does some simple marshalling tests
func TestHighAvailabilityVrrpGlobalParametersUnmarshalVyos(t *testing.T) {
	has := map[string]interface{}{
		"startup-delay": 12,
		"garp": map[string]interface{}{
			"interval": 0,
		},
	}

	want := &resourcemodel.HighAvailabilityVrrpGlobalParameters{
		ID: basetypes.NewStringValue("high-availability__vrrp__global-parameters"),
		LeafHighAvailabilityVrrpGlobalParametersStartupDelay: basetypes.NewNumberValue(big.NewFloat(12)),
		ExistsNodeHighAvailabilityVrrpGlobalParametersGarp:   true,
	}

	got := &resourcemodel.HighAvailabilityVrrpGlobalParameters{}

	err := helpers.UnmarshalVyos(context.Background(), has, got)
	if err != nil {
		t.Fatalf(`desired value can not be unmarshaled: %v`, err)
	}

	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("compare failed: %v", diff)
	}
}

// TestHighAvailabilityVrrpGlobalParametersGenerateVyosOps generates vyos ops to test global parameters
func TestHighAvailabilityVrrpGlobalParametersGenerateVyosOps(t *testing.T) {
	hasPath := []string{"high-availability", "vrrp"}
	hasData := map[string]interface{}{
		"global-parameters": map[string]interface{}{
			"startup-delay": float64(12),
		},
	}

	want := [][]string{
		{"high-availability", "vrrp", "global-parameters", "startup-delay", "12"},
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