// Package globalserviceidsddosprotectionthresholdicmp code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceidsddosprotectionthresholdicmp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceIDsDdosProtectionThresholdIcmp) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

*service*  
⯯  
Intrusion Detection System  
⯯  
FastNetMon detection and protection parameters  
⯯  
Attack limits thresholds  
⯯  
**ICMP threshold**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
