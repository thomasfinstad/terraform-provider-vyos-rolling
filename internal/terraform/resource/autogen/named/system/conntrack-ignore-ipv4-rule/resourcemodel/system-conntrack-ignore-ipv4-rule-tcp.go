/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &SystemConntrackIgnoreIPvfourRuleTCP{}

// SystemConntrackIgnoreIPvfourRuleTCP describes the resource data model.
type SystemConntrackIgnoreIPvfourRuleTCP struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeSystemConntrackIgnoreIPvfourRuleTCPFlags *SystemConntrackIgnoreIPvfourRuleTCPFlags `tfsdk:"flags" vyos:"flags,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemConntrackIgnoreIPvfourRuleTCP) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"flags": schema.SingleNestedAttribute{
			Attributes: SystemConntrackIgnoreIPvfourRuleTCPFlags{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `TCP flags to match

`,
			Description: `TCP flags to match

`,
		},
	}
}
