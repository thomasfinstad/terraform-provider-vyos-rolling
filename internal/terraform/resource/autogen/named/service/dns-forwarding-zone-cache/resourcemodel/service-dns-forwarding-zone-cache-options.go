// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceDNSForwardingZoneCacheOptions{}

// ServiceDNSForwardingZoneCacheOptions describes the resource data model.
type ServiceDNSForwardingZoneCacheOptions struct {
	// LeafNodes
	LeafServiceDNSForwardingZoneCacheOptionsTimeout       types.Number `tfsdk:"timeout" vyos:"timeout,omitempty"`
	LeafServiceDNSForwardingZoneCacheOptionsRetryInterval types.Number `tfsdk:"retry_interval" vyos:"retry-interval,omitempty"`
	LeafServiceDNSForwardingZoneCacheOptionsMaxZoneSize   types.Number `tfsdk:"max_zone_size" vyos:"max-zone-size,omitempty"`
	LeafServiceDNSForwardingZoneCacheOptionsZonemd        types.String `tfsdk:"zonemd" vyos:"zonemd,omitempty"`
	LeafServiceDNSForwardingZoneCacheOptionsDNSsec        types.String `tfsdk:"dnssec" vyos:"dnssec,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeServiceDNSForwardingZoneCacheOptionsRefresh *ServiceDNSForwardingZoneCacheOptionsRefresh `tfsdk:"refresh" vyos:"refresh,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingZoneCacheOptions) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"timeout": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Zone retrieval timeout

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  1-3600  |  Request timeout in seconds  |
`,
			Description: `Zone retrieval timeout

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  1-3600  |  Request timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`20`),
			Computed: true,
		},

		"retry_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Retry interval after zone retrieval errors

    |  Format   |  Description              |
    |-----------|---------------------------|
    |  1-86400  |  Retry period in seconds  |
`,
			Description: `Retry interval after zone retrieval errors

    |  Format   |  Description              |
    |-----------|---------------------------|
    |  1-86400  |  Retry period in seconds  |
`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		"max_zone_size": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum zone size in megabytes

    |  Format  |  Description        |
    |----------|---------------------|
    |  0       |  No restriction     |
    |  1-1024  |  Size in megabytes  |
`,
			Description: `Maximum zone size in megabytes

    |  Format  |  Description        |
    |----------|---------------------|
    |  0       |  No restriction     |
    |  1-1024  |  Size in megabytes  |
`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"zonemd": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Message Digest for DNS Zones (RFC 8976)

    |  Format    |  Description                                |
    |------------|---------------------------------------------|
    |  ignore    |  Ignore ZONEMD records                      |
    |  validate  |  Validate ZONEMD if present                 |
    |  require   |  Require valid ZONEMD record to be present  |
`,
			Description: `Message Digest for DNS Zones (RFC 8976)

    |  Format    |  Description                                |
    |------------|---------------------------------------------|
    |  ignore    |  Ignore ZONEMD records                      |
    |  validate  |  Validate ZONEMD if present                 |
    |  require   |  Require valid ZONEMD record to be present  |
`,

			// Default:          stringdefault.StaticString(`validate`),
			Computed: true,
		},

		"dnssec": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DNSSEC mode

    |  Format    |  Description                                                       |
    |------------|--------------------------------------------------------------------|
    |  ignore    |  Do not do DNSSEC validation                                       |
    |  validate  |  Reject zones with incorrect signatures but accept unsigned zones  |
    |  require   |  Require DNSSEC validation                                         |
`,
			Description: `DNSSEC mode

    |  Format    |  Description                                                       |
    |------------|--------------------------------------------------------------------|
    |  ignore    |  Do not do DNSSEC validation                                       |
    |  validate  |  Reject zones with incorrect signatures but accept unsigned zones  |
    |  require   |  Require DNSSEC validation                                         |
`,

			// Default:          stringdefault.StaticString(`validate`),
			Computed: true,
		},

		// Nodes

		"refresh": schema.SingleNestedAttribute{
			Attributes: ServiceDNSForwardingZoneCacheOptionsRefresh{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Zone caching options

`,
			Description: `Zone caching options

`,
		},
	}
}
