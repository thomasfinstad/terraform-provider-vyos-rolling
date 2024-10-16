// Package namedprotocolsopenfabricdomaininterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsopenfabricdomaininterface

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsOpenfabricDomainInterface) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*protocols*  
⯯  
OpenFabric protocol  
⯯  
OpenFabric process name  
⯯  
**Interface params**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
