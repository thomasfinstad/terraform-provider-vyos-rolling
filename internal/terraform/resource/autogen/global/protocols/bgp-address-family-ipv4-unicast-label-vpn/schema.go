// Package globalprotocolsbgpaddressfamilyipvfourunicastlabelvpn code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyipvfourunicastlabelvpn

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r protocolsBgpAddressFamilyIPvfourUnicastLabelVpn) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: `~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>protocols</i>

<br>
&darr;
<br>
Border Gateway Protocol (BGP)

<br>
&darr;
<br>
BGP address-family parameters

<br>
&darr;
<br>
IPv4 BGP settings

<br>
&darr;
<br>
Label value for VRF

<br>
&darr;
<br>
<b>
Between current address-family and VPN
</b>
</div>
`,
		Attributes: r.model.ResourceSchemaAttributes(),
	}
}
