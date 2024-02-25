// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// HighAvailabilityVrrpGlobalParameters describes the resource data model.
type HighAvailabilityVrrpGlobalParameters struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafHighAvailabilityVrrpGlobalParametersStartupDelay types.Number `tfsdk:"startup_delay" vyos:"startup-delay,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeHighAvailabilityVrrpGlobalParametersGarp bool `tfsdk:"-" vyos:"garp,child"`
}

// SetID configures the resource ID
func (o *HighAvailabilityVrrpGlobalParameters) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *HighAvailabilityVrrpGlobalParameters) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"high-availability",

		"vrrp",

		"global-parameters",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpGlobalParameters) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"startup_delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Time VRRP startup process (in seconds)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-600  &emsp; |  Interval in seconds  |

`,
		},
	}
}
