// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfive{}

// VrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfive describes the resource data model.
type VrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfive struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfiveKeyID bool `tfsdk:"key_id" vyos:"key-id,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspfAreaVirtualLinkAuthenticationMdfive) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{

		// LeafNodes

		// Nodes

	}
}
