// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesBondingMember{}

// InterfacesBondingMember describes the resource data model.
type InterfacesBondingMember struct {
	// LeafNodes
	LeafInterfacesBondingMemberInterface types.List `tfsdk:"interface" vyos:"interface,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBondingMember) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"interface": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Member interface name

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
			Description: `Member interface name

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
		},

		// Nodes

	}
}
