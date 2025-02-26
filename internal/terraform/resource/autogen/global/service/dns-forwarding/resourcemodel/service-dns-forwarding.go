/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (forwarding) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceDNSForwarding{}

// ServiceDNSForwarding describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ServiceDNSForwarding struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceDNSForwardingCacheSize                 types.Number `tfsdk:"cache_size" vyos:"cache-size,omitempty"`
	LeafServiceDNSForwardingDhcp                      types.List   `tfsdk:"dhcp" vyos:"dhcp,omitempty"`
	LeafServiceDNSForwardingDNSsixfourPrefix          types.String `tfsdk:"dns64_prefix" vyos:"dns64-prefix,omitempty"`
	LeafServiceDNSForwardingDNSsec                    types.String `tfsdk:"dnssec" vyos:"dnssec,omitempty"`
	LeafServiceDNSForwardingIgnoreHostsFile           types.Bool   `tfsdk:"ignore_hosts_file" vyos:"ignore-hosts-file,omitempty"`
	LeafServiceDNSForwardingNoServeRfconenineoneeight types.Bool   `tfsdk:"no_serve_rfc1918" vyos:"no-serve-rfc1918,omitempty"`
	LeafServiceDNSForwardingAllowFrom                 types.List   `tfsdk:"allow_from" vyos:"allow-from,omitempty"`
	LeafServiceDNSForwardingListenAddress             types.List   `tfsdk:"listen_address" vyos:"listen-address,omitempty"`
	LeafServiceDNSForwardingPort                      types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafServiceDNSForwardingNegativeTTL               types.Number `tfsdk:"negative_ttl" vyos:"negative-ttl,omitempty"`
	LeafServiceDNSForwardingServeStaleExtension       types.Number `tfsdk:"serve_stale_extension" vyos:"serve-stale-extension,omitempty"`
	LeafServiceDNSForwardingTimeout                   types.Number `tfsdk:"timeout" vyos:"timeout,omitempty"`
	LeafServiceDNSForwardingSourceAddress             types.List   `tfsdk:"source_address" vyos:"source-address,omitempty"`
	LeafServiceDNSForwardingSystem                    types.Bool   `tfsdk:"system" vyos:"system,omitempty"`
	LeafServiceDNSForwardingExcludeThroTTLeAddress    types.List   `tfsdk:"exclude_throttle_address" vyos:"exclude-throttle-address,omitempty"`

	// TagNodes

	ExistsTagServiceDNSForwardingDomain bool `tfsdk:"-" vyos:"domain,child"`

	ExistsTagServiceDNSForwardingAuthoritativeDomain bool `tfsdk:"-" vyos:"authoritative-domain,child"`

	ExistsTagServiceDNSForwardingNameServer bool `tfsdk:"-" vyos:"name-server,child"`

	ExistsTagServiceDNSForwardingZoneCache bool `tfsdk:"-" vyos:"zone-cache,child"`

	// Nodes

	ExistsNodeServiceDNSForwardingOptions bool `tfsdk:"-" vyos:"options,child"`
}

// SetID configures the resource ID
func (o *ServiceDNSForwarding) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceDNSForwarding) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceDNSForwarding) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDNSForwarding) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"forwarding",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceDNSForwarding) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (dns) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (service) */
		"service", // Node

		"dns", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceDNSForwarding) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (dns) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (service) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwarding) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"cache_size":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (cache-size) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `DNS forwarding cache size

    |  Format        |  Description                |
    |----------------|-----------------------------|
    |  0-2147483647  |  DNS forwarding cache size  |
`,
			Description: `DNS forwarding cache size

    |  Format        |  Description                |
    |----------------|-----------------------------|
    |  0-2147483647  |  DNS forwarding cache size  |
`,

			// Default:          stringdefault.StaticString(`10000`),
			Computed: true,
		},

		"dhcp":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (dhcp) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Interfaces whose DHCP client nameservers to forward requests to

`,
			Description: `Interfaces whose DHCP client nameservers to forward requests to

`,
		},

		"dns64_prefix":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (dns64-prefix) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Help to communicate between IPv6-only client and IPv4-only server

    |  Format   |  Description                              |
    |-----------|-------------------------------------------|
    |  ipv6net  |  IPv6 address and /96 only prefix length  |
`,
			Description: `Help to communicate between IPv6-only client and IPv4-only server

    |  Format   |  Description                              |
    |-----------|-------------------------------------------|
    |  ipv6net  |  IPv6 address and /96 only prefix length  |
`,
		},

		"dnssec":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (dnssec) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DNSSEC mode

    |  Format               |  Description                                                                                      |
    |-----------------------|---------------------------------------------------------------------------------------------------|
    |  off                  |  No DNSSEC processing whatsoever!                                                                 |
    |  process-no-validate  |  Respond with DNSSEC records to clients that ask for it. No validation done at all!               |
    |  process              |  Respond with DNSSEC records to clients that ask for it. Validation for clients that request it.  |
    |  log-fail             |  Similar behaviour to process, but validate RRSIGs on responses and log bogus responses.          |
    |  validate             |  Full blown DNSSEC validation. Send SERVFAIL to clients on bogus responses.                       |
`,
			Description: `DNSSEC mode

    |  Format               |  Description                                                                                      |
    |-----------------------|---------------------------------------------------------------------------------------------------|
    |  off                  |  No DNSSEC processing whatsoever!                                                                 |
    |  process-no-validate  |  Respond with DNSSEC records to clients that ask for it. No validation done at all!               |
    |  process              |  Respond with DNSSEC records to clients that ask for it. Validation for clients that request it.  |
    |  log-fail             |  Similar behaviour to process, but validate RRSIGs on responses and log bogus responses.          |
    |  validate             |  Full blown DNSSEC validation. Send SERVFAIL to clients on bogus responses.                       |
`,

			// Default:          stringdefault.StaticString(`process-no-validate`),
			Computed: true,
		},

		"ignore_hosts_file":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ignore-hosts-file) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not use local /etc/hosts file in name resolution

`,
			Description: `Do not use local /etc/hosts file in name resolution

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_serve_rfc1918":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (no-serve-rfc1918) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Makes the server authoritatively not aware of RFC1918 addresses

`,
			Description: `Makes the server authoritatively not aware of RFC1918 addresses

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"allow_from":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (allow-from) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Networks allowed to query this server

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IP address and prefix length    |
    |  ipv6net  |  IPv6 address and prefix length  |
`,
			Description: `Networks allowed to query this server

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IP address and prefix length    |
    |  ipv6net  |  IPv6 address and prefix length  |
`,
		},

		"listen_address":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (listen-address) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Local IP addresses to listen on

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  ipv4    |  IPv4 address to listen for incoming connections  |
    |  ipv6    |  IPv6 address to listen for incoming connections  |
`,
			Description: `Local IP addresses to listen on

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  ipv4    |  IPv4 address to listen for incoming connections  |
    |  ipv6    |  IPv6 address to listen for incoming connections  |
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

			// Default:          stringdefault.StaticString(`53`),
			Computed: true,
		},

		"negative_ttl":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (negative-ttl) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum amount of time negative entries are cached

    |  Format  |  Description                        |
    |----------|-------------------------------------|
    |  0-7200  |  Seconds to cache NXDOMAIN entries  |
`,
			Description: `Maximum amount of time negative entries are cached

    |  Format  |  Description                        |
    |----------|-------------------------------------|
    |  0-7200  |  Seconds to cache NXDOMAIN entries  |
`,

			// Default:          stringdefault.StaticString(`3600`),
			Computed: true,
		},

		"serve_stale_extension":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (serve-stale-extension) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of times the expired TTL of a record is extended by 30 seconds when serving stale

    |  Format   |  Description                        |
    |-----------|-------------------------------------|
    |  0-65535  |  Number of times to extend the TTL  |
`,
			Description: `Number of times the expired TTL of a record is extended by 30 seconds when serving stale

    |  Format   |  Description                        |
    |-----------|-------------------------------------|
    |  0-65535  |  Number of times to extend the TTL  |
`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"timeout":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (timeout) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of milliseconds to wait for a remote authoritative server to respond

    |  Format    |  Description                      |
    |------------|-----------------------------------|
    |  10-60000  |  Network timeout in milliseconds  |
`,
			Description: `Number of milliseconds to wait for a remote authoritative server to respond

    |  Format    |  Description                      |
    |------------|-----------------------------------|
    |  10-60000  |  Network timeout in milliseconds  |
`,

			// Default:          stringdefault.StaticString(`1500`),
			Computed: true,
		},

		"source_address":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (source-address) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
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

			// Default:          stringdefault.StaticString(`0.0.0.0 ::`),
			Computed: true,
		},

		"system":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (system) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use system name servers

`,
			Description: `Use system name servers

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"exclude_throttle_address":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (exclude-throttle-address) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address or subnet

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  ipv4     |  IPv4 address to match  |
    |  ipv4net  |  IPv4 prefix to match   |
    |  ipv6     |  IPv6 address           |
    |  ipv6net  |  IPv6 address           |
`,
			Description: `IP address or subnet

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  ipv4     |  IPv4 address to match  |
    |  ipv4net  |  IPv4 prefix to match   |
    |  ipv6     |  IPv6 address           |
    |  ipv6net  |  IPv6 address           |
`,
		},

		// TagNodes

		// Nodes

	}
}
