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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (ipv6) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameIPvsix{}

// VrfNameIPvsix describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameIPvsix struct {
	// LeafNodes
	LeafVrfNameIPvsixDisableForwarding types.Bool `tfsdk:"disable_forwarding" vyos:"disable-forwarding,omitempty"`

	// TagNodes

	ExistsTagVrfNameIPvsixProtocol bool `tfsdk:"-" vyos:"protocol,child"`

	// Nodes

	NodeVrfNameIPvsixNht *VrfNameIPvsixNht `tfsdk:"nht" vyos:"nht,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameIPvsix) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable_forwarding":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (disable-forwarding) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable IP forwarding on this interface

`,
			Description: `Disable IP forwarding on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

		"nht": schema.SingleNestedAttribute{
			Attributes: VrfNameIPvsixNht{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Filter Next Hop tracking route resolution

`,
			Description: `Filter Next Hop tracking route resolution

`,
		},
	}
}
