/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesVirtualEthernetDhcpvsixOptions{}

// InterfacesVirtualEthernetDhcpvsixOptions describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type InterfacesVirtualEthernetDhcpvsixOptions struct {
	// LeafNodes
	LeafInterfacesVirtualEthernetDhcpvsixOptionsDuID           types.String `tfsdk:"duid" vyos:"duid,omitempty"`
	LeafInterfacesVirtualEthernetDhcpvsixOptionsParametersOnly types.Bool   `tfsdk:"parameters_only" vyos:"parameters-only,omitempty"`
	LeafInterfacesVirtualEthernetDhcpvsixOptionsRAPIDCommit    types.Bool   `tfsdk:"rapid_commit" vyos:"rapid-commit,omitempty"`
	LeafInterfacesVirtualEthernetDhcpvsixOptionsTemporary      types.Bool   `tfsdk:"temporary" vyos:"temporary,omitempty"`
	LeafInterfacesVirtualEthernetDhcpvsixOptionsNoRelease      types.Bool   `tfsdk:"no_release" vyos:"no-release,omitempty"`

	// TagNodes

	ExistsTagInterfacesVirtualEthernetDhcpvsixOptionsPd bool `tfsdk:"-" vyos:"pd,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesVirtualEthernetDhcpvsixOptions) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"duid":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
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

		"parameters_only":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Acquire only config parameters, no address

`,
			Description: `Acquire only config parameters, no address

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rapid_commit":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Wait for immediate reply instead of advertisements

`,
			Description: `Wait for immediate reply instead of advertisements

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"temporary":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 temporary address

`,
			Description: `IPv6 temporary address

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_release":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not send a release message on client exit

`,
			Description: `Do not send a release message on client exit

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
