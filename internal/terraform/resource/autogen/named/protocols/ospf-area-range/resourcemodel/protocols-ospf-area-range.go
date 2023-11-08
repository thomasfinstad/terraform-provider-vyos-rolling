// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsOspfAreaRange describes the resource data model.
type ProtocolsOspfAreaRange struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"range_id" vyos:",self-id"`

	ParentIDProtocolsOspfArea types.String `tfsdk:"area" vyos:"area,parent-id"`

	// LeafNodes
	LeafProtocolsOspfAreaRangeCost         types.Number `tfsdk:"cost" vyos:"cost,omitempty"`
	LeafProtocolsOspfAreaRangeNotAdvertise types.Bool   `tfsdk:"not_advertise" vyos:"not-advertise,omitempty"`
	LeafProtocolsOspfAreaRangeSubstitute   types.String `tfsdk:"substitute" vyos:"substitute,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetID returns the resource ID
func (o ProtocolsOspfAreaRange) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o ProtocolsOspfAreaRange) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfAreaRange) GetVyosPath() []string {
	return []string{
		"protocols",

		"ospf",

		"area",
		o.ParentIDProtocolsOspfArea.ValueString(),

		"range",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfAreaRange) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"range_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Summarize routes matching a prefix (border routers only)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  Area range prefix  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"area_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `OSPF area settings

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  OSPF area number in decimal notation  |
    |  ipv4  &emsp; |  OSPF area number in dotted decimal notation  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"cost": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Metric for this range

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-16777215  &emsp; |  Metric for this range  |

`,
		},

		"not_advertise": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not advertise this range

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"substitute": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Advertise area range as another prefix

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  Advertise area range as another prefix  |

`,
		},

		// Nodes

	}
}
