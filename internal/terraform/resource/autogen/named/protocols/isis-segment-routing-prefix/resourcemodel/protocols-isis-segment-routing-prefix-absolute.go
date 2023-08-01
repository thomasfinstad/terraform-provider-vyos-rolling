// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsIsisSegmentRoutingPrefixAbsolute describes the resource data model.
type ProtocolsIsisSegmentRoutingPrefixAbsolute struct {
	// LeafNodes
	LeafProtocolsIsisSegmentRoutingPrefixAbsoluteValue        types.Number `tfsdk:"value" vyos:"value,omitempty"`
	LeafProtocolsIsisSegmentRoutingPrefixAbsoluteExplicitNull types.Bool   `tfsdk:"explicit_null" vyos:"explicit-null,omitempty"`
	LeafProtocolsIsisSegmentRoutingPrefixAbsoluteNoPhpFlag    types.Bool   `tfsdk:"no_php_flag" vyos:"no-php-flag,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsIsisSegmentRoutingPrefixAbsolute) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"value": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Specify the absolute value of prefix segment/label ID

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 16-1048575  &emsp; |  The absolute segment/label ID value  |

`,
		},

		"explicit_null": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Request upstream neighbor to replace segment/label with explicit null label

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_php_flag": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not request penultimate hop popping for segment/label

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsIsisSegmentRoutingPrefixAbsolute) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsIsisSegmentRoutingPrefixAbsolute) UnmarshalJSON(_ []byte) error {
	return nil
}