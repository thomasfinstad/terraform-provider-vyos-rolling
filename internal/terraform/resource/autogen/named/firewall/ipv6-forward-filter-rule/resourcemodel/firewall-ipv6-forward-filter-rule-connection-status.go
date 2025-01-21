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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (connection-status) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallIPvsixForwardFilterRuleConnectionStatus{}

// FirewallIPvsixForwardFilterRuleConnectionStatus describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type FirewallIPvsixForwardFilterRuleConnectionStatus struct {
	// LeafNodes
	LeafFirewallIPvsixForwardFilterRuleConnectionStatusNat types.String `tfsdk:"nat" vyos:"nat,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallIPvsixForwardFilterRuleConnectionStatus) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"nat":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (nat) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NAT connection status

    |  Format       |  Description                                            |
    |---------------|---------------------------------------------------------|
    |  destination  |  Match connections that are subject to destination NAT  |
    |  source       |  Match connections that are subject to source NAT       |
`,
			Description: `NAT connection status

    |  Format       |  Description                                            |
    |---------------|---------------------------------------------------------|
    |  destination  |  Match connections that are subject to destination NAT  |
    |  source       |  Match connections that are subject to source NAT       |
`,
		},

		// TagNodes

		// Nodes

	}
}
