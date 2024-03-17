// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndex{}

// VrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndex describes the resource data model.
type VrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndex struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Number `tfsdk:"index_id" vyos:"-,self-id"`

	ParentIDVrfName types.String `tfsdk:"name_id" vyos:"name,parent-id"`

	// LeafNodes
	LeafVrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndexLevelOne types.Bool `tfsdk:"level_1" vyos:"level-1,omitempty"`
	LeafVrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndexLevelTwo types.Bool `tfsdk:"level_2" vyos:"level-2,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *VrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndex) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndex) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"index",
		o.SelfIdentifier.ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndex) GetVyosParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"isis",

		"fast-reroute",

		"lfa",

		"local",

		"tiebreaker",

		"downstream",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndex) GetVyosNamedParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndex) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"index_id": schema.NumberAttribute{
			Required: true,
			MarkdownDescription: `Set preference order among tiebreakers

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  The index integer value  |

`,
			PlanModifiers: []planmodifier.Number{
				numberplanmodifier.RequiresReplace(),
			},
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in name_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  name_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		// LeafNodes

		"level_1": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Match on IS-IS level-1 routes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"level_2": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Match on IS-IS level-2 routes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
