/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedfirewallipvsixnamerule code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedfirewallipvsixnamerule

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (rule) */
// Metadata method to define the resource type name.
func (r firewallIPvsixNameRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall_ipv6_name_rule"
}
