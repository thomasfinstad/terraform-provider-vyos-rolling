// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceMonitoringTelegrafAzureDataExplorer describes the resource data model.
type ServiceMonitoringTelegrafAzureDataExplorer struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	// LeafNodes
	LeafServiceMonitoringTelegrafAzureDataExplorerDatabase     types.String `tfsdk:"database" vyos:"database,omitempty"`
	LeafServiceMonitoringTelegrafAzureDataExplorerGroupMetrics types.String `tfsdk:"group_metrics" vyos:"group-metrics,omitempty"`
	LeafServiceMonitoringTelegrafAzureDataExplorerTable        types.String `tfsdk:"table" vyos:"table,omitempty"`
	LeafServiceMonitoringTelegrafAzureDataExplorerURL          types.String `tfsdk:"url" vyos:"url,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeServiceMonitoringTelegrafAzureDataExplorerAuthentication bool `tfsdk:"-" vyos:"authentication,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceMonitoringTelegrafAzureDataExplorer) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceMonitoringTelegrafAzureDataExplorer) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"monitoring",

		"telegraf",

		"azure-data-explorer",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceMonitoringTelegrafAzureDataExplorer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"database": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Remote database name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Remote database name  |

`,
		},

		"group_metrics": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Type of metrics grouping when push to Azure Data Explorer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  single-table  &emsp; |  Metrics stores in one table  |
    |  table-per-metric  &emsp; |  One table per gorups of metric by the metric name  |

`,

			// Default:          stringdefault.StaticString(`table-per-metric`),
			Computed: true,
		},

		"table": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Name of the single table [Only if set group-metrics single-table]

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Table name  |

`,
		},

		"url": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Remote URL

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  url  &emsp; |  Remote URL  |

`,
		},
	}
}
