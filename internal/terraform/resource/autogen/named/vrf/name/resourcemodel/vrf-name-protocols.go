// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocols{}

// VrfNameProtocols describes the resource data model.
type VrfNameProtocols struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsBgp        *VrfNameProtocolsBgp        `tfsdk:"bgp" vyos:"bgp,omitempty"`
	NodeVrfNameProtocolsEigrp      *VrfNameProtocolsEigrp      `tfsdk:"eigrp" vyos:"eigrp,omitempty"`
	NodeVrfNameProtocolsIsis       *VrfNameProtocolsIsis       `tfsdk:"isis" vyos:"isis,omitempty"`
	NodeVrfNameProtocolsOspf       *VrfNameProtocolsOspf       `tfsdk:"ospf" vyos:"ospf,omitempty"`
	NodeVrfNameProtocolsOspfvthree *VrfNameProtocolsOspfvthree `tfsdk:"ospfv3" vyos:"ospfv3,omitempty"`
	NodeVrfNameProtocolsStatic     *VrfNameProtocolsStatic     `tfsdk:"static" vyos:"static,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocols) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"bgp": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgp{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Border Gateway Protocol (BGP)

`,
			Description: `Border Gateway Protocol (BGP)

`,
		},

		"eigrp": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsEigrp{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Enhanced Interior Gateway Routing Protocol (EIGRP)

`,
			Description: `Enhanced Interior Gateway Routing Protocol (EIGRP)

`,
		},

		"isis": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsis{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Intermediate System to Intermediate System (IS-IS)

`,
			Description: `Intermediate System to Intermediate System (IS-IS)

`,
		},

		"ospf": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspf{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Open Shortest Path First (OSPF)

`,
			Description: `Open Shortest Path First (OSPF)

`,
		},

		"ospfv3": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfvthree{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Open Shortest Path First (OSPF) for IPv6

`,
			Description: `Open Shortest Path First (OSPF) for IPv6

`,
		},

		"static": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsStatic{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Static Routing

`,
			Description: `Static Routing

`,
		},
	}
}
