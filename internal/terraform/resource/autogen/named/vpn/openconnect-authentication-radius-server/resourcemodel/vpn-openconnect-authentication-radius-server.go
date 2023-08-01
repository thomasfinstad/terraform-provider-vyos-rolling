// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VpnOpenconnectAuthenticationRadiusServer describes the resource data model.
type VpnOpenconnectAuthenticationRadiusServer struct {
	SelfIdentifier types.String `tfsdk:"server_id" vyos:",self-id"`

	// LeafNodes
	LeafVpnOpenconnectAuthenticationRadiusServerDisable types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafVpnOpenconnectAuthenticationRadiusServerKey     types.String `tfsdk:"key" vyos:"key,omitempty"`
	LeafVpnOpenconnectAuthenticationRadiusServerPort    types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnOpenconnectAuthenticationRadiusServer) GetVyosPath() []string {
	return []string{
		"vpn",

		"openconnect",

		"authentication",

		"radius",

		"server",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnOpenconnectAuthenticationRadiusServer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `server_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
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

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnOpenconnectAuthenticationRadiusServer) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VpnOpenconnectAuthenticationRadiusServer) UnmarshalJSON(_ []byte) error {
	return nil
}
