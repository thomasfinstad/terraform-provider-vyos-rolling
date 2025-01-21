/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (station-address) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesWirelessSecURItyStationAddress{}

// InterfacesWirelessSecURItyStationAddress describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type InterfacesWirelessSecURItyStationAddress struct {
	// LeafNodes
	LeafInterfacesWirelessSecURItyStationAddressMode types.String `tfsdk:"mode" vyos:"mode,omitempty"`

	// TagNodes

	// Nodes

	NodeInterfacesWirelessSecURItyStationAddressAccept *InterfacesWirelessSecURItyStationAddressAccept `tfsdk:"accept" vyos:"accept,omitempty"`

	NodeInterfacesWirelessSecURItyStationAddressDeny *InterfacesWirelessSecURItyStationAddressDeny `tfsdk:"deny" vyos:"deny,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessSecURItyStationAddress) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"mode":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mode) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Select security operation mode

    |  Format  |  Description                                   |
    |----------|------------------------------------------------|
    |  accept  |  Accept all clients unless found in deny list  |
    |  deny    |  Deny all clients unless found in accept list  |
`,
			Description: `Select security operation mode

    |  Format  |  Description                                   |
    |----------|------------------------------------------------|
    |  accept  |  Accept all clients unless found in deny list  |
    |  deny    |  Deny all clients unless found in accept list  |
`,

			// Default:          stringdefault.StaticString(`accept`),
			Computed: true,
		},

		// TagNodes

		// Nodes

		"accept": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessSecURItyStationAddressAccept{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Accept station MAC address

`,
			Description: `Accept station MAC address

`,
		},

		"deny": schema.SingleNestedAttribute{
			Attributes: InterfacesWirelessSecURItyStationAddressDeny{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Deny station MAC address

`,
			Description: `Deny station MAC address

`,
		},
	}
}
