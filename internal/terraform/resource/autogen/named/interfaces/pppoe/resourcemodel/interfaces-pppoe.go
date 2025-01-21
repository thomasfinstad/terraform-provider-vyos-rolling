/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (pppoe) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &InterfacesPppoe{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (pppoe) */
// InterfacesPppoeSelfIdentifier used by TagNodes to keep the id info
type InterfacesPppoeSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (pppoe) */

	InterfacesPppoe types.String `tfsdk:"pppoe" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (interfaces) */
}

// InterfacesPppoe describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type InterfacesPppoe struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *InterfacesPppoeSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafInterfacesPppoeAccessConcentrator   types.String `tfsdk:"access_concentrator" vyos:"access-concentrator,omitempty"`
	LeafInterfacesPppoeConnectOnDemand      types.Bool   `tfsdk:"connect_on_demand" vyos:"connect-on-demand,omitempty"`
	LeafInterfacesPppoeNoDefaultRoute       types.Bool   `tfsdk:"no_default_route" vyos:"no-default-route,omitempty"`
	LeafInterfacesPppoeDefaultRouteDistance types.Number `tfsdk:"default_route_distance" vyos:"default-route-distance,omitempty"`
	LeafInterfacesPppoeDescrIPtion          types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesPppoeDisable              types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesPppoeIDleTimeout          types.Number `tfsdk:"idle_timeout" vyos:"idle-timeout,omitempty"`
	LeafInterfacesPppoeHostUniq             types.String `tfsdk:"host_uniq" vyos:"host-uniq,omitempty"`
	LeafInterfacesPppoeHoldoff              types.Number `tfsdk:"holdoff" vyos:"holdoff,omitempty"`
	LeafInterfacesPppoeSourceInterface      types.String `tfsdk:"source_interface" vyos:"source-interface,omitempty"`
	LeafInterfacesPppoeLocalAddress         types.String `tfsdk:"local_address" vyos:"local-address,omitempty"`
	LeafInterfacesPppoeMtu                  types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesPppoeMru                  types.Number `tfsdk:"mru" vyos:"mru,omitempty"`
	LeafInterfacesPppoeNoPeerDNS            types.Bool   `tfsdk:"no_peer_dns" vyos:"no-peer-dns,omitempty"`
	LeafInterfacesPppoeRemoteAddress        types.String `tfsdk:"remote_address" vyos:"remote-address,omitempty"`
	LeafInterfacesPppoeServiceName          types.String `tfsdk:"service_name" vyos:"service-name,omitempty"`
	LeafInterfacesPppoeRedirect             types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`
	LeafInterfacesPppoeVrf                  types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes

	// Nodes

	NodeInterfacesPppoeAuthentication *InterfacesPppoeAuthentication `tfsdk:"authentication" vyos:"authentication,omitempty"`

	NodeInterfacesPppoeDhcpvsixOptions *InterfacesPppoeDhcpvsixOptions `tfsdk:"dhcpv6_options" vyos:"dhcpv6-options,omitempty"`

	NodeInterfacesPppoeIP *InterfacesPppoeIP `tfsdk:"ip" vyos:"ip,omitempty"`

	NodeInterfacesPppoeIPvsix *InterfacesPppoeIPvsix `tfsdk:"ipv6" vyos:"ipv6,omitempty"`

	NodeInterfacesPppoeMirror *InterfacesPppoeMirror `tfsdk:"mirror" vyos:"mirror,omitempty"`
}

// SetID configures the resource ID
func (o *InterfacesPppoe) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *InterfacesPppoe) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *InterfacesPppoe) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesPppoe) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"pppoe",
		o.SelfIdentifier.InterfacesPppoe.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *InterfacesPppoe) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (interfaces) */
		"interfaces", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *InterfacesPppoe) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (interfaces) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesPppoe) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"pppoe": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Point-to-Point Protocol over Ethernet (PPPoE) Interface

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  pppoeN  |  PPPoE dialer interface name  |
`,
					Description: `Point-to-Point Protocol over Ethernet (PPPoE) Interface

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  pppoeN  |  PPPoE dialer interface name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in pppoe, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  pppoe, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (interfaces) */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"access_concentrator":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (access-concentrator) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Access concentrator name

`,
			Description: `Access concentrator name

`,
		},

		"connect_on_demand":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (connect-on-demand) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Establishment connection automatically when traffic is sent

`,
			Description: `Establishment connection automatically when traffic is sent

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_default_route":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (no-default-route) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not install default route to system

`,
			Description: `Do not install default route to system

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"default_route_distance":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (default-route-distance) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for installed default route

    |  Format  |  Description                                              |
    |----------|-----------------------------------------------------------|
    |  1-255   |  Distance for the default route received from the server  |
`,
			Description: `Distance for installed default route

    |  Format  |  Description                                              |
    |----------|-----------------------------------------------------------|
    |  1-255   |  Distance for the default route received from the server  |
`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"description":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (description) */
		schema.StringAttribute{
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

		"disable":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (disable) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administratively disable interface

`,
			Description: `Administratively disable interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"idle_timeout":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (idle-timeout) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Delay before disconnecting idle session (in seconds)

    |  Format   |  Description              |
    |-----------|---------------------------|
    |  0-86400  |  Idle timeout in seconds  |
`,
			Description: `Delay before disconnecting idle session (in seconds)

    |  Format   |  Description              |
    |-----------|---------------------------|
    |  0-86400  |  Idle timeout in seconds  |
`,
		},

		"host_uniq":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (host-uniq) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `PPPoE RFC2516 host-uniq tag

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  txt     |  Host-uniq tag as byte string in HEX  |
`,
			Description: `PPPoE RFC2516 host-uniq tag

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  txt     |  Host-uniq tag as byte string in HEX  |
`,
		},

		"holdoff":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (holdoff) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Delay before re-dial to the access concentrator when PPP session terminated by peer (in seconds)

    |  Format   |  Description              |
    |-----------|---------------------------|
    |  0-86400  |  Holdoff time in seconds  |
`,
			Description: `Delay before re-dial to the access concentrator when PPP session terminated by peer (in seconds)

    |  Format   |  Description              |
    |-----------|---------------------------|
    |  0-86400  |  Holdoff time in seconds  |
`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"source_interface":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (source-interface) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface used to establish connection

    |  Format     |  Description     |
    |-------------|------------------|
    |  interface  |  Interface name  |
`,
			Description: `Interface used to establish connection

    |  Format     |  Description     |
    |-------------|------------------|
    |  interface  |  Interface name  |
`,
		},

		"local_address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (local-address) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 address of local end of the PPPoE link

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  ipv4    |  Address of local end of the PPPoE link  |
`,
			Description: `IPv4 address of local end of the PPPoE link

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  ipv4    |  Address of local end of the PPPoE link  |
`,
		},

		"mtu":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mtu) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

    |  Format   |  Description                        |
    |-----------|-------------------------------------|
    |  68-1500  |  Maximum Transmission Unit in byte  |
`,
			Description: `Maximum Transmission Unit (MTU)

    |  Format   |  Description                        |
    |-----------|-------------------------------------|
    |  68-1500  |  Maximum Transmission Unit in byte  |
`,

			// Default:          stringdefault.StaticString(`1492`),
			Computed: true,
		},

		"mru":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mru) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Receive Unit (MRU) (default: MTU value)

    |  Format     |  Description                   |
    |-------------|--------------------------------|
    |  128-16384  |  Maximum Receive Unit in byte  |
`,
			Description: `Maximum Receive Unit (MRU) (default: MTU value)

    |  Format     |  Description                   |
    |-------------|--------------------------------|
    |  128-16384  |  Maximum Receive Unit in byte  |
`,
		},

		"no_peer_dns":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (no-peer-dns) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not use DNS servers provided by the peer

`,
			Description: `Do not use DNS servers provided by the peer

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"remote_address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (remote-address) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 address of remote end of the PPPoE link

    |  Format  |  Description                              |
    |----------|-------------------------------------------|
    |  ipv4    |  Address of remote end of the PPPoE link  |
`,
			Description: `IPv4 address of remote end of the PPPoE link

    |  Format  |  Description                              |
    |----------|-------------------------------------------|
    |  ipv4    |  Address of remote end of the PPPoE link  |
`,
		},

		"service_name":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (service-name) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Service name, only connect to access concentrators advertising this

`,
			Description: `Service name, only connect to access concentrators advertising this

`,
		},

		"redirect":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (redirect) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Redirect incoming packet to destination

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  Destination interface name  |
`,
			Description: `Redirect incoming packet to destination

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  Destination interface name  |
`,
		},

		"vrf":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (vrf) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
			Description: `VRF instance name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
		},

		// TagNodes

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: InterfacesPppoeAuthentication{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Authentication settings

`,
			Description: `Authentication settings

`,
		},

		"dhcpv6_options": schema.SingleNestedAttribute{
			Attributes: InterfacesPppoeDhcpvsixOptions{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `DHCPv6 client settings/options

`,
			Description: `DHCPv6 client settings/options

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesPppoeIP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
			Description: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesPppoeIPvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
			Description: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesPppoeMirror{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
			Description: `Mirror ingress/egress packets

`,
		},
	}
}
