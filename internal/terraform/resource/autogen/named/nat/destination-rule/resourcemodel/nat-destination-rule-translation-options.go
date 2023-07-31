// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NatDestinationRuleTranSLAtionOptions describes the resource data model.
type NatDestinationRuleTranSLAtionOptions struct {
	// LeafNodes
	LeafNatDestinationRuleTranSLAtionOptionsAddressMapping types.String `tfsdk:"address_mapping" vyos:"address-mapping,omitempty"`
	LeafNatDestinationRuleTranSLAtionOptionsPortMapping    types.String `tfsdk:"port_mapping" vyos:"port-mapping,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NatDestinationRuleTranSLAtionOptions) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address_mapping": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Address mapping options

    |  Format  |  Description  |
    |----------|---------------|
    |  persistent  |  Gives a client the same source or destination-address for each connection  |
    |  random  |  Random source or destination address allocation for each connection  |

`,

			// Default:          stringdefault.StaticString(`random`),
			Computed: true,
		},

		"port_mapping": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Port mapping options

    |  Format  |  Description  |
    |----------|---------------|
    |  random  |  Randomize source port mapping  |
    |  fully-random  |  Full port randomization  |
    |  none  |  Do not apply port randomization  |

`,

			// Default:          stringdefault.StaticString(`none`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *NatDestinationRuleTranSLAtionOptions) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *NatDestinationRuleTranSLAtionOptions) UnmarshalJSON(_ []byte) error {
	return nil
}
