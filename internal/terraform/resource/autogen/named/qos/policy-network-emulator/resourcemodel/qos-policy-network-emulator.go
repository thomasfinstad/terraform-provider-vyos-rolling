/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (network-emulator) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &QosPolicyNetworkEmulator{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (network-emulator) */
// QosPolicyNetworkEmulatorSelfIdentifier used by TagNodes to keep the id info
type QosPolicyNetworkEmulatorSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (network-emulator) */

	QosPolicyNetworkEmulator types.String `tfsdk:"network_emulator" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (policy) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (qos) */
}

// QosPolicyNetworkEmulator describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type QosPolicyNetworkEmulator struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *QosPolicyNetworkEmulatorSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafQosPolicyNetworkEmulatorDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyNetworkEmulatorBandwIDth   types.String `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`
	LeafQosPolicyNetworkEmulatorDelay       types.String `tfsdk:"delay" vyos:"delay,omitempty"`
	LeafQosPolicyNetworkEmulatorCorruption  types.String `tfsdk:"corruption" vyos:"corruption,omitempty"`
	LeafQosPolicyNetworkEmulatorDuplicate   types.String `tfsdk:"duplicate" vyos:"duplicate,omitempty"`
	LeafQosPolicyNetworkEmulatorLoss        types.String `tfsdk:"loss" vyos:"loss,omitempty"`
	LeafQosPolicyNetworkEmulatorReordering  types.String `tfsdk:"reordering" vyos:"reordering,omitempty"`
	LeafQosPolicyNetworkEmulatorQueueLimit  types.Number `tfsdk:"queue_limit" vyos:"queue-limit,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *QosPolicyNetworkEmulator) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *QosPolicyNetworkEmulator) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *QosPolicyNetworkEmulator) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyNetworkEmulator) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"network-emulator",
		o.SelfIdentifier.QosPolicyNetworkEmulator.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *QosPolicyNetworkEmulator) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (policy) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (qos) */
		"qos", // Node

		"policy", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *QosPolicyNetworkEmulator) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (policy) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (qos) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyNetworkEmulator) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"network_emulator": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Network emulator policy

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Policy name  |
`,
					Description: `Network emulator policy

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Policy name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in network_emulator, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  network_emulator, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (policy) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (qos) */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"description":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (description) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
			Description: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
		},

		"bandwidth":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (bandwidth) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Available bandwidth for this policy

    |  Format        |  Description                         |
    |----------------|--------------------------------------|
    |  <number>      |  Bits per second                     |
    |  <number>bit   |  Bits per second                     |
    |  <number>kbit  |  Kilobits per second                 |
    |  <number>mbit  |  Megabits per second                 |
    |  <number>gbit  |  Gigabits per second                 |
    |  <number>tbit  |  Terabits per second                 |
    |  <number>%%    |  Percentage of interface link speed  |
`,
			Description: `Available bandwidth for this policy

    |  Format        |  Description                         |
    |----------------|--------------------------------------|
    |  <number>      |  Bits per second                     |
    |  <number>bit   |  Bits per second                     |
    |  <number>kbit  |  Kilobits per second                 |
    |  <number>mbit  |  Megabits per second                 |
    |  <number>gbit  |  Gigabits per second                 |
    |  <number>tbit  |  Terabits per second                 |
    |  <number>%%    |  Percentage of interface link speed  |
`,
		},

		"delay":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (delay) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Adds delay to packets outgoing to chosen network interface

    |  Format    |  Description           |
    |------------|------------------------|
    |  <number>  |  Time in milliseconds  |
`,
			Description: `Adds delay to packets outgoing to chosen network interface

    |  Format    |  Description           |
    |------------|------------------------|
    |  <number>  |  Time in milliseconds  |
`,
		},

		"corruption":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (corruption) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Introducing error in a random position for chosen percent of packets

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  <number>  |  Percentage of packets affected  |
`,
			Description: `Introducing error in a random position for chosen percent of packets

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  <number>  |  Percentage of packets affected  |
`,
		},

		"duplicate":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (duplicate) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Cosen percent of packets is duplicated before queuing them

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  <number>  |  Percentage of packets affected  |
`,
			Description: `Cosen percent of packets is duplicated before queuing them

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  <number>  |  Percentage of packets affected  |
`,
		},

		"loss":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (loss) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Add independent loss probability to the packets outgoing to chosen network interface

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  <number>  |  Percentage of packets affected  |
`,
			Description: `Add independent loss probability to the packets outgoing to chosen network interface

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  <number>  |  Percentage of packets affected  |
`,
		},

		"reordering":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (reordering) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Emulated packet reordering percentage

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  <number>  |  Percentage of packets affected  |
`,
			Description: `Emulated packet reordering percentage

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  <number>  |  Percentage of packets affected  |
`,
		},

		"queue_limit":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (queue-limit) */
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

		// TagNodes

		// Nodes

	}
}
