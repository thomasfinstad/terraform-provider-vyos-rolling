// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceLldpInterfaceLocation describes the resource data model.
type ServiceLldpInterfaceLocation struct {
	// LeafNodes
	LeafServiceLldpInterfaceLocationElin types.Number `tfsdk:"elin" vyos:"elin,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeServiceLldpInterfaceLocationCoordinateBased *ServiceLldpInterfaceLocationCoordinateBased `tfsdk:"coordinate_based" vyos:"coordinate-based,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceLldpInterfaceLocation) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"elin": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `ECS ELIN (Emergency location identifier number)

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0-9999999999  |  Emergency Call Service ELIN number (between 10-25 numbers)  |

`,
		},

		// Nodes

		"coordinate_based": schema.SingleNestedAttribute{
			Attributes: ServiceLldpInterfaceLocationCoordinateBased{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Coordinate based location

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceLldpInterfaceLocation) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceLldpInterfaceLocation) UnmarshalJSON(_ []byte) error {
	return nil
}
