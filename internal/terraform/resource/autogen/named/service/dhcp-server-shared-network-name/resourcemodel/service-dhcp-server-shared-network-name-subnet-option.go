/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (option) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceDhcpServerSharedNetworkNameSubnetOption{}

// ServiceDhcpServerSharedNetworkNameSubnetOption describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ServiceDhcpServerSharedNetworkNameSubnetOption struct {
	// LeafNodes
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionCaptivePortal       types.String `tfsdk:"captive_portal" vyos:"captive-portal,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionDomainName          types.String `tfsdk:"domain_name" vyos:"domain-name,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionDomainSearch        types.List   `tfsdk:"domain_search" vyos:"domain-search,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionNtpServer           types.List   `tfsdk:"ntp_server" vyos:"ntp-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionNameServer          types.List   `tfsdk:"name_server" vyos:"name-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionBootfileName        types.String `tfsdk:"bootfile_name" vyos:"bootfile-name,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionBootfileServer      types.String `tfsdk:"bootfile_server" vyos:"bootfile-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionBootfileSize        types.Number `tfsdk:"bootfile_size" vyos:"bootfile-size,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionClientPrefixLength  types.Number `tfsdk:"client_prefix_length" vyos:"client-prefix-length,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionDefaultRouter       types.String `tfsdk:"default_router" vyos:"default-router,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionIPForwarding        types.Bool   `tfsdk:"ip_forwarding" vyos:"ip-forwarding,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionIPvsixOnlyPreferred types.Number `tfsdk:"ipv6_only_preferred" vyos:"ipv6-only-preferred,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionPopServer           types.List   `tfsdk:"pop_server" vyos:"pop-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionServerIDentifier    types.String `tfsdk:"server_identifier" vyos:"server-identifier,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionSMTPServer          types.List   `tfsdk:"smtp_server" vyos:"smtp-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionTftpServerName      types.String `tfsdk:"tftp_server_name" vyos:"tftp-server-name,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionTimeOffset          types.String `tfsdk:"time_offset" vyos:"time-offset,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionTimeServer          types.List   `tfsdk:"time_server" vyos:"time-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionTimeZone            types.String `tfsdk:"time_zone" vyos:"time-zone,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionWinsServer          types.List   `tfsdk:"wins_server" vyos:"wins-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetOptionWpadURL             types.String `tfsdk:"wpad_url" vyos:"wpad-url,omitempty"`

	// TagNodes

	ExistsTagServiceDhcpServerSharedNetworkNameSubnetOptionStaticRoute bool `tfsdk:"-" vyos:"static-route,child"`

	// Nodes

	NodeServiceDhcpServerSharedNetworkNameSubnetOptionVendorOption *ServiceDhcpServerSharedNetworkNameSubnetOptionVendorOption `tfsdk:"vendor_option" vyos:"vendor-option,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameSubnetOption) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"captive_portal":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (captive-portal) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Captive portal API endpoint

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  txt     |  Captive portal API endpoint  |
`,
			Description: `Captive portal API endpoint

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  txt     |  Captive portal API endpoint  |
`,
		},

		"domain_name":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (domain-name) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client Domain Name

`,
			Description: `Client Domain Name

`,
		},

		"domain_search":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (domain-search) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Client Domain Name search list

`,
			Description: `Client Domain Name search list

`,
		},

		"ntp_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (ntp-server) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address of NTP server

    |  Format  |  Description              |
    |----------|---------------------------|
    |  ipv4    |  NTP server IPv4 address  |
`,
			Description: `IP address of NTP server

    |  Format  |  Description              |
    |----------|---------------------------|
    |  ipv4    |  NTP server IPv4 address  |
`,
		},

		"name_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (name-server) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Domain Name Servers (DNS) addresses

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  ipv4    |  Domain Name Server (DNS) IPv4 address  |
`,
			Description: `Domain Name Servers (DNS) addresses

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  ipv4    |  Domain Name Server (DNS) IPv4 address  |
`,
		},

		"bootfile_name":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (bootfile-name) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Bootstrap file name

`,
			Description: `Bootstrap file name

`,
		},

		"bootfile_server":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (bootfile-server) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Server from which the initial boot file is to be loaded

    |  Format    |  Description                   |
    |------------|--------------------------------|
    |  ipv4      |  Bootfile server IPv4 address  |
    |  hostname  |  Bootfile server FQDN          |
`,
			Description: `Server from which the initial boot file is to be loaded

    |  Format    |  Description                   |
    |------------|--------------------------------|
    |  ipv4      |  Bootfile server IPv4 address  |
    |  hostname  |  Bootfile server FQDN          |
`,
		},

		"bootfile_size":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (bootfile-size) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Bootstrap file size

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  1-16    |  Bootstrap file size in 512 byte blocks  |
`,
			Description: `Bootstrap file size

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  1-16    |  Bootstrap file size in 512 byte blocks  |
`,
		},

		"client_prefix_length":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (client-prefix-length) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Specifies the clients subnet mask as per RFC 950. If unset, subnet declaration is used.

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  0-32    |  DHCP client prefix length must be 0 to 32  |
`,
			Description: `Specifies the clients subnet mask as per RFC 950. If unset, subnet declaration is used.

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  0-32    |  DHCP client prefix length must be 0 to 32  |
`,
		},

		"default_router":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (default-router) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address of default router

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  ipv4    |  Default router IPv4 address  |
`,
			Description: `IP address of default router

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  ipv4    |  Default router IPv4 address  |
`,
		},

		"ip_forwarding":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ip-forwarding) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable IP forwarding on client

`,
			Description: `Enable IP forwarding on client

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ipv6_only_preferred":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ipv6-only-preferred) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Disable IPv4 on IPv6 only hosts (RFC 8925)

    |  Format  |  Description  |
    |----------|---------------|
    |  u32     |  Seconds      |
`,
			Description: `Disable IPv4 on IPv6 only hosts (RFC 8925)

    |  Format  |  Description  |
    |----------|---------------|
    |  u32     |  Seconds      |
`,
		},

		"pop_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (pop-server) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address of POP3 server

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  POP3 server IPv4 address  |
`,
			Description: `IP address of POP3 server

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  POP3 server IPv4 address  |
`,
		},

		"server_identifier":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (server-identifier) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Address for DHCP server identifier

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  ipv4    |  DHCP server identifier IPv4 address  |
`,
			Description: `Address for DHCP server identifier

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  ipv4    |  DHCP server identifier IPv4 address  |
`,
		},

		"smtp_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (smtp-server) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address of SMTP server

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  SMTP server IPv4 address  |
`,
			Description: `IP address of SMTP server

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  SMTP server IPv4 address  |
`,
		},

		"tftp_server_name":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (tftp-server-name) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `TFTP server name

    |  Format    |  Description               |
    |------------|----------------------------|
    |  ipv4      |  TFTP server IPv4 address  |
    |  hostname  |  TFTP server FQDN          |
`,
			Description: `TFTP server name

    |  Format    |  Description               |
    |------------|----------------------------|
    |  ipv4      |  TFTP server IPv4 address  |
    |  hostname  |  TFTP server FQDN          |
`,
		},

		"time_offset":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (time-offset) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client subnet offset in seconds from Coordinated Universal Time (UTC)

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  [-]N    |  Time offset (number, may be negative)  |
`,
			Description: `Client subnet offset in seconds from Coordinated Universal Time (UTC)

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  [-]N    |  Time offset (number, may be negative)  |
`,
		},

		"time_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (time-server) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address of time server

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  Time server IPv4 address  |
`,
			Description: `IP address of time server

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  Time server IPv4 address  |
`,
		},

		"time_zone":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (time-zone) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time zone to send to clients. Uses RFC4833 options 100 and 101

`,
			Description: `Time zone to send to clients. Uses RFC4833 options 100 and 101

`,
		},

		"wins_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (wins-server) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address for Windows Internet Name Service (WINS) server

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  WINS server IPv4 address  |
`,
			Description: `IP address for Windows Internet Name Service (WINS) server

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  WINS server IPv4 address  |
`,
		},

		"wpad_url":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (wpad-url) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Web Proxy Autodiscovery (WPAD) URL

`,
			Description: `Web Proxy Autodiscovery (WPAD) URL

`,
		},

		// TagNodes

		// Nodes

		"vendor_option": schema.SingleNestedAttribute{
			Attributes: ServiceDhcpServerSharedNetworkNameSubnetOptionVendorOption{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Vendor Specific Options

`,
			Description: `Vendor Specific Options

`,
		},
	}
}
