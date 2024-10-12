// Package globalprotocolsmplsldpexportipvfourexportfilter code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsmplsldpexportipvfourexportfilter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsMplsLdpExportIPvfourExportFilter) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

*protocols*  
⯯  
Multiprotocol Label Switching (MPLS)  
⯯  
Label Distribution Protocol (LDP)  
⯯  
Export parameters  
⯯  
IPv4 parameters  
⯯  
**Forwarding equivalence class export filter**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
