// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ContainerNameEnvironment describes the resource data model.
type ContainerNameEnvironment struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	ParentIDContainerName types.String `tfsdk:"name" vyos:"name_identifier,parent-id"`

	// LeafNodes
	LeafContainerNameEnvironmentValue types.String `tfsdk:"value" vyos:"value,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ContainerNameEnvironment) GetVyosPath() []string {
	return []string{
		"container",

		"name",
		o.ParentIDContainerName.ValueString(),

		"environment",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ContainerNameEnvironment) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Add custom environment variables

`,
		},

		"name_identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Container name

`,
		},

		// LeafNodes

		"value": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set environment option value

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Set environment option value  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ContainerNameEnvironment) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ContainerNameEnvironment) UnmarshalJSON(_ []byte) error {
	return nil
}
