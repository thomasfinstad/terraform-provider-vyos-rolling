// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VpnSstpAuthenticationLocalUsersUsername describes the resource data model.
type VpnSstpAuthenticationLocalUsersUsername struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	// LeafNodes
	LeafVpnSstpAuthenticationLocalUsersUsernameDisable  types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafVpnSstpAuthenticationLocalUsersUsernamePassword types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafVpnSstpAuthenticationLocalUsersUsernameStaticIP types.String `tfsdk:"static_ip" vyos:"static-ip,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVpnSstpAuthenticationLocalUsersUsernameRateLimit *VpnSstpAuthenticationLocalUsersUsernameRateLimit `tfsdk:"rate_limit" vyos:"rate-limit,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnSstpAuthenticationLocalUsersUsername) GetVyosPath() []string {
	return []string{
		"vpn",

		"sstp",

		"authentication",

		"local-users",

		"username",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnSstpAuthenticationLocalUsersUsername) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `User name for authentication

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

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Password for authentication

`,
		},

		"static_ip": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Static client IP address

`,

			// Default:          stringdefault.StaticString(`&`),
			Computed: true,
		},

		// Nodes

		"rate_limit": schema.SingleNestedAttribute{
			Attributes: VpnSstpAuthenticationLocalUsersUsernameRateLimit{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Upload/Download speed limits

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnSstpAuthenticationLocalUsersUsername) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VpnSstpAuthenticationLocalUsersUsername) UnmarshalJSON(_ []byte) error {
	return nil
}
