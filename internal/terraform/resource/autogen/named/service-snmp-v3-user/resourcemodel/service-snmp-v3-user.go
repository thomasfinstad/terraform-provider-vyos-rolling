// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// ServiceSnmpVthreeUser describes the resource data model.
type ServiceSnmpVthreeUser struct {
	// LeafNodes
	ServiceSnmpVthreeUserGroup customtypes.CustomStringValue `tfsdk:"group" json:"group,omitempty"`
	ServiceSnmpVthreeUserMode  customtypes.CustomStringValue `tfsdk:"mode" json:"mode,omitempty"`

	// TagNodes

	// Nodes
	ServiceSnmpVthreeUserAuth    types.Object `tfsdk:"auth" json:"auth,omitempty"`
	ServiceSnmpVthreeUserPrivacy types.Object `tfsdk:"privacy" json:"privacy,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o ServiceSnmpVthreeUser) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"group": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Specifies group for user name

`,
		},

		"mode": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Define access permission

|  Format  |  Description  |
|----------|---------------|
|  ro  |  Read-Only  |
|  rw  |  read write  |

`,

			// Default:          stringdefault.StaticString(`ro`),
			Computed: true,
		},

		// TagNodes

		// Nodes

		"auth": schema.SingleNestedAttribute{
			Attributes: ServiceSnmpVthreeUserAuth{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Specifies the auth

`,
		},

		"privacy": schema.SingleNestedAttribute{
			Attributes: ServiceSnmpVthreeUserPrivacy{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Defines the privacy

`,
		},
	}
}
