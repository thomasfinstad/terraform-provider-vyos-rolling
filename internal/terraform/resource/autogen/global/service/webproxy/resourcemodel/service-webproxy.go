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

var _ helpers.VyosTopResourceDataModel = &ServiceWebproxy{}

// ServiceWebproxy describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ServiceWebproxy struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceWebproxySafePorts         types.List   `tfsdk:"safe_ports" vyos:"safe-ports,omitempty"`
	LeafServiceWebproxySslSafePorts      types.List   `tfsdk:"ssl_safe_ports" vyos:"ssl-safe-ports,omitempty"`
	LeafServiceWebproxyAppendDomain      types.String `tfsdk:"append_domain" vyos:"append-domain,omitempty"`
	LeafServiceWebproxyCacheSize         types.String `tfsdk:"cache_size" vyos:"cache-size,omitempty"`
	LeafServiceWebproxyDefaultPort       types.Number `tfsdk:"default_port" vyos:"default-port,omitempty"`
	LeafServiceWebproxyDisableAccessLog  types.Bool   `tfsdk:"disable_access_log" vyos:"disable-access-log,omitempty"`
	LeafServiceWebproxyDomainBlock       types.List   `tfsdk:"domain_block" vyos:"domain-block,omitempty"`
	LeafServiceWebproxyDomainNoncache    types.List   `tfsdk:"domain_noncache" vyos:"domain-noncache,omitempty"`
	LeafServiceWebproxyMaximumObjectSize types.Number `tfsdk:"maximum_object_size" vyos:"maximum-object-size,omitempty"`
	LeafServiceWebproxyMemCacheSize      types.Number `tfsdk:"mem_cache_size" vyos:"mem-cache-size,omitempty"`
	LeafServiceWebproxyMinimumObjectSize types.Number `tfsdk:"minimum_object_size" vyos:"minimum-object-size,omitempty"`
	LeafServiceWebproxyOutgoingAddress   types.String `tfsdk:"outgoing_address" vyos:"outgoing-address,omitempty"`
	LeafServiceWebproxyReplyBlockMime    types.List   `tfsdk:"reply_block_mime" vyos:"reply-block-mime,omitempty"`
	LeafServiceWebproxyReplyBodyMaxSize  types.Number `tfsdk:"reply_body_max_size" vyos:"reply-body-max-size,omitempty"`

	// TagNodes

	ExistsTagServiceWebproxyCachePeer bool `tfsdk:"-" vyos:"cache-peer,child"`

	ExistsTagServiceWebproxyListenAddress bool `tfsdk:"-" vyos:"listen-address,child"`

	// Nodes

	ExistsNodeServiceWebproxyAuthentication bool `tfsdk:"-" vyos:"authentication,child"`

	ExistsNodeServiceWebproxyURLFiltering bool `tfsdk:"-" vyos:"url-filtering,child"`
}

// SetID configures the resource ID
func (o *ServiceWebproxy) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceWebproxy) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceWebproxy) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceWebproxy) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"webproxy",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceWebproxy) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"service", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceWebproxy) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceWebproxy) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"safe_ports":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.NumberType,
			Optional:    true,
			MarkdownDescription: `Safe port ACL

    |  Format  |  Description                                                                              |
    |----------|-------------------------------------------------------------------------------------------|
    |  1-1024  |  Port number. Ports included by default: 21,70,80,210,280,443,488,591,777,873,1025-65535  |
`,
			Description: `Safe port ACL

    |  Format  |  Description                                                                              |
    |----------|-------------------------------------------------------------------------------------------|
    |  1-1024  |  Port number. Ports included by default: 21,70,80,210,280,443,488,591,777,873,1025-65535  |
`,
		},

		"ssl_safe_ports":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.NumberType,
			Optional:    true,
			MarkdownDescription: `SSL safe port

    |  Format   |  Description                                  |
    |-----------|-----------------------------------------------|
    |  1-65535  |  Port number. Ports included by default: 443  |
`,
			Description: `SSL safe port

    |  Format   |  Description                                  |
    |-----------|-----------------------------------------------|
    |  1-65535  |  Port number. Ports included by default: 443  |
`,
		},

		"append_domain":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default domain name

    |  Format  |  Description                                       |
    |----------|----------------------------------------------------|
    |  domain  |  Domain to use for urls that do not contain a '.'  |
`,
			Description: `Default domain name

    |  Format  |  Description                                       |
    |----------|----------------------------------------------------|
    |  domain  |  Domain to use for urls that do not contain a '.'  |
`,
		},

		"cache_size":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disk cache size in MB

    |  Format  |  Description            |
    |----------|-------------------------|
    |  u32     |  Disk cache size in MB  |
    |  0       |  Disable disk caching   |
`,
			Description: `Disk cache size in MB

    |  Format  |  Description            |
    |----------|-------------------------|
    |  u32     |  Disk cache size in MB  |
    |  0       |  Disable disk caching   |
`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		"default_port":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Default Proxy Port

    |  Format      |  Description          |
    |--------------|-----------------------|
    |  1025-65535  |  Default port number  |
`,
			Description: `Default Proxy Port

    |  Format      |  Description          |
    |--------------|-----------------------|
    |  1025-65535  |  Default port number  |
`,

			// Default:          stringdefault.StaticString(`3128`),
			Computed: true,
		},

		"disable_access_log":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable logging of HTTP accesses

`,
			Description: `Disable logging of HTTP accesses

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"domain_block":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Domain name to block

`,
			Description: `Domain name to block

`,
		},

		"domain_noncache":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Domain name to access without caching

`,
			Description: `Domain name to access without caching

`,
		},

		"maximum_object_size":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum size of object to be stored in cache in kilobytes

    |  Format  |  Description        |
    |----------|---------------------|
    |  u32     |  Object size in KB  |
`,
			Description: `Maximum size of object to be stored in cache in kilobytes

    |  Format  |  Description        |
    |----------|---------------------|
    |  u32     |  Object size in KB  |
`,
		},

		"mem_cache_size":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Memory cache size in MB

    |  Format  |  Description               |
    |----------|----------------------------|
    |  u32     |  Memory cache size in MB   |
`,
			Description: `Memory cache size in MB

    |  Format  |  Description               |
    |----------|----------------------------|
    |  u32     |  Memory cache size in MB   |
`,

			// Default:          stringdefault.StaticString(`20`),
			Computed: true,
		},

		"minimum_object_size":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum size of object to be stored in cache in kilobytes

    |  Format  |  Description        |
    |----------|---------------------|
    |  u32     |  Object size in KB  |
`,
			Description: `Maximum size of object to be stored in cache in kilobytes

    |  Format  |  Description        |
    |----------|---------------------|
    |  u32     |  Object size in KB  |
`,
		},

		"outgoing_address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Outgoing IP address for webproxy

`,
			Description: `Outgoing IP address for webproxy

`,
		},

		"reply_block_mime":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `MIME type to block

`,
			Description: `MIME type to block

`,
		},

		"reply_body_max_size":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum reply body size in KB

    |  Format  |  Description       |
    |----------|--------------------|
    |  u32     |  Reply size in KB  |
`,
			Description: `Maximum reply body size in KB

    |  Format  |  Description       |
    |----------|--------------------|
    |  u32     |  Reply size in KB  |
`,
		},

		// TagNodes

		// Nodes

	}
}
