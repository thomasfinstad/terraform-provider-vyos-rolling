// Package globalhighavailabilityvrrp code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalhighavailabilityvrrp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r highAvailabilityVrrp) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_high_availability_vrrp"
}
