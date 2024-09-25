// Package namedservicemonitoringzabbixagentserveractive code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicemonitoringzabbixagentserveractive

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceMonitoringZabbixAgentServerActive) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*service*  
⯯  
Monitoring services  
⯯  
Zabbix-agent settings  
⯯  
**Remote server address to get active checks from**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
