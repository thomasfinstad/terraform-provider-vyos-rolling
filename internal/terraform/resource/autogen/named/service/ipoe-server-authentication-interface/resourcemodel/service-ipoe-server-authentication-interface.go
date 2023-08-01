// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceIPoeServerAuthenticationInterface describes the resource data model.
type ServiceIPoeServerAuthenticationInterface struct {
	SelfIdentifier types.String `tfsdk:"interface_id" vyos:",self-id"`

	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagServiceIPoeServerAuthenticationInterfaceMac bool `tfsdk:"mac" vyos:"mac,child"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceIPoeServerAuthenticationInterface) GetVyosPath() []string {
	return []string{
		"service",

		"ipoe-server",

		"authentication",

		"interface",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIPoeServerAuthenticationInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Network interface for client MAC addresses

`,
		},

		// LeafNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceIPoeServerAuthenticationInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceIPoeServerAuthenticationInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
