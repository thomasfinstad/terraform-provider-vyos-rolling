/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsis{}

// VrfNameProtocolsIsis describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsIsis struct {
	// LeafNodes
	LeafVrfNameProtocolsIsisAdvertiseHighMetrics types.Bool   `tfsdk:"advertise_high_metrics" vyos:"advertise-high-metrics,omitempty"`
	LeafVrfNameProtocolsIsisAdvertisePassiveOnly types.Bool   `tfsdk:"advertise_passive_only" vyos:"advertise-passive-only,omitempty"`
	LeafVrfNameProtocolsIsisDynamicHostname      types.Bool   `tfsdk:"dynamic_hostname" vyos:"dynamic-hostname,omitempty"`
	LeafVrfNameProtocolsIsisLevel                types.String `tfsdk:"level" vyos:"level,omitempty"`
	LeafVrfNameProtocolsIsisLogAdjacencyChanges  types.Bool   `tfsdk:"log_adjacency_changes" vyos:"log-adjacency-changes,omitempty"`
	LeafVrfNameProtocolsIsisLspGenInterval       types.Number `tfsdk:"lsp_gen_interval" vyos:"lsp-gen-interval,omitempty"`
	LeafVrfNameProtocolsIsisLspMtu               types.Number `tfsdk:"lsp_mtu" vyos:"lsp-mtu,omitempty"`
	LeafVrfNameProtocolsIsisLspRefreshInterval   types.Number `tfsdk:"lsp_refresh_interval" vyos:"lsp-refresh-interval,omitempty"`
	LeafVrfNameProtocolsIsisMaxLspLifetime       types.Number `tfsdk:"max_lsp_lifetime" vyos:"max-lsp-lifetime,omitempty"`
	LeafVrfNameProtocolsIsisMetricStyle          types.String `tfsdk:"metric_style" vyos:"metric-style,omitempty"`
	LeafVrfNameProtocolsIsisTopology             types.String `tfsdk:"topology" vyos:"topology,omitempty"`
	LeafVrfNameProtocolsIsisNet                  types.String `tfsdk:"net" vyos:"net,omitempty"`
	LeafVrfNameProtocolsIsisPurgeOriginator      types.Bool   `tfsdk:"purge_originator" vyos:"purge-originator,omitempty"`
	LeafVrfNameProtocolsIsisSetAttachedBit       types.Bool   `tfsdk:"set_attached_bit" vyos:"set-attached-bit,omitempty"`
	LeafVrfNameProtocolsIsisSetOverloadBit       types.Bool   `tfsdk:"set_overload_bit" vyos:"set-overload-bit,omitempty"`
	LeafVrfNameProtocolsIsisSpfInterval          types.Number `tfsdk:"spf_interval" vyos:"spf-interval,omitempty"`

	// TagNodes

	ExistsTagVrfNameProtocolsIsisInterface bool `tfsdk:"-" vyos:"interface,child"`

	// Nodes

	NodeVrfNameProtocolsIsisAreaPassword *VrfNameProtocolsIsisAreaPassword `tfsdk:"area_password" vyos:"area-password,omitempty"`

	NodeVrfNameProtocolsIsisDefaultInformation *VrfNameProtocolsIsisDefaultInformation `tfsdk:"default_information" vyos:"default-information,omitempty"`

	NodeVrfNameProtocolsIsisDomainPassword *VrfNameProtocolsIsisDomainPassword `tfsdk:"domain_password" vyos:"domain-password,omitempty"`

	NodeVrfNameProtocolsIsisLdpSync *VrfNameProtocolsIsisLdpSync `tfsdk:"ldp_sync" vyos:"ldp-sync,omitempty"`

	NodeVrfNameProtocolsIsisFastReroute *VrfNameProtocolsIsisFastReroute `tfsdk:"fast_reroute" vyos:"fast-reroute,omitempty"`

	NodeVrfNameProtocolsIsisTrafficEngineering *VrfNameProtocolsIsisTrafficEngineering `tfsdk:"traffic_engineering" vyos:"traffic-engineering,omitempty"`

	NodeVrfNameProtocolsIsisSegmentRouting *VrfNameProtocolsIsisSegmentRouting `tfsdk:"segment_routing" vyos:"segment-routing,omitempty"`

	NodeVrfNameProtocolsIsisRedistribute *VrfNameProtocolsIsisRedistribute `tfsdk:"redistribute" vyos:"redistribute,omitempty"`

	NodeVrfNameProtocolsIsisSpfDelayIetf *VrfNameProtocolsIsisSpfDelayIetf `tfsdk:"spf_delay_ietf" vyos:"spf-delay-ietf,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsis) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"advertise_high_metrics":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise high metric value on all interfaces

`,
			Description: `Advertise high metric value on all interfaces

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"advertise_passive_only":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Advertise prefixes of passive interfaces only

`,
			Description: `Advertise prefixes of passive interfaces only

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"dynamic_hostname":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Dynamic hostname for IS-IS

`,
			Description: `Dynamic hostname for IS-IS

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"level":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IS-IS level number

    |  Format     |  Description                               |
    |-------------|--------------------------------------------|
    |  level-1    |  Act as a station router                   |
    |  level-1-2  |  Act as both a station and an area router  |
    |  level-2    |  Act as an area router                     |
`,
			Description: `IS-IS level number

    |  Format     |  Description                               |
    |-------------|--------------------------------------------|
    |  level-1    |  Act as a station router                   |
    |  level-1-2  |  Act as both a station and an area router  |
    |  level-2    |  Act as an area router                     |
`,
		},

		"log_adjacency_changes":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log changes in adjacency state

`,
			Description: `Log changes in adjacency state

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"lsp_gen_interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum interval between regenerating same LSP

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  1-120   |  Minimum interval in seconds  |
`,
			Description: `Minimum interval between regenerating same LSP

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  1-120   |  Minimum interval in seconds  |
`,
		},

		"lsp_mtu":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Configure the maximum size of generated LSPs

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  128-4352  |  Maximum size of generated LSPs  |
`,
			Description: `Configure the maximum size of generated LSPs

    |  Format    |  Description                     |
    |------------|----------------------------------|
    |  128-4352  |  Maximum size of generated LSPs  |
`,

			// Default:          stringdefault.StaticString(`1497`),
			Computed: true,
		},

		"lsp_refresh_interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `LSP refresh interval

    |  Format   |  Description                      |
    |-----------|-----------------------------------|
    |  1-65235  |  LSP refresh interval in seconds  |
`,
			Description: `LSP refresh interval

    |  Format   |  Description                      |
    |-----------|-----------------------------------|
    |  1-65235  |  LSP refresh interval in seconds  |
`,
		},

		"max_lsp_lifetime":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum LSP lifetime

    |  Format     |  Description              |
    |-------------|---------------------------|
    |  350-65535  |  LSP lifetime in seconds  |
`,
			Description: `Maximum LSP lifetime

    |  Format     |  Description              |
    |-------------|---------------------------|
    |  350-65535  |  LSP lifetime in seconds  |
`,
		},

		"metric_style":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use old-style (ISO 10589) or new-style packet formats

    |  Format      |  Description                                            |
    |--------------|---------------------------------------------------------|
    |  narrow      |  Use old style of TLVs with narrow metric               |
    |  transition  |  Send and accept both styles of TLVs during transition  |
    |  wide        |  Use new style of TLVs to carry wider metric            |
`,
			Description: `Use old-style (ISO 10589) or new-style packet formats

    |  Format      |  Description                                            |
    |--------------|---------------------------------------------------------|
    |  narrow      |  Use old style of TLVs with narrow metric               |
    |  transition  |  Send and accept both styles of TLVs during transition  |
    |  wide        |  Use new style of TLVs to carry wider metric            |
`,
		},

		"topology":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Configure IS-IS topologies

    |  Format          |  Description                   |
    |------------------|--------------------------------|
    |  ipv4-multicast  |  Use IPv4 multicast topology   |
    |  ipv4-mgmt       |  Use IPv4 management topology  |
    |  ipv6-unicast    |  Use IPv6 unicast topology     |
    |  ipv6-multicast  |  Use IPv6 multicast topology   |
    |  ipv6-mgmt       |  Use IPv6 management topology  |
    |  ipv6-dstsrc     |  Use IPv6 dst-src topology     |
`,
			Description: `Configure IS-IS topologies

    |  Format          |  Description                   |
    |------------------|--------------------------------|
    |  ipv4-multicast  |  Use IPv4 multicast topology   |
    |  ipv4-mgmt       |  Use IPv4 management topology  |
    |  ipv6-unicast    |  Use IPv6 unicast topology     |
    |  ipv6-multicast  |  Use IPv6 multicast topology   |
    |  ipv6-mgmt       |  Use IPv6 management topology  |
    |  ipv6-dstsrc     |  Use IPv6 dst-src topology     |
`,
		},

		"net":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `A Network Entity Title for the process (ISO only)

    |  Format                |  Description                 |
    |------------------------|------------------------------|
    |  XX.XXXX. ... .XXX.XX  |  Network entity title (NET)  |
`,
			Description: `A Network Entity Title for the process (ISO only)

    |  Format                |  Description                 |
    |------------------------|------------------------------|
    |  XX.XXXX. ... .XXX.XX  |  Network entity title (NET)  |
`,
		},

		"purge_originator":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use the RFC 6232 purge-originator

`,
			Description: `Use the RFC 6232 purge-originator

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"set_attached_bit":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Set attached bit to identify as L1/L2 router for inter-area traffic

`,
			Description: `Set attached bit to identify as L1/L2 router for inter-area traffic

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"set_overload_bit":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Set overload bit to avoid any transit traffic

`,
			Description: `Set overload bit to avoid any transit traffic

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"spf_interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum interval between SPF calculations

    |  Format  |  Description          |
    |----------|-----------------------|
    |  1-120   |  Interval in seconds  |
`,
			Description: `Minimum interval between SPF calculations

    |  Format  |  Description          |
    |----------|-----------------------|
    |  1-120   |  Interval in seconds  |
`,
		},

		// TagNodes

		// Nodes

		"area_password": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisAreaPassword{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Configure the authentication password for an area

`,
			Description: `Configure the authentication password for an area

`,
		},

		"default_information": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisDefaultInformation{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Control distribution of default information

`,
			Description: `Control distribution of default information

`,
		},

		"domain_password": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisDomainPassword{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Set the authentication password for a routing domain

`,
			Description: `Set the authentication password for a routing domain

`,
		},

		"ldp_sync": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisLdpSync{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Protocol wide LDP-IGP synchronization configuration

`,
			Description: `Protocol wide LDP-IGP synchronization configuration

`,
		},

		"fast_reroute": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisFastReroute{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IS-IS fast reroute configuration

`,
			Description: `IS-IS fast reroute configuration

`,
		},

		"traffic_engineering": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisTrafficEngineering{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IS-IS traffic engineering extensions

`,
			Description: `IS-IS traffic engineering extensions

`,
		},

		"segment_routing": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisSegmentRouting{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Segment-Routing (SPRING) settings

`,
			Description: `Segment-Routing (SPRING) settings

`,
		},

		"redistribute": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisRedistribute{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute information from another routing protocol

`,
			Description: `Redistribute information from another routing protocol

`,
		},

		"spf_delay_ietf": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisSpfDelayIetf{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IETF SPF delay algorithm

`,
			Description: `IETF SPF delay algorithm

`,
		},
	}
}
