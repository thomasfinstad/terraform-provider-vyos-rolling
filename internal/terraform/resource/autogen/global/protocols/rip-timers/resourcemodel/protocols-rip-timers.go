// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsRIPTimers describes the resource data model.
type ProtocolsRIPTimers struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafProtocolsRIPTimersGarbageCollection types.Number `tfsdk:"garbage_collection" vyos:"garbage-collection,omitempty"`
	LeafProtocolsRIPTimersTimeout           types.Number `tfsdk:"timeout" vyos:"timeout,omitempty"`
	LeafProtocolsRIPTimersUpdate            types.Number `tfsdk:"update" vyos:"update,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsRIPTimers) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsRIPTimers) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"rip",

		"timers",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPTimers) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"garbage_collection": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Garbage collection timer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 5-2147483647  &emsp; |  Garbage colletion time  |

`,

			// Default:          stringdefault.StaticString(`120`),
			Computed: true,
		},

		"timeout": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Routing information timeout timer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 5-2147483647  &emsp; |  Routing information timeout timer  |

`,

			// Default:          stringdefault.StaticString(`180`),
			Computed: true,
		},

		"update": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Routing table update timer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 5-2147483647  &emsp; |  Routing table update timer in seconds  |

`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},
	}
}