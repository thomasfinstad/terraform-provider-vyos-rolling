// Package globalserviceipoeserverauthenticationradiusdynamicauthor code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceipoeserverauthenticationradiusdynamicauthor

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r serviceIPoeServerAuthenticationRadiusDynamicAuthor) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_ipoe_server_authentication_radius_dynamic_author"
}