// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// HighAvailabilityVrrpGroupAddress describes the resource data model.
type HighAvailabilityVrrpGroupAddress struct {
	SelfIdentifier types.String `tfsdk:"address_id" vyos:",self-id"`

	ParentIDHighAvailabilityVrrpGroup types.String `tfsdk:"group" vyos:"group,parent-id"`

	// LeafNodes
	LeafHighAvailabilityVrrpGroupAddressInterface types.String `tfsdk:"interface" vyos:"interface,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *HighAvailabilityVrrpGroupAddress) GetVyosPath() []string {
	return []string{
		"high-availability",

		"vrrp",

		"group",
		o.ParentIDHighAvailabilityVrrpGroup.ValueString(),

		"address",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpGroupAddress) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"address_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual IP address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 address and prefix length  |
    |  ipv6net  &emsp; |  IPv6 address and prefix length  |

`,
		},

		"group_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `VRRP group

`,
		},

		// LeafNodes

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface Name to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *HighAvailabilityVrrpGroupAddress) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *HighAvailabilityVrrpGroupAddress) UnmarshalJSON(_ []byte) error {
	return nil
}