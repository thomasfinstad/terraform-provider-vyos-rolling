// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAllowasIn{}

// VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAllowasIn describes the resource data model.
type VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAllowasIn struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAllowasInNumber types.Number `tfsdk:"number" vyos:"number,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpNeighborAddressFamilyLtwovpnEvpnAllowasIn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"number": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of occurrences of AS number

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  1-10    |  Number of times AS is allowed in path  |
`,
			Description: `Number of occurrences of AS number

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  1-10    |  Number of times AS is allowed in path  |
`,
		},

		// Nodes

	}
}
