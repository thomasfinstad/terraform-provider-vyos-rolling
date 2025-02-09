/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (ospf) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsOspf{}

// VrfNameProtocolsOspf describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsOspf struct {
	// LeafNodes
	LeafVrfNameProtocolsOspfDefaultMetric    types.Number `tfsdk:"default_metric" vyos:"default-metric,omitempty"`
	LeafVrfNameProtocolsOspfMaximumPaths     types.Number `tfsdk:"maximum_paths" vyos:"maximum-paths,omitempty"`
	LeafVrfNameProtocolsOspfPassiveInterface types.String `tfsdk:"passive_interface" vyos:"passive-interface,omitempty"`

	// TagNodes

	ExistsTagVrfNameProtocolsOspfAccessList bool `tfsdk:"-" vyos:"access-list,child"`

	ExistsTagVrfNameProtocolsOspfArea bool `tfsdk:"-" vyos:"area,child"`

	ExistsTagVrfNameProtocolsOspfInterface bool `tfsdk:"-" vyos:"interface,child"`

	ExistsTagVrfNameProtocolsOspfNeighbor bool `tfsdk:"-" vyos:"neighbor,child"`

	ExistsTagVrfNameProtocolsOspfSummaryAddress bool `tfsdk:"-" vyos:"summary-address,child"`

	// Nodes

	NodeVrfNameProtocolsOspfAggregation *VrfNameProtocolsOspfAggregation `tfsdk:"aggregation" vyos:"aggregation,omitempty"`

	NodeVrfNameProtocolsOspfAutoCost *VrfNameProtocolsOspfAutoCost `tfsdk:"auto_cost" vyos:"auto-cost,omitempty"`

	NodeVrfNameProtocolsOspfCapability *VrfNameProtocolsOspfCapability `tfsdk:"capability" vyos:"capability,omitempty"`

	NodeVrfNameProtocolsOspfDefaultInformation *VrfNameProtocolsOspfDefaultInformation `tfsdk:"default_information" vyos:"default-information,omitempty"`

	NodeVrfNameProtocolsOspfGracefulRestart *VrfNameProtocolsOspfGracefulRestart `tfsdk:"graceful_restart" vyos:"graceful-restart,omitempty"`

	NodeVrfNameProtocolsOspfLdpSync *VrfNameProtocolsOspfLdpSync `tfsdk:"ldp_sync" vyos:"ldp-sync,omitempty"`

	NodeVrfNameProtocolsOspfDistance *VrfNameProtocolsOspfDistance `tfsdk:"distance" vyos:"distance,omitempty"`

	NodeVrfNameProtocolsOspfLogAdjacencyChanges *VrfNameProtocolsOspfLogAdjacencyChanges `tfsdk:"log_adjacency_changes" vyos:"log-adjacency-changes,omitempty"`

	NodeVrfNameProtocolsOspfMaxMetric *VrfNameProtocolsOspfMaxMetric `tfsdk:"max_metric" vyos:"max-metric,omitempty"`

	NodeVrfNameProtocolsOspfMplsTe *VrfNameProtocolsOspfMplsTe `tfsdk:"mpls_te" vyos:"mpls-te,omitempty"`

	NodeVrfNameProtocolsOspfParameters *VrfNameProtocolsOspfParameters `tfsdk:"parameters" vyos:"parameters,omitempty"`

	NodeVrfNameProtocolsOspfSegmentRouting *VrfNameProtocolsOspfSegmentRouting `tfsdk:"segment_routing" vyos:"segment-routing,omitempty"`

	NodeVrfNameProtocolsOspfRedistribute *VrfNameProtocolsOspfRedistribute `tfsdk:"redistribute" vyos:"redistribute,omitempty"`

	NodeVrfNameProtocolsOspfRefresh *VrfNameProtocolsOspfRefresh `tfsdk:"refresh" vyos:"refresh,omitempty"`

	NodeVrfNameProtocolsOspfTimers *VrfNameProtocolsOspfTimers `tfsdk:"timers" vyos:"timers,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsOspf) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"default_metric":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (default-metric) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Metric of redistributed routes

    |  Format      |  Description                     |
    |--------------|----------------------------------|
    |  0-16777214  |  Metric of redistributed routes  |
`,
			Description: `Metric of redistributed routes

    |  Format      |  Description                     |
    |--------------|----------------------------------|
    |  0-16777214  |  Metric of redistributed routes  |
`,
		},

		"maximum_paths":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (maximum-paths) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum multiple paths (ECMP)

    |  Format  |  Description                    |
    |----------|---------------------------------|
    |  1-64    |  Maximum multiple paths (ECMP)  |
`,
			Description: `Maximum multiple paths (ECMP)

    |  Format  |  Description                    |
    |----------|---------------------------------|
    |  1-64    |  Maximum multiple paths (ECMP)  |
`,
		},

		"passive_interface":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (passive-interface) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Suppress routing updates on an interface

    |  Format   |  Description                                            |
    |-----------|---------------------------------------------------------|
    |  default  |  Default to suppress routing updates on all interfaces  |
`,
			Description: `Suppress routing updates on an interface

    |  Format   |  Description                                            |
    |-----------|---------------------------------------------------------|
    |  default  |  Default to suppress routing updates on all interfaces  |
`,
		},

		// TagNodes

		// Nodes

		"aggregation": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfAggregation{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `External route aggregation

`,
			Description: `External route aggregation

`,
		},

		"auto_cost": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfAutoCost{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Calculate interface cost according to bandwidth

`,
			Description: `Calculate interface cost according to bandwidth

`,
		},

		"capability": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfCapability{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Enable specific OSPF features

`,
			Description: `Enable specific OSPF features

`,
		},

		"default_information": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfDefaultInformation{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Default route advertisment settings

`,
			Description: `Default route advertisment settings

`,
		},

		"graceful_restart": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfGracefulRestart{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Graceful Restart

`,
			Description: `Graceful Restart

`,
		},

		"ldp_sync": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfLdpSync{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Protocol wide LDP-IGP synchronization configuration

`,
			Description: `Protocol wide LDP-IGP synchronization configuration

`,
		},

		"distance": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfDistance{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Administrative distance

`,
			Description: `Administrative distance

`,
		},

		"log_adjacency_changes": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfLogAdjacencyChanges{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Log adjacency state changes

`,
			Description: `Log adjacency state changes

`,
		},

		"max_metric": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfMaxMetric{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `OSPF maximum and infinite-distance metric

`,
			Description: `OSPF maximum and infinite-distance metric

`,
		},

		"mpls_te": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfMplsTe{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `MultiProtocol Label Switching-Traffic Engineering (MPLS-TE) parameters

`,
			Description: `MultiProtocol Label Switching-Traffic Engineering (MPLS-TE) parameters

`,
		},

		"parameters": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfParameters{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `OSPF specific parameters

`,
			Description: `OSPF specific parameters

`,
		},

		"segment_routing": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfSegmentRouting{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Segment-Routing (SPRING) settings

`,
			Description: `Segment-Routing (SPRING) settings

`,
		},

		"redistribute": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfRedistribute{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute information from another routing protocol

`,
			Description: `Redistribute information from another routing protocol

`,
		},

		"refresh": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfRefresh{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Adjust refresh parameters

`,
			Description: `Adjust refresh parameters

`,
		},

		"timers": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsOspfTimers{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Adjust routing timers

`,
			Description: `Adjust routing timers

`,
		},
	}
}
