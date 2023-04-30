// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// InterfacesBondingVifDhcpvsixOptionsPdInterface describes the resource data model.
type InterfacesBondingVifDhcpvsixOptionsPdInterface struct {
	// LeafNodes
	InterfacesBondingVifDhcpvsixOptionsPdInterfaceAddress customtypes.CustomStringValue `tfsdk:"address" json:"address,omitempty"`
	InterfacesBondingVifDhcpvsixOptionsPdInterfaceSLAID   customtypes.CustomStringValue `tfsdk:"sla_id" json:"sla-id,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o InterfacesBondingVifDhcpvsixOptionsPdInterface) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Local interface address assigned to interface (default: EUI-64)

|  Format  |  Description  |
|----------|---------------|
|  >0  |  Used to form IPv6 interface address  |

`,
		},

		"sla_id": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Interface site-Level aggregator (SLA)

|  Format  |  Description  |
|----------|---------------|
|  u32:0-65535  |  Decimal integer which fits in the length of SLA IDs  |

`,
		},

		// TagNodes

		// Nodes

	}
}
