// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VpnIPsecProfile describes the resource data model.
type VpnIPsecProfile struct {
	SelfIdentifier types.String `tfsdk:"profile_id" vyos:",self-id"`

	// LeafNodes
	LeafVpnIPsecProfileDisable  types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafVpnIPsecProfileEspGroup types.String `tfsdk:"esp_group" vyos:"esp-group,omitempty"`
	LeafVpnIPsecProfileIkeGroup types.String `tfsdk:"ike_group" vyos:"ike-group,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVpnIPsecProfileAuthentication *VpnIPsecProfileAuthentication `tfsdk:"authentication" vyos:"authentication,omitempty"`
	NodeVpnIPsecProfileBind           *VpnIPsecProfileBind           `tfsdk:"bind" vyos:"bind,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnIPsecProfile) GetVyosPath() []string {
	return []string{
		"vpn",

		"ipsec",

		"profile",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecProfile) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"profile_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `VPN IPsec profile

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Profile name  |

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

		"esp_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Encapsulating Security Payloads (ESP) group name

`,
		},

		"ike_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Internet Key Exchange (IKE) group name

`,
		},

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: VpnIPsecProfileAuthentication{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication

`,
		},

		"bind": schema.SingleNestedAttribute{
			Attributes: VpnIPsecProfileBind{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `DMVPN tunnel configuration

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnIPsecProfile) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VpnIPsecProfile) UnmarshalJSON(_ []byte) error {
	return nil
}