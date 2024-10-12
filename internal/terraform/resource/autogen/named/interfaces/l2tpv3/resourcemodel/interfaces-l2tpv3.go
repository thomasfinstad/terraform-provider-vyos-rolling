// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

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

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &InterfacesLtwotpvthree{}

// InterfacesLtwotpvthree describes the resource data model.
type InterfacesLtwotpvthree struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafInterfacesLtwotpvthreeAddress         types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesLtwotpvthreeDescrIPtion     types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesLtwotpvthreeDestinationPort types.Number `tfsdk:"destination_port" vyos:"destination-port,omitempty"`
	LeafInterfacesLtwotpvthreeDisable         types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesLtwotpvthreeEncapsulation   types.String `tfsdk:"encapsulation" vyos:"encapsulation,omitempty"`
	LeafInterfacesLtwotpvthreeSourceAddress   types.String `tfsdk:"source_address" vyos:"source-address,omitempty"`
	LeafInterfacesLtwotpvthreeMtu             types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesLtwotpvthreePeerSessionID   types.Number `tfsdk:"peer_session_id" vyos:"peer-session-id,omitempty"`
	LeafInterfacesLtwotpvthreePeerTunnelID    types.Number `tfsdk:"peer_tunnel_id" vyos:"peer-tunnel-id,omitempty"`
	LeafInterfacesLtwotpvthreeRemote          types.String `tfsdk:"remote" vyos:"remote,omitempty"`
	LeafInterfacesLtwotpvthreeSessionID       types.Number `tfsdk:"session_id" vyos:"session-id,omitempty"`
	LeafInterfacesLtwotpvthreeSourcePort      types.Number `tfsdk:"source_port" vyos:"source-port,omitempty"`
	LeafInterfacesLtwotpvthreeTunnelID        types.Number `tfsdk:"tunnel_id" vyos:"tunnel-id,omitempty"`
	LeafInterfacesLtwotpvthreeVrf             types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeInterfacesLtwotpvthreeIP     *InterfacesLtwotpvthreeIP     `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesLtwotpvthreeIPvsix *InterfacesLtwotpvthreeIPvsix `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
	NodeInterfacesLtwotpvthreeMirror *InterfacesLtwotpvthreeMirror `tfsdk:"mirror" vyos:"mirror,omitempty"`
}

// SetID configures the resource ID
func (o *InterfacesLtwotpvthree) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *InterfacesLtwotpvthree) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *InterfacesLtwotpvthree) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesLtwotpvthree) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"l2tpv3",
		o.SelfIdentifier.Attributes()["l2tpv3"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *InterfacesLtwotpvthree) GetVyosParentPath() []string {
	return []string{
		"interfaces",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *InterfacesLtwotpvthree) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesLtwotpvthree) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"l2tpv3": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Layer 2 Tunnel Protocol Version 3 (L2TPv3) Interface

    |  Format    |  Description            |
    |------------|-------------------------|
    |  l2tpethN  |  L2TPv3 interface name  |
`,
					Description: `Layer 2 Tunnel Protocol Version 3 (L2TPv3) Interface

    |  Format    |  Description            |
    |------------|-------------------------|
    |  l2tpethN  |  L2TPv3 interface name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in l2tpv3, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  l2tpv3, value must match: ^[a-zA-Z0-9-_]*$",
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

		"address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IPv4 address and prefix length  |
    |  ipv6net  |  IPv6 address and prefix length  |
`,
			Description: `IP address

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IPv4 address and prefix length  |
    |  ipv6net  |  IPv6 address and prefix length  |
`,
		},

		"description": schema.StringAttribute{
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

		"destination_port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `UDP destination port for L2TPv3 tunnel

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
			Description: `UDP destination port for L2TPv3 tunnel

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,

			// Default:          stringdefault.StaticString(`5000`),
			Computed: true,
		},

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administratively disable interface

`,
			Description: `Administratively disable interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"encapsulation": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Encapsulation type

    |  Format  |  Description        |
    |----------|---------------------|
    |  udp     |  UDP encapsulation  |
    |  ip      |  IP encapsulation   |
`,
			Description: `Encapsulation type

    |  Format  |  Description        |
    |----------|---------------------|
    |  udp     |  UDP encapsulation  |
    |  ip      |  IP encapsulation   |
`,

			// Default:          stringdefault.StaticString(`udp`),
			Computed: true,
		},

		"source_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source IP address used to initiate connection

    |  Format  |  Description          |
    |----------|-----------------------|
    |  ipv4    |  IPv4 source address  |
    |  ipv6    |  IPv6 source address  |
`,
			Description: `Source IP address used to initiate connection

    |  Format  |  Description          |
    |----------|-----------------------|
    |  ipv4    |  IPv4 source address  |
    |  ipv6    |  IPv6 source address  |
`,
		},

		"mtu": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

    |  Format    |  Description                        |
    |------------|-------------------------------------|
    |  68-16000  |  Maximum Transmission Unit in byte  |
`,
			Description: `Maximum Transmission Unit (MTU)

    |  Format    |  Description                        |
    |------------|-------------------------------------|
    |  68-16000  |  Maximum Transmission Unit in byte  |
`,

			// Default:          stringdefault.StaticString(`1488`),
			Computed: true,
		},

		"peer_session_id": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Peer session identifier

    |  Format       |  Description                     |
    |---------------|----------------------------------|
    |  1-429496729  |  L2TPv3 peer session identifier  |
`,
			Description: `Peer session identifier

    |  Format       |  Description                     |
    |---------------|----------------------------------|
    |  1-429496729  |  L2TPv3 peer session identifier  |
`,
		},

		"peer_tunnel_id": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Peer tunnel identifier

    |  Format       |  Description                    |
    |---------------|---------------------------------|
    |  1-429496729  |  L2TPv3 peer tunnel identifier  |
`,
			Description: `Peer tunnel identifier

    |  Format       |  Description                    |
    |---------------|---------------------------------|
    |  1-429496729  |  L2TPv3 peer tunnel identifier  |
`,
		},

		"remote": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Tunnel remote address

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  ipv4    |  Tunnel remote IPv4 address  |
    |  ipv6    |  Tunnel remote IPv6 address  |
`,
			Description: `Tunnel remote address

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  ipv4    |  Tunnel remote IPv4 address  |
    |  ipv6    |  Tunnel remote IPv6 address  |
`,
		},

		"session_id": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Session identifier

    |  Format       |  Description                |
    |---------------|-----------------------------|
    |  1-429496729  |  L2TPv3 session identifier  |
`,
			Description: `Session identifier

    |  Format       |  Description                |
    |---------------|-----------------------------|
    |  1-429496729  |  L2TPv3 session identifier  |
`,
		},

		"source_port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `UDP source port for L2TPv3 tunnel

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
			Description: `UDP source port for L2TPv3 tunnel

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,

			// Default:          stringdefault.StaticString(`5000`),
			Computed: true,
		},

		"tunnel_id": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Local tunnel identifier

    |  Format       |  Description                     |
    |---------------|----------------------------------|
    |  1-429496729  |  L2TPv3 local tunnel identifier  |
`,
			Description: `Local tunnel identifier

    |  Format       |  Description                     |
    |---------------|----------------------------------|
    |  1-429496729  |  L2TPv3 local tunnel identifier  |
`,
		},

		"vrf": schema.StringAttribute{
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

		// Nodes

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesLtwotpvthreeIP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
			Description: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesLtwotpvthreeIPvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
			Description: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesLtwotpvthreeMirror{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
			Description: `Mirror ingress/egress packets

`,
		},
	}
}
