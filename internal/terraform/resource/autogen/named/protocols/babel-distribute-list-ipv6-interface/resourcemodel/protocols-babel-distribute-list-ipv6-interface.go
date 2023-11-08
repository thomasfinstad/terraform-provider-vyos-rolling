// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBabelDistributeListIPvsixInterface describes the resource data model.
type ProtocolsBabelDistributeListIPvsixInterface struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"interface_id" vyos:",self-id"`

	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeProtocolsBabelDistributeListIPvsixInterfaceAccessList *ProtocolsBabelDistributeListIPvsixInterfaceAccessList `tfsdk:"access_list" vyos:"access-list,omitempty"`
	NodeProtocolsBabelDistributeListIPvsixInterfacePrefixList *ProtocolsBabelDistributeListIPvsixInterfacePrefixList `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`
}

// GetID returns the resource ID
func (o ProtocolsBabelDistributeListIPvsixInterface) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o ProtocolsBabelDistributeListIPvsixInterface) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBabelDistributeListIPvsixInterface) GetVyosPath() []string {
	return []string{
		"protocols",

		"babel",

		"distribute-list",

		"ipv6",

		"interface",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBabelDistributeListIPvsixInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Apply filtering to an interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Apply filtering to an interface  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		// Nodes

		"access_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBabelDistributeListIPvsixInterfaceAccessList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Access-list

`,
		},

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsBabelDistributeListIPvsixInterfacePrefixList{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Prefix-list

`,
		},
	}
}
