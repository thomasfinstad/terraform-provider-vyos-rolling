// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleMatchEvpn{}

// PolicyRouteMapRuleMatchEvpn describes the resource data model.
type PolicyRouteMapRuleMatchEvpn struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchEvpnDefaultRoute types.Bool   `tfsdk:"default_route" vyos:"default-route,omitempty"`
	LeafPolicyRouteMapRuleMatchEvpnRd           types.String `tfsdk:"rd" vyos:"rd,omitempty"`
	LeafPolicyRouteMapRuleMatchEvpnRouteType    types.String `tfsdk:"route_type" vyos:"route-type,omitempty"`
	LeafPolicyRouteMapRuleMatchEvpnVni          types.Number `tfsdk:"vni" vyos:"vni,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatchEvpn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"default_route": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Default EVPN type-5 route

`,
			Description: `Default EVPN type-5 route

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rd": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Distinguisher

    |  Format                   |  Description                                   |
    |---------------------------|------------------------------------------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
`,
			Description: `Route Distinguisher

    |  Format                   |  Description                                   |
    |---------------------------|------------------------------------------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
`,
		},

		"route_type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match route-type

    |  Format     |  Description   |
    |-------------|----------------|
    |  macip      |  mac-ip route  |
    |  multicast  |  IMET route    |
    |  prefix     |  Prefix route  |
`,
			Description: `Match route-type

    |  Format     |  Description   |
    |-------------|----------------|
    |  macip      |  mac-ip route  |
    |  multicast  |  IMET route    |
    |  prefix     |  Prefix route  |
`,
		},

		"vni": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Virtual Network Identifier

    |  Format      |  Description                       |
    |--------------|------------------------------------|
    |  0-16777214  |  VXLAN virtual network identifier  |
`,
			Description: `Virtual Network Identifier

    |  Format      |  Description                       |
    |--------------|------------------------------------|
    |  0-16777214  |  VXLAN virtual network identifier  |
`,
		},

		// Nodes

	}
}
