// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsIgmpInterfaceJoin describes the resource data model.
type ProtocolsIgmpInterfaceJoin struct {
	SelfIdentifier types.String `tfsdk:"join_id" vyos:",self-id"`

	ParentIDProtocolsIgmpInterface types.String `tfsdk:"interface" vyos:"interface,parent-id"`

	// LeafNodes
	LeafProtocolsIgmpInterfaceJoinSource types.List `tfsdk:"source" vyos:"source,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsIgmpInterfaceJoin) GetVyosPath() []string {
	return []string{
		"protocols",

		"igmp",

		"interface",
		o.ParentIDProtocolsIgmpInterface.ValueString(),

		"join",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsIgmpInterfaceJoin) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"join_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IGMP join multicast group

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Multicast group address  |

`,
		},

		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IGMP interface

`,
		},

		// LeafNodes

		"source": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Source address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Source address  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsIgmpInterfaceJoin) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsIgmpInterfaceJoin) UnmarshalJSON(_ []byte) error {
	return nil
}
