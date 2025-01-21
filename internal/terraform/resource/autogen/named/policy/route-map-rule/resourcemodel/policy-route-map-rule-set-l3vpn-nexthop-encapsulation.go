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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (encapsulation) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleSetLthreevpnNexthopEncapsulation{}

// PolicyRouteMapRuleSetLthreevpnNexthopEncapsulation describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type PolicyRouteMapRuleSetLthreevpnNexthopEncapsulation struct {
	// LeafNodes
	LeafPolicyRouteMapRuleSetLthreevpnNexthopEncapsulationGre types.Bool `tfsdk:"gre" vyos:"gre,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleSetLthreevpnNexthopEncapsulation) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"gre":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (gre) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Accept L3VPN traffic over GRE encapsulation

`,
			Description: `Accept L3VPN traffic over GRE encapsulation

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
