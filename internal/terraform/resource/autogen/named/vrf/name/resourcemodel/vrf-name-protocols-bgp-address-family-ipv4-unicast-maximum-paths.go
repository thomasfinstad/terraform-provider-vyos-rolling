// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastMaximumPaths describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastMaximumPaths struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastMaximumPathsEbgp types.Number `tfsdk:"ebgp" vyos:"ebgp,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourUnicastMaximumPathsIbgp types.Number `tfsdk:"ibgp" vyos:"ibgp,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastMaximumPaths) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ebgp": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `eBGP maximum paths

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-256  &emsp; |  Number of paths to consider  |

`,
		},

		"ibgp": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `iBGP maximum paths

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-256  &emsp; |  Number of paths to consider  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastMaximumPaths) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourUnicastMaximumPaths) UnmarshalJSON(_ []byte) error {
	return nil
}