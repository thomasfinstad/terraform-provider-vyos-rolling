// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

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

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &QosPolicyCake{}

// QosPolicyCake describes the resource data model.
type QosPolicyCake struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafQosPolicyCakeDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyCakeBandwIDth   types.String `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`
	LeafQosPolicyCakeRtt         types.Number `tfsdk:"rtt" vyos:"rtt,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeQosPolicyCakeFlowIsolation *QosPolicyCakeFlowIsolation `tfsdk:"flow_isolation" vyos:"flow-isolation,omitempty"`
}

// SetID configures the resource ID
func (o *QosPolicyCake) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *QosPolicyCake) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *QosPolicyCake) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyCake) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"cake",
		o.SelfIdentifier.Attributes()["cake"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *QosPolicyCake) GetVyosParentPath() []string {
	return []string{
		"qos",

		"policy",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *QosPolicyCake) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyCake) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"cake": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Common Applications Kept Enhanced (CAKE)

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Policy name  |
`,
					Description: `Common Applications Kept Enhanced (CAKE)

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
									"double underscores in cake, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  cake, value must match: ^[a-zA-Z0-9-_]*$",
							),
						),
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"description": schema.StringAttribute{
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

		"bandwidth": schema.StringAttribute{
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

		"rtt": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Round-Trip-Time for Active Queue Management (AQM)

    |  Format     |  Description  |
    |-------------|---------------|
    |  1-3600000  |  RTT in ms    |
`,
			Description: `Round-Trip-Time for Active Queue Management (AQM)

    |  Format     |  Description  |
    |-------------|---------------|
    |  1-3600000  |  RTT in ms    |
`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		// Nodes

		"flow_isolation": schema.SingleNestedAttribute{
			Attributes: QosPolicyCakeFlowIsolation{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Flow isolation settings

`,
			Description: `Flow isolation settings

`,
		},
	}
}
