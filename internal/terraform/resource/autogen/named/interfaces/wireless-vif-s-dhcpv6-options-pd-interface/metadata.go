/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedinterfaceswirelessvifsdhcpvsixoptionspdinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswirelessvifsdhcpvsixoptionspdinterface

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (interface) */
// Metadata method to define the resource type name.
func (r interfacesWirelessVifSDhcpvsixOptionsPdInterface) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interfaces_wireless_vif_s_dhcpv6_options_pd_interface"
}
