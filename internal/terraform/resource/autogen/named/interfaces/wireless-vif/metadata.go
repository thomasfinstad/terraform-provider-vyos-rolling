// Package namedinterfaceswirelessvif code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswirelessvif

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r interfacesWirelessVif) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interfaces_wireless_vif"
}