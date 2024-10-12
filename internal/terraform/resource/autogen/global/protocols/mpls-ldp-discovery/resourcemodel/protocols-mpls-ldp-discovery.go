// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance
var _ helpers.VyosTopResourceDataModel = &ProtocolsMplsLdpDiscovery{}

// ProtocolsMplsLdpDiscovery describes the resource data model.
type ProtocolsMplsLdpDiscovery struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsMplsLdpDiscoveryHelloIPvfourHoldtime    types.Number `tfsdk:"hello_ipv4_holdtime" vyos:"hello-ipv4-holdtime,omitempty"`
	LeafProtocolsMplsLdpDiscoveryHelloIPvfourInterval    types.Number `tfsdk:"hello_ipv4_interval" vyos:"hello-ipv4-interval,omitempty"`
	LeafProtocolsMplsLdpDiscoveryHelloIPvsixHoldtime     types.Number `tfsdk:"hello_ipv6_holdtime" vyos:"hello-ipv6-holdtime,omitempty"`
	LeafProtocolsMplsLdpDiscoveryHelloIPvsixInterval     types.Number `tfsdk:"hello_ipv6_interval" vyos:"hello-ipv6-interval,omitempty"`
	LeafProtocolsMplsLdpDiscoverySessionIPvfourHoldtime  types.Number `tfsdk:"session_ipv4_holdtime" vyos:"session-ipv4-holdtime,omitempty"`
	LeafProtocolsMplsLdpDiscoverySessionIPvsixHoldtime   types.Number `tfsdk:"session_ipv6_holdtime" vyos:"session-ipv6-holdtime,omitempty"`
	LeafProtocolsMplsLdpDiscoveryTransportIPvfourAddress types.String `tfsdk:"transport_ipv4_address" vyos:"transport-ipv4-address,omitempty"`
	LeafProtocolsMplsLdpDiscoveryTransportIPvsixAddress  types.String `tfsdk:"transport_ipv6_address" vyos:"transport-ipv6-address,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsMplsLdpDiscovery) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsMplsLdpDiscovery) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsMplsLdpDiscovery) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsMplsLdpDiscovery) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"discovery",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsMplsLdpDiscovery) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"mpls",

		"ldp",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ProtocolsMplsLdpDiscovery) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsMplsLdpDiscovery) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"hello_ipv4_holdtime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hello IPv4 hold time

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Time in seconds  |
`,
			Description: `Hello IPv4 hold time

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Time in seconds  |
`,
		},

		"hello_ipv4_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hello IPv4 interval

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Time in seconds  |
`,
			Description: `Hello IPv4 interval

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Time in seconds  |
`,
		},

		"hello_ipv6_holdtime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hello IPv6 hold time

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Time in seconds  |
`,
			Description: `Hello IPv6 hold time

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Time in seconds  |
`,
		},

		"hello_ipv6_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hello IPv6 interval

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Time in seconds  |
`,
			Description: `Hello IPv6 interval

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Time in seconds  |
`,
		},

		"session_ipv4_holdtime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Session IPv4 hold time

    |  Format    |  Description      |
    |------------|-------------------|
    |  15-65535  |  Time in seconds  |
`,
			Description: `Session IPv4 hold time

    |  Format    |  Description      |
    |------------|-------------------|
    |  15-65535  |  Time in seconds  |
`,
		},

		"session_ipv6_holdtime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Session IPv6 hold time

    |  Format    |  Description      |
    |------------|-------------------|
    |  15-65535  |  Time in seconds  |
`,
			Description: `Session IPv6 hold time

    |  Format    |  Description      |
    |------------|-------------------|
    |  15-65535  |  Time in seconds  |
`,
		},

		"transport_ipv4_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Transport IPv4 address

    |  Format  |  Description             |
    |----------|--------------------------|
    |  ipv4    |  IPv4 bind as transport  |
`,
			Description: `Transport IPv4 address

    |  Format  |  Description             |
    |----------|--------------------------|
    |  ipv4    |  IPv4 bind as transport  |
`,
		},

		"transport_ipv6_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Transport IPv6 address

    |  Format  |  Description             |
    |----------|--------------------------|
    |  ipv6    |  IPv6 bind as transport  |
`,
			Description: `Transport IPv6 address

    |  Format  |  Description             |
    |----------|--------------------------|
    |  ipv6    |  IPv6 bind as transport  |
`,
		},
	}
}
