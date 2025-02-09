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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (openvpn) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &InterfacesOpenvpn{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (openvpn) */
// InterfacesOpenvpnSelfIdentifier used by TagNodes to keep the id info
type InterfacesOpenvpnSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (openvpn) */

	InterfacesOpenvpn types.String `tfsdk:"openvpn" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (interfaces) */
}

// InterfacesOpenvpn describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type InterfacesOpenvpn struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *InterfacesOpenvpnSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafInterfacesOpenvpnDescrIPtion       types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesOpenvpnDeviceType        types.String `tfsdk:"device_type" vyos:"device-type,omitempty"`
	LeafInterfacesOpenvpnDisable           types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesOpenvpnHash              types.String `tfsdk:"hash" vyos:"hash,omitempty"`
	LeafInterfacesOpenvpnLocalHost         types.String `tfsdk:"local_host" vyos:"local-host,omitempty"`
	LeafInterfacesOpenvpnLocalPort         types.Number `tfsdk:"local_port" vyos:"local-port,omitempty"`
	LeafInterfacesOpenvpnMode              types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafInterfacesOpenvpnOpenvpnOption     types.List   `tfsdk:"openvpn_option" vyos:"openvpn-option,omitempty"`
	LeafInterfacesOpenvpnPersistentTunnel  types.Bool   `tfsdk:"persistent_tunnel" vyos:"persistent-tunnel,omitempty"`
	LeafInterfacesOpenvpnProtocol          types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafInterfacesOpenvpnIPVersion         types.String `tfsdk:"ip_version" vyos:"ip-version,omitempty"`
	LeafInterfacesOpenvpnRemoteAddress     types.List   `tfsdk:"remote_address" vyos:"remote-address,omitempty"`
	LeafInterfacesOpenvpnRemoteHost        types.List   `tfsdk:"remote_host" vyos:"remote-host,omitempty"`
	LeafInterfacesOpenvpnRemotePort        types.Number `tfsdk:"remote_port" vyos:"remote-port,omitempty"`
	LeafInterfacesOpenvpnSharedSecretKey   types.String `tfsdk:"shared_secret_key" vyos:"shared-secret-key,omitempty"`
	LeafInterfacesOpenvpnUseLzoCompression types.Bool   `tfsdk:"use_lzo_compression" vyos:"use-lzo-compression,omitempty"`
	LeafInterfacesOpenvpnRedirect          types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`
	LeafInterfacesOpenvpnVrf               types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes

	TagInterfacesOpenvpnLocalAddress map[string]*InterfacesOpenvpnLocalAddress `tfsdk:"local_address" vyos:"local-address,omitempty"`

	// Nodes

	NodeInterfacesOpenvpnAuthentication *InterfacesOpenvpnAuthentication `tfsdk:"authentication" vyos:"authentication,omitempty"`

	NodeInterfacesOpenvpnEncryption *InterfacesOpenvpnEncryption `tfsdk:"encryption" vyos:"encryption,omitempty"`

	NodeInterfacesOpenvpnIP *InterfacesOpenvpnIP `tfsdk:"ip" vyos:"ip,omitempty"`

	NodeInterfacesOpenvpnIPvsix *InterfacesOpenvpnIPvsix `tfsdk:"ipv6" vyos:"ipv6,omitempty"`

	NodeInterfacesOpenvpnMirror *InterfacesOpenvpnMirror `tfsdk:"mirror" vyos:"mirror,omitempty"`

	NodeInterfacesOpenvpnKeepAlive *InterfacesOpenvpnKeepAlive `tfsdk:"keep_alive" vyos:"keep-alive,omitempty"`

	NodeInterfacesOpenvpnOffload *InterfacesOpenvpnOffload `tfsdk:"offload" vyos:"offload,omitempty"`

	NodeInterfacesOpenvpnReplaceDefaultRoute *InterfacesOpenvpnReplaceDefaultRoute `tfsdk:"replace_default_route" vyos:"replace-default-route,omitempty"`

	NodeInterfacesOpenvpnServer *InterfacesOpenvpnServer `tfsdk:"server" vyos:"server,omitempty"`

	NodeInterfacesOpenvpnTLS *InterfacesOpenvpnTLS `tfsdk:"tls" vyos:"tls,omitempty"`
}

// SetID configures the resource ID
func (o *InterfacesOpenvpn) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *InterfacesOpenvpn) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *InterfacesOpenvpn) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesOpenvpn) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"openvpn",
		o.SelfIdentifier.InterfacesOpenvpn.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *InterfacesOpenvpn) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (interfaces) */
		"interfaces", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *InterfacesOpenvpn) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (interfaces) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpn) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"openvpn": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `OpenVPN Tunnel Interface

    |  Format  |  Description             |
    |----------|--------------------------|
    |  vtunN   |  OpenVPN interface name  |
`,
					Description: `OpenVPN Tunnel Interface

    |  Format  |  Description             |
    |----------|--------------------------|
    |  vtunN   |  OpenVPN interface name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in openvpn, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  openvpn, value must match: ^[.:a-zA-Z0-9-_/]+$",
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

		"device_type":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (device-type) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `OpenVPN interface device-type

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  tun     |  TUN device, required for OSI layer 3  |
    |  tap     |  TAP device, required for OSI layer 2  |
`,
			Description: `OpenVPN interface device-type

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  tun     |  TUN device, required for OSI layer 3  |
    |  tap     |  TAP device, required for OSI layer 2  |
`,

			// Default:          stringdefault.StaticString(`tun`),
			Computed: true,
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

		"hash":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (hash) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Hashing Algorithm

    |  Format  |  Description        |
    |----------|---------------------|
    |  md5     |  MD5 algorithm      |
    |  sha1    |  SHA-1 algorithm    |
    |  sha256  |  SHA-256 algorithm  |
    |  sha384  |  SHA-384 algorithm  |
    |  sha512  |  SHA-512 algorithm  |
`,
			Description: `Hashing Algorithm

    |  Format  |  Description        |
    |----------|---------------------|
    |  md5     |  MD5 algorithm      |
    |  sha1    |  SHA-1 algorithm    |
    |  sha256  |  SHA-256 algorithm  |
    |  sha384  |  SHA-384 algorithm  |
    |  sha512  |  SHA-512 algorithm  |
`,
		},

		"local_host":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (local-host) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local IP address to accept connections (all if not set)

    |  Format  |  Description         |
    |----------|----------------------|
    |  ipv4    |  Local IPv4 address  |
    |  ipv6    |  Local IPv6 address  |
`,
			Description: `Local IP address to accept connections (all if not set)

    |  Format  |  Description         |
    |----------|----------------------|
    |  ipv4    |  Local IPv4 address  |
    |  ipv6    |  Local IPv6 address  |
`,
		},

		"local_port":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (local-port) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Local port number to accept connections

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
			Description: `Local port number to accept connections

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
		},

		"mode":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mode) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `OpenVPN mode of operation

    |  Format        |  Description                   |
    |----------------|--------------------------------|
    |  site-to-site  |  Site-to-site mode             |
    |  client        |  Client in client-server mode  |
    |  server        |  Server in client-server mode  |
`,
			Description: `OpenVPN mode of operation

    |  Format        |  Description                   |
    |----------------|--------------------------------|
    |  site-to-site  |  Site-to-site mode             |
    |  client        |  Client in client-server mode  |
    |  server        |  Server in client-server mode  |
`,
		},

		"openvpn_option":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (openvpn-option) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Additional OpenVPN options. You must use the syntax of openvpn.conf in this text-field. Using this without proper knowledge may result in a crashed OpenVPN server. Check system log to look for errors.

`,
			Description: `Additional OpenVPN options. You must use the syntax of openvpn.conf in this text-field. Using this without proper knowledge may result in a crashed OpenVPN server. Check system log to look for errors.

`,
		},

		"persistent_tunnel":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (persistent-tunnel) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not close and reopen interface (TUN/TAP device) on client restarts

`,
			Description: `Do not close and reopen interface (TUN/TAP device) on client restarts

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"protocol":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (protocol) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `OpenVPN communication protocol

    |  Format       |  Description                             |
    |---------------|------------------------------------------|
    |  udp          |  UDP                                     |
    |  tcp-passive  |  TCP and accepts connections passively   |
    |  tcp-active   |  TCP and initiates connections actively  |
`,
			Description: `OpenVPN communication protocol

    |  Format       |  Description                             |
    |---------------|------------------------------------------|
    |  udp          |  UDP                                     |
    |  tcp-passive  |  TCP and accepts connections passively   |
    |  tcp-active   |  TCP and initiates connections actively  |
`,

			// Default:          stringdefault.StaticString(`udp`),
			Computed: true,
		},

		"ip_version":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ip-version) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Force OpenVPN to use a specific IP protocol version

    |  Format      |  Description                                                                          |
    |--------------|---------------------------------------------------------------------------------------|
    |  auto        |  Select one IP protocol to use based on local or remote host                          |
    |  _ipv4       |  Accept connections on or initate connections to IPv4 addresses only                  |
    |  _ipv6       |  Accept connections on or initate connections to IPv6 addresses only                  |
    |  dual-stack  |  Accept connections on both protocols simultaneously (only supported in server mode)  |
`,
			Description: `Force OpenVPN to use a specific IP protocol version

    |  Format      |  Description                                                                          |
    |--------------|---------------------------------------------------------------------------------------|
    |  auto        |  Select one IP protocol to use based on local or remote host                          |
    |  _ipv4       |  Accept connections on or initate connections to IPv4 addresses only                  |
    |  _ipv6       |  Accept connections on or initate connections to IPv6 addresses only                  |
    |  dual-stack  |  Accept connections on both protocols simultaneously (only supported in server mode)  |
`,

			// Default:          stringdefault.StaticString(`auto`),
			Computed: true,
		},

		"remote_address":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (remote-address) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address of remote end of tunnel

    |  Format  |  Description              |
    |----------|---------------------------|
    |  ipv4    |  Remote end IPv4 address  |
    |  ipv6    |  Remote end IPv6 address  |
`,
			Description: `IP address of remote end of tunnel

    |  Format  |  Description              |
    |----------|---------------------------|
    |  ipv4    |  Remote end IPv4 address  |
    |  ipv6    |  Remote end IPv6 address  |
`,
		},

		"remote_host":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (remote-host) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Remote host to connect to (dynamic if not set)

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  ipv4    |  IPv4 address of remote host  |
    |  ipv6    |  IPv6 address of remote host  |
    |  txt     |  Hostname of remote host      |
`,
			Description: `Remote host to connect to (dynamic if not set)

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  ipv4    |  IPv4 address of remote host  |
    |  ipv6    |  IPv6 address of remote host  |
    |  txt     |  Hostname of remote host      |
`,
		},

		"remote_port":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (remote-port) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Remote port number to connect to

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
			Description: `Remote port number to connect to

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
		},

		"shared_secret_key":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (shared-secret-key) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Secret key shared with remote end of tunnel

`,
			Description: `Secret key shared with remote end of tunnel

`,
		},

		"use_lzo_compression":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (use-lzo-compression) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use fast LZO compression on this TUN/TAP interface

`,
			Description: `Use fast LZO compression on this TUN/TAP interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
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

		"local_address": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: InterfacesOpenvpnLocalAddress{}.ResourceSchemaAttributes(ctx),
			},
			Optional: true,
			MarkdownDescription: `Local IP address of tunnel (IPv4 or IPv6)

`,
			Description: `Local IP address of tunnel (IPv4 or IPv6)

`,
		},

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnAuthentication{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Authentication settings

`,
			Description: `Authentication settings

`,
		},

		"encryption": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnEncryption{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Data Encryption settings

`,
			Description: `Data Encryption settings

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnIP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
			Description: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnIPvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
			Description: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnMirror{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
			Description: `Mirror ingress/egress packets

`,
		},

		"keep_alive": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnKeepAlive{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Keepalive helper options

`,
			Description: `Keepalive helper options

`,
		},

		"offload": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnOffload{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Configurable offload options

`,
			Description: `Configurable offload options

`,
		},

		"replace_default_route": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnReplaceDefaultRoute{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `OpenVPN tunnel to be used as the default route

`,
			Description: `OpenVPN tunnel to be used as the default route

`,
		},

		"server": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnServer{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Server-mode options

`,
			Description: `Server-mode options

`,
		},

		"tls": schema.SingleNestedAttribute{
			Attributes: InterfacesOpenvpnTLS{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Transport Layer Security (TLS) options

`,
			Description: `Transport Layer Security (TLS) options

`,
		},
	}
}
