// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VpnSstpAuthenticationRadiusServer describes the resource data model.
type VpnSstpAuthenticationRadiusServer struct {
	// LeafNodes
	VpnSstpAuthenticationRadiusServerDisable           customtypes.CustomStringValue `tfsdk:"disable" json:"disable,omitempty"`
	VpnSstpAuthenticationRadiusServerKey               customtypes.CustomStringValue `tfsdk:"key" json:"key,omitempty"`
	VpnSstpAuthenticationRadiusServerPort              customtypes.CustomStringValue `tfsdk:"port" json:"port,omitempty"`
	VpnSstpAuthenticationRadiusServerAcctPort          customtypes.CustomStringValue `tfsdk:"acct_port" json:"acct-port,omitempty"`
	VpnSstpAuthenticationRadiusServerDisableAccounting customtypes.CustomStringValue `tfsdk:"disable_accounting" json:"disable-accounting,omitempty"`
	VpnSstpAuthenticationRadiusServerFailTime          customtypes.CustomStringValue `tfsdk:"fail_time" json:"fail-time,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VpnSstpAuthenticationRadiusServer) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable instance

`,
		},

		"key": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Shared secret key

`,
		},

		"port": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Authentication port

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`1812`),
			Computed: true,
		},

		"acct_port": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Accounting port

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`1813`),
			Computed: true,
		},

		"disable_accounting": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable accounting

`,
		},

		"fail_time": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Mark server unavailable for <n> seconds on failure

|  Format  |  Description  |
|----------|---------------|
|  u32:0-600  |  Fail time penalty  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
