// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VpnSstpAuthenticationRadiusServer describes the resource data model.
type VpnSstpAuthenticationRadiusServer struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"server_id" vyos:",self-id"`

	// LeafNodes
	LeafVpnSstpAuthenticationRadiusServerDisable           types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafVpnSstpAuthenticationRadiusServerKey               types.String `tfsdk:"key" vyos:"key,omitempty"`
	LeafVpnSstpAuthenticationRadiusServerPort              types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafVpnSstpAuthenticationRadiusServerAcctPort          types.Number `tfsdk:"acct_port" vyos:"acct-port,omitempty"`
	LeafVpnSstpAuthenticationRadiusServerDisableAccounting types.Bool   `tfsdk:"disable_accounting" vyos:"disable-accounting,omitempty"`
	LeafVpnSstpAuthenticationRadiusServerFailTime          types.Number `tfsdk:"fail_time" vyos:"fail-time,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetID returns the resource ID
func (o VpnSstpAuthenticationRadiusServer) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o VpnSstpAuthenticationRadiusServer) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnSstpAuthenticationRadiusServer) GetVyosPath() []string {
	return []string{
		"vpn",

		"sstp",

		"authentication",

		"radius",

		"server",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnSstpAuthenticationRadiusServer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"server_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `RADIUS server configuration

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  RADIUS server IPv4 address  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Shared secret key

`,
		},

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Authentication port

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`1812`),
			Computed: true,
		},

		"acct_port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Accounting port

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`1813`),
			Computed: true,
		},

		"disable_accounting": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable accounting

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"fail_time": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Mark server unavailable for <n> seconds on failure

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-600  &emsp; |  Fail time penalty  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		// Nodes

	}
}
