// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsOspfAccessList describes the resource data model.
type ProtocolsOspfAccessList struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"access_list_id" vyos:",self-id"`

	// LeafNodes
	LeafProtocolsOspfAccessListExport types.List `tfsdk:"export" vyos:"export,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetID returns the resource ID
func (o ProtocolsOspfAccessList) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o ProtocolsOspfAccessList) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfAccessList) GetVyosPath() []string {
	return []string{
		"protocols",

		"ospf",

		"access-list",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfAccessList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"access_list_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Access list to filter networks in routing updates

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Access-list number  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"export": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Filter for outgoing routing update

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  bgp  &emsp; |  Filter BGP routes  |
    |  connected  &emsp; |  Filter connected routes  |
    |  isis  &emsp; |  Filter IS-IS routes  |
    |  kernel  &emsp; |  Filter Kernel routes  |
    |  rip  &emsp; |  Filter RIP routes  |
    |  static  &emsp; |  Filter static routes  |

`,
		},

		// Nodes

	}
}
