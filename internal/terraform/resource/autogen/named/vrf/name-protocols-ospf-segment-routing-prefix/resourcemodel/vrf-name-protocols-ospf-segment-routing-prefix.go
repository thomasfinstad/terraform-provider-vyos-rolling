// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsOspfSegmentRoutingPrefix describes the resource data model.
type VrfNameProtocolsOspfSegmentRoutingPrefix struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDVrfName types.String `tfsdk:"name" vyos:"name_identifier,parent-id"`

	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsOspfSegmentRoutingPrefixIndex *VrfNameProtocolsOspfSegmentRoutingPrefixIndex `tfsdk:"index" vyos:"index,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsOspfSegmentRoutingPrefix) GetVyosPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"ospf",

		"segment-routing",

		"prefix",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfSegmentRoutingPrefix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Static IPv4 prefix segment/label mapping

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4net  |  IPv4 prefix segment  |

`,
		},

		"name_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  VRF instance name  |

`,
		},

		// LeafNodes

		// Nodes

		"index": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfSegmentRoutingPrefixIndex{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Specify the index value of prefix segment/label ID

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfSegmentRoutingPrefix) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsOspfSegmentRoutingPrefix) UnmarshalJSON(_ []byte) error {
	return nil
}
