// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsOspfDistance describes the resource data model.
type ProtocolsOspfDistance struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	// LeafNodes
	LeafProtocolsOspfDistanceGlobal types.Number `tfsdk:"global" vyos:"global,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeProtocolsOspfDistanceOspf bool `tfsdk:"-" vyos:"ospf,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsOspfDistance) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfDistance) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"ospf",

		"distance",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfDistance) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"global": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Administrative distance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Administrative distance  |

`,
		},
	}
}
