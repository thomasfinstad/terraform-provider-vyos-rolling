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

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &FirewallZoneFrom{}

// FirewallZoneFrom describes the resource data model.
type FirewallZoneFrom struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"from_id" vyos:"-,self-id"`

	ParentIDFirewallZone types.String `tfsdk:"zone_id" vyos:"zone,parent-id"`

	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeFirewallZoneFromFirewall *FirewallZoneFromFirewall `tfsdk:"firewall" vyos:"firewall,omitempty"`
}

// SetID configures the resource ID
func (o *FirewallZoneFrom) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallZoneFrom) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"firewall",

		"zone",
		o.ParentIDFirewallZone.ValueString(),

		"from",
		o.SelfIdentifier.ValueString(),
	}
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *FirewallZoneFrom) GetVyosParentPath() []string {
	return []string{
		"firewall",

		"zone",
		o.ParentIDFirewallZone.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallZoneFrom) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"from_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Zone from which to filter traffic

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in from_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  from_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"zone_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Zone-policy

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Zone name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in zone_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  zone_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		// LeafNodes

		// Nodes

		"firewall": schema.SingleNestedAttribute{
			Attributes: FirewallZoneFromFirewall{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Firewall options

`,
		},
	}
}
