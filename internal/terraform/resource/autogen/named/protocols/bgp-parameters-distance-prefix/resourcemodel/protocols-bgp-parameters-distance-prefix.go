// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpParametersDistancePrefix describes the resource data model.
type ProtocolsBgpParametersDistancePrefix struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"prefix_id" vyos:",self-id"`

	// LeafNodes
	LeafProtocolsBgpParametersDistancePrefixDistance types.Number `tfsdk:"distance" vyos:"distance,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetID returns the resource ID
func (o ProtocolsBgpParametersDistancePrefix) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o ProtocolsBgpParametersDistancePrefix) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpParametersDistancePrefix) GetVyosPath() []string {
	return []string{
		"protocols",

		"bgp",

		"parameters",

		"distance",

		"prefix",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpParametersDistancePrefix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"prefix_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Administrative distance for a specific BGP prefix

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  Administrative distance for a specific BGP prefix  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"distance": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Administrative distance for prefix

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Administrative distance for external BGP routes  |

`,
		},

		// Nodes

	}
}
