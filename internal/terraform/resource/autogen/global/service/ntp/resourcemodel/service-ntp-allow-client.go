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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (allow-client) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceNtpAllowClient{}

// ServiceNtpAllowClient describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ServiceNtpAllowClient struct {
	// LeafNodes
	LeafServiceNtpAllowClientAddress types.List `tfsdk:"address" vyos:"address,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceNtpAllowClient) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (address) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Allowed IP client addresses

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4     |  IPv4 address                    |
    |  ipv6     |  IPv6 address                    |
    |  ipv4net  |  IPv4 address and prefix length  |
    |  ipv6net  |  IPv6 address and prefix length  |
`,
			Description: `Allowed IP client addresses

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4     |  IPv4 address                    |
    |  ipv6     |  IPv6 address                    |
    |  ipv4net  |  IPv4 address and prefix length  |
    |  ipv6net  |  IPv6 address and prefix length  |
`,
		},

		// TagNodes

		// Nodes

	}
}
