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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (transition-script) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &HighAvailabilityVrrpSyncGroupTransitionScrIPt{}

// HighAvailabilityVrrpSyncGroupTransitionScrIPt describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type HighAvailabilityVrrpSyncGroupTransitionScrIPt struct {
	// LeafNodes
	LeafHighAvailabilityVrrpSyncGroupTransitionScrIPtMaster types.String `tfsdk:"master" vyos:"master,omitempty"`
	LeafHighAvailabilityVrrpSyncGroupTransitionScrIPtBackup types.String `tfsdk:"backup" vyos:"backup,omitempty"`
	LeafHighAvailabilityVrrpSyncGroupTransitionScrIPtFault  types.String `tfsdk:"fault" vyos:"fault,omitempty"`
	LeafHighAvailabilityVrrpSyncGroupTransitionScrIPtStop   types.String `tfsdk:"stop" vyos:"stop,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpSyncGroupTransitionScrIPt) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"master":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (master) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Script to run on VRRP state transition to master

`,
			Description: `Script to run on VRRP state transition to master

`,
		},

		"backup":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (backup) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Script to run on VRRP state transition to backup

`,
			Description: `Script to run on VRRP state transition to backup

`,
		},

		"fault":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (fault) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Script to run on VRRP state transition to fault

`,
			Description: `Script to run on VRRP state transition to fault

`,
		},

		"stop":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (stop) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Script to run on VRRP state transition to stop

`,
			Description: `Script to run on VRRP state transition to stop

`,
		},

		// TagNodes

		// Nodes

	}
}
