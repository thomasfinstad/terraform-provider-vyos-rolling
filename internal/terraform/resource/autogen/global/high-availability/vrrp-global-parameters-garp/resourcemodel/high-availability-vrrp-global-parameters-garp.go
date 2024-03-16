// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance
var _ helpers.VyosTopResourceDataModel = &HighAvailabilityVrrpGlobalParametersGarp{}

// HighAvailabilityVrrpGlobalParametersGarp describes the resource data model.
type HighAvailabilityVrrpGlobalParametersGarp struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafHighAvailabilityVrrpGlobalParametersGarpInterval            types.String `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafHighAvailabilityVrrpGlobalParametersGarpMasterDelay         types.Number `tfsdk:"master_delay" vyos:"master-delay,omitempty"`
	LeafHighAvailabilityVrrpGlobalParametersGarpMasterRefresh       types.Number `tfsdk:"master_refresh" vyos:"master-refresh,omitempty"`
	LeafHighAvailabilityVrrpGlobalParametersGarpMasterRefreshRepeat types.Number `tfsdk:"master_refresh_repeat" vyos:"master-refresh-repeat,omitempty"`
	LeafHighAvailabilityVrrpGlobalParametersGarpMasterRepeat        types.Number `tfsdk:"master_repeat" vyos:"master-repeat,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *HighAvailabilityVrrpGlobalParametersGarp) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *HighAvailabilityVrrpGlobalParametersGarp) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"high-availability",

		"vrrp",

		"global-parameters",

		"garp",
	}
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *HighAvailabilityVrrpGlobalParametersGarp) GetVyosParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpGlobalParametersGarp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

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
	}
}
