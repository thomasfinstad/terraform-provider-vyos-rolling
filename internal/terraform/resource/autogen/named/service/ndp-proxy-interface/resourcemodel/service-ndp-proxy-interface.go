// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceNdpProxyInterface{}

// ServiceNdpProxyInterface describes the resource data model.
type ServiceNdpProxyInterface struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceNdpProxyInterfaceDisable         types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafServiceNdpProxyInterfaceEnableRouterBit types.Bool   `tfsdk:"enable_router_bit" vyos:"enable-router-bit,omitempty"`
	LeafServiceNdpProxyInterfaceTimeout         types.Number `tfsdk:"timeout" vyos:"timeout,omitempty"`
	LeafServiceNdpProxyInterfaceTTL             types.Number `tfsdk:"ttl" vyos:"ttl,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagServiceNdpProxyInterfacePrefix bool `tfsdk:"-" vyos:"prefix,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceNdpProxyInterface) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceNdpProxyInterface) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceNdpProxyInterface) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceNdpProxyInterface) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"interface",
		o.SelfIdentifier.Attributes()["interface"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceNdpProxyInterface) GetVyosParentPath() []string {
	return []string{
		"service",

		"ndp-proxy",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceNdpProxyInterface) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceNdpProxyInterface) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"interface": schema.StringAttribute{
						Required: true,
						MarkdownDescription: `NDP proxy listener interface

`,
						Description: `NDP proxy listener interface

`,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						}, Validators: []validator.String{
							stringvalidator.All(
								helpers.StringNot(
									stringvalidator.RegexMatches(
										regexp.MustCompile(`^.*__.*$`),
										"double underscores in interface, conflicts with the internal resource id",
									),
								),
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  interface, value must match: ^[a-zA-Z0-9-_]*$",
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

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"enable_router_bit": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable router bit in Neighbor Advertisement messages

`,
			Description: `Enable router bit in Neighbor Advertisement messages

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"timeout": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Timeout for Neighbor Advertisement after Neighbor Solicitation message

    |  Format      |  Description              |
    |--------------|---------------------------|
    |  500-120000  |  Timeout in milliseconds  |
`,
			Description: `Timeout for Neighbor Advertisement after Neighbor Solicitation message

    |  Format      |  Description              |
    |--------------|---------------------------|
    |  500-120000  |  Timeout in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`500`),
			Computed: true,
		},

		"ttl": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Proxy entry cache Time-To-Live

    |  Format        |  Description           |
    |----------------|------------------------|
    |  10000-120000  |  Time in milliseconds  |
`,
			Description: `Proxy entry cache Time-To-Live

    |  Format        |  Description           |
    |----------------|------------------------|
    |  10000-120000  |  Time in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`30000`),
			Computed: true,
		},

		// Nodes

	}
}
