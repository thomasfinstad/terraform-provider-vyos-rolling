// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesWirelessVifSMirror{}

// InterfacesWirelessVifSMirror describes the resource data model.
type InterfacesWirelessVifSMirror struct {
	// LeafNodes
	LeafInterfacesWirelessVifSMirrorIngress types.String `tfsdk:"ingress" vyos:"ingress,omitempty"`
	LeafInterfacesWirelessVifSMirrorEgress  types.String `tfsdk:"egress" vyos:"egress,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesWirelessVifSMirror) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ingress": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Mirror ingress traffic to destination interface

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  Destination interface name  |
`,
			Description: `Mirror ingress traffic to destination interface

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  Destination interface name  |
`,
		},

		"egress": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Mirror egress traffic to destination interface

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  Destination interface name  |
`,
			Description: `Mirror egress traffic to destination interface

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  Destination interface name  |
`,
		},

		// Nodes

	}
}
