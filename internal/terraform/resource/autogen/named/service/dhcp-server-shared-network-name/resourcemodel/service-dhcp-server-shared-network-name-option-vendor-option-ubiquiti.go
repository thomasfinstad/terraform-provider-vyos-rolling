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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (ubiquiti) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceDhcpServerSharedNetworkNameOptionVendorOptionUbiquiti{}

// ServiceDhcpServerSharedNetworkNameOptionVendorOptionUbiquiti describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ServiceDhcpServerSharedNetworkNameOptionVendorOptionUbiquiti struct {
	// LeafNodes
	LeafServiceDhcpServerSharedNetworkNameOptionVendorOptionUbiquitiUnifiController types.String `tfsdk:"unifi_controller" vyos:"unifi-controller,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameOptionVendorOptionUbiquiti) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"unifi_controller":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (unifi-controller) */
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

		// TagNodes

		// Nodes

	}
}
