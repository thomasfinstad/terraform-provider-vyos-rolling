package helpers_test

import (
	"context"
	"testing"

	"github.com/go-test/deep"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	ntpResourceModel "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/service/ntp/resourcemodel"
)

// TestMarshalToVyosNamedResMergedIntoGlobal
// Verify that what would have been a named resource that has to be merged
// in under a global resource results in a map of config parameters and works as expected
func TestMarshalToVyosNamedResMergedIntoGlobal(t *testing.T) {
	ctx := context.Background()

	has := &ntpResourceModel.ServiceNtp{
		// ID: basetypes.NewStringNull(),
		// Timeouts:                    basetypes,
		// LeafServiceNtpInterface:     {},
		// LeafServiceNtpListenAddress: {},
		// LeafServiceNtpVrf:           {},
		// LeafServiceNtpLeapSecond:    {},
		TagServiceNtpServer: map[string]*ntpResourceModel.ServiceNtpServer{
			"no.pool.ntp.org": {
				LeafServiceNtpServerNoselect:   basetypes.NewBoolValue(false),
				LeafServiceNtpServerNts:        basetypes.NewBoolValue(false),
				LeafServiceNtpServerPool:       basetypes.NewBoolValue(true),
				LeafServiceNtpServerPrefer:     basetypes.NewBoolValue(true),
				LeafServiceNtpServerPtp:        basetypes.NewBoolValue(false),
				LeafServiceNtpServerInterleave: basetypes.NewBoolValue(false),
			},
		},
		// NodeServiceNtpAllowClient: nil,
		ExistsNodeServiceNtpPtp: false,
	}

	want := map[string]any{
		"server": map[string]any{
			"no.pool.ntp.org": map[string]any{
				"pool":   map[string]any{},
				"prefer": map[string]any{},
			},
		},
	}

	got, err := helpers.MarshalVyos(ctx, has)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}

	// Diff results
	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("Want: %v", want)
		t.Errorf("Got: %v", got)
		t.Errorf("compare failed: %v", diff)
	}
}
