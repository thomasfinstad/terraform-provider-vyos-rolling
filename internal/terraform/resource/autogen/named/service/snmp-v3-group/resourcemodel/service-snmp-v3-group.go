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

var _ helpers.VyosTopResourceDataModel = &ServiceSnmpVthreeGroup{}

// ServiceSnmpVthreeGroup describes the resource data model.
type ServiceSnmpVthreeGroup struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceSnmpVthreeGroupMode     types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafServiceSnmpVthreeGroupSeclevel types.String `tfsdk:"seclevel" vyos:"seclevel,omitempty"`
	LeafServiceSnmpVthreeGroupView     types.String `tfsdk:"view" vyos:"view,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceSnmpVthreeGroup) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceSnmpVthreeGroup) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceSnmpVthreeGroup) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceSnmpVthreeGroup) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"group",
		o.SelfIdentifier.Attributes()["group"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceSnmpVthreeGroup) GetVyosParentPath() []string {
	return []string{
		"service",

		"snmp",

		"v3",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceSnmpVthreeGroup) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSnmpVthreeGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"group": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Specifies the group with name groupname

`,
					Description: `Specifies the group with name groupname

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in group, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  group, value must match: ^[.:a-zA-Z0-9-_]+$",
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

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Define access permission

    |  Format  |  Description  |
    |----------|---------------|
    |  ro      |  Read-Only    |
    |  rw      |  read write   |
`,
			Description: `Define access permission

    |  Format  |  Description  |
    |----------|---------------|
    |  ro      |  Read-Only    |
    |  rw      |  read write   |
`,

			// Default:          stringdefault.StaticString(`ro`),
			Computed: true,
		},

		"seclevel": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Security levels

    |  Format  |  Description                                                  |
    |----------|---------------------------------------------------------------|
    |  noauth  |  Messages not authenticated and not encrypted (noAuthNoPriv)  |
    |  auth    |  Messages are authenticated but not encrypted (authNoPriv)    |
    |  priv    |  Messages are authenticated and encrypted (authPriv)          |
`,
			Description: `Security levels

    |  Format  |  Description                                                  |
    |----------|---------------------------------------------------------------|
    |  noauth  |  Messages not authenticated and not encrypted (noAuthNoPriv)  |
    |  auth    |  Messages are authenticated but not encrypted (authNoPriv)    |
    |  priv    |  Messages are authenticated and encrypted (authPriv)          |
`,

			// Default:          stringdefault.StaticString(`auth`),
			Computed: true,
		},

		"view": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Defines the name of view

`,
			Description: `Defines the name of view

`,
		},

		// Nodes

	}
}
