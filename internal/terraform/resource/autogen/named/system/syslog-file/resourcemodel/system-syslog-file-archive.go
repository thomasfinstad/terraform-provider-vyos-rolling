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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (archive) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &SystemSyslogFileArchive{}

// SystemSyslogFileArchive describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type SystemSyslogFileArchive struct {
	// LeafNodes
	LeafSystemSyslogFileArchiveFile types.String `tfsdk:"file" vyos:"file,omitempty"`
	LeafSystemSyslogFileArchiveSize types.String `tfsdk:"size" vyos:"size,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemSyslogFileArchive) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"file":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (file) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Number of saved files

`,
			Description: `Number of saved files

`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		"size":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (size) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Size of log files in kbytes

`,
			Description: `Size of log files in kbytes

`,

			// Default:          stringdefault.StaticString(`256`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
