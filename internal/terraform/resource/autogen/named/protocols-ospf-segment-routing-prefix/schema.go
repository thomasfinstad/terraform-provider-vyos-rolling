// Package namedprotocolsospfsegmentroutingprefix code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfsegmentroutingprefix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsOspfSegmentRoutingPrefix) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `Open Shortest Path First (OSPF)

Segment-Routing (SPRING) settings

Static IPv4 prefix segment/label mapping

|  Format  |  Description  |
|----------|---------------|
|  ipv4net  |  IPv4 prefix segment  |

`,
		Attributes: r.model.ResourceAttributes(),
	}
}
