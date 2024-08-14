// Package namedsystemconntracktimeoutcustomipvfourrule code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemconntracktimeoutcustomipvfourrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r systemConntrackTimeoutCustomIPvfourRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*system*  
⯯  
Connection Tracking Engine Options  
⯯  
Connection timeout options  
⯯  
Define custom timeouts per connection  
⯯  
IPv4 rules  
⯯  
**Rule number**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}