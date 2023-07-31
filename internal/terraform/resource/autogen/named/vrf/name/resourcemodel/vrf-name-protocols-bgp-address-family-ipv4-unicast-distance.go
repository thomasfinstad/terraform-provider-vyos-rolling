// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceExternal types.Number `tfsdk:"external" vyos:"external,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceInternal types.Number `tfsdk:"internal" vyos:"internal,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistanceLocal    types.Number `tfsdk:"local" vyos:"local,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagVrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistancePrefix bool `tfsdk:"prefix" vyos:"prefix,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"external": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `eBGP routes administrative distance

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-255  |  eBGP routes administrative distance  |

`,
		},

		"internal": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `iBGP routes administrative distance

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-255  |  iBGP routes administrative distance  |

`,
		},

		"local": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Locally originated BGP routes administrative distance

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:1-255  |  Locally originated BGP routes administrative distance  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistance) UnmarshalJSON(_ []byte) error {
	return nil
}
