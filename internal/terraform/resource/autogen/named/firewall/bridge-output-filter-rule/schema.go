// Package namedfirewallbridgeoutputfilterrule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallbridgeoutputfilterrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r firewallBrIDgeOutputFilterRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Firewall  
⯯  
Bridge firewall  
⯯  
Bridge output firewall  
⯯  
Bridge firewall output filter  
⯯  
**Bridge Firewall output filter rule number**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
