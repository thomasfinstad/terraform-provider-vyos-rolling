// Package globalfirewallipvfouroutputfilter code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallipvfouroutputfilter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r firewallIPvfourOutputFilter) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

Firewall  
⯯  
IPv4 firewall  
⯯  
IPv4 output firewall  
⯯  
**IPv4 firewall output filter**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
