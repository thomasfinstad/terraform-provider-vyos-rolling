/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalhighavailabilityvrrpglobalparameters code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalhighavailabilityvrrpglobalparameters

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (global-parameters) */
// Metadata method to define the resource type name.
func (r highAvailabilityVrrpGlobalParameters) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_high_availability_vrrp_global_parameters"
}
