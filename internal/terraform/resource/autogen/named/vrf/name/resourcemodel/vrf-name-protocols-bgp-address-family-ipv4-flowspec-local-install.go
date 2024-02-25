// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstall describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstall struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstallInterface types.List `tfsdk:"interface" vyos:"interface,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourFlowspecLocalInstall) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"interface": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Interface to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |

`,
		},

		// Nodes

	}
}
