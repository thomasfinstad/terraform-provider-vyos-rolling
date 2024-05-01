// Package namedvrfnameprotocolsbgpaddressfamilyipvfourunicastdistanceprefix code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvfourunicastdistanceprefix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vrfNameProtocolsBgpAddressFamilyIPvfourUnicastDistancePrefix) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Border Gateway Protocol (BGP)  
⯯  
BGP address-family parameters  
⯯  
IPv4 BGP settings  
⯯  
Administrative distances for BGP routes  
⯯  
**Administrative distance for a specific BGP prefix**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
