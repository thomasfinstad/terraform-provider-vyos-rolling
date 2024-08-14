// Package globalserviceconfigsyncsectionsystem code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceconfigsyncsectionsystem

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceConfigSyncSectionSystem) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Configuration synchronization  
⯯  
Section for synchronization  
⯯  
**System parameters**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}