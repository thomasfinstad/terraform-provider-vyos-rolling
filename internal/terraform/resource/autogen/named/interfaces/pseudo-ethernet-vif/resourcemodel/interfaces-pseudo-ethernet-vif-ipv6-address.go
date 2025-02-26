/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (address) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesPseudoEthernetVifIPvsixAddress{}

// InterfacesPseudoEthernetVifIPvsixAddress describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type InterfacesPseudoEthernetVifIPvsixAddress struct {
	// LeafNodes
	LeafInterfacesPseudoEthernetVifIPvsixAddressAutoconf           types.Bool `tfsdk:"autoconf" vyos:"autoconf,omitempty"`
	LeafInterfacesPseudoEthernetVifIPvsixAddressEuisixfour         types.List `tfsdk:"eui64" vyos:"eui64,omitempty"`
	LeafInterfacesPseudoEthernetVifIPvsixAddressNoDefaultLinkLocal types.Bool `tfsdk:"no_default_link_local" vyos:"no-default-link-local,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPseudoEthernetVifIPvsixAddress) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"autoconf":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (autoconf) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)

`,
			Description: `Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"eui64":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (eui64) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Prefix for IPv6 address with MAC-based EUI-64

    |  Format                |  Description       |
    |------------------------|--------------------|
    |  <h:h:h:h:h:h:h:h/64>  |  IPv6 /64 network  |
`,
			Description: `Prefix for IPv6 address with MAC-based EUI-64

    |  Format                |  Description       |
    |------------------------|--------------------|
    |  <h:h:h:h:h:h:h:h/64>  |  IPv6 /64 network  |
`,
		},

		"no_default_link_local":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (no-default-link-local) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Remove the default link-local address from the interface

`,
			Description: `Remove the default link-local address from the interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
