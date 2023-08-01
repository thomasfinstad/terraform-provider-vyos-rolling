// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NetnsName describes the resource data model.
type NetnsName struct {
	SelfIdentifier types.String `tfsdk:"name_id" vyos:",self-id"`

	// LeafNodes
	LeafNetnsNameDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *NetnsName) GetVyosPath() []string {
	return []string{
		"netns",

		"name",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o NetnsName) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Network namespace name

`,
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *NetnsName) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *NetnsName) UnmarshalJSON(_ []byte) error {
	return nil
}