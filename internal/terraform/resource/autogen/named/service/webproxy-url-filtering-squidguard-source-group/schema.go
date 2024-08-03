// Package namedservicewebproxyurlfilteringsquidguardsourcegroup code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicewebproxyurlfilteringsquidguardsourcegroup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceWebproxyURLFilteringSquIDguardSourceGroup) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `*service*  
⯯  
Webproxy service settings  
⯯  
URL filtering settings  
⯯  
URL filtering via squidGuard redirector  
⯯  
**Source group name**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
