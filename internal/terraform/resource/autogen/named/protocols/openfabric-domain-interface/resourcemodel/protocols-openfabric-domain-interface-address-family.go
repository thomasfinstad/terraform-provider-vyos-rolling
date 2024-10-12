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

var _ helpers.VyosResourceDataModel = &ProtocolsOpenfabricDomainInterfaceAddressFamily{}

// ProtocolsOpenfabricDomainInterfaceAddressFamily describes the resource data model.
type ProtocolsOpenfabricDomainInterfaceAddressFamily struct {
	// LeafNodes
	LeafProtocolsOpenfabricDomainInterfaceAddressFamilyIPvfour types.Bool `tfsdk:"ipv4" vyos:"ipv4,omitempty"`
	LeafProtocolsOpenfabricDomainInterfaceAddressFamilyIPvsix  types.Bool `tfsdk:"ipv6" vyos:"ipv6,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOpenfabricDomainInterfaceAddressFamily) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ipv4": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 OpenFabric

`,
			Description: `IPv4 OpenFabric

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ipv6": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `IPv6 OpenFabric

`,
			Description: `IPv6 OpenFabric

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
