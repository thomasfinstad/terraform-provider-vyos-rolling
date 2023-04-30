// Package namedservicewebproxyurlfilteringsquidguardrule code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicewebproxyurlfilteringsquidguardrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceWebproxyURLFilteringSquIDguardRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Webproxy service settings

URL filtering settings

URL filtering via squidGuard redirector

URL filter rule for a source-group

|  Format  |  Description  |
|----------|---------------|
|  u32:1-1024  |  Rule Number  |

`,
		Attributes: r.model.ResourceAttributes(),
	}
}
