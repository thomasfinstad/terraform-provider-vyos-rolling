// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesOpenvpnServerClientIPPool{}

// InterfacesOpenvpnServerClientIPPool describes the resource data model.
type InterfacesOpenvpnServerClientIPPool struct {
	// LeafNodes
	LeafInterfacesOpenvpnServerClientIPPoolDisable    types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesOpenvpnServerClientIPPoolStart      types.String `tfsdk:"start" vyos:"start,omitempty"`
	LeafInterfacesOpenvpnServerClientIPPoolStop       types.String `tfsdk:"stop" vyos:"stop,omitempty"`
	LeafInterfacesOpenvpnServerClientIPPoolSubnetMask types.String `tfsdk:"subnet_mask" vyos:"subnet-mask,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnServerClientIPPool) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"start": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `First IP address in the pool

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
`,
			Description: `First IP address in the pool

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
`,
		},

		"stop": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Last IP address in the pool

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
`,
			Description: `Last IP address in the pool

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
`,
		},

		"subnet_mask": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Subnet mask pushed to dynamic clients. If not set the server subnet mask will be used. Only used with topology subnet or device type tap. Not used with bridged interfaces.

    |  Format  |  Description       |
    |----------|--------------------|
    |  ipv4    |  IPv4 subnet mask  |
`,
			Description: `Subnet mask pushed to dynamic clients. If not set the server subnet mask will be used. Only used with topology subnet or device type tap. Not used with bridged interfaces.

    |  Format  |  Description       |
    |----------|--------------------|
    |  ipv4    |  IPv4 subnet mask  |
`,
		},

		// Nodes

	}
}
