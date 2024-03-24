// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &QosPolicyPriorityQueueClassMatch{}

// QosPolicyPriorityQueueClassMatch describes the resource data model.
type QosPolicyPriorityQueueClassMatch struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"match_id" vyos:"-,self-id"`

	ParentIDQosPolicyPriorityQueue types.String `tfsdk:"priority_queue_id" vyos:"priority-queue,parent-id"`

	ParentIDQosPolicyPriorityQueueClass types.Number `tfsdk:"class_id" vyos:"class,parent-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafQosPolicyPriorityQueueClassMatchDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyPriorityQueueClassMatchInterface   types.String `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafQosPolicyPriorityQueueClassMatchMark        types.Number `tfsdk:"mark" vyos:"mark,omitempty"`
	LeafQosPolicyPriorityQueueClassMatchVif         types.Number `tfsdk:"vif" vyos:"vif,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeQosPolicyPriorityQueueClassMatchEther  *QosPolicyPriorityQueueClassMatchEther  `tfsdk:"ether" vyos:"ether,omitempty"`
	NodeQosPolicyPriorityQueueClassMatchIP     *QosPolicyPriorityQueueClassMatchIP     `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeQosPolicyPriorityQueueClassMatchIPvsix *QosPolicyPriorityQueueClassMatchIPvsix `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
}

// SetID configures the resource ID
func (o *QosPolicyPriorityQueueClassMatch) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *QosPolicyPriorityQueueClassMatch) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *QosPolicyPriorityQueueClassMatch) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyPriorityQueueClassMatch) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"match",
		o.SelfIdentifier.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *QosPolicyPriorityQueueClassMatch) GetVyosParentPath() []string {
	return []string{
		"qos",

		"policy",

		"priority-queue",
		o.ParentIDQosPolicyPriorityQueue.ValueString(),

		"class",
		o.ParentIDQosPolicyPriorityQueueClass.ValueBigFloat().String(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *QosPolicyPriorityQueueClassMatch) GetVyosNamedParentPath() []string {
	return []string{
		"qos",

		"policy",

		"priority-queue",
		o.ParentIDQosPolicyPriorityQueue.ValueString(),

		"class",
		o.ParentIDQosPolicyPriorityQueueClass.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyPriorityQueueClassMatch) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"match_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Class matching rule name

`,
			Description: `Class matching rule name

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in match_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illegal character in  match_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"priority_queue_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Priority queuing based policy

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Policy name  |
`,
			Description: `Priority queuing based policy

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
							"double underscores in priority_queue_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illegal character in  priority_queue_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"class_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `Class Handle

    |  Format  |  Description  |
    |----------|---------------|
    |  1-7     |  Priority     |
`,
			Description: `Class Handle

    |  Format  |  Description  |
    |----------|---------------|
    |  1-7     |  Priority     |
`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
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

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface to use

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
			Description: `Interface to use

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
		},

		"mark": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Match on mark applied by firewall

    |  Format  |  Description       |
    |----------|--------------------|
    |  u32     |  FW mark to match  |
`,
			Description: `Match on mark applied by firewall

    |  Format  |  Description       |
    |----------|--------------------|
    |  u32     |  FW mark to match  |
`,
		},

		"vif": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Virtual Local Area Network (VLAN) ID for this match

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  0-4095  |  Virtual Local Area Network (VLAN) tag   |
`,
			Description: `Virtual Local Area Network (VLAN) ID for this match

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  0-4095  |  Virtual Local Area Network (VLAN) tag   |
`,
		},

		// Nodes

		"ether": schema.SingleNestedAttribute{
			Attributes: QosPolicyPriorityQueueClassMatchEther{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Ethernet header match

`,
			Description: `Ethernet header match

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: QosPolicyPriorityQueueClassMatchIP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Match IP protocol header

`,
			Description: `Match IP protocol header

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: QosPolicyPriorityQueueClassMatchIPvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Match IPv6 protocol header

`,
			Description: `Match IPv6 protocol header

`,
		},
	}
}
