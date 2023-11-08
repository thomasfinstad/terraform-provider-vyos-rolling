// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsIsisSpfDelayIetf describes the resource data model.
type VrfNameProtocolsIsisSpfDelayIetf struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisSpfDelayIetfInitDelay   types.Number `tfsdk:"init_delay" vyos:"init-delay,omitempty"`
	LeafVrfNameProtocolsIsisSpfDelayIetfShortDelay  types.Number `tfsdk:"short_delay" vyos:"short-delay,omitempty"`
	LeafVrfNameProtocolsIsisSpfDelayIetfLongDelay   types.Number `tfsdk:"long_delay" vyos:"long-delay,omitempty"`
	LeafVrfNameProtocolsIsisSpfDelayIetfHolddown    types.Number `tfsdk:"holddown" vyos:"holddown,omitempty"`
	LeafVrfNameProtocolsIsisSpfDelayIetfTimeToLearn types.Number `tfsdk:"time_to_learn" vyos:"time-to-learn,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisSpfDelayIetf) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"init_delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Delay used while in QUIET state

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-60000  &emsp; |  Delay used while in QUIET state (in ms)  |

`,
		},

		"short_delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Delay used while in SHORT_WAIT state

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-60000  &emsp; |  Delay used while in SHORT_WAIT state (in ms)  |

`,
		},

		"long_delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Delay used while in LONG_WAIT

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-60000  &emsp; |  Delay used while in LONG_WAIT state in ms  |

`,
		},

		"holddown": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Time with no received IGP events before considering IGP stable

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-60000  &emsp; |  Time with no received IGP events before considering IGP stable in ms  |

`,
		},

		"time_to_learn": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum duration needed to learn all the events related to a single failure

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-60000  &emsp; |  Maximum duration needed to learn all the events related to a single failure in ms  |

`,
		},

		// Nodes

	}
}
