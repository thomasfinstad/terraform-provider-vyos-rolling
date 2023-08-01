// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesGeneveParametersIP describes the resource data model.
type InterfacesGeneveParametersIP struct {
	// LeafNodes
	LeafInterfacesGeneveParametersIPDf  types.String `tfsdk:"df" vyos:"df,omitempty"`
	LeafInterfacesGeneveParametersIPTos types.Number `tfsdk:"tos" vyos:"tos,omitempty"`
	LeafInterfacesGeneveParametersIPTTL types.Number `tfsdk:"ttl" vyos:"ttl,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesGeneveParametersIP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"df": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Usage of the DF (don't Fragment) bit in outgoing packets

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  set  &emsp; |  Always set DF (don't fragment) bit  |
    |  unset  &emsp; |  Always unset DF (don't fragment) bit  |
    |  inherit  &emsp; |  Copy from the original IP header  |

`,

			// Default:          stringdefault.StaticString(`unset`),
			Computed: true,
		},

		"tos": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Specifies TOS value to use in outgoing packets

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-99  &emsp; |  Type of Service (TOS)  |

`,

			// Default:          stringdefault.StaticString(`inherit`),
			Computed: true,
		},

		"ttl": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Specifies TTL value to use in outgoing packets

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0  &emsp; |  Inherit - copy value from original IP header  |
    |  number: 1-255  &emsp; |  Time to Live  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesGeneveParametersIP) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesGeneveParametersIP) UnmarshalJSON(_ []byte) error {
	return nil
}
