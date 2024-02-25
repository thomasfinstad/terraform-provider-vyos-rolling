// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// HighAvailabilityVrrpGroup describes the resource data model.
type HighAvailabilityVrrpGroup struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"group_id" vyos:"-,self-id"`

	// LeafNodes
	LeafHighAvailabilityVrrpGroupInterface                          types.String `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafHighAvailabilityVrrpGroupAdvertiseInterval                  types.Number `tfsdk:"advertise_interval" vyos:"advertise-interval,omitempty"`
	LeafHighAvailabilityVrrpGroupDescrIPtion                        types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafHighAvailabilityVrrpGroupDisable                            types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafHighAvailabilityVrrpGroupHelloSourceAddress                 types.String `tfsdk:"hello_source_address" vyos:"hello-source-address,omitempty"`
	LeafHighAvailabilityVrrpGroupPeerAddress                        types.String `tfsdk:"peer_address" vyos:"peer-address,omitempty"`
	LeafHighAvailabilityVrrpGroupNoPreempt                          types.Bool   `tfsdk:"no_preempt" vyos:"no-preempt,omitempty"`
	LeafHighAvailabilityVrrpGroupPreemptDelay                       types.Number `tfsdk:"preempt_delay" vyos:"preempt-delay,omitempty"`
	LeafHighAvailabilityVrrpGroupPriority                           types.Number `tfsdk:"priority" vyos:"priority,omitempty"`
	LeafHighAvailabilityVrrpGroupRfcthreesevensixeightCompatibility types.Bool   `tfsdk:"rfc3768_compatibility" vyos:"rfc3768-compatibility,omitempty"`
	LeafHighAvailabilityVrrpGroupExcludedAddress                    types.List   `tfsdk:"excluded_address" vyos:"excluded-address,omitempty"`
	LeafHighAvailabilityVrrpGroupVrID                               types.Number `tfsdk:"vrid" vyos:"vrid,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagHighAvailabilityVrrpGroupAddress bool `tfsdk:"-" vyos:"address,child"`

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

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *HighAvailabilityVrrpGroup) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"high-availability",

		"vrrp",

		"group",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"group_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `VRRP group

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |

`,
		},

		"advertise_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Advertise interval

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Advertise interval in seconds  |

`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"hello_source_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRRP hello source address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 hello source address  |
    |  ipv6  &emsp; |  IPv6 hello source address  |

`,
		},

		"peer_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Unicast VRRP peer address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 unicast peer address  |
    |  ipv6  &emsp; |  IPv6 unicast peer address  |

`,
		},

		"no_preempt": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable master preemption

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"preempt_delay": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Preempt delay (in seconds)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-1000  &emsp; |  preempt delay  |

`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"priority": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Router priority

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Router priority  |

`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		"rfc3768_compatibility": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use VRRP virtual MAC address as per RFC3768

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"excluded_address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Virtual address (If you need additional IPv4 and IPv6 in same group)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IP address  |
    |  ipv6  &emsp; |  IPv6 address  |

`,
		},

		"vrid": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Virtual router identifier

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Virtual router identifier  |

`,
		},

		// Nodes

		"garp": schema.SingleNestedAttribute{
			Attributes: HighAvailabilityVrrpGroupGarp{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Gratuitous ARP parameters

`,
		},

		"authentication": schema.SingleNestedAttribute{
			Attributes: HighAvailabilityVrrpGroupAuthentication{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `VRRP authentication

`,
		},

		"health_check": schema.SingleNestedAttribute{
			Attributes: HighAvailabilityVrrpGroupHealthCheck{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Health check

`,
		},

		"track": schema.SingleNestedAttribute{
			Attributes: HighAvailabilityVrrpGroupTrack{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Track settings

`,
		},

		"transition_script": schema.SingleNestedAttribute{
			Attributes: HighAvailabilityVrrpGroupTransitionScrIPt{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `VRRP transition scripts

`,
		},
	}
}
