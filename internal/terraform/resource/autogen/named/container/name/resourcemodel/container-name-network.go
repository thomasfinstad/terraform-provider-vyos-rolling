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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (network) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ContainerNameNetwork{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (network) */
// ContainerNameNetworkSelfIdentifier used by TagNodes to keep the id info
type ContainerNameNetworkSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (network) */

	ContainerNameNetwork types.String `tfsdk:"network" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (name) */

	ContainerName types.String `tfsdk:"name" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (container) */
}

// ContainerNameNetwork describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ContainerNameNetwork struct {
	// LeafNodes
	LeafContainerNameNetworkAddress types.List `tfsdk:"address" vyos:"address,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ContainerNameNetwork) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"address":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (address) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Assign static IP address to container

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
    |  ipv6    |  IPv6 address  |
`,
			Description: `Assign static IP address to container

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
    |  ipv6    |  IPv6 address  |
`,
		},

		// TagNodes

		// Nodes

	}
}
