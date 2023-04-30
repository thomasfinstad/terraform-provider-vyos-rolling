// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ProtocolsStaticArpInterfaceAddress describes the resource data model.
type ProtocolsStaticArpInterfaceAddress struct {
	// LeafNodes
	ProtocolsStaticArpInterfaceAddressDescrIPtion customtypes.CustomStringValue `tfsdk:"description" json:"description,omitempty"`
	ProtocolsStaticArpInterfaceAddressMac         customtypes.CustomStringValue `tfsdk:"mac" json:"mac,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ProtocolsStaticArpInterfaceAddress) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"description": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Description

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Description  |

`,
		},

		"mac": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Media Access Control (MAC) address

|  Format  |  Description  |
|----------|---------------|
|  macaddr  |  Hardware (MAC) address  |

`,
		},

		// TagNodes

		// Nodes

	}
}
