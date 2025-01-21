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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (evpn) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesEthernetEvpn{}

// InterfacesEthernetEvpn describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type InterfacesEthernetEvpn struct {
	// LeafNodes
	LeafInterfacesEthernetEvpnUplink types.Bool `tfsdk:"uplink" vyos:"uplink,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernetEvpn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"uplink":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (uplink) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Uplink to the VXLAN core

`,
			Description: `Uplink to the VXLAN core

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
