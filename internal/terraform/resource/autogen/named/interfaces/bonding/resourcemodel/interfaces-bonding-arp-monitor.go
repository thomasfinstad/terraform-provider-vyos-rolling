/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (arp-monitor) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesBondingArpMonitor{}

// InterfacesBondingArpMonitor describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type InterfacesBondingArpMonitor struct {
	// LeafNodes
	LeafInterfacesBondingArpMonitorInterval types.Number `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafInterfacesBondingArpMonitorTarget   types.List   `tfsdk:"target" vyos:"target,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBondingArpMonitor) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (interval) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `ARP link monitoring interval

    |  Format  |  Description                                                  |
    |----------|---------------------------------------------------------------|
    |  u32     |  Specifies the ARP link monitoring frequency in milliseconds  |
`,
			Description: `ARP link monitoring interval

    |  Format  |  Description                                                  |
    |----------|---------------------------------------------------------------|
    |  u32     |  Specifies the ARP link monitoring frequency in milliseconds  |
`,
		},

		"target":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (target) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address used for ARP monitoring

    |  Format  |  Description                                                    |
    |----------|-----------------------------------------------------------------|
    |  ipv4    |  Specify IPv4 address of ARP requests when interval is enabled  |
`,
			Description: `IP address used for ARP monitoring

    |  Format  |  Description                                                    |
    |----------|-----------------------------------------------------------------|
    |  ipv4    |  Specify IPv4 address of ARP requests when interval is enabled  |
`,
		},

		// TagNodes

		// Nodes

	}
}
