// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
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

var _ helpers.VyosTopResourceDataModel = &ServiceSURIcataPortGroup{}

// ServiceSURIcataPortGroup describes the resource data model.
type ServiceSURIcataPortGroup struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceSURIcataPortGroupPort  types.List `tfsdk:"port" vyos:"port,omitempty"`
	LeafServiceSURIcataPortGroupGroup types.List `tfsdk:"group" vyos:"group,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceSURIcataPortGroup) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceSURIcataPortGroup) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceSURIcataPortGroup) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceSURIcataPortGroup) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"port-group",
		o.SelfIdentifier.Attributes()["port_group"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceSURIcataPortGroup) GetVyosParentPath() []string {
	return []string{
		"service",

		"suricata",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceSURIcataPortGroup) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSURIcataPortGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"port_group": schema.StringAttribute{
						Required: true,
						MarkdownDescription: `Port group name

`,
						Description: `Port group name

`,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						}, Validators: []validator.String{
							stringvalidator.All(
								helpers.StringNot(
									stringvalidator.RegexMatches(
										regexp.MustCompile(`^.*__.*$`),
										"double underscores in port_group, conflicts with the internal resource id",
									),
								),
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  port_group, value must match: ^[a-zA-Z0-9-_]*$",
								),
							),
						},
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"port": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Port number

    |  Format      |  Description                                                    |
    |--------------|-----------------------------------------------------------------|
    |  1-65535     |  Numeric port to match                                          |
    |  !1-65535    |  Numeric port to exclude from matches                           |
    |  start-end   |  Numbered port range (e.g. 1001-1005) to match                  |
    |  !start-end  |  Numbered port range (e.g. !1001-1005) to exclude from matches  |
`,
			Description: `Port number

    |  Format      |  Description                                                    |
    |--------------|-----------------------------------------------------------------|
    |  1-65535     |  Numeric port to match                                          |
    |  !1-65535    |  Numeric port to exclude from matches                           |
    |  start-end   |  Numbered port range (e.g. 1001-1005) to match                  |
    |  !start-end  |  Numbered port range (e.g. !1001-1005) to exclude from matches  |
`,
		},

		"group": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Port group

    |  Format  |  Description                                    |
    |----------|-------------------------------------------------|
    |  txt     |  Port group to match                            |
    |  !txt    |  Exclude the specified port group from matches  |
`,
			Description: `Port group

    |  Format  |  Description                                    |
    |----------|-------------------------------------------------|
    |  txt     |  Port group to match                            |
    |  !txt    |  Exclude the specified port group from matches  |
`,
		},

		// Nodes

	}
}