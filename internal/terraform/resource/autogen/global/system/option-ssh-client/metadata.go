// Package globalsystemoptiontcpclient code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemoptiontcpclient

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r systemOptionTCPClient) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_option_ssh_client"
}
