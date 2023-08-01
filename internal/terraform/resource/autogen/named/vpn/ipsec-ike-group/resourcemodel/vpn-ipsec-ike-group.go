// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VpnIPsecIkeGroup describes the resource data model.
type VpnIPsecIkeGroup struct {
	SelfIdentifier types.String `tfsdk:"ike_group_id" vyos:",self-id"`

	// LeafNodes
	LeafVpnIPsecIkeGroupCloseAction   types.String `tfsdk:"close_action" vyos:"close-action,omitempty"`
	LeafVpnIPsecIkeGroupIkevtwoReauth types.Bool   `tfsdk:"ikev2_reauth" vyos:"ikev2-reauth,omitempty"`
	LeafVpnIPsecIkeGroupKeyExchange   types.String `tfsdk:"key_exchange" vyos:"key-exchange,omitempty"`
	LeafVpnIPsecIkeGroupLifetime      types.Number `tfsdk:"lifetime" vyos:"lifetime,omitempty"`
	LeafVpnIPsecIkeGroupDisableMobike types.Bool   `tfsdk:"disable_mobike" vyos:"disable-mobike,omitempty"`
	LeafVpnIPsecIkeGroupMode          types.String `tfsdk:"mode" vyos:"mode,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVpnIPsecIkeGroupProposal bool `tfsdk:"proposal" vyos:"proposal,child"`

	// Nodes
	NodeVpnIPsecIkeGroupDeadPeerDetection *VpnIPsecIkeGroupDeadPeerDetection `tfsdk:"dead_peer_detection" vyos:"dead-peer-detection,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnIPsecIkeGroup) GetVyosPath() []string {
	return []string{
		"vpn",

		"ipsec",

		"ike-group",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecIkeGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"ike_group_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Internet Key Exchange (IKE) group name

`,
		},

		// LeafNodes

		"close_action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Action to take if a child SA is unexpectedly closed

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  none  &emsp; |  Do nothing  |
    |  hold  &emsp; |  Attempt to re-negotiate when matching traffic is seen  |
    |  restart  &emsp; |  Attempt to re-negotiate the connection immediately  |

`,

			// Default:          stringdefault.StaticString(`none`),
			Computed: true,
		},

		"ikev2_reauth": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Re-authentication of the remote peer during an IKE re-key (IKEv2 only)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"key_exchange": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IKE version

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ikev1  &emsp; |  Use IKEv1 for key exchange  |
    |  ikev2  &emsp; |  Use IKEv2 for key exchange  |

`,
		},

		"lifetime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `IKE lifetime

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 30-86400  &emsp; |  IKE lifetime in seconds  |

`,

			// Default:          stringdefault.StaticString(`28800`),
			Computed: true,
		},

		"disable_mobike": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable MOBIKE Support (IKEv2 only)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IKEv1 phase 1 mode

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  main  &emsp; |  Use the main mode (recommended)  |
    |  aggressive  &emsp; |  Use the aggressive mode (insecure, not recommended)  |

`,

			// Default:          stringdefault.StaticString(`main`),
			Computed: true,
		},

		// Nodes

		"dead_peer_detection": schema.SingleNestedAttribute{
			Attributes: VpnIPsecIkeGroupDeadPeerDetection{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Dead Peer Detection (DPD)

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnIPsecIkeGroup) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VpnIPsecIkeGroup) UnmarshalJSON(_ []byte) error {
	return nil
}