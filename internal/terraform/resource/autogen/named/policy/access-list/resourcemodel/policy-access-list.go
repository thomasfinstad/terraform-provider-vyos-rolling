// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &PolicyAccessList{}

// PolicyAccessList describes the resource data model.
type PolicyAccessList struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafPolicyAccessListDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagPolicyAccessListRule bool `tfsdk:"-" vyos:"rule,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *PolicyAccessList) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *PolicyAccessList) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *PolicyAccessList) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyAccessList) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"access-list",
		o.SelfIdentifier.Attributes()["access_list"].(types.Number).ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *PolicyAccessList) GetVyosParentPath() []string {
	return []string{
		"policy",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *PolicyAccessList) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyAccessList) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"access_list": schema.NumberAttribute{
					Required: true,
					MarkdownDescription: `IP access-list filter

    |  Format     |  Description                               |
    |-------------|--------------------------------------------|
    |  1-99       |  IP standard access list                   |
    |  100-199    |  IP extended access list                   |
    |  1300-1999  |  IP standard access list (expanded range)  |
    |  2000-2699  |  IP extended access list (expanded range)  |
`,
					Description: `IP access-list filter

    |  Format     |  Description                               |
    |-------------|--------------------------------------------|
    |  1-99       |  IP standard access list                   |
    |  100-199    |  IP extended access list                   |
    |  1300-1999  |  IP standard access list (expanded range)  |
    |  2000-2699  |  IP extended access list (expanded range)  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
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

		// Nodes

	}
}
