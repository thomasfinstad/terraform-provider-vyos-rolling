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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (match) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRouteMapRuleMatch{}

// PolicyRouteMapRuleMatch describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type PolicyRouteMapRuleMatch struct {
	// LeafNodes
	LeafPolicyRouteMapRuleMatchAsPath          types.String `tfsdk:"as_path" vyos:"as-path,omitempty"`
	LeafPolicyRouteMapRuleMatchExtcommunity    types.String `tfsdk:"extcommunity" vyos:"extcommunity,omitempty"`
	LeafPolicyRouteMapRuleMatchInterface       types.String `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafPolicyRouteMapRuleMatchLocalPreference types.Number `tfsdk:"local_preference" vyos:"local-preference,omitempty"`
	LeafPolicyRouteMapRuleMatchMetric          types.Number `tfsdk:"metric" vyos:"metric,omitempty"`
	LeafPolicyRouteMapRuleMatchOrigin          types.String `tfsdk:"origin" vyos:"origin,omitempty"`
	LeafPolicyRouteMapRuleMatchPeer            types.String `tfsdk:"peer" vyos:"peer,omitempty"`
	LeafPolicyRouteMapRuleMatchProtocol        types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafPolicyRouteMapRuleMatchRpki            types.String `tfsdk:"rpki" vyos:"rpki,omitempty"`
	LeafPolicyRouteMapRuleMatchTag             types.Number `tfsdk:"tag" vyos:"tag,omitempty"`

	// TagNodes

	// Nodes

	NodePolicyRouteMapRuleMatchCommunity *PolicyRouteMapRuleMatchCommunity `tfsdk:"community" vyos:"community,omitempty"`

	NodePolicyRouteMapRuleMatchEvpn *PolicyRouteMapRuleMatchEvpn `tfsdk:"evpn" vyos:"evpn,omitempty"`

	NodePolicyRouteMapRuleMatchIP *PolicyRouteMapRuleMatchIP `tfsdk:"ip" vyos:"ip,omitempty"`

	NodePolicyRouteMapRuleMatchIPvsix *PolicyRouteMapRuleMatchIPvsix `tfsdk:"ipv6" vyos:"ipv6,omitempty"`

	NodePolicyRouteMapRuleMatchLargeCommunity *PolicyRouteMapRuleMatchLargeCommunity `tfsdk:"large_community" vyos:"large-community,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRouteMapRuleMatch) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"as_path":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (as-path) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP as-path-list to match

`,
			Description: `BGP as-path-list to match

`,
		},

		"extcommunity":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (extcommunity) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP extended community to match

`,
			Description: `BGP extended community to match

`,
		},

		"interface":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (interface) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
			Description: `Interface

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
		},

		"local_preference":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (local-preference) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Local Preference

    |  Format        |  Description       |
    |----------------|--------------------|
    |  0-4294967295  |  Local Preference  |
`,
			Description: `Local Preference

    |  Format        |  Description       |
    |----------------|--------------------|
    |  0-4294967295  |  Local Preference  |
`,
		},

		"metric":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (metric) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Metric of route to match

    |  Format   |  Description   |
    |-----------|----------------|
    |  1-65535  |  Route metric  |
`,
			Description: `Metric of route to match

    |  Format   |  Description   |
    |-----------|----------------|
    |  1-65535  |  Route metric  |
`,
		},

		"origin":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (origin) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP origin code to match

    |  Format      |  Description                       |
    |--------------|------------------------------------|
    |  egp         |  Exterior gateway protocol origin  |
    |  igp         |  Interior gateway protocol origin  |
    |  incomplete  |  Incomplete origin                 |
`,
			Description: `BGP origin code to match

    |  Format      |  Description                       |
    |--------------|------------------------------------|
    |  egp         |  Exterior gateway protocol origin  |
    |  igp         |  Interior gateway protocol origin  |
    |  incomplete  |  Incomplete origin                 |
`,
		},

		"peer":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (peer) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Peer address to match

    |  Format  |  Description        |
    |----------|---------------------|
    |  ipv4    |  Peer IP address    |
    |  ipv6    |  Peer IPv6 address  |
`,
			Description: `Peer address to match

    |  Format  |  Description        |
    |----------|---------------------|
    |  ipv4    |  Peer IP address    |
    |  ipv6    |  Peer IPv6 address  |
`,
		},

		"protocol":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (protocol) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match protocol via which the route was learnt

    |  Format     |  Description                                                  |
    |-------------|---------------------------------------------------------------|
    |  babel      |  Babel routing protocol (Babel)                               |
    |  bgp        |  Border Gateway Protocol (BGP)                                |
    |  connected  |  Connected routes (directly attached subnet or host)          |
    |  isis       |  Intermediate System to Intermediate System (IS-IS)           |
    |  kernel     |  Kernel routes                                                |
    |  ospf       |  Open Shortest Path First (OSPFv2)                            |
    |  ospfv3     |  Open Shortest Path First (IPv6) (OSPFv3)                     |
    |  rip        |  Routing Information Protocol (RIP)                           |
    |  ripng      |  Routing Information Protocol next-generation (IPv6) (RIPng)  |
    |  static     |  Statically configured routes                                 |
    |  table      |  Non-main Kernel Routing Table                                |
    |  vnc        |  Virtual Network Control (VNC)                                |
`,
			Description: `Match protocol via which the route was learnt

    |  Format     |  Description                                                  |
    |-------------|---------------------------------------------------------------|
    |  babel      |  Babel routing protocol (Babel)                               |
    |  bgp        |  Border Gateway Protocol (BGP)                                |
    |  connected  |  Connected routes (directly attached subnet or host)          |
    |  isis       |  Intermediate System to Intermediate System (IS-IS)           |
    |  kernel     |  Kernel routes                                                |
    |  ospf       |  Open Shortest Path First (OSPFv2)                            |
    |  ospfv3     |  Open Shortest Path First (IPv6) (OSPFv3)                     |
    |  rip        |  Routing Information Protocol (RIP)                           |
    |  ripng      |  Routing Information Protocol next-generation (IPv6) (RIPng)  |
    |  static     |  Statically configured routes                                 |
    |  table      |  Non-main Kernel Routing Table                                |
    |  vnc        |  Virtual Network Control (VNC)                                |
`,
		},

		"rpki":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (rpki) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match RPKI validation result

    |  Format    |  Description             |
    |------------|--------------------------|
    |  invalid   |  Match invalid entries   |
    |  notfound  |  Match notfound entries  |
    |  valid     |  Match valid entries     |
`,
			Description: `Match RPKI validation result

    |  Format    |  Description             |
    |------------|--------------------------|
    |  invalid   |  Match invalid entries   |
    |  notfound  |  Match notfound entries  |
    |  valid     |  Match valid entries     |
`,
		},

		"tag":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (tag) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Route tag value

    |  Format   |  Description  |
    |-----------|---------------|
    |  1-65535  |  Route tag    |
`,
			Description: `Route tag value

    |  Format   |  Description  |
    |-----------|---------------|
    |  1-65535  |  Route tag    |
`,
		},

		// TagNodes

		// Nodes

		"community": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleMatchCommunity{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `BGP community-list to match

`,
			Description: `BGP community-list to match

`,
		},

		"evpn": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleMatchEvpn{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Ethernet Virtual Private Network

`,
			Description: `Ethernet Virtual Private Network

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleMatchIP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IP prefix parameters to match

`,
			Description: `IP prefix parameters to match

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleMatchIPvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 prefix parameters to match

`,
			Description: `IPv6 prefix parameters to match

`,
		},

		"large_community": schema.SingleNestedAttribute{
			Attributes: PolicyRouteMapRuleMatchLargeCommunity{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Match BGP large communities

`,
			Description: `Match BGP large communities

`,
		},
	}
}
