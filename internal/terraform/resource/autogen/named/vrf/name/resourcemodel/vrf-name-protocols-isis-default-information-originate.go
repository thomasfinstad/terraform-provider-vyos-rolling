// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// VrfNameProtocolsIsisDefaultInformationOriginate describes the resource data model.
type VrfNameProtocolsIsisDefaultInformationOriginate struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvfour *VrfNameProtocolsIsisDefaultInformationOriginateIPvfour `tfsdk:"ipv4" vyos:"ipv4,omitempty"`
	NodeVrfNameProtocolsIsisDefaultInformationOriginateIPvsix  *VrfNameProtocolsIsisDefaultInformationOriginateIPvsix  `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisDefaultInformationOriginate) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"ipv4": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisDefaultInformationOriginateIPvfour{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Distribute default route for IPv4

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisDefaultInformationOriginateIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Distribute default route for IPv6

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *VrfNameProtocolsIsisDefaultInformationOriginate) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *VrfNameProtocolsIsisDefaultInformationOriginate) UnmarshalJSON(_ []byte) error {
	return nil
}
