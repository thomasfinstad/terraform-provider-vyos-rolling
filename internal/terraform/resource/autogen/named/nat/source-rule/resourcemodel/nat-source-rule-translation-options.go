// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &NatSourceRuleTranSLAtionOptions{}

// NatSourceRuleTranSLAtionOptions describes the resource data model.
type NatSourceRuleTranSLAtionOptions struct {
	// LeafNodes
	LeafNatSourceRuleTranSLAtionOptionsAddressMapping types.String `tfsdk:"address_mapping" vyos:"address-mapping,omitempty"`
	LeafNatSourceRuleTranSLAtionOptionsPortMapping    types.String `tfsdk:"port_mapping" vyos:"port-mapping,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatSourceRuleTranSLAtionOptions) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address_mapping": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Address mapping options

    |  Format      |  Description                                                                |
    |--------------|-----------------------------------------------------------------------------|
    |  persistent  |  Gives a client the same source or destination-address for each connection  |
    |  random      |  Random source or destination address allocation for each connection        |
`,
			Description: `Address mapping options

    |  Format      |  Description                                                                |
    |--------------|-----------------------------------------------------------------------------|
    |  persistent  |  Gives a client the same source or destination-address for each connection  |
    |  random      |  Random source or destination address allocation for each connection        |
`,

			// Default:          stringdefault.StaticString(`random`),
			Computed: true,
		},

		"port_mapping": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port mapping options

    |  Format        |  Description                      |
    |----------------|-----------------------------------|
    |  random        |  Randomize source port mapping    |
    |  fully-random  |  Full port randomization          |
    |  none          |  Do not apply port randomization  |
`,
			Description: `Port mapping options

    |  Format        |  Description                      |
    |----------------|-----------------------------------|
    |  random        |  Randomize source port mapping    |
    |  fully-random  |  Full port randomization          |
    |  none          |  Do not apply port randomization  |
`,

			// Default:          stringdefault.StaticString(`none`),
			Computed: true,
		},

		// Nodes

	}
}
