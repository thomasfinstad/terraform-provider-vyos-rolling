// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceDhcpServerSharedNetworkNameSubnetRangeOptionVendorOptionUbiquiti{}

// ServiceDhcpServerSharedNetworkNameSubnetRangeOptionVendorOptionUbiquiti describes the resource data model.
type ServiceDhcpServerSharedNetworkNameSubnetRangeOptionVendorOptionUbiquiti struct {
	// LeafNodes
	LeafServiceDhcpServerSharedNetworkNameSubnetRangeOptionVendorOptionUbiquitiUnifiController types.String `tfsdk:"unifi_controller" vyos:"unifi-controller,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameSubnetRangeOptionVendorOptionUbiquiti) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"unifi_controller": schema.StringAttribute{
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
