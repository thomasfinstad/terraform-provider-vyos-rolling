// Package namedfirewallipvfouroutputrawrule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvfouroutputrawrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r firewallIPvfourOutputRawRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Firewall  
⯯  
IPv4 firewall  
⯯  
IPv4 output firewall  
⯯  
IPv4 firewall output raw  
⯯  
**IPv4 Firewall output raw rule number**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
