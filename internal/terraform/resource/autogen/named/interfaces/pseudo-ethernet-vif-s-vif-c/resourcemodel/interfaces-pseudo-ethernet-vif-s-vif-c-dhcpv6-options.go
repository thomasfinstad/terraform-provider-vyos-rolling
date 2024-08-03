// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesPseudoEthernetVifSVifCDhcpvsixOptions{}

// InterfacesPseudoEthernetVifSVifCDhcpvsixOptions describes the resource data model.
type InterfacesPseudoEthernetVifSVifCDhcpvsixOptions struct {
	// LeafNodes
	LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsDuID           types.String `tfsdk:"duid" vyos:"duid,omitempty"`
	LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsParametersOnly types.Bool   `tfsdk:"parameters_only" vyos:"parameters-only,omitempty"`
	LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsRAPIDCommit    types.Bool   `tfsdk:"rapid_commit" vyos:"rapid-commit,omitempty"`
	LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsTemporary      types.Bool   `tfsdk:"temporary" vyos:"temporary,omitempty"`
	LeafInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsNoRelease      types.Bool   `tfsdk:"no_release" vyos:"no-release,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagInterfacesPseudoEthernetVifSVifCDhcpvsixOptionsPd bool `tfsdk:"pd" vyos:"pd,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPseudoEthernetVifSVifCDhcpvsixOptions) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"duid": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DHCP unique identifier (DUID) to be sent by client

    |  Format  |  Description             |
    |----------|--------------------------|
    |  duid    |  DHCP unique identifier  |
`,
			Description: `DHCP unique identifier (DUID) to be sent by client

    |  Format  |  Description             |
    |----------|--------------------------|
    |  duid    |  DHCP unique identifier  |
`,
		},

		"parameters_only": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Acquire only config parameters, no address

`,
			Description: `Acquire only config parameters, no address

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rapid_commit": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Wait for immediate reply instead of advertisements

`,
			Description: `Wait for immediate reply instead of advertisements

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"temporary": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 temporary address

`,
			Description: `IPv6 temporary address

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_release": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not send a release message on client exit

`,
			Description: `Do not send a release message on client exit

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
