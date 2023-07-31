// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBabelDistributeListIPvfourInterface describes the resource data model.
type ProtocolsBabelDistributeListIPvfourInterface struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeProtocolsBabelDistributeListIPvfourInterfaceAccessList *ProtocolsBabelDistributeListIPvfourInterfaceAccessList `tfsdk:"access_list" vyos:"access-list,omitempty"`
	NodeProtocolsBabelDistributeListIPvfourInterfacePrefixList *ProtocolsBabelDistributeListIPvfourInterfacePrefixList `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBabelDistributeListIPvfourInterface) GetVyosPath() []string {
	return []string{
		"protocols",

		"babel",

		"distribute-list",

		"ipv4",

		"interface",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBabelDistributeListIPvfourInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Apply filtering to an interface

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Apply filtering to an interface  |

`,
		},

		// LeafNodes

		// Nodes

		"access_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBabelDistributeListIPvfourInterfaceAccessList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Access-list

`,
		},

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBabelDistributeListIPvfourInterfacePrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Prefix-list

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBabelDistributeListIPvfourInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBabelDistributeListIPvfourInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
