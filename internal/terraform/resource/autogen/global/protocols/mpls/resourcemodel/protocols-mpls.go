// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsMpls describes the resource data model.
type ProtocolsMpls struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	// LeafNodes
	LeafProtocolsMplsInterface types.List `tfsdk:"interface" vyos:"interface,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeProtocolsMplsLdp        bool `tfsdk:"-" vyos:"ldp,ignore,omitempty"`
	ExistsNodeProtocolsMplsParameters bool `tfsdk:"-" vyos:"parameters,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsMpls) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsMpls) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"mpls",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsMpls) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"interface": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Interface to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |

`,
		},
	}
}
