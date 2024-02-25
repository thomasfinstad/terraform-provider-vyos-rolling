// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceWebproxy describes the resource data model.
type ServiceWebproxy struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

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

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagServiceWebproxyCachePeer     bool `tfsdk:"-" vyos:"cache-peer,child"`
	ExistsTagServiceWebproxyListenAddress bool `tfsdk:"-" vyos:"listen-address,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeServiceWebproxyAuthentication bool `tfsdk:"-" vyos:"authentication,child"`
	ExistsNodeServiceWebproxyURLFiltering   bool `tfsdk:"-" vyos:"url-filtering,child"`
}

// SetID configures the resource ID
func (o *ServiceWebproxy) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceWebproxy) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"webproxy",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceWebproxy) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"safe_ports": schema.ListAttribute{
			ElementType: types.NumberType,
			Optional:    true,
			MarkdownDescription: `Safe port ACL

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-1024  &emsp; |  Port number. Ports included by default: 21,70,80,210,280,443,488,591,777,873,1025-65535  |

`,
		},

		"ssl_safe_ports": schema.ListAttribute{
			ElementType: types.NumberType,
			Optional:    true,
			MarkdownDescription: `SSL safe port

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Port number. Ports included by default: 443  |

`,
		},

		"append_domain": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default domain name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  domain  &emsp; |  Domain to use for urls that do not contain a '.'  |

`,
		},

		"cache_size": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Disk cache size in MB

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Disk cache size in MB  |
    |  0  &emsp; |  Disable disk caching  |

`,

			// Default:          stringdefault.StaticString(`100`),
			Computed: true,
		},

		"default_port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Default Proxy Port

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1025-65535  &emsp; |  Default port number  |

`,

			// Default:          stringdefault.StaticString(`3128`),
			Computed: true,
		},

		"disable_access_log": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable logging of HTTP accesses

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"domain_block": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Domain name to block

`,
		},

		"domain_noncache": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Domain name to access without caching

`,
		},

		"maximum_object_size": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum size of object to be stored in cache in kilobytes

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Object size in KB  |

`,
		},

		"mem_cache_size": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Memory cache size in MB

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Memory cache size in MB   |

`,

			// Default:          stringdefault.StaticString(`20`),
			Computed: true,
		},

		"minimum_object_size": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum size of object to be stored in cache in kilobytes

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Object size in KB  |

`,
		},

		"outgoing_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Outgoing IP address for webproxy

`,
		},

		"reply_block_mime": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `MIME type to block

`,
		},

		"reply_body_max_size": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum reply body size in KB

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Reply size in KB  |

`,
		},
	}
}
