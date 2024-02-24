// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceMonitoringTelegrafPrometheusClient describes the resource data model.
type ServiceMonitoringTelegrafPrometheusClient struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafServiceMonitoringTelegrafPrometheusClientAllowFrom     types.List   `tfsdk:"allow_from" vyos:"allow-from,omitempty"`
	LeafServiceMonitoringTelegrafPrometheusClientListenAddress types.String `tfsdk:"listen_address" vyos:"listen-address,omitempty"`
	LeafServiceMonitoringTelegrafPrometheusClientMetricVersion types.Number `tfsdk:"metric_version" vyos:"metric-version,omitempty"`
	LeafServiceMonitoringTelegrafPrometheusClientPort          types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeServiceMonitoringTelegrafPrometheusClientAuthentication bool `tfsdk:"-" vyos:"authentication,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceMonitoringTelegrafPrometheusClient) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceMonitoringTelegrafPrometheusClient) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"monitoring",

		"telegraf",

		"prometheus-client",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceMonitoringTelegrafPrometheusClient) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"allow_from": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Networks allowed to query this server

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IP address and prefix length  |
    |  ipv6net  &emsp; |  IPv6 address and prefix length  |

`,
		},

		"listen_address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local IP addresses to listen on

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address to listen for incoming connections  |
    |  ipv6  &emsp; |  IPv6 address to listen for incoming connections  |

`,
		},

		"metric_version": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Metric version control mapping from Telegraf to Prometheus format

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-2  &emsp; |  Metric version (default: 2)  |

`,

			// Default:          stringdefault.StaticString(`2`),
			Computed: true,
		},

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |

`,

			// Default:          stringdefault.StaticString(`9273`),
			Computed: true,
		},
	}
}