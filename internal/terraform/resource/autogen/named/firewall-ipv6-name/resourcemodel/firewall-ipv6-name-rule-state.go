// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// FirewallIPvsixNameRuleState describes the resource data model.
type FirewallIPvsixNameRuleState struct {
	// LeafNodes
	FirewallIPvsixNameRuleStateEstablished customtypes.CustomStringValue `tfsdk:"established" json:"established,omitempty"`
	FirewallIPvsixNameRuleStateInvalID     customtypes.CustomStringValue `tfsdk:"invalid" json:"invalid,omitempty"`
	FirewallIPvsixNameRuleStateNew         customtypes.CustomStringValue `tfsdk:"new" json:"new,omitempty"`
	FirewallIPvsixNameRuleStateRelated     customtypes.CustomStringValue `tfsdk:"related" json:"related,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o FirewallIPvsixNameRuleState) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"established": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Established state

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable  |
|  disable  |  Disable  |

`,
		},

		"invalid": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Invalid state

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable  |
|  disable  |  Disable  |

`,
		},

		"new": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `New state

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable  |
|  disable  |  Disable  |

`,
		},

		"related": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Related state

|  Format  |  Description  |
|----------|---------------|
|  enable  |  Enable  |
|  disable  |  Disable  |

`,
		},

		// TagNodes

		// Nodes

	}
}
