// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// QosPolicyRoundRobinClassMatch describes the resource data model.
type QosPolicyRoundRobinClassMatch struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"match_id" vyos:"-,self-id"`

	ParentIDQosPolicyRoundRobin types.String `tfsdk:"round_robin_id" vyos:"round-robin,parent-id"`

	ParentIDQosPolicyRoundRobinClass types.String `tfsdk:"class_id" vyos:"class,parent-id"`

	// LeafNodes
	LeafQosPolicyRoundRobinClassMatchDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyRoundRobinClassMatchInterface   types.String `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafQosPolicyRoundRobinClassMatchMark        types.Number `tfsdk:"mark" vyos:"mark,omitempty"`
	LeafQosPolicyRoundRobinClassMatchVif         types.Number `tfsdk:"vif" vyos:"vif,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeQosPolicyRoundRobinClassMatchEther  *QosPolicyRoundRobinClassMatchEther  `tfsdk:"ether" vyos:"ether,omitempty"`
	NodeQosPolicyRoundRobinClassMatchIP     *QosPolicyRoundRobinClassMatchIP     `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeQosPolicyRoundRobinClassMatchIPvsix *QosPolicyRoundRobinClassMatchIPvsix `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
}

// SetID configures the resource ID
func (o *QosPolicyRoundRobinClassMatch) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyRoundRobinClassMatch) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"qos",

		"policy",

		"round-robin",
		o.ParentIDQosPolicyRoundRobin.ValueString(),

		"class",
		o.ParentIDQosPolicyRoundRobinClass.ValueString(),

		"match",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyRoundRobinClassMatch) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"match_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Class matching rule name

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
						"illigal character in  match_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"round_robin_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Deficit Round Robin Scheduler

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Policy name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
			Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in round_robin_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  round_robin_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"class_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Class ID

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4095  &emsp; |  Class Identifier  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
			Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in class_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  class_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |

`,
		},

		"mark": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Match on mark applied by firewall

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  FW mark to match  |

`,
		},

		"vif": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Virtual Local Area Network (VLAN) ID for this match

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-4095  &emsp; |  Virtual Local Area Network (VLAN) tag   |

`,
		},

		// Nodes

		"ether": schema.SingleNestedAttribute{
			Attributes: QosPolicyRoundRobinClassMatchEther{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Ethernet header match

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: QosPolicyRoundRobinClassMatchIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match IP protocol header

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: QosPolicyRoundRobinClassMatchIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match IPv6 protocol header

`,
		},
	}
}
