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

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &SystemConntrackTimeoutCustomIPvsixRuleProtocolTCP{}

// SystemConntrackTimeoutCustomIPvsixRuleProtocolTCP describes the resource data model.
type SystemConntrackTimeoutCustomIPvsixRuleProtocolTCP struct {
	// LeafNodes
	LeafSystemConntrackTimeoutCustomIPvsixRuleProtocolTCPCloseWait   types.Number `tfsdk:"close_wait" vyos:"close-wait,omitempty"`
	LeafSystemConntrackTimeoutCustomIPvsixRuleProtocolTCPClose       types.Number `tfsdk:"close" vyos:"close,omitempty"`
	LeafSystemConntrackTimeoutCustomIPvsixRuleProtocolTCPEstablished types.Number `tfsdk:"established" vyos:"established,omitempty"`
	LeafSystemConntrackTimeoutCustomIPvsixRuleProtocolTCPFinWait     types.Number `tfsdk:"fin_wait" vyos:"fin-wait,omitempty"`
	LeafSystemConntrackTimeoutCustomIPvsixRuleProtocolTCPLastAck     types.Number `tfsdk:"last_ack" vyos:"last-ack,omitempty"`
	LeafSystemConntrackTimeoutCustomIPvsixRuleProtocolTCPSynRecv     types.Number `tfsdk:"syn_recv" vyos:"syn-recv,omitempty"`
	LeafSystemConntrackTimeoutCustomIPvsixRuleProtocolTCPSynSent     types.Number `tfsdk:"syn_sent" vyos:"syn-sent,omitempty"`
	LeafSystemConntrackTimeoutCustomIPvsixRuleProtocolTCPTimeWait    types.Number `tfsdk:"time_wait" vyos:"time-wait,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemConntrackTimeoutCustomIPvsixRuleProtocolTCP) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"close_wait":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP CLOSE-WAIT timeout in seconds

    |  Format      |  Description                        |
    |--------------|-------------------------------------|
    |  1-21474836  |  TCP CLOSE-WAIT timeout in seconds  |
`,
			Description: `TCP CLOSE-WAIT timeout in seconds

    |  Format      |  Description                        |
    |--------------|-------------------------------------|
    |  1-21474836  |  TCP CLOSE-WAIT timeout in seconds  |
`,
		},

		"close":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP CLOSE timeout in seconds

    |  Format      |  Description                   |
    |--------------|--------------------------------|
    |  1-21474836  |  TCP CLOSE timeout in seconds  |
`,
			Description: `TCP CLOSE timeout in seconds

    |  Format      |  Description                   |
    |--------------|--------------------------------|
    |  1-21474836  |  TCP CLOSE timeout in seconds  |
`,
		},

		"established":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP ESTABLISHED timeout in seconds

    |  Format      |  Description                         |
    |--------------|--------------------------------------|
    |  1-21474836  |  TCP ESTABLISHED timeout in seconds  |
`,
			Description: `TCP ESTABLISHED timeout in seconds

    |  Format      |  Description                         |
    |--------------|--------------------------------------|
    |  1-21474836  |  TCP ESTABLISHED timeout in seconds  |
`,
		},

		"fin_wait":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP FIN-WAIT timeout in seconds

    |  Format      |  Description                      |
    |--------------|-----------------------------------|
    |  1-21474836  |  TCP FIN-WAIT timeout in seconds  |
`,
			Description: `TCP FIN-WAIT timeout in seconds

    |  Format      |  Description                      |
    |--------------|-----------------------------------|
    |  1-21474836  |  TCP FIN-WAIT timeout in seconds  |
`,
		},

		"last_ack":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP LAST-ACK timeout in seconds

    |  Format      |  Description                      |
    |--------------|-----------------------------------|
    |  1-21474836  |  TCP LAST-ACK timeout in seconds  |
`,
			Description: `TCP LAST-ACK timeout in seconds

    |  Format      |  Description                      |
    |--------------|-----------------------------------|
    |  1-21474836  |  TCP LAST-ACK timeout in seconds  |
`,
		},

		"syn_recv":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP SYN-RECEIVED timeout in seconds

    |  Format      |  Description                          |
    |--------------|---------------------------------------|
    |  1-21474836  |  TCP SYN-RECEIVED timeout in seconds  |
`,
			Description: `TCP SYN-RECEIVED timeout in seconds

    |  Format      |  Description                          |
    |--------------|---------------------------------------|
    |  1-21474836  |  TCP SYN-RECEIVED timeout in seconds  |
`,
		},

		"syn_sent":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP SYN-SENT timeout in seconds

    |  Format      |  Description                      |
    |--------------|-----------------------------------|
    |  1-21474836  |  TCP SYN-SENT timeout in seconds  |
`,
			Description: `TCP SYN-SENT timeout in seconds

    |  Format      |  Description                      |
    |--------------|-----------------------------------|
    |  1-21474836  |  TCP SYN-SENT timeout in seconds  |
`,
		},

		"time_wait":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP TIME-WAIT timeout in seconds

    |  Format      |  Description                       |
    |--------------|------------------------------------|
    |  1-21474836  |  TCP TIME-WAIT timeout in seconds  |
`,
			Description: `TCP TIME-WAIT timeout in seconds

    |  Format      |  Description                       |
    |--------------|------------------------------------|
    |  1-21474836  |  TCP TIME-WAIT timeout in seconds  |
`,
		},

		// Nodes

	}
}
