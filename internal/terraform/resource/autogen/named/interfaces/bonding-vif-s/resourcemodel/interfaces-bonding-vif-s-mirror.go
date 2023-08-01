// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesBondingVifSMirror describes the resource data model.
type InterfacesBondingVifSMirror struct {
	// LeafNodes
	LeafInterfacesBondingVifSMirrorIngress types.String `tfsdk:"ingress" vyos:"ingress,omitempty"`
	LeafInterfacesBondingVifSMirrorEgress  types.String `tfsdk:"egress" vyos:"egress,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBondingVifSMirror) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ingress": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Mirror ingress traffic to destination interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Destination interface name  |

`,
		},

		"egress": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Mirror egress traffic to destination interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Destination interface name  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesBondingVifSMirror) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesBondingVifSMirror) UnmarshalJSON(_ []byte) error {
	return nil
}