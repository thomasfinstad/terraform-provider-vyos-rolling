// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VpnOpenconnectAuthentication describes the resource data model.
type VpnOpenconnectAuthentication struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafVpnOpenconnectAuthenticationGroup types.List `tfsdk:"group" vyos:"group,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeVpnOpenconnectAuthenticationMode       bool `tfsdk:"-" vyos:"mode,child"`
	ExistsNodeVpnOpenconnectAuthenticationLocalUsers bool `tfsdk:"-" vyos:"local-users,child"`
	ExistsNodeVpnOpenconnectAuthenticationRadius     bool `tfsdk:"-" vyos:"radius,child"`
}

// SetID configures the resource ID
func (o *VpnOpenconnectAuthentication) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnOpenconnectAuthentication) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vpn",

		"openconnect",

		"authentication",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnOpenconnectAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"group": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Group that a client is allowed to select (from a list). Maps to RADIUS Class attribute.

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Group string. The group may be followed by a user-friendly name in brackets: group1[First Group]  |

`,
		},
	}
}
