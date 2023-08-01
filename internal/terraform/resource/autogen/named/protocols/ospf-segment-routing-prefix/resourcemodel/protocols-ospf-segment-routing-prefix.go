// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsOspfSegmentRoutingPrefix describes the resource data model.
type ProtocolsOspfSegmentRoutingPrefix struct {
	SelfIdentifier types.String `tfsdk:"prefix_id" vyos:",self-id"`

	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeProtocolsOspfSegmentRoutingPrefixIndex *ProtocolsOspfSegmentRoutingPrefixIndex `tfsdk:"index" vyos:"index,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfSegmentRoutingPrefix) GetVyosPath() []string {
	return []string{
		"protocols",

		"ospf",

		"segment-routing",

		"prefix",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfSegmentRoutingPrefix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"prefix_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Static IPv4 prefix segment/label mapping

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 prefix segment  |

`,
		},

		// LeafNodes

		// Nodes

		"index": schema.SingleNestedAttribute{
			Attributes: ProtocolsOspfSegmentRoutingPrefixIndex{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Specify the index value of prefix segment/label ID

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsOspfSegmentRoutingPrefix) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsOspfSegmentRoutingPrefix) UnmarshalJSON(_ []byte) error {
	return nil
}
