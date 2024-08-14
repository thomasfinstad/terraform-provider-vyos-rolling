// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixOutputFilterRuleGreFlagsChecksum{}

// FirewallIPvsixOutputFilterRuleGreFlagsChecksum describes the resource data model.
type FirewallIPvsixOutputFilterRuleGreFlagsChecksum struct {
	// LeafNodes
	LeafFirewallIPvsixOutputFilterRuleGreFlagsChecksumUnset types.Bool `tfsdk:"unset" vyos:"unset,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixOutputFilterRuleGreFlagsChecksum) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"unset": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Header does not include optional checksum

`,
			Description: `Header does not include optional checksum

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
