// Package globalfirewallbridgeinputfilter code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallbridgeinputfilter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r firewallBrIDgeInputFilter) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

Firewall  
⯯  
Bridge firewall  
⯯  
Bridge input firewall  
⯯  
**Bridge firewall input filter**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
