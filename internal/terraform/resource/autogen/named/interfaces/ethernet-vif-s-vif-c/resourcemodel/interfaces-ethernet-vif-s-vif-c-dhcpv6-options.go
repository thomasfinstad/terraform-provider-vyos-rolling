// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesEthernetVifSVifCDhcpvsixOptions describes the resource data model.
type InterfacesEthernetVifSVifCDhcpvsixOptions struct {
	// LeafNodes
	LeafInterfacesEthernetVifSVifCDhcpvsixOptionsDuID           types.String `tfsdk:"duid" vyos:"duid,omitempty"`
	LeafInterfacesEthernetVifSVifCDhcpvsixOptionsParametersOnly types.Bool   `tfsdk:"parameters_only" vyos:"parameters-only,omitempty"`
	LeafInterfacesEthernetVifSVifCDhcpvsixOptionsRAPIDCommit    types.Bool   `tfsdk:"rapid_commit" vyos:"rapid-commit,omitempty"`
	LeafInterfacesEthernetVifSVifCDhcpvsixOptionsTemporary      types.Bool   `tfsdk:"temporary" vyos:"temporary,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagInterfacesEthernetVifSVifCDhcpvsixOptionsPd bool `tfsdk:"pd" vyos:"pd,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernetVifSVifCDhcpvsixOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"duid": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DHCP unique identifier (DUID) to be sent by dhcpv6 client

    |  Format  |  Description  |
    |----------|---------------|
    |  duid  |  DHCP unique identifier (DUID)  |

`,
		},

		"parameters_only": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Acquire only config parameters, no address

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rapid_commit": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Wait for immediate reply instead of advertisements

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"temporary": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 temporary address

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesEthernetVifSVifCDhcpvsixOptions) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesEthernetVifSVifCDhcpvsixOptions) UnmarshalJSON(_ []byte) error {
	return nil
}
