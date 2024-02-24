// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsRIPngDistributeListPrefixList describes the resource data model.
type ProtocolsRIPngDistributeListPrefixList struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	// LeafNodes
	LeafProtocolsRIPngDistributeListPrefixListIn  types.String `tfsdk:"in" vyos:"in,omitempty"`
	LeafProtocolsRIPngDistributeListPrefixListOut types.String `tfsdk:"out" vyos:"out,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsRIPngDistributeListPrefixList) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsRIPngDistributeListPrefixList) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"ripng",

		"distribute-list",

		"prefix-list",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPngDistributeListPrefixList) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"in": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix-list to apply to input packets

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Prefix-list to apply to input packets  |

`,
		},

		"out": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Prefix-list to apply to output packets

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Prefix-list to apply to output packets  |

`,
		},
	}
}
