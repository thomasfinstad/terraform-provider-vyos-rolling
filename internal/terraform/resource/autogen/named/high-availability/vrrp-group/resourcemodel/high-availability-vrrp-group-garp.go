// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// HighAvailabilityVrrpGroupGarp describes the resource data model.
type HighAvailabilityVrrpGroupGarp struct {
	// LeafNodes
	LeafHighAvailabilityVrrpGroupGarpInterval            types.String `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafHighAvailabilityVrrpGroupGarpMasterDelay         types.Number `tfsdk:"master_delay" vyos:"master-delay,omitempty"`
	LeafHighAvailabilityVrrpGroupGarpMasterRefresh       types.Number `tfsdk:"master_refresh" vyos:"master-refresh,omitempty"`
	LeafHighAvailabilityVrrpGroupGarpMasterRefreshRepeat types.Number `tfsdk:"master_refresh_repeat" vyos:"master-refresh-repeat,omitempty"`
	LeafHighAvailabilityVrrpGroupGarpMasterRepeat        types.Number `tfsdk:"master_repeat" vyos:"master-repeat,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpGroupGarp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"interval": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interval between Gratuitous ARP

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <0.000-1000>  &emsp; |  Interval in seconds, resolution microseconds  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"master_delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Delay for second set of gratuitous ARPs after transition to master

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-1000  &emsp; |  Delay in seconds  |

`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		"master_refresh": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum time interval for refreshing gratuitous ARPs while beeing master

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0  &emsp; |  No refresh  |
    |  number: 1-255  &emsp; |  Interval in seconds  |

`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		"master_refresh_repeat": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of gratuitous ARP messages to send at a time while beeing master

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Number of gratuitous ARP messages  |

`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"master_repeat": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of gratuitous ARP messages to send at a time after transition to master

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Number of gratuitous ARP messages  |

`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *HighAvailabilityVrrpGroupGarp) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *HighAvailabilityVrrpGroupGarp) UnmarshalJSON(_ []byte) error {
	return nil
}
