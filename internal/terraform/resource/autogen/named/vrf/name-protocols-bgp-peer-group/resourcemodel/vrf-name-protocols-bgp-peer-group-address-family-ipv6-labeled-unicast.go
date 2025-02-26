/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (ipv6-labeled-unicast) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicast{}

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicast describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicast struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastAddpathTxAll         types.Bool   `tfsdk:"addpath_tx_all" vyos:"addpath-tx-all,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastAddpathTxPerAs       types.Bool   `tfsdk:"addpath_tx_per_as" vyos:"addpath-tx-per-as,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastAsOverrIDe           types.Bool   `tfsdk:"as_override" vyos:"as-override,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastMaximumPrefix        types.Number `tfsdk:"maximum_prefix" vyos:"maximum-prefix,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastMaximumPrefixOut     types.Number `tfsdk:"maximum_prefix_out" vyos:"maximum-prefix-out,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastRouteReflectorClient types.Bool   `tfsdk:"route_reflector_client" vyos:"route-reflector-client,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastRouteServerClient    types.Bool   `tfsdk:"route_server_client" vyos:"route-server-client,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastUnsuppressMap        types.String `tfsdk:"unsuppress_map" vyos:"unsuppress-map,omitempty"`
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastWeight               types.Number `tfsdk:"weight" vyos:"weight,omitempty"`

	// TagNodes

	// Nodes

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastCapability *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastCapability `tfsdk:"capability" vyos:"capability,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastNexthopLocal *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastNexthopLocal `tfsdk:"nexthop_local" vyos:"nexthop-local,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastPrefixList *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastPrefixList `tfsdk:"prefix_list" vyos:"prefix-list,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastConditionallyAdvertise *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastConditionallyAdvertise `tfsdk:"conditionally_advertise" vyos:"conditionally-advertise,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastAllowasIn *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastAllowasIn `tfsdk:"allowas_in" vyos:"allowas-in,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastAttributeUnchanged *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastAttributeUnchanged `tfsdk:"attribute_unchanged" vyos:"attribute-unchanged,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastDisableSendCommunity *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastDisableSendCommunity `tfsdk:"disable_send_community" vyos:"disable-send-community,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastDistributeList *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastDistributeList `tfsdk:"distribute_list" vyos:"distribute-list,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastFilterList *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastFilterList `tfsdk:"filter_list" vyos:"filter-list,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastNexthopSelf *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastNexthopSelf `tfsdk:"nexthop_self" vyos:"nexthop-self,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastRemovePrivateAs *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastRemovePrivateAs `tfsdk:"remove_private_as" vyos:"remove-private-as,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastRouteMap *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastRouteMap `tfsdk:"route_map" vyos:"route-map,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastSoftReconfiguration *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastSoftReconfiguration `tfsdk:"soft_reconfiguration" vyos:"soft-reconfiguration,omitempty"`

	NodeVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastDefaultOriginate *VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastDefaultOriginate `tfsdk:"default_originate" vyos:"default-originate,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicast) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"addpath_tx_all":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (addpath-tx-all) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use addpath to advertise all paths to a neighbor

`,
			Description: `Use addpath to advertise all paths to a neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"addpath_tx_per_as":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (addpath-tx-per-as) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use addpath to advertise the bestpath per each neighboring AS

`,
			Description: `Use addpath to advertise the bestpath per each neighboring AS

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"as_override":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (as-override) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Override ASN in outbound updates to configured neighbor local-as

`,
			Description: `Override ASN in outbound updates to configured neighbor local-as

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"maximum_prefix":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (maximum-prefix) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of prefixes to accept from this peer

    |  Format        |  Description   |
    |----------------|----------------|
    |  1-4294967295  |  Prefix limit  |
`,
			Description: `Maximum number of prefixes to accept from this peer

    |  Format        |  Description   |
    |----------------|----------------|
    |  1-4294967295  |  Prefix limit  |
`,
		},

		"maximum_prefix_out":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (maximum-prefix-out) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of prefixes to be sent to this peer

    |  Format        |  Description   |
    |----------------|----------------|
    |  1-4294967295  |  Prefix limit  |
`,
			Description: `Maximum number of prefixes to be sent to this peer

    |  Format        |  Description   |
    |----------------|----------------|
    |  1-4294967295  |  Prefix limit  |
`,
		},

		"route_reflector_client":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (route-reflector-client) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route reflector client

`,
			Description: `Peer is a route reflector client

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"route_server_client":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (route-server-client) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Peer is a route server client

`,
			Description: `Peer is a route server client

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"unsuppress_map":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (unsuppress-map) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route-map to selectively unsuppress suppressed routes

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
			Description: `Route-map to selectively unsuppress suppressed routes

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
		},

		"weight":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (weight) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Default weight for routes from this peer

    |  Format   |  Description     |
    |-----------|------------------|
    |  1-65535  |  Default weight  |
`,
			Description: `Default weight for routes from this peer

    |  Format   |  Description     |
    |-----------|------------------|
    |  1-65535  |  Default weight  |
`,
		},

		// TagNodes

		// Nodes

		"capability": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastCapability{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Advertise capabilities to this neighbor (IPv6)

`,
			Description: `Advertise capabilities to this neighbor (IPv6)

`,
		},

		"nexthop_local": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastNexthopLocal{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Nexthop attributes

`,
			Description: `Nexthop attributes

`,
		},

		"prefix_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastPrefixList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Prefix-list to filter route updates to/from this peer

`,
			Description: `Prefix-list to filter route updates to/from this peer

`,
		},

		"conditionally_advertise": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastConditionallyAdvertise{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Use route-map to conditionally advertise routes

`,
			Description: `Use route-map to conditionally advertise routes

`,
		},

		"allowas_in": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastAllowasIn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Accept route that contains the local-as in the as-path

`,
			Description: `Accept route that contains the local-as in the as-path

`,
		},

		"attribute_unchanged": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastAttributeUnchanged{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `BGP attributes are sent unchanged

`,
			Description: `BGP attributes are sent unchanged

`,
		},

		"disable_send_community": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastDisableSendCommunity{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Disable sending community attributes to this peer

`,
			Description: `Disable sending community attributes to this peer

`,
		},

		"distribute_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastDistributeList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Access-list to filter route updates to/from this peer-group

`,
			Description: `Access-list to filter route updates to/from this peer-group

`,
		},

		"filter_list": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastFilterList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `as-path-list to filter route updates to/from this peer

`,
			Description: `as-path-list to filter route updates to/from this peer

`,
		},

		"nexthop_self": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastNexthopSelf{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Disable the next hop calculation for this peer

`,
			Description: `Disable the next hop calculation for this peer

`,
		},

		"remove_private_as": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastRemovePrivateAs{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Remove private AS numbers from AS path in outbound route updates

`,
			Description: `Remove private AS numbers from AS path in outbound route updates

`,
		},

		"route_map": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastRouteMap{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Route-map to filter route updates to/from this peer

`,
			Description: `Route-map to filter route updates to/from this peer

`,
		},

		"soft_reconfiguration": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastSoftReconfiguration{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Soft reconfiguration for peer

`,
			Description: `Soft reconfiguration for peer

`,
		},

		"default_originate": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastDefaultOriginate{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Originate default route to this peer

`,
			Description: `Originate default route to this peer

`,
		},
	}
}
