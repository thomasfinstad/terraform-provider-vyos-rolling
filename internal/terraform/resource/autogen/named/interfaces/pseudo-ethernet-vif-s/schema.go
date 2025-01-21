/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedinterfacespseudoethernetvifs code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacespseudoethernetvifs

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl #schema (vif-s) */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r interfacesPseudoEthernetVifS) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `
*interfaces*  
⯯  
Pseudo Ethernet Interface (Macvlan)  
⯯  
**QinQ TAG-S Virtual Local Area Network (VLAN) ID**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
