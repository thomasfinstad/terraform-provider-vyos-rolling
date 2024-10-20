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

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisFastRerouteLfaLocalLoadSharing{}

// VrfNameProtocolsIsisFastRerouteLfaLocalLoadSharing describes the resource data model.
type VrfNameProtocolsIsisFastRerouteLfaLocalLoadSharing struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeVrfNameProtocolsIsisFastRerouteLfaLocalLoadSharingDisable *VrfNameProtocolsIsisFastRerouteLfaLocalLoadSharingDisable `tfsdk:"disable" vyos:"disable,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisFastRerouteLfaLocalLoadSharing) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"disable": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisFastRerouteLfaLocalLoadSharingDisable{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Disable load sharing

`,
			Description: `Disable load sharing

`,
		},
	}
}
