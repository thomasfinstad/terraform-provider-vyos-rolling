/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsbgpaddressfamilyipvfourflowspeclocalinstall code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvfourflowspeclocalinstall

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

/* tools/generate-terraform-resource-full/templates/resources/common/schema.gotmpl #schema (local-install) */
// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsBgpAddressFamilyIPvfourFlowspecLocalInstall) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

*protocols*  
⯯  
Border Gateway Protocol (BGP)  
⯯  
BGP address-family parameters  
⯯  
Flowspec IPv4 BGP settings  
⯯  
**Apply local policy routing to interface**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
