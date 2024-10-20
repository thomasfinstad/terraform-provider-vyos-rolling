/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &SystemSyslogHostFormat{}

// SystemSyslogHostFormat describes the resource data model.
type SystemSyslogHostFormat struct {
	// LeafNodes
	LeafSystemSyslogHostFormatOctetCounted    types.Bool `tfsdk:"octet_counted" vyos:"octet-counted,omitempty"`
	LeafSystemSyslogHostFormatIncludeTimezone types.Bool `tfsdk:"include_timezone" vyos:"include-timezone,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemSyslogHostFormat) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"octet_counted":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Allows for the transmission of all characters inside a syslog message

`,
			Description: `Allows for the transmission of all characters inside a syslog message

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"include_timezone":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Include system timezone in syslog message

`,
			Description: `Include system timezone in syslog message

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
