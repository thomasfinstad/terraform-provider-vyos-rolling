// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// ProtocolsOspfInterfaceAuthenticationMdfive describes the resource data model.
type ProtocolsOspfInterfaceAuthenticationMdfive struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsOspfInterfaceAuthenticationMdfiveKeyID bool `tfsdk:"key_id" vyos:"key-id,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfInterfaceAuthenticationMdfive) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{

		// LeafNodes

		// Nodes

	}
}
