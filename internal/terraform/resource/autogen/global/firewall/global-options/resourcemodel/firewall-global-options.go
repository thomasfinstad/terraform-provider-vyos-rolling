/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &FirewallGlobalOptions{}

// FirewallGlobalOptions describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type FirewallGlobalOptions struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafFirewallGlobalOptionsAllPing                types.String `tfsdk:"all_ping" vyos:"all-ping,omitempty"`
	LeafFirewallGlobalOptionsBroadcastPing          types.String `tfsdk:"broadcast_ping" vyos:"broadcast-ping,omitempty"`
	LeafFirewallGlobalOptionsDirectedBroadcast      types.String `tfsdk:"directed_broadcast" vyos:"directed-broadcast,omitempty"`
	LeafFirewallGlobalOptionsIPSrcRoute             types.String `tfsdk:"ip_src_route" vyos:"ip-src-route,omitempty"`
	LeafFirewallGlobalOptionsLogMartians            types.String `tfsdk:"log_martians" vyos:"log-martians,omitempty"`
	LeafFirewallGlobalOptionsReceiveRedirects       types.String `tfsdk:"receive_redirects" vyos:"receive-redirects,omitempty"`
	LeafFirewallGlobalOptionsResolverCache          types.Bool   `tfsdk:"resolver_cache" vyos:"resolver-cache,omitempty"`
	LeafFirewallGlobalOptionsResolverInterval       types.Number `tfsdk:"resolver_interval" vyos:"resolver-interval,omitempty"`
	LeafFirewallGlobalOptionsSendRedirects          types.String `tfsdk:"send_redirects" vyos:"send-redirects,omitempty"`
	LeafFirewallGlobalOptionsSourceValIDation       types.String `tfsdk:"source_validation" vyos:"source-validation,omitempty"`
	LeafFirewallGlobalOptionsSynCookies             types.String `tfsdk:"syn_cookies" vyos:"syn-cookies,omitempty"`
	LeafFirewallGlobalOptionsTwaHazardsProtection   types.String `tfsdk:"twa_hazards_protection" vyos:"twa-hazards-protection,omitempty"`
	LeafFirewallGlobalOptionsIPvsixReceiveRedirects types.String `tfsdk:"ipv6_receive_redirects" vyos:"ipv6-receive-redirects,omitempty"`
	LeafFirewallGlobalOptionsIPvsixSourceValIDation types.String `tfsdk:"ipv6_source_validation" vyos:"ipv6-source-validation,omitempty"`
	LeafFirewallGlobalOptionsIPvsixSrcRoute         types.String `tfsdk:"ipv6_src_route" vyos:"ipv6-src-route,omitempty"`

	// TagNodes

	// Nodes

	ExistsNodeFirewallGlobalOptionsApplyToBrIDgedTraffic bool `tfsdk:"-" vyos:"apply-to-bridged-traffic,child"`

	NodeFirewallGlobalOptionsStatePolicy *FirewallGlobalOptionsStatePolicy `tfsdk:"state_policy" vyos:"state-policy,omitempty"`

	ExistsNodeFirewallGlobalOptionsTimeout bool `tfsdk:"-" vyos:"timeout,child"`
}

// SetID configures the resource ID
func (o *FirewallGlobalOptions) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *FirewallGlobalOptions) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *FirewallGlobalOptions) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallGlobalOptions) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"global-options",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *FirewallGlobalOptions) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"firewall", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *FirewallGlobalOptions) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallGlobalOptions) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"all_ping":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for handling of all IPv4 ICMP echo requests

    |  Format   |  Description                                        |
    |-----------|-----------------------------------------------------|
    |  enable   |  Enable processing of all IPv4 ICMP echo requests   |
    |  disable  |  Disable processing of all IPv4 ICMP echo requests  |
`,
			Description: `Policy for handling of all IPv4 ICMP echo requests

    |  Format   |  Description                                        |
    |-----------|-----------------------------------------------------|
    |  enable   |  Enable processing of all IPv4 ICMP echo requests   |
    |  disable  |  Disable processing of all IPv4 ICMP echo requests  |
`,

			// Default:          stringdefault.StaticString(`enable`),
			Computed: true,
		},

		"broadcast_ping":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for handling broadcast IPv4 ICMP echo and timestamp requests

    |  Format   |  Description                                                        |
    |-----------|---------------------------------------------------------------------|
    |  enable   |  Enable processing of broadcast IPv4 ICMP echo/timestamp requests   |
    |  disable  |  Disable processing of broadcast IPv4 ICMP echo/timestamp requests  |
`,
			Description: `Policy for handling broadcast IPv4 ICMP echo and timestamp requests

    |  Format   |  Description                                                        |
    |-----------|---------------------------------------------------------------------|
    |  enable   |  Enable processing of broadcast IPv4 ICMP echo/timestamp requests   |
    |  disable  |  Disable processing of broadcast IPv4 ICMP echo/timestamp requests  |
`,

			// Default:          stringdefault.StaticString(`disable`),
			Computed: true,
		},

		"directed_broadcast":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for handling IPv4 directed broadcast forwarding on all interfaces

    |  Format   |  Description                                                   |
    |-----------|----------------------------------------------------------------|
    |  enable   |  Enable IPv4 directed broadcast forwarding on all interfaces   |
    |  disable  |  Disable IPv4 directed broadcast forwarding on all interfaces  |
`,
			Description: `Policy for handling IPv4 directed broadcast forwarding on all interfaces

    |  Format   |  Description                                                   |
    |-----------|----------------------------------------------------------------|
    |  enable   |  Enable IPv4 directed broadcast forwarding on all interfaces   |
    |  disable  |  Disable IPv4 directed broadcast forwarding on all interfaces  |
`,

			// Default:          stringdefault.StaticString(`enable`),
			Computed: true,
		},

		"ip_src_route":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for handling IPv4 packets with source route option

    |  Format   |  Description                                                  |
    |-----------|---------------------------------------------------------------|
    |  enable   |  Enable processing of IPv4 packets with source route option   |
    |  disable  |  Disable processing of IPv4 packets with source route option  |
`,
			Description: `Policy for handling IPv4 packets with source route option

    |  Format   |  Description                                                  |
    |-----------|---------------------------------------------------------------|
    |  enable   |  Enable processing of IPv4 packets with source route option   |
    |  disable  |  Disable processing of IPv4 packets with source route option  |
`,

			// Default:          stringdefault.StaticString(`disable`),
			Computed: true,
		},

		"log_martians":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for logging IPv4 packets with invalid addresses

    |  Format   |  Description                                             |
    |-----------|----------------------------------------------------------|
    |  enable   |  Enable logging of IPv4 packets with invalid addresses   |
    |  disable  |  Disable logging of Ipv4 packets with invalid addresses  |
`,
			Description: `Policy for logging IPv4 packets with invalid addresses

    |  Format   |  Description                                             |
    |-----------|----------------------------------------------------------|
    |  enable   |  Enable logging of IPv4 packets with invalid addresses   |
    |  disable  |  Disable logging of Ipv4 packets with invalid addresses  |
`,

			// Default:          stringdefault.StaticString(`enable`),
			Computed: true,
		},

		"receive_redirects":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for handling received IPv4 ICMP redirect messages

    |  Format   |  Description                                                 |
    |-----------|--------------------------------------------------------------|
    |  enable   |  Enable processing of received IPv4 ICMP redirect messages   |
    |  disable  |  Disable processing of received IPv4 ICMP redirect messages  |
`,
			Description: `Policy for handling received IPv4 ICMP redirect messages

    |  Format   |  Description                                                 |
    |-----------|--------------------------------------------------------------|
    |  enable   |  Enable processing of received IPv4 ICMP redirect messages   |
    |  disable  |  Disable processing of received IPv4 ICMP redirect messages  |
`,

			// Default:          stringdefault.StaticString(`disable`),
			Computed: true,
		},

		"resolver_cache":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Retains last successful value if domain resolution fails

`,
			Description: `Retains last successful value if domain resolution fails

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"resolver_interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Domain resolver update interval

    |  Format   |  Description         |
    |-----------|----------------------|
    |  10-3600  |  Interval (seconds)  |
`,
			Description: `Domain resolver update interval

    |  Format   |  Description         |
    |-----------|----------------------|
    |  10-3600  |  Interval (seconds)  |
`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"send_redirects":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for sending IPv4 ICMP redirect messages

    |  Format   |  Description                                  |
    |-----------|-----------------------------------------------|
    |  enable   |  Enable sending IPv4 ICMP redirect messages   |
    |  disable  |  Disable sending IPv4 ICMP redirect messages  |
`,
			Description: `Policy for sending IPv4 ICMP redirect messages

    |  Format   |  Description                                  |
    |-----------|-----------------------------------------------|
    |  enable   |  Enable sending IPv4 ICMP redirect messages   |
    |  disable  |  Disable sending IPv4 ICMP redirect messages  |
`,

			// Default:          stringdefault.StaticString(`enable`),
			Computed: true,
		},

		"source_validation":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for IPv4 source validation by reversed path, as specified in RFC3704

    |  Format   |  Description                                                       |
    |-----------|--------------------------------------------------------------------|
    |  strict   |  Enable IPv4 Strict Reverse Path Forwarding as defined in RFC3704  |
    |  loose    |  Enable IPv4 Loose Reverse Path Forwarding as defined in RFC3704   |
    |  disable  |  No IPv4 source validation                                         |
`,
			Description: `Policy for IPv4 source validation by reversed path, as specified in RFC3704

    |  Format   |  Description                                                       |
    |-----------|--------------------------------------------------------------------|
    |  strict   |  Enable IPv4 Strict Reverse Path Forwarding as defined in RFC3704  |
    |  loose    |  Enable IPv4 Loose Reverse Path Forwarding as defined in RFC3704   |
    |  disable  |  No IPv4 source validation                                         |
`,

			// Default:          stringdefault.StaticString(`disable`),
			Computed: true,
		},

		"syn_cookies":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for using TCP SYN cookies with IPv4

    |  Format   |  Description                               |
    |-----------|--------------------------------------------|
    |  enable   |  Enable use of TCP SYN cookies with IPv4   |
    |  disable  |  Disable use of TCP SYN cookies with IPv4  |
`,
			Description: `Policy for using TCP SYN cookies with IPv4

    |  Format   |  Description                               |
    |-----------|--------------------------------------------|
    |  enable   |  Enable use of TCP SYN cookies with IPv4   |
    |  disable  |  Disable use of TCP SYN cookies with IPv4  |
`,

			// Default:          stringdefault.StaticString(`enable`),
			Computed: true,
		},

		"twa_hazards_protection":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `RFC1337 TCP TIME-WAIT assasination hazards protection

    |  Format   |  Description                                   |
    |-----------|------------------------------------------------|
    |  enable   |  Enable RFC1337 TIME-WAIT hazards protection   |
    |  disable  |  Disable RFC1337 TIME-WAIT hazards protection  |
`,
			Description: `RFC1337 TCP TIME-WAIT assasination hazards protection

    |  Format   |  Description                                   |
    |-----------|------------------------------------------------|
    |  enable   |  Enable RFC1337 TIME-WAIT hazards protection   |
    |  disable  |  Disable RFC1337 TIME-WAIT hazards protection  |
`,

			// Default:          stringdefault.StaticString(`disable`),
			Computed: true,
		},

		"ipv6_receive_redirects":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for handling received ICMPv6 redirect messages

    |  Format   |  Description                                              |
    |-----------|-----------------------------------------------------------|
    |  enable   |  Enable processing of received ICMPv6 redirect messages   |
    |  disable  |  Disable processing of received ICMPv6 redirect messages  |
`,
			Description: `Policy for handling received ICMPv6 redirect messages

    |  Format   |  Description                                              |
    |-----------|-----------------------------------------------------------|
    |  enable   |  Enable processing of received ICMPv6 redirect messages   |
    |  disable  |  Disable processing of received ICMPv6 redirect messages  |
`,

			// Default:          stringdefault.StaticString(`disable`),
			Computed: true,
		},

		"ipv6_source_validation":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for IPv6 source validation by reversed path, as specified in RFC3704

    |  Format   |  Description                                                       |
    |-----------|--------------------------------------------------------------------|
    |  strict   |  Enable IPv6 Strict Reverse Path Forwarding as defined in RFC3704  |
    |  loose    |  Enable IPv6 Loose Reverse Path Forwarding as defined in RFC3704   |
    |  disable  |  No IPv6 source validation                                         |
`,
			Description: `Policy for IPv6 source validation by reversed path, as specified in RFC3704

    |  Format   |  Description                                                       |
    |-----------|--------------------------------------------------------------------|
    |  strict   |  Enable IPv6 Strict Reverse Path Forwarding as defined in RFC3704  |
    |  loose    |  Enable IPv6 Loose Reverse Path Forwarding as defined in RFC3704   |
    |  disable  |  No IPv6 source validation                                         |
`,

			// Default:          stringdefault.StaticString(`disable`),
			Computed: true,
		},

		"ipv6_src_route":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for handling IPv6 packets with routing extension header

    |  Format   |  Description                                                   |
    |-----------|----------------------------------------------------------------|
    |  enable   |  Enable processing of IPv6 packets with routing header type 2  |
    |  disable  |  Disable processing of IPv6 packets with routing header        |
`,
			Description: `Policy for handling IPv6 packets with routing extension header

    |  Format   |  Description                                                   |
    |-----------|----------------------------------------------------------------|
    |  enable   |  Enable processing of IPv6 packets with routing header type 2  |
    |  disable  |  Disable processing of IPv6 packets with routing header        |
`,

			// Default:          stringdefault.StaticString(`disable`),
			Computed: true,
		},

		// TagNodes

		// Nodes

		"state_policy": schema.SingleNestedAttribute{
			Attributes: FirewallGlobalOptionsStatePolicy{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Global firewall state-policy

`,
			Description: `Global firewall state-policy

`,
		},
	}
}
