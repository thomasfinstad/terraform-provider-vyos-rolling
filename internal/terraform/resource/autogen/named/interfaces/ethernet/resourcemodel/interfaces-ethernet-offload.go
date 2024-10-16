// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesEthernetOffload{}

// InterfacesEthernetOffload describes the resource data model.
type InterfacesEthernetOffload struct {
	// LeafNodes
	LeafInterfacesEthernetOffloadGro         types.Bool `tfsdk:"gro" vyos:"gro,omitempty"`
	LeafInterfacesEthernetOffloadGso         types.Bool `tfsdk:"gso" vyos:"gso,omitempty"`
	LeafInterfacesEthernetOffloadHwTcOffload types.Bool `tfsdk:"hw_tc_offload" vyos:"hw-tc-offload,omitempty"`
	LeafInterfacesEthernetOffloadLro         types.Bool `tfsdk:"lro" vyos:"lro,omitempty"`
	LeafInterfacesEthernetOffloadRps         types.Bool `tfsdk:"rps" vyos:"rps,omitempty"`
	LeafInterfacesEthernetOffloadRfs         types.Bool `tfsdk:"rfs" vyos:"rfs,omitempty"`
	LeafInterfacesEthernetOffloadSg          types.Bool `tfsdk:"sg" vyos:"sg,omitempty"`
	LeafInterfacesEthernetOffloadTso         types.Bool `tfsdk:"tso" vyos:"tso,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernetOffload) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"gro": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable Generic Receive Offload

`,
			Description: `Enable Generic Receive Offload

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"gso": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable Generic Segmentation Offload

`,
			Description: `Enable Generic Segmentation Offload

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"hw_tc_offload": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable Hardware Flow Offload

`,
			Description: `Enable Hardware Flow Offload

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"lro": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable Large Receive Offload

`,
			Description: `Enable Large Receive Offload

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rps": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable Receive Packet Steering

`,
			Description: `Enable Receive Packet Steering

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rfs": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable Receive Flow Steering

`,
			Description: `Enable Receive Flow Steering

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"sg": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable Scatter-Gather

`,
			Description: `Enable Scatter-Gather

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"tso": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable TCP Segmentation Offloading

`,
			Description: `Enable TCP Segmentation Offloading

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
