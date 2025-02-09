/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (vrrp) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceConntrackSyncFailoverMechanismVrrp{}

// ServiceConntrackSyncFailoverMechanismVrrp describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ServiceConntrackSyncFailoverMechanismVrrp struct {
	// LeafNodes
	LeafServiceConntrackSyncFailoverMechanismVrrpSyncGroup types.String `tfsdk:"sync_group" vyos:"sync-group,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceConntrackSyncFailoverMechanismVrrp) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"sync_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (sync-group) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRRP sync group

`,
			Description: `VRRP sync group

`,
		},

		// TagNodes

		// Nodes

	}
}
