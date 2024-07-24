// Package namedvrfnameprotocolsbgpaddressfamilyipvsixmulticastaggregateaddress code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyipvsixmulticastaggregateaddress

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r vrfNameProtocolsBgpAddressFamilyIPvsixMulticastAggregateAddress) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
Multicast IPv6 BGP settings  
⯯  
**BGP aggregate network/prefix**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
