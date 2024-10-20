/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceDhcpServerSharedNetworkNameSubnetStaticMappingOption{}

// ServiceDhcpServerSharedNetworkNameSubnetStaticMappingOption describes the resource data model.
type ServiceDhcpServerSharedNetworkNameSubnetStaticMappingOption struct {
	// LeafNodes
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionCaptivePortal       types.String `tfsdk:"captive_portal" vyos:"captive-portal,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionDomainName          types.String `tfsdk:"domain_name" vyos:"domain-name,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionDomainSearch        types.List   `tfsdk:"domain_search" vyos:"domain-search,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionNtpServer           types.List   `tfsdk:"ntp_server" vyos:"ntp-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionNameServer          types.List   `tfsdk:"name_server" vyos:"name-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionBootfileName        types.String `tfsdk:"bootfile_name" vyos:"bootfile-name,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionBootfileServer      types.String `tfsdk:"bootfile_server" vyos:"bootfile-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionBootfileSize        types.Number `tfsdk:"bootfile_size" vyos:"bootfile-size,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionClientPrefixLength  types.Number `tfsdk:"client_prefix_length" vyos:"client-prefix-length,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionDefaultRouter       types.String `tfsdk:"default_router" vyos:"default-router,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionIPForwarding        types.Bool   `tfsdk:"ip_forwarding" vyos:"ip-forwarding,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionIPvsixOnlyPreferred types.Number `tfsdk:"ipv6_only_preferred" vyos:"ipv6-only-preferred,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionPopServer           types.List   `tfsdk:"pop_server" vyos:"pop-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionServerIDentifier    types.String `tfsdk:"server_identifier" vyos:"server-identifier,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionSMTPServer          types.List   `tfsdk:"smtp_server" vyos:"smtp-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionTftpServerName      types.String `tfsdk:"tftp_server_name" vyos:"tftp-server-name,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionTimeOffset          types.String `tfsdk:"time_offset" vyos:"time-offset,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionTimeServer          types.List   `tfsdk:"time_server" vyos:"time-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionTimeZone            types.String `tfsdk:"time_zone" vyos:"time-zone,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionWinsServer          types.List   `tfsdk:"wins_server" vyos:"wins-server,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionWpadURL             types.String `tfsdk:"wpad_url" vyos:"wpad-url,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionStaticRoute bool `tfsdk:"-" vyos:"static-route,child"`

	// Nodes
	NodeServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionVendorOption *ServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionVendorOption `tfsdk:"vendor_option" vyos:"vendor-option,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameSubnetStaticMappingOption) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"captive_portal":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client Domain Name

`,
			Description: `Client Domain Name

`,
		},

		"domain_search":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype-multi.gotmpl */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Client Domain Name search list

`,
			Description: `Client Domain Name search list

`,
		},

		"ntp_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype-multi.gotmpl */
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
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype-multi.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Bootstrap file name

`,
			Description: `Bootstrap file name

`,
		},

		"bootfile_server":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype-multi.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype-multi.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
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
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype-multi.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time zone to send to clients. Uses RFC4833 options 100 and 101

`,
			Description: `Time zone to send to clients. Uses RFC4833 options 100 and 101

`,
		},

		"wins_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype-multi.gotmpl */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Web Proxy Autodiscovery (WPAD) URL

`,
			Description: `Web Proxy Autodiscovery (WPAD) URL

`,
		},

		// Nodes

		"vendor_option": schema.SingleNestedAttribute{
			Attributes: ServiceDhcpServerSharedNetworkNameSubnetStaticMappingOptionVendorOption{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Vendor Specific Options

`,
			Description: `Vendor Specific Options

`,
		},
	}
}
