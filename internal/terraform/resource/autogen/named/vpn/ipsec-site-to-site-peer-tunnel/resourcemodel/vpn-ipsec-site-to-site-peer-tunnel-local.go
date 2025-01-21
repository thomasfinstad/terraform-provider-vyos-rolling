/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (local) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VpnIPsecSiteToSitePeerTunnelLocal{}

// VpnIPsecSiteToSitePeerTunnelLocal describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VpnIPsecSiteToSitePeerTunnelLocal struct {
	// LeafNodes
	LeafVpnIPsecSiteToSitePeerTunnelLocalPort   types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafVpnIPsecSiteToSitePeerTunnelLocalPrefix types.List   `tfsdk:"prefix" vyos:"prefix,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecSiteToSitePeerTunnelLocal) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"port":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (port) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
			Description: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
		},

		"prefix":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (prefix) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Local IPv4 or IPv6 prefix

    |  Format   |  Description        |
    |-----------|---------------------|
    |  ipv4net  |  Local IPv4 prefix  |
    |  ipv6net  |  Local IPv6 prefix  |
`,
			Description: `Local IPv4 or IPv6 prefix

    |  Format   |  Description        |
    |-----------|---------------------|
    |  ipv4net  |  Local IPv4 prefix  |
    |  ipv6net  |  Local IPv6 prefix  |
`,
		},

		// TagNodes

		// Nodes

	}
}
