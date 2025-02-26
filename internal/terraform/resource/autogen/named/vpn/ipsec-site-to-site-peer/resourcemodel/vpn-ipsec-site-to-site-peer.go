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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (peer) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VpnIPsecSiteToSitePeer{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (peer) */
// VpnIPsecSiteToSitePeerSelfIdentifier used by TagNodes to keep the id info
type VpnIPsecSiteToSitePeerSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (peer) */

	VpnIPsecSiteToSitePeer types.String `tfsdk:"peer" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (site-to-site) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (ipsec) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (vpn) */
}

// VpnIPsecSiteToSitePeer describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type VpnIPsecSiteToSitePeer struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *VpnIPsecSiteToSitePeerSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafVpnIPsecSiteToSitePeerDisable               types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafVpnIPsecSiteToSitePeerConnectionType        types.String `tfsdk:"connection_type" vyos:"connection-type,omitempty"`
	LeafVpnIPsecSiteToSitePeerDefaultEspGroup       types.String `tfsdk:"default_esp_group" vyos:"default-esp-group,omitempty"`
	LeafVpnIPsecSiteToSitePeerDescrIPtion           types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafVpnIPsecSiteToSitePeerDhcpInterface         types.String `tfsdk:"dhcp_interface" vyos:"dhcp-interface,omitempty"`
	LeafVpnIPsecSiteToSitePeerForceUDPEncapsulation types.Bool   `tfsdk:"force_udp_encapsulation" vyos:"force-udp-encapsulation,omitempty"`
	LeafVpnIPsecSiteToSitePeerIkeGroup              types.String `tfsdk:"ike_group" vyos:"ike-group,omitempty"`
	LeafVpnIPsecSiteToSitePeerIkevtwoReauth         types.String `tfsdk:"ikev2_reauth" vyos:"ikev2-reauth,omitempty"`
	LeafVpnIPsecSiteToSitePeerLocalAddress          types.String `tfsdk:"local_address" vyos:"local-address,omitempty"`
	LeafVpnIPsecSiteToSitePeerRemoteAddress         types.List   `tfsdk:"remote_address" vyos:"remote-address,omitempty"`
	LeafVpnIPsecSiteToSitePeerReplayWindow          types.Number `tfsdk:"replay_window" vyos:"replay-window,omitempty"`
	LeafVpnIPsecSiteToSitePeerVirtualAddress        types.List   `tfsdk:"virtual_address" vyos:"virtual-address,omitempty"`

	// TagNodes

	ExistsTagVpnIPsecSiteToSitePeerTunnel bool `tfsdk:"-" vyos:"tunnel,child"`

	// Nodes

	NodeVpnIPsecSiteToSitePeerAuthentication *VpnIPsecSiteToSitePeerAuthentication `tfsdk:"authentication" vyos:"authentication,omitempty"`

	NodeVpnIPsecSiteToSitePeerVti *VpnIPsecSiteToSitePeerVti `tfsdk:"vti" vyos:"vti,omitempty"`
}

// SetID configures the resource ID
func (o *VpnIPsecSiteToSitePeer) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnIPsecSiteToSitePeer) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnIPsecSiteToSitePeer) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnIPsecSiteToSitePeer) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"peer",
		o.SelfIdentifier.VpnIPsecSiteToSitePeer.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnIPsecSiteToSitePeer) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (site-to-site) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (ipsec) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (vpn) */
		"vpn", // Node

		"ipsec", // Node

		"site-to-site", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VpnIPsecSiteToSitePeer) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (site-to-site) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (ipsec) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (vpn) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecSiteToSitePeer) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"peer": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Connection name of the peer

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  txt     |  Connection name of the peer  |
`,
					Description: `Connection name of the peer

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  txt     |  Connection name of the peer  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in peer, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  peer, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (site-to-site) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (ipsec) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (vpn) */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"disable":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (disable) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"connection_type":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (connection-type) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Connection type

    |  Format    |  Description                                   |
    |------------|------------------------------------------------|
    |  initiate  |  Bring the connection up immediately           |
    |  respond   |  Wait for the peer to initiate the connection  |
    |  none      |  Load the connection only                      |
`,
			Description: `Connection type

    |  Format    |  Description                                   |
    |------------|------------------------------------------------|
    |  initiate  |  Bring the connection up immediately           |
    |  respond   |  Wait for the peer to initiate the connection  |
    |  none      |  Load the connection only                      |
`,
		},

		"default_esp_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (default-esp-group) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Defult ESP group name

`,
			Description: `Defult ESP group name

`,
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

		"dhcp_interface":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (dhcp-interface) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DHCP interface supplying next-hop IP address

    |  Format  |  Description          |
    |----------|-----------------------|
    |  txt     |  DHCP interface name  |
`,
			Description: `DHCP interface supplying next-hop IP address

    |  Format  |  Description          |
    |----------|-----------------------|
    |  txt     |  DHCP interface name  |
`,
		},

		"force_udp_encapsulation":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (force-udp-encapsulation) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Force UDP encapsulation

`,
			Description: `Force UDP encapsulation

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ike_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ike-group) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Internet Key Exchange (IKE) group name

`,
			Description: `Internet Key Exchange (IKE) group name

`,
		},

		"ikev2_reauth":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ikev2-reauth) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Re-authentication of the remote peer during an IKE re-key (IKEv2 only)

    |  Format   |  Description                                                                                          |
    |-----------|-------------------------------------------------------------------------------------------------------|
    |  yes      |  Enable remote host re-autentication during an IKE re-key. Currently broken due to a strong swan bug  |
    |  no       |  Disable remote host re-authenticaton during an IKE re-key.                                           |
    |  inherit  |  Inherit the reauth configuration form your IKE-group                                                 |
`,
			Description: `Re-authentication of the remote peer during an IKE re-key (IKEv2 only)

    |  Format   |  Description                                                                                          |
    |-----------|-------------------------------------------------------------------------------------------------------|
    |  yes      |  Enable remote host re-autentication during an IKE re-key. Currently broken due to a strong swan bug  |
    |  no       |  Disable remote host re-authenticaton during an IKE re-key.                                           |
    |  inherit  |  Inherit the reauth configuration form your IKE-group                                                 |
`,
		},

		"local_address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (local-address) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 or IPv6 address of a local interface to use for VPN

    |  Format  |  Description                                                      |
    |----------|-------------------------------------------------------------------|
    |  ipv4    |  IPv4 address of a local interface for VPN                        |
    |  ipv6    |  IPv6 address of a local interface for VPN                        |
    |  any     |  Allow any IPv4 address present on the system to be used for VPN  |
`,
			Description: `IPv4 or IPv6 address of a local interface to use for VPN

    |  Format  |  Description                                                      |
    |----------|-------------------------------------------------------------------|
    |  ipv4    |  IPv4 address of a local interface for VPN                        |
    |  ipv6    |  IPv6 address of a local interface for VPN                        |
    |  any     |  Allow any IPv4 address present on the system to be used for VPN  |
`,
		},

		"remote_address":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (remote-address) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IPv4 or IPv6 address of the remote peer

    |  Format    |  Description                                     |
    |------------|--------------------------------------------------|
    |  ipv4      |  IPv4 address of the remote peer                 |
    |  ipv6      |  IPv6 address of the remote peer                 |
    |  hostname  |  Fully qualified domain name of the remote peer  |
    |  any       |  Allow any IP address of the remote peer         |
`,
			Description: `IPv4 or IPv6 address of the remote peer

    |  Format    |  Description                                     |
    |------------|--------------------------------------------------|
    |  ipv4      |  IPv4 address of the remote peer                 |
    |  ipv6      |  IPv6 address of the remote peer                 |
    |  hostname  |  Fully qualified domain name of the remote peer  |
    |  any       |  Allow any IP address of the remote peer         |
`,
		},

		"replay_window":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (replay-window) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `IPsec replay window to configure for this CHILD_SA

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  0       |  Disable IPsec replay protection  |
    |  1-2040  |  Replay window size in packets    |
`,
			Description: `IPsec replay window to configure for this CHILD_SA

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  0       |  Disable IPsec replay protection  |
    |  1-2040  |  Replay window size in packets    |
`,

			// Default:          stringdefault.StaticString(`32`),
			Computed: true,
		},

		"virtual_address":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (virtual-address) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Initiator request virtual-address from peer

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  ipv4    |  Request IPv4 address from peer  |
    |  ipv6    |  Request IPv6 address from peer  |
`,
			Description: `Initiator request virtual-address from peer

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  ipv4    |  Request IPv4 address from peer  |
    |  ipv6    |  Request IPv6 address from peer  |
`,
		},

		// TagNodes

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: VpnIPsecSiteToSitePeerAuthentication{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Peer authentication

`,
			Description: `Peer authentication

`,
		},

		"vti": schema.SingleNestedAttribute{
			Attributes: VpnIPsecSiteToSitePeerVti{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Virtual tunnel interface

`,
			Description: `Virtual tunnel interface

`,
		},
	}
}
