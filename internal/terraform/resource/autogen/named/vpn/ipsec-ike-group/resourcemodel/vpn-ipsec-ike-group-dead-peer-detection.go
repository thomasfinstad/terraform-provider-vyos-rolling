// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VpnIPsecIkeGroupDeadPeerDetection describes the resource data model.
type VpnIPsecIkeGroupDeadPeerDetection struct {
	// LeafNodes
	LeafVpnIPsecIkeGroupDeadPeerDetectionAction   types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafVpnIPsecIkeGroupDeadPeerDetectionInterval types.Number `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafVpnIPsecIkeGroupDeadPeerDetectionTimeout  types.Number `tfsdk:"timeout" vyos:"timeout,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecIkeGroupDeadPeerDetection) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Keep-alive failure action

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  hold  &emsp; |  Attempt to re-negotiate the connection when matching traffic is seen  |
    |  clear  &emsp; |  Remove the connection immediately  |
    |  restart  &emsp; |  Attempt to re-negotiate the connection immediately  |

`,

			// Default:          stringdefault.StaticString(`clear`),
			Computed: true,
		},

		"interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Keep-alive interval

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 2-86400  &emsp; |  Keep-alive interval in seconds  |

`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"timeout": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Dead Peer Detection keep-alive timeout (IKEv1 only)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 2-86400  &emsp; |  Keep-alive timeout in seconds  |

`,

			// Default:          stringdefault.StaticString(`120`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VpnIPsecIkeGroupDeadPeerDetection) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VpnIPsecIkeGroupDeadPeerDetection) UnmarshalJSON(_ []byte) error {
	return nil
}
