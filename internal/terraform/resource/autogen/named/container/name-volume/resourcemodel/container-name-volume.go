// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ContainerNameVolume describes the resource data model.
type ContainerNameVolume struct {
	SelfIdentifier types.String `tfsdk:"volume_id" vyos:",self-id"`

	ParentIDContainerName types.String `tfsdk:"name" vyos:"name,parent-id"`

	// LeafNodes
	LeafContainerNameVolumeSource      types.String `tfsdk:"source" vyos:"source,omitempty"`
	LeafContainerNameVolumeDestination types.String `tfsdk:"destination" vyos:"destination,omitempty"`
	LeafContainerNameVolumeMode        types.String `tfsdk:"mode" vyos:"mode,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ContainerNameVolume) GetVyosPath() []string {
	return []string{
		"container",

		"name",
		o.ParentIDContainerName.ValueString(),

		"volume",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ContainerNameVolume) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"volume_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Mount a volume into the container

`,
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Container name

`,
		},

		// LeafNodes

		"source": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source host directory

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Source host directory  |

`,
		},

		"destination": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Destination container directory

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Destination container directory  |

`,
		},

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Volume access mode ro/rw

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ro  &emsp; |  Volume mounted into the container as read-only  |
    |  rw  &emsp; |  Volume mounted into the container as read-write  |

`,

			// Default:          stringdefault.StaticString(`rw`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ContainerNameVolume) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ContainerNameVolume) UnmarshalJSON(_ []byte) error {
	return nil
}
