// Package namedinterfaceswirelessdhcpvsixoptionspdinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswirelessdhcpvsixoptionspdinterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r interfacesWirelessDhcpvsixOptionsPdInterface) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interfaces_wireless_dhcpv6_options_pd_interface"
}
