// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceMonitoringTelegrafInfluxdbAuthentication describes the resource data model.
type ServiceMonitoringTelegrafInfluxdbAuthentication struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafServiceMonitoringTelegrafInfluxdbAuthenticationOrganization types.String `tfsdk:"organization" vyos:"organization,omitempty"`
	LeafServiceMonitoringTelegrafInfluxdbAuthenticationToken        types.String `tfsdk:"token" vyos:"token,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ServiceMonitoringTelegrafInfluxdbAuthentication) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceMonitoringTelegrafInfluxdbAuthentication) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"monitoring",

		"telegraf",

		"influxdb",

		"authentication",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceMonitoringTelegrafInfluxdbAuthentication) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"organization": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Authentication organization for InfluxDB v2

`,
		},

		"token": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Authentication token for InfluxDB v2

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Authentication token  |

`,
		},
	}
}