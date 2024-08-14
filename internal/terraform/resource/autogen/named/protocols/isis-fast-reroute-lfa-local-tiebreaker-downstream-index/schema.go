// Package namedprotocolsisisfastreroutelfalocaltiebreakerdownstreamindex code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsisisfastreroutelfalocaltiebreakerdownstreamindex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndex) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*protocols*  
⯯  
Intermediate System to Intermediate System (IS-IS)  
⯯  
IS-IS fast reroute configuration  
⯯  
Loop free alternate functionality  
⯯  
Local loop free alternate options  
⯯  
Configure tiebreaker for multiple backups  
⯯  
Prefer backup path via downstream node  
⯯  
**Set preference order among tiebreakers**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}