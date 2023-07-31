// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesEthernetVifSVifCMirror describes the resource data model.
type InterfacesEthernetVifSVifCMirror struct {
	// LeafNodes
	LeafInterfacesEthernetVifSVifCMirrorIngress types.String `tfsdk:"ingress" vyos:"ingress,omitempty"`
	LeafInterfacesEthernetVifSVifCMirrorEgress  types.String `tfsdk:"egress" vyos:"egress,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernetVifSVifCMirror) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ingress": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Mirror ingress traffic to destination interface

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Destination interface name  |

`,
		},

		"egress": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Mirror egress traffic to destination interface

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Destination interface name  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesEthernetVifSVifCMirror) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesEthernetVifSVifCMirror) UnmarshalJSON(_ []byte) error {
	return nil
}
