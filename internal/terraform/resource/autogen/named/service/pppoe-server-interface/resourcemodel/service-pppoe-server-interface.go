// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServicePppoeServerInterface describes the resource data model.
type ServicePppoeServerInterface struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	// LeafNodes
	LeafServicePppoeServerInterfaceVlan types.List `tfsdk:"vlan" vyos:"vlan,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServicePppoeServerInterface) GetVyosPath() []string {
	return []string{
		"service",

		"pppoe-server",

		"interface",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServicePppoeServerInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `interface(s) to listen on

`,
		},

		// LeafNodes

		"vlan": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `VLAN monitor for automatic creation of VLAN interfaces

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-4094  |  VLAN for automatic creation  |
    |  start-end  |  VLAN range for automatic creation (e.g. 1-4094)  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServicePppoeServerInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServicePppoeServerInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
