/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedservicerouteradvertinterfaceroute code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicerouteradvertinterfaceroute

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (route) */
// Metadata method to define the resource type name.
func (r serviceRouterAdvertInterfaceRoute) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_router_advert_interface_route"
}
