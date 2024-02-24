// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// VpnOpenconnectAuthenticationRadius describes the resource data model.
type VpnOpenconnectAuthenticationRadius struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	// LeafNodes
	LeafVpnOpenconnectAuthenticationRadiusSourceAddress types.String `tfsdk:"source_address" vyos:"source-address,omitempty"`
	LeafVpnOpenconnectAuthenticationRadiusTimeout       types.Number `tfsdk:"timeout" vyos:"timeout,omitempty"`
	LeafVpnOpenconnectAuthenticationRadiusGroupconfig   types.String `tfsdk:"groupconfig" vyos:"groupconfig,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVpnOpenconnectAuthenticationRadiusServer bool `tfsdk:"-" vyos:"server,ignore,child"`

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *VpnOpenconnectAuthenticationRadius) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnOpenconnectAuthenticationRadius) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"vpn",

		"openconnect",

		"authentication",

		"radius",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnOpenconnectAuthenticationRadius) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"source_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 source address used to initiate connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 source address  |

`,
		},

		"timeout": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Session timeout

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-240  &emsp; |  Session timeout in seconds (default: 2)  |

`,

			// Default:          stringdefault.StaticString(`2`),
			Computed: true,
		},

		"groupconfig": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `If the groupconfig option is set, then config-per-user will be overriden, and all configuration will be read from RADIUS.

`,
		},
	}
}
