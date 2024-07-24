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

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastSoftReconfiguration{}

// VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastSoftReconfiguration describes the resource data model.
type VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastSoftReconfiguration struct {
	// LeafNodes
	LeafVrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastSoftReconfigurationInbound types.Bool `tfsdk:"inbound" vyos:"inbound,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupAddressFamilyIPvsixLabeledUnicastSoftReconfiguration) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"inbound": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable inbound soft reconfiguration

`,
			Description: `Enable inbound soft reconfiguration

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
