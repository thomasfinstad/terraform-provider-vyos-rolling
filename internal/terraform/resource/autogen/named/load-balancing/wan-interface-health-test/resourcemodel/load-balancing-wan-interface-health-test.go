// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
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

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &LoadBalancingWanInterfaceHealthTest{}

// LoadBalancingWanInterfaceHealthTest describes the resource data model.
type LoadBalancingWanInterfaceHealthTest struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafLoadBalancingWanInterfaceHealthTestRespTime   types.Number `tfsdk:"resp_time" vyos:"resp-time,omitempty"`
	LeafLoadBalancingWanInterfaceHealthTestTarget     types.String `tfsdk:"target" vyos:"target,omitempty"`
	LeafLoadBalancingWanInterfaceHealthTestTestScrIPt types.String `tfsdk:"test_script" vyos:"test-script,omitempty"`
	LeafLoadBalancingWanInterfaceHealthTestTTLLimit   types.Number `tfsdk:"ttl_limit" vyos:"ttl-limit,omitempty"`
	LeafLoadBalancingWanInterfaceHealthTestType       types.String `tfsdk:"type" vyos:"type,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *LoadBalancingWanInterfaceHealthTest) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *LoadBalancingWanInterfaceHealthTest) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *LoadBalancingWanInterfaceHealthTest) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *LoadBalancingWanInterfaceHealthTest) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"test",
		o.SelfIdentifier.Attributes()["test"].(types.Number).ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *LoadBalancingWanInterfaceHealthTest) GetVyosParentPath() []string {
	return []string{
		"load-balancing",

		"wan",

		"interface-health",

		o.SelfIdentifier.Attributes()["interface_health"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *LoadBalancingWanInterfaceHealthTest) GetVyosNamedParentPath() []string {
	return []string{
		"load-balancing",

		"wan",

		"interface-health",

		o.SelfIdentifier.Attributes()["interface_health"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o LoadBalancingWanInterfaceHealthTest) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"test": schema.NumberAttribute{
					Required: true,
					MarkdownDescription: `Rule number

    |  Format        |  Description  |
    |----------------|---------------|
    |  0-4294967295  |  Rule number  |
`,
					Description: `Rule number

    |  Format        |  Description  |
    |----------------|---------------|
    |  0-4294967295  |  Rule number  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
					},
				},

				"interface_health": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Interface name

`,
					Description: `Interface name

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in interface_health, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  interface_health, value must match: ^[a-zA-Z0-9-_]*$",
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

		"resp_time": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Ping response time (seconds)

    |  Format  |  Description              |
    |----------|---------------------------|
    |  1-30    |  Response time (seconds)  |
`,
			Description: `Ping response time (seconds)

    |  Format  |  Description              |
    |----------|---------------------------|
    |  1-30    |  Response time (seconds)  |
`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		"target": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Health target address

    |  Format  |  Description            |
    |----------|-------------------------|
    |  ipv4    |  Health target address  |
`,
			Description: `Health target address

    |  Format  |  Description            |
    |----------|-------------------------|
    |  ipv4    |  Health target address  |
`,
		},

		"test_script": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Path to user-defined script

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  txt     |  Script in /config/scripts  |
`,
			Description: `Path to user-defined script

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  txt     |  Script in /config/scripts  |
`,
		},

		"ttl_limit": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TTL limit (hop count)

    |  Format  |  Description     |
    |----------|------------------|
    |  1-254   |  Number of hops  |
`,
			Description: `TTL limit (hop count)

    |  Format  |  Description     |
    |----------|------------------|
    |  1-254   |  Number of hops  |
`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `WLB test type

    |  Format        |  Description                         |
    |----------------|--------------------------------------|
    |  ping          |  Test with ICMP echo response        |
    |  ttl           |  Test with UDP TTL expired response  |
    |  user-defined  |  User-defined test script            |
`,
			Description: `WLB test type

    |  Format        |  Description                         |
    |----------------|--------------------------------------|
    |  ping          |  Test with ICMP echo response        |
    |  ttl           |  Test with UDP TTL expired response  |
    |  user-defined  |  User-defined test script            |
`,

			// Default:          stringdefault.StaticString(`ping`),
			Computed: true,
		},

		// Nodes

	}
}
