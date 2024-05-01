// Package namedpkikeypair code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpkikeypair

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r pkiKeyPair) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Public key infrastructure (PKI)  
⯯  
**Public and private keys**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
