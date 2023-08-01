// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsOspfAreaRange describes the resource data model.
type VrfNameProtocolsOspfAreaRange struct {
	SelfIdentifier types.String `tfsdk:"range_id" vyos:",self-id"`

	ParentIDVrfName types.String `tfsdk:"name" vyos:"name,parent-id"`

	ParentIDVrfNameProtocolsOspfArea types.String `tfsdk:"area" vyos:"area,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsOspfAreaRangeCost         types.Number `tfsdk:"cost" vyos:"cost,omitempty"`
	LeafVrfNameProtocolsOspfAreaRangeNotAdvertise types.Bool   `tfsdk:"not_advertise" vyos:"not-advertise,omitempty"`
	LeafVrfNameProtocolsOspfAreaRangeSubstitute   types.String `tfsdk:"substitute" vyos:"substitute,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsOspfAreaRange) GetVyosPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"ospf",

		"area",
		o.ParentIDVrfNameProtocolsOspfArea.ValueString(),

		"range",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfAreaRange) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"range_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Summarize routes matching a prefix (border routers only)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  Area range prefix  |

`,
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

`,
		},

		"area_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `OSPF area settings

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  OSPF area number in decimal notation  |
    |  ipv4  &emsp; |  OSPF area number in dotted decimal notation  |

`,
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

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsOspfAreaRange) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsOspfAreaRange) UnmarshalJSON(_ []byte) error {
	return nil
}
