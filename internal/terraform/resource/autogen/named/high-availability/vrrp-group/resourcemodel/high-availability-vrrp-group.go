// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
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

var _ helpers.VyosTopResourceDataModel = &HighAvailabilityVrrpGroup{}

// HighAvailabilityVrrpGroup describes the resource data model.
type HighAvailabilityVrrpGroup struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"group_id" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafHighAvailabilityVrrpGroupInterface                          types.String `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafHighAvailabilityVrrpGroupAdvertiseInterval                  types.Number `tfsdk:"advertise_interval" vyos:"advertise-interval,omitempty"`
	LeafHighAvailabilityVrrpGroupDescrIPtion                        types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafHighAvailabilityVrrpGroupDisable                            types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafHighAvailabilityVrrpGroupHelloSourceAddress                 types.String `tfsdk:"hello_source_address" vyos:"hello-source-address,omitempty"`
	LeafHighAvailabilityVrrpGroupPeerAddress                        types.List   `tfsdk:"peer_address" vyos:"peer-address,omitempty"`
	LeafHighAvailabilityVrrpGroupNoPreempt                          types.Bool   `tfsdk:"no_preempt" vyos:"no-preempt,omitempty"`
	LeafHighAvailabilityVrrpGroupPreemptDelay                       types.Number `tfsdk:"preempt_delay" vyos:"preempt-delay,omitempty"`
	LeafHighAvailabilityVrrpGroupPriority                           types.Number `tfsdk:"priority" vyos:"priority,omitempty"`
	LeafHighAvailabilityVrrpGroupRfcthreesevensixeightCompatibility types.Bool   `tfsdk:"rfc3768_compatibility" vyos:"rfc3768-compatibility,omitempty"`
	LeafHighAvailabilityVrrpGroupVrID                               types.Number `tfsdk:"vrid" vyos:"vrid,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagHighAvailabilityVrrpGroupAddress         bool `tfsdk:"-" vyos:"address,child"`
	ExistsTagHighAvailabilityVrrpGroupExcludedAddress bool `tfsdk:"-" vyos:"excluded-address,child"`

	// Nodes
	NodeHighAvailabilityVrrpGroupGarp             *HighAvailabilityVrrpGroupGarp             `tfsdk:"garp" vyos:"garp,omitempty"`
	NodeHighAvailabilityVrrpGroupAuthentication   *HighAvailabilityVrrpGroupAuthentication   `tfsdk:"authentication" vyos:"authentication,omitempty"`
	NodeHighAvailabilityVrrpGroupHealthCheck      *HighAvailabilityVrrpGroupHealthCheck      `tfsdk:"health_check" vyos:"health-check,omitempty"`
	NodeHighAvailabilityVrrpGroupTrack            *HighAvailabilityVrrpGroupTrack            `tfsdk:"track" vyos:"track,omitempty"`
	NodeHighAvailabilityVrrpGroupTransitionScrIPt *HighAvailabilityVrrpGroupTransitionScrIPt `tfsdk:"transition_script" vyos:"transition-script,omitempty"`
}

// SetID configures the resource ID
func (o *HighAvailabilityVrrpGroup) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *HighAvailabilityVrrpGroup) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *HighAvailabilityVrrpGroup) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *HighAvailabilityVrrpGroup) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"group",
		o.SelfIdentifier.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *HighAvailabilityVrrpGroup) GetVyosParentPath() []string {
	return []string{
		"high-availability",

		"vrrp",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *HighAvailabilityVrrpGroup) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"group_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `VRRP group

`,
			Description: `VRRP group

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in group_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illegal character in  group_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

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

		"advertise_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Advertise interval

    |  Format  |  Description                    |
    |----------|---------------------------------|
    |  1-255   |  Advertise interval in seconds  |
`,
			Description: `Advertise interval

    |  Format  |  Description                    |
    |----------|---------------------------------|
    |  1-255   |  Advertise interval in seconds  |
`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

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

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"hello_source_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRRP hello source address

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  ipv4    |  IPv4 hello source address  |
    |  ipv6    |  IPv6 hello source address  |
`,
			Description: `VRRP hello source address

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  ipv4    |  IPv4 hello source address  |
    |  ipv6    |  IPv6 hello source address  |
`,
		},

		"peer_address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Unicast VRRP peer address

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  ipv4    |  IPv4 unicast peer address  |
    |  ipv6    |  IPv6 unicast peer address  |
`,
			Description: `Unicast VRRP peer address

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  ipv4    |  IPv4 unicast peer address  |
    |  ipv6    |  IPv6 unicast peer address  |
`,
		},

		"no_preempt": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable master preemption

`,
			Description: `Disable master preemption

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"preempt_delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Preempt delay (in seconds)

    |  Format  |  Description    |
    |----------|-----------------|
    |  0-1000  |  preempt delay  |
`,
			Description: `Preempt delay (in seconds)

    |  Format  |  Description    |
    |----------|-----------------|
    |  0-1000  |  preempt delay  |
`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"priority": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Router priority

    |  Format  |  Description      |
    |----------|-------------------|
    |  1-255   |  Router priority  |
`,
			Description: `Router priority

    |  Format  |  Description      |
    |----------|-------------------|
    |  1-255   |  Router priority  |
`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		"rfc3768_compatibility": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use VRRP virtual MAC address as per RFC3768

`,
			Description: `Use VRRP virtual MAC address as per RFC3768

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"vrid": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Virtual router identifier

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  1-255   |  Virtual router identifier  |
`,
			Description: `Virtual router identifier

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  1-255   |  Virtual router identifier  |
`,
		},

		// Nodes

		"garp": schema.SingleNestedAttribute{
			Attributes: HighAvailabilityVrrpGroupGarp{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Gratuitous ARP parameters

`,
			Description: `Gratuitous ARP parameters

`,
		},

		"authentication": schema.SingleNestedAttribute{
			Attributes: HighAvailabilityVrrpGroupAuthentication{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `VRRP authentication

`,
			Description: `VRRP authentication

`,
		},

		"health_check": schema.SingleNestedAttribute{
			Attributes: HighAvailabilityVrrpGroupHealthCheck{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Health check

`,
			Description: `Health check

`,
		},

		"track": schema.SingleNestedAttribute{
			Attributes: HighAvailabilityVrrpGroupTrack{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Track settings

`,
			Description: `Track settings

`,
		},

		"transition_script": schema.SingleNestedAttribute{
			Attributes: HighAvailabilityVrrpGroupTransitionScrIPt{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `VRRP transition scripts

`,
			Description: `VRRP transition scripts

`,
		},
	}
}
