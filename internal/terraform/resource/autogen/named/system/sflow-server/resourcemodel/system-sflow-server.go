// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SystemSflowServer describes the resource data model.
type SystemSflowServer struct {
	SelfIdentifier types.String `tfsdk:"server_id" vyos:",self-id"`

	// LeafNodes
	LeafSystemSflowServerPort types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemSflowServer) GetVyosPath() []string {
	return []string{
		"system",

		"sflow",

		"server",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemSflowServer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"server_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `sFlow destination server

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 server to export sFlow  |
    |  ipv6  &emsp; |  IPv6 server to export sFlow  |

`,
		},

		// LeafNodes

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`6343`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *SystemSflowServer) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *SystemSflowServer) UnmarshalJSON(_ []byte) error {
	return nil
}