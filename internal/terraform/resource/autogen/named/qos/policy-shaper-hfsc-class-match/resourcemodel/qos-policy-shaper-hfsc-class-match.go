// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"regexp"
	"strings"

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

var _ helpers.VyosTopResourceDataModel = &QosPolicyShaperHfscClassMatch{}

// QosPolicyShaperHfscClassMatch describes the resource data model.
type QosPolicyShaperHfscClassMatch struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"match_id" vyos:"-,self-id"`

	ParentIDQosPolicyShaperHfsc types.String `tfsdk:"shaper_hfsc_id" vyos:"shaper-hfsc,parent-id"`

	ParentIDQosPolicyShaperHfscClass types.Number `tfsdk:"class_id" vyos:"class,parent-id"`

	// LeafNodes
	LeafQosPolicyShaperHfscClassMatchDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafQosPolicyShaperHfscClassMatchInterface   types.String `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafQosPolicyShaperHfscClassMatchMark        types.Number `tfsdk:"mark" vyos:"mark,omitempty"`
	LeafQosPolicyShaperHfscClassMatchVif         types.Number `tfsdk:"vif" vyos:"vif,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeQosPolicyShaperHfscClassMatchEther  *QosPolicyShaperHfscClassMatchEther  `tfsdk:"ether" vyos:"ether,omitempty"`
	NodeQosPolicyShaperHfscClassMatchIP     *QosPolicyShaperHfscClassMatchIP     `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeQosPolicyShaperHfscClassMatchIPvsix *QosPolicyShaperHfscClassMatchIPvsix `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
}

// SetID configures the resource ID
func (o *QosPolicyShaperHfscClassMatch) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *QosPolicyShaperHfscClassMatch) GetVyosPath() []string {
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
func (o *QosPolicyShaperHfscClassMatch) GetVyosParentPath() []string {
	return []string{
		"qos",

		"policy",

		"shaper-hfsc",
		o.ParentIDQosPolicyShaperHfsc.ValueString(),

		"class",
		o.ParentIDQosPolicyShaperHfscClass.ValueBigFloat().String(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *QosPolicyShaperHfscClassMatch) GetVyosNamedParentPath() []string {
	return []string{
		"qos",

		"policy",

		"shaper-hfsc",
		o.ParentIDQosPolicyShaperHfsc.ValueString(),

		"class",
		o.ParentIDQosPolicyShaperHfscClass.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfscClassMatch) ResourceSchemaAttributes() map[string]schema.Attribute {
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

		"shaper_hfsc_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Hierarchical Fair Service Curve's policy

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Policy name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in shaper_hfsc_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  shaper_hfsc_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"class_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `Class ID

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4095  &emsp; |  Class Identifier  |

`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
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
			Attributes: QosPolicyShaperHfscClassMatchEther{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Ethernet header match

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscClassMatchIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match IP protocol header

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscClassMatchIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match IPv6 protocol header

`,
		},
	}
}
