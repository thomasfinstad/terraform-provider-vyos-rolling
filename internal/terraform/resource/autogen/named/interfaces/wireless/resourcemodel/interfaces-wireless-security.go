// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesWirelessSecURIty{}

// InterfacesWirelessSecURIty describes the resource data model.
type InterfacesWirelessSecURIty struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeInterfacesWirelessSecURItyStationAddress *InterfacesWirelessSecURItyStationAddress `tfsdk:"station_address" vyos:"station-address,omitempty"`
	NodeInterfacesWirelessSecURItyWep            *InterfacesWirelessSecURItyWep            `tfsdk:"wep" vyos:"wep,omitempty"`
	NodeInterfacesWirelessSecURItyWpa            *InterfacesWirelessSecURItyWpa            `tfsdk:"wpa" vyos:"wpa,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessSecURIty) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"station_address": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessSecURItyStationAddress{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Station MAC address based authentication

`,
			Description: `Station MAC address based authentication

`,
		},

		"wep": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessSecURItyWep{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Wired Equivalent Privacy (WEP) parameters

`,
			Description: `Wired Equivalent Privacy (WEP) parameters

`,
		},

		"wpa": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessSecURItyWpa{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Wifi Protected Access (WPA) parameters

`,
			Description: `Wifi Protected Access (WPA) parameters

`,
		},
	}
}