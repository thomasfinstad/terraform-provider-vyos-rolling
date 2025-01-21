/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (option) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceDhcpvsixServerSharedNetworkNameOption{}

// ServiceDhcpvsixServerSharedNetworkNameOption describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ServiceDhcpvsixServerSharedNetworkNameOption struct {
	// LeafNodes
	LeafServiceDhcpvsixServerSharedNetworkNameOptionCaptivePortal   types.String `tfsdk:"captive_portal" vyos:"captive-portal,omitempty"`
	LeafServiceDhcpvsixServerSharedNetworkNameOptionDomainSearch    types.List   `tfsdk:"domain_search" vyos:"domain-search,omitempty"`
	LeafServiceDhcpvsixServerSharedNetworkNameOptionNameServer      types.List   `tfsdk:"name_server" vyos:"name-server,omitempty"`
	LeafServiceDhcpvsixServerSharedNetworkNameOptionNisDomain       types.String `tfsdk:"nis_domain" vyos:"nis-domain,omitempty"`
	LeafServiceDhcpvsixServerSharedNetworkNameOptionNisServer       types.List   `tfsdk:"nis_server" vyos:"nis-server,omitempty"`
	LeafServiceDhcpvsixServerSharedNetworkNameOptionNisplusDomain   types.String `tfsdk:"nisplus_domain" vyos:"nisplus-domain,omitempty"`
	LeafServiceDhcpvsixServerSharedNetworkNameOptionNisplusServer   types.List   `tfsdk:"nisplus_server" vyos:"nisplus-server,omitempty"`
	LeafServiceDhcpvsixServerSharedNetworkNameOptionSIPServer       types.List   `tfsdk:"sip_server" vyos:"sip-server,omitempty"`
	LeafServiceDhcpvsixServerSharedNetworkNameOptionSntpServer      types.List   `tfsdk:"sntp_server" vyos:"sntp-server,omitempty"`
	LeafServiceDhcpvsixServerSharedNetworkNameOptionInfoRefreshTime types.Number `tfsdk:"info_refresh_time" vyos:"info-refresh-time,omitempty"`

	// TagNodes

	// Nodes

	NodeServiceDhcpvsixServerSharedNetworkNameOptionVendorOption *ServiceDhcpvsixServerSharedNetworkNameOptionVendorOption `tfsdk:"vendor_option" vyos:"vendor-option,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpvsixServerSharedNetworkNameOption) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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

		"name_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (name-server) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Domain Name Servers (DNS) addresses

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  ipv6    |  Domain Name Server (DNS) IPv6 address  |
`,
			Description: `Domain Name Servers (DNS) addresses

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  ipv6    |  Domain Name Server (DNS) IPv6 address  |
`,
		},

		"nis_domain":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (nis-domain) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NIS domain name for client to use

`,
			Description: `NIS domain name for client to use

`,
		},

		"nis_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (nis-server) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IPv6 address of a NIS Server

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  ipv6    |  IPv6 address of NIS server  |
`,
			Description: `IPv6 address of a NIS Server

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  ipv6    |  IPv6 address of NIS server  |
`,
		},

		"nisplus_domain":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (nisplus-domain) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NIS+ domain name for client to use

`,
			Description: `NIS+ domain name for client to use

`,
		},

		"nisplus_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (nisplus-server) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IPv6 address of a NIS+ Server

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  ipv6    |  IPv6 address of NIS+ server  |
`,
			Description: `IPv6 address of a NIS+ Server

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  ipv6    |  IPv6 address of NIS+ server  |
`,
		},

		"sip_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (sip-server) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IPv6 address of SIP server

    |  Format    |  Description                 |
    |------------|------------------------------|
    |  ipv6      |  IPv6 address of SIP server  |
    |  hostname  |  FQDN of SIP server          |
`,
			Description: `IPv6 address of SIP server

    |  Format    |  Description                 |
    |------------|------------------------------|
    |  ipv6      |  IPv6 address of SIP server  |
    |  hostname  |  FQDN of SIP server          |
`,
		},

		"sntp_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (sntp-server) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IPv6 address of an SNTP server for client to use

`,
			Description: `IPv6 address of an SNTP server for client to use

`,
		},

		"info_refresh_time":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (info-refresh-time) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Time (in seconds) that stateless clients should wait between refreshing the information they were given

    |  Format        |  Description                      |
    |----------------|-----------------------------------|
    |  1-4294967295  |  DHCPv6 information refresh time  |
`,
			Description: `Time (in seconds) that stateless clients should wait between refreshing the information they were given

    |  Format        |  Description                      |
    |----------------|-----------------------------------|
    |  1-4294967295  |  DHCPv6 information refresh time  |
`,
		},

		// TagNodes

		// Nodes

		"vendor_option": schema.SingleNestedAttribute{
			Attributes: ServiceDhcpvsixServerSharedNetworkNameOptionVendorOption{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Vendor Specific Options

`,
			Description: `Vendor Specific Options

`,
		},
	}
}
