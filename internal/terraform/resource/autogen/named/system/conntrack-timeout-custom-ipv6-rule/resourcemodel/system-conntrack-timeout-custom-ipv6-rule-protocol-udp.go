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

var _ helpers.VyosResourceDataModel = &SystemConntrackTimeoutCustomIPvsixRuleProtocolUDP{}

// SystemConntrackTimeoutCustomIPvsixRuleProtocolUDP describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type SystemConntrackTimeoutCustomIPvsixRuleProtocolUDP struct {
	// LeafNodes
	LeafSystemConntrackTimeoutCustomIPvsixRuleProtocolUDPReplied   types.Number `tfsdk:"replied" vyos:"replied,omitempty"`
	LeafSystemConntrackTimeoutCustomIPvsixRuleProtocolUDPUnreplied types.Number `tfsdk:"unreplied" vyos:"unreplied,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemConntrackTimeoutCustomIPvsixRuleProtocolUDP) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"replied":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Timeout for UDP connection seen in both directions

    |  Format      |  Description                                         |
    |--------------|------------------------------------------------------|
    |  1-21474836  |  Timeout for UDP connection seen in both directions  |
`,
			Description: `Timeout for UDP connection seen in both directions

    |  Format      |  Description                                         |
    |--------------|------------------------------------------------------|
    |  1-21474836  |  Timeout for UDP connection seen in both directions  |
`,
		},

		"unreplied":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Timeout for unreplied UDP

    |  Format      |  Description                |
    |--------------|-----------------------------|
    |  1-21474836  |  Timeout for unreplied UDP  |
`,
			Description: `Timeout for unreplied UDP

    |  Format      |  Description                |
    |--------------|-----------------------------|
    |  1-21474836  |  Timeout for unreplied UDP  |
`,
		},

		// TagNodes

		// Nodes

	}
}
