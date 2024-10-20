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

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesVxlanParameters{}

// InterfacesVxlanParameters describes the resource data model.
type InterfacesVxlanParameters struct {
	// LeafNodes
	LeafInterfacesVxlanParametersExternal         types.Bool `tfsdk:"external" vyos:"external,omitempty"`
	LeafInterfacesVxlanParametersNolearning       types.Bool `tfsdk:"nolearning" vyos:"nolearning,omitempty"`
	LeafInterfacesVxlanParametersNeighborSuppress types.Bool `tfsdk:"neighbor_suppress" vyos:"neighbor-suppress,omitempty"`
	LeafInterfacesVxlanParametersVniFilter        types.Bool `tfsdk:"vni_filter" vyos:"vni-filter,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeInterfacesVxlanParametersIP     *InterfacesVxlanParametersIP     `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesVxlanParametersIPvsix *InterfacesVxlanParametersIPvsix `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesVxlanParameters) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"external":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use external control plane

`,
			Description: `Use external control plane

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"nolearning":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not add unknown addresses into forwarding database

`,
			Description: `Do not add unknown addresses into forwarding database

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"neighbor_suppress":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable neighbor discovery (ARP and ND) suppression

`,
			Description: `Enable neighbor discovery (ARP and ND) suppression

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"vni_filter":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable VNI filter support

`,
			Description: `Enable VNI filter support

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesVxlanParametersIP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 specific tunnel parameters

`,
			Description: `IPv4 specific tunnel parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesVxlanParametersIPvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 specific tunnel parameters

`,
			Description: `IPv6 specific tunnel parameters

`,
		},
	}
}
