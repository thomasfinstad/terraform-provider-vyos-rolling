/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceDhcpServerSharedNetworkNameOptionVendorOptionUbiquiti{}

// ServiceDhcpServerSharedNetworkNameOptionVendorOptionUbiquiti describes the resource data model.
type ServiceDhcpServerSharedNetworkNameOptionVendorOptionUbiquiti struct {
	// LeafNodes
	LeafServiceDhcpServerSharedNetworkNameOptionVendorOptionUbiquitiUnifiController types.String `tfsdk:"unifi_controller" vyos:"unifi-controller,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameOptionVendorOptionUbiquiti) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"unifi_controller":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Address of UniFi controller

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  ipv4    |  IP address of UniFi controller  |
`,
			Description: `Address of UniFi controller

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  ipv4    |  IP address of UniFi controller  |
`,
		},

		// Nodes

	}
}
