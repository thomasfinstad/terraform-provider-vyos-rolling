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

var _ helpers.VyosTopResourceDataModel = &InterfacesWireguardPeer{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (peer) */
// InterfacesWireguardPeerSelfIdentifier used by TagNodes to keep the id info
type InterfacesWireguardPeerSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (peer) */

	InterfacesWireguardPeer types.String `tfsdk:"peer" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (wireguard) */

	InterfacesWireguard types.String `tfsdk:"wireguard" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (interfaces) */
}

// InterfacesWireguardPeer describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type InterfacesWireguardPeer struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *InterfacesWireguardPeerSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafInterfacesWireguardPeerDisable             types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesWireguardPeerDescrIPtion         types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesWireguardPeerPublicKey           types.String `tfsdk:"public_key" vyos:"public-key,omitempty"`
	LeafInterfacesWireguardPeerPresharedKey        types.String `tfsdk:"preshared_key" vyos:"preshared-key,omitempty"`
	LeafInterfacesWireguardPeerAllowedIPs          types.List   `tfsdk:"allowed_ips" vyos:"allowed-ips,omitempty"`
	LeafInterfacesWireguardPeerAddress             types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesWireguardPeerHostName            types.String `tfsdk:"host_name" vyos:"host-name,omitempty"`
	LeafInterfacesWireguardPeerPort                types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafInterfacesWireguardPeerPersistentKeepalive types.Number `tfsdk:"persistent_keepalive" vyos:"persistent-keepalive,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *InterfacesWireguardPeer) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *InterfacesWireguardPeer) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *InterfacesWireguardPeer) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesWireguardPeer) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"peer",
		o.SelfIdentifier.InterfacesWireguardPeer.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *InterfacesWireguardPeer) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (wireguard) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (interfaces) */
		"interfaces", // Node

		"wireguard",
		o.SelfIdentifier.InterfacesWireguard.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *InterfacesWireguardPeer) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (wireguard) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (wireguard) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (interfaces) */
		"interfaces", // Node

		"wireguard",
		o.SelfIdentifier.InterfacesWireguard.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWireguardPeer) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
					MarkdownDescription: `peer alias

`,
					Description: `peer alias

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

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (wireguard) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (interfaces) */

				"wireguard": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `WireGuard Interface

    |  Format  |  Description               |
    |----------|----------------------------|
    |  wgN     |  WireGuard interface name  |
`,
					Description: `WireGuard Interface

    |  Format  |  Description               |
    |----------|----------------------------|
    |  wgN     |  WireGuard interface name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in wireguard, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  wireguard, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},
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

		"public_key":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (public-key) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `base64 encoded public key

`,
			Description: `base64 encoded public key

`,
		},

		"preshared_key":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (preshared-key) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `base64 encoded preshared key

`,
			Description: `base64 encoded preshared key

`,
		},

		"allowed_ips":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (allowed-ips) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP addresses allowed to traverse the peer

`,
			Description: `IP addresses allowed to traverse the peer

`,
		},

		"address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (address) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address of tunnel endpoint

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  ipv4    |  IPv4 address of remote tunnel endpoint  |
    |  ipv6    |  IPv6 address of remote tunnel endpoint  |
`,
			Description: `IP address of tunnel endpoint

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  ipv4    |  IPv4 address of remote tunnel endpoint  |
    |  ipv6    |  IPv6 address of remote tunnel endpoint  |
`,
		},

		"host_name":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (host-name) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Hostname of tunnel endpoint

    |  Format    |  Description                 |
    |------------|------------------------------|
    |  hostname  |  FQDN of WireGuard endpoint  |
`,
			Description: `Hostname of tunnel endpoint

    |  Format    |  Description                 |
    |------------|------------------------------|
    |  hostname  |  FQDN of WireGuard endpoint  |
`,
		},

		"port":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (port) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
			Description: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
		},

		"persistent_keepalive":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (persistent-keepalive) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval to send keepalive messages

    |  Format   |  Description          |
    |-----------|-----------------------|
    |  1-65535  |  Interval in seconds  |
`,
			Description: `Interval to send keepalive messages

    |  Format   |  Description          |
    |-----------|-----------------------|
    |  1-65535  |  Interval in seconds  |
`,
		},

		// TagNodes

		// Nodes

	}
}
