// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameIPNht{}

// VrfNameIPNht describes the resource data model.
type VrfNameIPNht struct {
	// LeafNodes
	LeafVrfNameIPNhtNoResolveViaDefault types.Bool `tfsdk:"no_resolve_via_default" vyos:"no-resolve-via-default,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameIPNht) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"no_resolve_via_default": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not resolve via default route

`,
			Description: `Do not resolve via default route

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
