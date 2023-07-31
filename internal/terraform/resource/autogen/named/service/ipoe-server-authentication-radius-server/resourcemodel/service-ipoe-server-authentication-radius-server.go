// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceIPoeServerAuthenticationRadiusServer describes the resource data model.
type ServiceIPoeServerAuthenticationRadiusServer struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	// LeafNodes
	LeafServiceIPoeServerAuthenticationRadiusServerDisable           types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusServerKey               types.String `tfsdk:"key" vyos:"key,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusServerPort              types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusServerAcctPort          types.Number `tfsdk:"acct_port" vyos:"acct-port,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusServerDisableAccounting types.Bool   `tfsdk:"disable_accounting" vyos:"disable-accounting,omitempty"`
	LeafServiceIPoeServerAuthenticationRadiusServerFailTime          types.Number `tfsdk:"fail_time" vyos:"fail-time,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceIPoeServerAuthenticationRadiusServer) GetVyosPath() []string {
	return []string{
		"service",

		"ipoe-server",

		"authentication",

		"radius",

		"server",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIPoeServerAuthenticationRadiusServer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `RADIUS server configuration

    |  Format  |  Description  |
    |----------|---------------|
    |  ipv4  |  RADIUS server IPv4 address  |

`,
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

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65535  |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`1812`),
			Computed: true,
		},

		"acct_port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Accounting port

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-65535  |  Numeric IP port  |

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

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0-600  |  Fail time penalty  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceIPoeServerAuthenticationRadiusServer) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceIPoeServerAuthenticationRadiusServer) UnmarshalJSON(_ []byte) error {
	return nil
}
