/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &HighAvailabilityVrrpGroupTransitionScrIPt{}

// HighAvailabilityVrrpGroupTransitionScrIPt describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type HighAvailabilityVrrpGroupTransitionScrIPt struct {
	// LeafNodes
	LeafHighAvailabilityVrrpGroupTransitionScrIPtMaster types.String `tfsdk:"master" vyos:"master,omitempty"`
	LeafHighAvailabilityVrrpGroupTransitionScrIPtBackup types.String `tfsdk:"backup" vyos:"backup,omitempty"`
	LeafHighAvailabilityVrrpGroupTransitionScrIPtFault  types.String `tfsdk:"fault" vyos:"fault,omitempty"`
	LeafHighAvailabilityVrrpGroupTransitionScrIPtStop   types.String `tfsdk:"stop" vyos:"stop,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpGroupTransitionScrIPt) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"master":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Script to run on VRRP state transition to master

`,
			Description: `Script to run on VRRP state transition to master

`,
		},

		"backup":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Script to run on VRRP state transition to backup

`,
			Description: `Script to run on VRRP state transition to backup

`,
		},

		"fault":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Script to run on VRRP state transition to fault

`,
			Description: `Script to run on VRRP state transition to fault

`,
		},

		"stop":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
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
