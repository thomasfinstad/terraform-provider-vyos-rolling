// Package namedservicemonitoringzabbixagentserveractive code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicemonitoringzabbixagentserveractive

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r serviceMonitoringZabbixAgentServerActive) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_monitoring_zabbix_agent_server_active"
}