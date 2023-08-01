// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VpnSstpClientIPvsixPoolPrefix describes the resource data model.
type VpnSstpClientIPvsixPoolPrefix struct {
	SelfIdentifier types.String `tfsdk:"prefix_id" vyos:",self-id"`

	// LeafNodes
	LeafVpnSstpClientIPvsixPoolPrefixMask types.Number `tfsdk:"mask" vyos:"mask,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnSstpClientIPvsixPoolPrefix) GetVyosPath() []string {
	return []string{
		"vpn",

		"sstp",

		"client-ipv6-pool",

		"prefix",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnSstpClientIPvsixPoolPrefix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"prefix_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Pool of addresses used to assign to clients

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6net  &emsp; |  IPv6 address and prefix length  |

`,
		},

		// LeafNodes

		"mask": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Prefix length used for individual client

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 48-128  &emsp; |  Client prefix length  |

`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnSstpClientIPvsixPoolPrefix) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VpnSstpClientIPvsixPoolPrefix) UnmarshalJSON(_ []byte) error {
	return nil
}
