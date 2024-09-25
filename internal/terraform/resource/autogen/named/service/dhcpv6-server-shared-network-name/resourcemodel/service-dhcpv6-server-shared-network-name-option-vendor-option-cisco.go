// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceDhcpvsixServerSharedNetworkNameOptionVendorOptionCisco{}

// ServiceDhcpvsixServerSharedNetworkNameOptionVendorOptionCisco describes the resource data model.
type ServiceDhcpvsixServerSharedNetworkNameOptionVendorOptionCisco struct {
	// LeafNodes
	LeafServiceDhcpvsixServerSharedNetworkNameOptionVendorOptionCiscoTftpServer types.List `tfsdk:"tftp_server" vyos:"tftp-server,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpvsixServerSharedNetworkNameOptionVendorOptionCisco) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"tftp_server": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `TFTP server name

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv6    |  TFTP server IPv6 address  |
`,
			Description: `TFTP server name

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv6    |  TFTP server IPv6 address  |
`,
		},

		// Nodes

	}
}
