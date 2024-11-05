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

var _ helpers.VyosResourceDataModel = &QosPolicyPriorityQueueDefault{}

// QosPolicyPriorityQueueDefault describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type QosPolicyPriorityQueueDefault struct {
	// LeafNodes
	LeafQosPolicyPriorityQueueDefaultCodelQuantum types.Number `tfsdk:"codel_quantum" vyos:"codel-quantum,omitempty"`
	LeafQosPolicyPriorityQueueDefaultFlows        types.Number `tfsdk:"flows" vyos:"flows,omitempty"`
	LeafQosPolicyPriorityQueueDefaultInterval     types.Number `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafQosPolicyPriorityQueueDefaultQueueLimit   types.Number `tfsdk:"queue_limit" vyos:"queue-limit,omitempty"`
	LeafQosPolicyPriorityQueueDefaultQueueType    types.String `tfsdk:"queue_type" vyos:"queue-type,omitempty"`
	LeafQosPolicyPriorityQueueDefaultTarget       types.Number `tfsdk:"target" vyos:"target,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyPriorityQueueDefault) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"codel_quantum":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Deficit in the fair queuing algorithm

    |  Format     |  Description                        |
    |-------------|-------------------------------------|
    |  0-1048576  |  Number of bytes used as 'deficit'  |
`,
			Description: `Deficit in the fair queuing algorithm

    |  Format     |  Description                        |
    |-------------|-------------------------------------|
    |  0-1048576  |  Number of bytes used as 'deficit'  |
`,

			// Default:          stringdefault.StaticString(`1514`),
			Computed: true,
		},

		"flows":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of flows into which the incoming packets are classified

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65536  |  Number of flows  |
`,
			Description: `Number of flows into which the incoming packets are classified

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65536  |  Number of flows  |
`,

			// Default:          stringdefault.StaticString(`1024`),
			Computed: true,
		},

		"interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval used to measure the delay

    |  Format  |  Description               |
    |----------|----------------------------|
    |  u32     |  Interval in milliseconds  |
`,
			Description: `Interval used to measure the delay

    |  Format  |  Description               |
    |----------|----------------------------|
    |  u32     |  Interval in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		"queue_limit":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum queue size

    |  Format        |  Description            |
    |----------------|-------------------------|
    |  1-4294967295  |  Queue size in packets  |
`,
			Description: `Maximum queue size

    |  Format        |  Description            |
    |----------------|-------------------------|
    |  1-4294967295  |  Queue size in packets  |
`,
		},

		"queue_type":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Queue type for default traffic

    |  Format         |  Description                   |
    |-----------------|--------------------------------|
    |  drop-tail      |  First-In-First-Out (FIFO)     |
    |  fair-queue     |  Stochastic Fair Queue (SFQ)   |
    |  fq-codel       |  Fair Queue Codel              |
    |  priority       |  Priority queuing              |
    |  random-detect  |  Random Early Detection (RED)  |
`,
			Description: `Queue type for default traffic

    |  Format         |  Description                   |
    |-----------------|--------------------------------|
    |  drop-tail      |  First-In-First-Out (FIFO)     |
    |  fair-queue     |  Stochastic Fair Queue (SFQ)   |
    |  fq-codel       |  Fair Queue Codel              |
    |  priority       |  Priority queuing              |
    |  random-detect  |  Random Early Detection (RED)  |
`,

			// Default:          stringdefault.StaticString(`drop-tail`),
			Computed: true,
		},

		"target":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Acceptable minimum standing/persistent queue delay

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  u32     |  Queue delay in milliseconds  |
`,
			Description: `Acceptable minimum standing/persistent queue delay

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  u32     |  Queue delay in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
