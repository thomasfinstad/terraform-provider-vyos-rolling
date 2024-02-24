// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// SystemConntrackTimeoutTCP describes the resource data model.
type SystemConntrackTimeoutTCP struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	// LeafNodes
	LeafSystemConntrackTimeoutTCPCloseWait   types.Number `tfsdk:"close_wait" vyos:"close-wait,omitempty"`
	LeafSystemConntrackTimeoutTCPClose       types.Number `tfsdk:"close" vyos:"close,omitempty"`
	LeafSystemConntrackTimeoutTCPEstablished types.Number `tfsdk:"established" vyos:"established,omitempty"`
	LeafSystemConntrackTimeoutTCPFinWait     types.Number `tfsdk:"fin_wait" vyos:"fin-wait,omitempty"`
	LeafSystemConntrackTimeoutTCPLastAck     types.Number `tfsdk:"last_ack" vyos:"last-ack,omitempty"`
	LeafSystemConntrackTimeoutTCPSynRecv     types.Number `tfsdk:"syn_recv" vyos:"syn-recv,omitempty"`
	LeafSystemConntrackTimeoutTCPSynSent     types.Number `tfsdk:"syn_sent" vyos:"syn-sent,omitempty"`
	LeafSystemConntrackTimeoutTCPTimeWait    types.Number `tfsdk:"time_wait" vyos:"time-wait,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *SystemConntrackTimeoutTCP) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemConntrackTimeoutTCP) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"system",

		"conntrack",

		"timeout",

		"tcp",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemConntrackTimeoutTCP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"close_wait": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP CLOSE-WAIT timeout in seconds

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-21474836  &emsp; |  TCP CLOSE-WAIT timeout in seconds  |

`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		"close": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP CLOSE timeout in seconds

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-21474836  &emsp; |  TCP CLOSE timeout in seconds  |

`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"established": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP ESTABLISHED timeout in seconds

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-21474836  &emsp; |  TCP ESTABLISHED timeout in seconds  |

`,

			// Default:          stringdefault.StaticString(`432000`),
			Computed: true,
		},

		"fin_wait": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP FIN-WAIT timeout in seconds

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-21474836  &emsp; |  TCP FIN-WAIT timeout in seconds  |

`,

			// Default:          stringdefault.StaticString(`120`),
			Computed: true,
		},

		"last_ack": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP LAST-ACK timeout in seconds

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-21474836  &emsp; |  TCP LAST-ACK timeout in seconds  |

`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"syn_recv": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP SYN-RECEIVED timeout in seconds

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-21474836  &emsp; |  TCP SYN-RECEIVED timeout in seconds  |

`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		"syn_sent": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP SYN-SENT timeout in seconds

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-21474836  &emsp; |  TCP SYN-SENT timeout in seconds  |

`,

			// Default:          stringdefault.StaticString(`120`),
			Computed: true,
		},

		"time_wait": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP TIME-WAIT timeout in seconds

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-21474836  &emsp; |  TCP TIME-WAIT timeout in seconds  |

`,

			// Default:          stringdefault.StaticString(`120`),
			Computed: true,
		},
	}
}
