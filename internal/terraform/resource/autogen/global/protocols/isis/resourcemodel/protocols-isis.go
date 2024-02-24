// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsIsis describes the resource data model.
type ProtocolsIsis struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	// LeafNodes
	LeafProtocolsIsisDynamicHostname     types.Bool   `tfsdk:"dynamic_hostname" vyos:"dynamic-hostname,omitempty"`
	LeafProtocolsIsisLevel               types.String `tfsdk:"level" vyos:"level,omitempty"`
	LeafProtocolsIsisLogAdjacencyChanges types.Bool   `tfsdk:"log_adjacency_changes" vyos:"log-adjacency-changes,omitempty"`
	LeafProtocolsIsisLspGenInterval      types.Number `tfsdk:"lsp_gen_interval" vyos:"lsp-gen-interval,omitempty"`
	LeafProtocolsIsisLspMtu              types.Number `tfsdk:"lsp_mtu" vyos:"lsp-mtu,omitempty"`
	LeafProtocolsIsisLspRefreshInterval  types.Number `tfsdk:"lsp_refresh_interval" vyos:"lsp-refresh-interval,omitempty"`
	LeafProtocolsIsisMaxLspLifetime      types.Number `tfsdk:"max_lsp_lifetime" vyos:"max-lsp-lifetime,omitempty"`
	LeafProtocolsIsisMetricStyle         types.String `tfsdk:"metric_style" vyos:"metric-style,omitempty"`
	LeafProtocolsIsisNet                 types.String `tfsdk:"net" vyos:"net,omitempty"`
	LeafProtocolsIsisPurgeOriginator     types.Bool   `tfsdk:"purge_originator" vyos:"purge-originator,omitempty"`
	LeafProtocolsIsisSetAttachedBit      types.Bool   `tfsdk:"set_attached_bit" vyos:"set-attached-bit,omitempty"`
	LeafProtocolsIsisSetOverloadBit      types.Bool   `tfsdk:"set_overload_bit" vyos:"set-overload-bit,omitempty"`
	LeafProtocolsIsisSpfInterval         types.Number `tfsdk:"spf_interval" vyos:"spf-interval,omitempty"`
	LeafProtocolsIsisRouteMap            types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsIsisInterface bool `tfsdk:"-" vyos:"interface,ignore,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeProtocolsIsisAreaPassword       bool `tfsdk:"-" vyos:"area-password,ignore,omitempty"`
	ExistsNodeProtocolsIsisDefaultInformation bool `tfsdk:"-" vyos:"default-information,ignore,omitempty"`
	ExistsNodeProtocolsIsisDomainPassword     bool `tfsdk:"-" vyos:"domain-password,ignore,omitempty"`
	ExistsNodeProtocolsIsisTrafficEngineering bool `tfsdk:"-" vyos:"traffic-engineering,ignore,omitempty"`
	ExistsNodeProtocolsIsisSegmentRouting     bool `tfsdk:"-" vyos:"segment-routing,ignore,omitempty"`
	ExistsNodeProtocolsIsisRedistribute       bool `tfsdk:"-" vyos:"redistribute,ignore,omitempty"`
	ExistsNodeProtocolsIsisSpfDelayIetf       bool `tfsdk:"-" vyos:"spf-delay-ietf,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsIsis) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsIsis) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"isis",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsIsis) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"dynamic_hostname": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Dynamic hostname for IS-IS

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"level": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IS-IS level number

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  level-1  &emsp; |  Act as a station router  |
    |  level-1-2  &emsp; |  Act as both a station and an area router  |
    |  level-2  &emsp; |  Act as an area router  |

`,
		},

		"log_adjacency_changes": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log adjacency state changes

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"lsp_gen_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum interval between regenerating same LSP

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-120  &emsp; |  Minimum interval in seconds  |

`,
		},

		"lsp_mtu": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Configure the maximum size of generated LSPs

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 128-4352  &emsp; |  Maximum size of generated LSPs  |

`,

			// Default:          stringdefault.StaticString(`1497`),
			Computed: true,
		},

		"lsp_refresh_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `LSP refresh interval

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65235  &emsp; |  LSP refresh interval in seconds  |

`,
		},

		"max_lsp_lifetime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum LSP lifetime

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 350-65535  &emsp; |  LSP lifetime in seconds  |

`,
		},

		"metric_style": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use old-style (ISO 10589) or new-style packet formats

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  narrow  &emsp; |  Use old style of TLVs with narrow metric  |
    |  transition  &emsp; |  Send and accept both styles of TLVs during transition  |
    |  wide  &emsp; |  Use new style of TLVs to carry wider metric  |

`,
		},

		"net": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `A Network Entity Title for this process (ISO only)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  XX.XXXX. ... .XXX.XX  &emsp; |  Network entity title (NET)  |

`,
		},

		"purge_originator": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use the RFC 6232 purge-originator

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"set_attached_bit": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Set attached bit to identify as L1/L2 router for inter-area traffic

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"set_overload_bit": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Set overload bit to avoid any transit traffic

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"spf_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum interval between SPF calculations

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-120  &emsp; |  Interval in seconds  |

`,
		},

		"route_map": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Route map name  |

`,
		},
	}
}
