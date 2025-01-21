/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (tcp) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &QosTrafficMatchGroupMatchIPTCP{}

// QosTrafficMatchGroupMatchIPTCP describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type QosTrafficMatchGroupMatchIPTCP struct {
	// LeafNodes
	LeafQosTrafficMatchGroupMatchIPTCPAck types.Bool `tfsdk:"ack" vyos:"ack,omitempty"`
	LeafQosTrafficMatchGroupMatchIPTCPSyn types.Bool `tfsdk:"syn" vyos:"syn,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosTrafficMatchGroupMatchIPTCP) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ack":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ack) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Match TCP ACK

`,
			Description: `Match TCP ACK

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"syn":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (syn) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Match TCP SYN

`,
			Description: `Match TCP SYN

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
