// Package globalserviceconntracksyncfailovermechanismvrrp code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceconntracksyncfailovermechanismvrrp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r serviceConntrackSyncFailoverMechanismVrrp) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

*service*  
⯯  
Connection tracking synchronization  
⯯  
Failover mechanism to use for conntrack-sync  
⯯  
**VRRP as failover-mechanism to use for conntrack-sync**
`,
		Attributes: r.model.ResourceSchemaAttributes(ctx),
	}
}
