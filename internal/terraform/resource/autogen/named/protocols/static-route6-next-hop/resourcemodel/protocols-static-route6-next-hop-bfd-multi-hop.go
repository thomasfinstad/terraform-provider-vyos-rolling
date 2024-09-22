// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsStaticRoutesixNextHopBfdMultiHop{}

// ProtocolsStaticRoutesixNextHopBfdMultiHop describes the resource data model.
type ProtocolsStaticRoutesixNextHopBfdMultiHop struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagProtocolsStaticRoutesixNextHopBfdMultiHopSource bool `tfsdk:"-" vyos:"source,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsStaticRoutesixNextHopBfdMultiHop) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{

		// LeafNodes

		// Nodes

	}
}
