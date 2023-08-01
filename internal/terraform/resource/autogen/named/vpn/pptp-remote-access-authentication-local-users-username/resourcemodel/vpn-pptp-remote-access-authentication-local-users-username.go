// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VpnPptpRemoteAccessAuthenticationLocalUsersUsername describes the resource data model.
type VpnPptpRemoteAccessAuthenticationLocalUsersUsername struct {
	SelfIdentifier types.String `tfsdk:"username_id" vyos:",self-id"`

	// LeafNodes
	LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernameDisable  types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernamePassword types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafVpnPptpRemoteAccessAuthenticationLocalUsersUsernameStaticIP types.String `tfsdk:"static_ip" vyos:"static-ip,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnPptpRemoteAccessAuthenticationLocalUsersUsername) GetVyosPath() []string {
	return []string{
		"vpn",

		"pptp",

		"remote-access",

		"authentication",

		"local-users",

		"username",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnPptpRemoteAccessAuthenticationLocalUsersUsername) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"username_id": schema.StringAttribute{
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
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnPptpRemoteAccessAuthenticationLocalUsersUsername) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VpnPptpRemoteAccessAuthenticationLocalUsersUsername) UnmarshalJSON(_ []byte) error {
	return nil
}
