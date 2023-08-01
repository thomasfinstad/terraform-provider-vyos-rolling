// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ContainerNamePort describes the resource data model.
type ContainerNamePort struct {
	SelfIdentifier types.String `tfsdk:"port_id" vyos:",self-id"`

	ParentIDContainerName types.String `tfsdk:"name" vyos:"name,parent-id"`

	// LeafNodes
	LeafContainerNamePortSource      types.String `tfsdk:"source" vyos:"source,omitempty"`
	LeafContainerNamePortDestination types.String `tfsdk:"destination" vyos:"destination,omitempty"`
	LeafContainerNamePortProtocol    types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ContainerNamePort) GetVyosPath() []string {
	return []string{
		"container",

		"name",
		o.ParentIDContainerName.ValueString(),

		"port",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ContainerNamePort) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"port_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Publish port to the container

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
			MarkdownDescription: `Source host port

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Source host port  |
    |  start-end  &emsp; |  Source host port range (e.g. 10025-10030)  |

`,
		},

		"destination": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Destination container port

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Destination container port  |
    |  start-end  &emsp; |  Destination container port range (e.g. 10025-10030)  |

`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Transport protocol used for port mapping

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  tcp  &emsp; |  Use Transmission Control Protocol for given port  |
    |  udp  &emsp; |  Use User Datagram Protocol for given port  |

`,

			// Default:          stringdefault.StaticString(`tcp`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ContainerNamePort) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ContainerNamePort) UnmarshalJSON(_ []byte) error {
	return nil
}