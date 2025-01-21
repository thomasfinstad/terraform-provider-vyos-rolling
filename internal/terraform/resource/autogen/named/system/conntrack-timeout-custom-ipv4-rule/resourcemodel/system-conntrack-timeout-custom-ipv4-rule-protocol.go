/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (protocol) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &SystemConntrackTimeoutCustomIPvfourRuleProtocol{}

// SystemConntrackTimeoutCustomIPvfourRuleProtocol describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type SystemConntrackTimeoutCustomIPvfourRuleProtocol struct {
	// LeafNodes

	// TagNodes

	// Nodes

	NodeSystemConntrackTimeoutCustomIPvfourRuleProtocolTCP *SystemConntrackTimeoutCustomIPvfourRuleProtocolTCP `tfsdk:"tcp" vyos:"tcp,omitempty"`

	NodeSystemConntrackTimeoutCustomIPvfourRuleProtocolUDP *SystemConntrackTimeoutCustomIPvfourRuleProtocolUDP `tfsdk:"udp" vyos:"udp,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemConntrackTimeoutCustomIPvfourRuleProtocol) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// TagNodes

		// Nodes

		"tcp": schema.SingleNestedAttribute{
			Attributes: SystemConntrackTimeoutCustomIPvfourRuleProtocolTCP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `TCP connection timeout options

`,
			Description: `TCP connection timeout options

`,
		},

		"udp": schema.SingleNestedAttribute{
			Attributes: SystemConntrackTimeoutCustomIPvfourRuleProtocolUDP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `UDP timeout options

`,
			Description: `UDP timeout options

`,
		},
	}
}
