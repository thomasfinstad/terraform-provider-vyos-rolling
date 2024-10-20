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

var _ helpers.VyosResourceDataModel = &QosPolicyShaperClassMatchIP{}

// QosPolicyShaperClassMatchIP describes the resource data model.
type QosPolicyShaperClassMatchIP struct {
	// LeafNodes
	LeafQosPolicyShaperClassMatchIPDscp      types.String `tfsdk:"dscp" vyos:"dscp,omitempty"`
	LeafQosPolicyShaperClassMatchIPMaxLength types.Number `tfsdk:"max_length" vyos:"max-length,omitempty"`
	LeafQosPolicyShaperClassMatchIPProtocol  types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeQosPolicyShaperClassMatchIPDestination *QosPolicyShaperClassMatchIPDestination `tfsdk:"destination" vyos:"destination,omitempty"`
	NodeQosPolicyShaperClassMatchIPSource      *QosPolicyShaperClassMatchIPSource      `tfsdk:"source" vyos:"source,omitempty"`
	NodeQosPolicyShaperClassMatchIPTCP         *QosPolicyShaperClassMatchIPTCP         `tfsdk:"tcp" vyos:"tcp,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperClassMatchIP) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"dscp":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match on Differentiated Services Codepoint (DSCP)

    |  Format          |  Description                                      |
    |------------------|---------------------------------------------------|
    |  0-63            |  Differentiated Services Codepoint (DSCP) value   |
    |  default         |  match DSCP (000000)                              |
    |  reliability     |  match DSCP (000001)                              |
    |  throughput      |  match DSCP (000010)                              |
    |  lowdelay        |  match DSCP (000100)                              |
    |  priority        |  match DSCP (001000)                              |
    |  immediate       |  match DSCP (010000)                              |
    |  flash           |  match DSCP (011000)                              |
    |  flash-override  |  match DSCP (100000)                              |
    |  critical        |  match DSCP (101000)                              |
    |  internet        |  match DSCP (110000)                              |
    |  network         |  match DSCP (111000)                              |
    |  AF11            |  High-throughput data                             |
    |  AF12            |  High-throughput data                             |
    |  AF13            |  High-throughput data                             |
    |  AF21            |  Low-latency data                                 |
    |  AF22            |  Low-latency data                                 |
    |  AF23            |  Low-latency data                                 |
    |  AF31            |  Multimedia streaming                             |
    |  AF32            |  Multimedia streaming                             |
    |  AF33            |  Multimedia streaming                             |
    |  AF41            |  Multimedia conferencing                          |
    |  AF42            |  Multimedia conferencing                          |
    |  AF43            |  Multimedia conferencing                          |
    |  CS1             |  Low-priority data                                |
    |  CS2             |  OAM                                              |
    |  CS3             |  Broadcast video                                  |
    |  CS4             |  Real-time interactive                            |
    |  CS5             |  Signaling                                        |
    |  CS6             |  Network control                                  |
    |  CS7             |  N/A                                              |
    |  EF              |  Expedited Forwarding                             |
`,
			Description: `Match on Differentiated Services Codepoint (DSCP)

    |  Format          |  Description                                      |
    |------------------|---------------------------------------------------|
    |  0-63            |  Differentiated Services Codepoint (DSCP) value   |
    |  default         |  match DSCP (000000)                              |
    |  reliability     |  match DSCP (000001)                              |
    |  throughput      |  match DSCP (000010)                              |
    |  lowdelay        |  match DSCP (000100)                              |
    |  priority        |  match DSCP (001000)                              |
    |  immediate       |  match DSCP (010000)                              |
    |  flash           |  match DSCP (011000)                              |
    |  flash-override  |  match DSCP (100000)                              |
    |  critical        |  match DSCP (101000)                              |
    |  internet        |  match DSCP (110000)                              |
    |  network         |  match DSCP (111000)                              |
    |  AF11            |  High-throughput data                             |
    |  AF12            |  High-throughput data                             |
    |  AF13            |  High-throughput data                             |
    |  AF21            |  Low-latency data                                 |
    |  AF22            |  Low-latency data                                 |
    |  AF23            |  Low-latency data                                 |
    |  AF31            |  Multimedia streaming                             |
    |  AF32            |  Multimedia streaming                             |
    |  AF33            |  Multimedia streaming                             |
    |  AF41            |  Multimedia conferencing                          |
    |  AF42            |  Multimedia conferencing                          |
    |  AF43            |  Multimedia conferencing                          |
    |  CS1             |  Low-priority data                                |
    |  CS2             |  OAM                                              |
    |  CS3             |  Broadcast video                                  |
    |  CS4             |  Real-time interactive                            |
    |  CS5             |  Signaling                                        |
    |  CS6             |  Network control                                  |
    |  CS7             |  N/A                                              |
    |  EF              |  Expedited Forwarding                             |
`,
		},

		"max_length":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum packet length

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  1-65535  |  Maximum packet/payload length  |
`,
			Description: `Maximum packet length

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  1-65535  |  Maximum packet/payload length  |
`,
		},

		"protocol":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol

    |  Format  |  Description    |
    |----------|-----------------|
    |  txt     |  Protocol name  |
`,
			Description: `Protocol

    |  Format  |  Description    |
    |----------|-----------------|
    |  txt     |  Protocol name  |
`,
		},

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperClassMatchIPDestination{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Match on destination port or address

`,
			Description: `Match on destination port or address

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperClassMatchIPSource{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Match on source port or address

`,
			Description: `Match on source port or address

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperClassMatchIPTCP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `TCP Flags matching

`,
			Description: `TCP Flags matching

`,
		},
	}
}
