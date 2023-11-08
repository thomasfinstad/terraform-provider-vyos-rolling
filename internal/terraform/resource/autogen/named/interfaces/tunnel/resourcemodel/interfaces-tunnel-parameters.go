// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// InterfacesTunnelParameters describes the resource data model.
type InterfacesTunnelParameters struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeInterfacesTunnelParametersErspan *InterfacesTunnelParametersErspan `tfsdk:"erspan" vyos:"erspan,omitempty"`
	NodeInterfacesTunnelParametersIP     *InterfacesTunnelParametersIP     `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesTunnelParametersIPvsix *InterfacesTunnelParametersIPvsix `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesTunnelParameters) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"erspan": schema.SingleNestedAttribute{
			Attributes: InterfacesTunnelParametersErspan{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `ERSPAN tunnel parameters

`,
		},

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesTunnelParametersIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4-specific tunnel parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesTunnelParametersIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6-specific tunnel parameters

`,
		},
	}
}
