// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// SystemConntrackLogTCP describes the resource data model.
type SystemConntrackLogTCP struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafSystemConntrackLogTCPDestroy types.Bool `tfsdk:"destroy" vyos:"destroy,omitempty"`
	LeafSystemConntrackLogTCPNew     types.Bool `tfsdk:"new" vyos:"new,omitempty"`
	LeafSystemConntrackLogTCPUpdate  types.Bool `tfsdk:"update" vyos:"update,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *SystemConntrackLogTCP) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemConntrackLogTCP) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"system",

		"conntrack",

		"log",

		"tcp",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemConntrackLogTCP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"destroy": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log connection deletion

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"new": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log connection creation

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"update": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log connection updates

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}