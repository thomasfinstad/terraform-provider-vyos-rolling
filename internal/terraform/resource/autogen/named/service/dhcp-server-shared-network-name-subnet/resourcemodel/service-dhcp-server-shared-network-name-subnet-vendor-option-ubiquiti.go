// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceDhcpServerSharedNetworkNameSubnetVendorOptionUbiquiti describes the resource data model.
type ServiceDhcpServerSharedNetworkNameSubnetVendorOptionUbiquiti struct {
	// LeafNodes
	LeafServiceDhcpServerSharedNetworkNameSubnetVendorOptionUbiquitiUnifiController types.String `tfsdk:"unifi_controller" vyos:"unifi-controller,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameSubnetVendorOptionUbiquiti) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"unifi_controller": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Address of UniFi controller

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IP address of UniFi controller  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceDhcpServerSharedNetworkNameSubnetVendorOptionUbiquiti) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceDhcpServerSharedNetworkNameSubnetVendorOptionUbiquiti) UnmarshalJSON(_ []byte) error {
	return nil
}