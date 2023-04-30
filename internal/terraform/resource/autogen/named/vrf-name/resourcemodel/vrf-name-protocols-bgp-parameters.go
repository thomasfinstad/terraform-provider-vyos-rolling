// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpParameters describes the resource data model.
type VrfNameProtocolsBgpParameters struct {
	// LeafNodes
	VrfNameProtocolsBgpParametersAlwaysCompareMed                  customtypes.CustomStringValue `tfsdk:"always_compare_med" json:"always-compare-med,omitempty"`
	VrfNameProtocolsBgpParametersClusterID                         customtypes.CustomStringValue `tfsdk:"cluster_id" json:"cluster-id,omitempty"`
	VrfNameProtocolsBgpParametersDeterministicMed                  customtypes.CustomStringValue `tfsdk:"deterministic_med" json:"deterministic-med,omitempty"`
	VrfNameProtocolsBgpParametersEbgpRequiresPolicy                customtypes.CustomStringValue `tfsdk:"ebgp_requires_policy" json:"ebgp-requires-policy,omitempty"`
	VrfNameProtocolsBgpParametersFastConvergence                   customtypes.CustomStringValue `tfsdk:"fast_convergence" json:"fast-convergence,omitempty"`
	VrfNameProtocolsBgpParametersGracefulShutdown                  customtypes.CustomStringValue `tfsdk:"graceful_shutdown" json:"graceful-shutdown,omitempty"`
	VrfNameProtocolsBgpParametersLogNeighborChanges                customtypes.CustomStringValue `tfsdk:"log_neighbor_changes" json:"log-neighbor-changes,omitempty"`
	VrfNameProtocolsBgpParametersMinimumHoldtime                   customtypes.CustomStringValue `tfsdk:"minimum_holdtime" json:"minimum-holdtime,omitempty"`
	VrfNameProtocolsBgpParametersNetworkImportCheck                customtypes.CustomStringValue `tfsdk:"network_import_check" json:"network-import-check,omitempty"`
	VrfNameProtocolsBgpParametersRouteReflectorAllowOutboundPolicy customtypes.CustomStringValue `tfsdk:"route_reflector_allow_outbound_policy" json:"route-reflector-allow-outbound-policy,omitempty"`
	VrfNameProtocolsBgpParametersNoClientToClientReflection        customtypes.CustomStringValue `tfsdk:"no_client_to_client_reflection" json:"no-client-to-client-reflection,omitempty"`
	VrfNameProtocolsBgpParametersNoFastExternalFailover            customtypes.CustomStringValue `tfsdk:"no_fast_external_failover" json:"no-fast-external-failover,omitempty"`
	VrfNameProtocolsBgpParametersNoSuppressDuplicates              customtypes.CustomStringValue `tfsdk:"no_suppress_duplicates" json:"no-suppress-duplicates,omitempty"`
	VrfNameProtocolsBgpParametersRejectAsSets                      customtypes.CustomStringValue `tfsdk:"reject_as_sets" json:"reject-as-sets,omitempty"`
	VrfNameProtocolsBgpParametersShutdown                          customtypes.CustomStringValue `tfsdk:"shutdown" json:"shutdown,omitempty"`
	VrfNameProtocolsBgpParametersSuppressFibPending                customtypes.CustomStringValue `tfsdk:"suppress_fib_pending" json:"suppress-fib-pending,omitempty"`
	VrfNameProtocolsBgpParametersRouterID                          customtypes.CustomStringValue `tfsdk:"router_id" json:"router-id,omitempty"`

	// TagNodes

	// Nodes
	VrfNameProtocolsBgpParametersBestpath                 types.Object `tfsdk:"bestpath" json:"bestpath,omitempty"`
	VrfNameProtocolsBgpParametersConfederation            types.Object `tfsdk:"confederation" json:"confederation,omitempty"`
	VrfNameProtocolsBgpParametersConditionalAdvertisement types.Object `tfsdk:"conditional_advertisement" json:"conditional-advertisement,omitempty"`
	VrfNameProtocolsBgpParametersDampening                types.Object `tfsdk:"dampening" json:"dampening,omitempty"`
	VrfNameProtocolsBgpParametersDefault                  types.Object `tfsdk:"default" json:"default,omitempty"`
	VrfNameProtocolsBgpParametersDistance                 types.Object `tfsdk:"distance" json:"distance,omitempty"`
	VrfNameProtocolsBgpParametersGracefulRestart          types.Object `tfsdk:"graceful_restart" json:"graceful-restart,omitempty"`
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpParameters) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"always_compare_med": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Always compare MEDs from different neighbors

`,
		},

		"cluster_id": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Route-reflector cluster-id

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Route-reflector cluster-id  |

`,
		},

		"deterministic_med": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Compare MEDs between different peers in the same AS

`,
		},

		"ebgp_requires_policy": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Require in and out policy for eBGP peers (RFC8212)

`,
		},

		"fast_convergence": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Teardown sessions immediately whenever peer becomes unreachable

`,
		},

		"graceful_shutdown": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Graceful shutdown

`,
		},

		"log_neighbor_changes": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Log neighbor up/down changes and reset reason

`,
		},

		"minimum_holdtime": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `BGP minimum holdtime

|  Format  |  Description  |
|----------|---------------|
|  u32:1-65535  |  Minimum holdtime in seconds  |

`,
		},

		"network_import_check": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Enable IGP route check for network statements

`,
		},

		"route_reflector_allow_outbound_policy": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Route reflector client allow policy outbound

`,
		},

		"no_client_to_client_reflection": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable client to client route reflection

`,
		},

		"no_fast_external_failover": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable immediate session reset on peer link down event

`,
		},

		"no_suppress_duplicates": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Disable suppress duplicate updates if the route actually not changed

`,
		},

		"reject_as_sets": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Reject routes with AS_SET or AS_CONFED_SET flag

`,
		},

		"shutdown": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Administrative shutdown of the BGP instance

`,
		},

		"suppress_fib_pending": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Advertise only routes that are programmed in kernel to peers

`,
		},

		"router_id": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Override default router identifier

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Router-ID in IP address format  |

`,
		},

		// TagNodes

		// Nodes

		"bestpath": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersBestpath{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Default bestpath selection mechanism

`,
		},

		"confederation": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersConfederation{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `AS confederation parameters

`,
		},

		"conditional_advertisement": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersConditionalAdvertisement{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Conditional advertisement settings

`,
		},

		"dampening": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersDampening{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Enable route-flap dampening

`,
		},

		"default": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersDefault{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `BGP defaults

`,
		},

		"distance": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersDistance{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Administratives distances for BGP routes

`,
		},

		"graceful_restart": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpParametersGracefulRestart{}.ResourceAttributes(),
			Optional:   true,
			MarkdownDescription: `Graceful restart capability parameters

`,
		},
	}
}
