// Package namedservicendpproxyinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicendpproxyinterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r serviceNdpProxyInterface) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_ndp_proxy_interface"
}
