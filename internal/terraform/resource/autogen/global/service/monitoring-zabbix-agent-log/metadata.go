/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalservicemonitoringzabbixagentlog code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicemonitoringzabbixagentlog

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (log) */
// Metadata method to define the resource type name.
func (r serviceMonitoringZabbixAgentLog) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_monitoring_zabbix_agent_log"
}
