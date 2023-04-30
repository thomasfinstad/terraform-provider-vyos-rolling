// Package namedsystemsyslogglobalfacility code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemsyslogglobalfacility

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r systemSyslogGlobalFacility) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `System logging

Logging to system standard location

Facility for logging

|  Format  |  Description  |
|----------|---------------|
|  all  |  All facilities excluding "mark"  |
|  auth  |  Authentication and authorization  |
|  authpriv  |  Non-system authorization  |
|  cron  |  Cron daemon  |
|  daemon  |  System daemons  |
|  kern  |  Kernel  |
|  lpr  |  Line printer spooler  |
|  mail  |  Mail subsystem  |
|  mark  |  Timestamp  |
|  news  |  USENET subsystem  |
|  protocols  |  depricated will be set to local7  |
|  security  |  depricated will be set to auth  |
|  syslog  |  Authentication and authorization  |
|  user  |  Application processes  |
|  uucp  |  UUCP subsystem  |
|  local0  |  Local facility 0  |
|  local1  |  Local facility 1  |
|  local2  |  Local facility 2  |
|  local3  |  Local facility 3  |
|  local4  |  Local facility 4  |
|  local5  |  Local facility 5  |
|  local6  |  Local facility 6  |
|  local7  |  Local facility 7  |

`,
		Attributes: r.model.ResourceAttributes(),
	}
}
