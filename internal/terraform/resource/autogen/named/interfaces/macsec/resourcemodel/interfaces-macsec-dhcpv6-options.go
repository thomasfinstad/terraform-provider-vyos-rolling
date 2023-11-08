// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesMacsecDhcpvsixOptions describes the resource data model.
type InterfacesMacsecDhcpvsixOptions struct {
	// LeafNodes
	LeafInterfacesMacsecDhcpvsixOptionsDuID           types.String `tfsdk:"duid" vyos:"duid,omitempty"`
	LeafInterfacesMacsecDhcpvsixOptionsParametersOnly types.Bool   `tfsdk:"parameters_only" vyos:"parameters-only,omitempty"`
	LeafInterfacesMacsecDhcpvsixOptionsRAPIDCommit    types.Bool   `tfsdk:"rapid_commit" vyos:"rapid-commit,omitempty"`
	LeafInterfacesMacsecDhcpvsixOptionsTemporary      types.Bool   `tfsdk:"temporary" vyos:"temporary,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagInterfacesMacsecDhcpvsixOptionsPd bool `tfsdk:"pd" vyos:"pd,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesMacsecDhcpvsixOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"duid": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DHCP unique identifier (DUID) to be sent by dhcpv6 client

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  duid  &emsp; |  DHCP unique identifier (DUID)  |

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
