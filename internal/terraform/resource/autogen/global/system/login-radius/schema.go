// Package globalsystemloginradius code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemloginradius

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r systemLoginRadius) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

*system*  
⯯  
System User Login Configuration  
⯯  
**RADIUS based user authentication**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
