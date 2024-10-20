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

var _ helpers.VyosResourceDataModel = &SystemConntrackTimeoutCustomIPvfourRuleProtocol{}

// SystemConntrackTimeoutCustomIPvfourRuleProtocol describes the resource data model.
type SystemConntrackTimeoutCustomIPvfourRuleProtocol struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeSystemConntrackTimeoutCustomIPvfourRuleProtocolTCP *SystemConntrackTimeoutCustomIPvfourRuleProtocolTCP `tfsdk:"tcp" vyos:"tcp,omitempty"`
	NodeSystemConntrackTimeoutCustomIPvfourRuleProtocolUDP *SystemConntrackTimeoutCustomIPvfourRuleProtocolUDP `tfsdk:"udp" vyos:"udp,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemConntrackTimeoutCustomIPvfourRuleProtocol) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

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
