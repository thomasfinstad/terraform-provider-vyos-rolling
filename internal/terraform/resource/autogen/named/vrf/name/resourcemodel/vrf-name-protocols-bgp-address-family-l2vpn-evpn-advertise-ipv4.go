/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvfour{}

// VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvfour describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvfour struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeVrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvfourUnicast *VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvfourUnicast `tfsdk:"unicast" vyos:"unicast,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvfour) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"unicast": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpAddressFamilyLtwovpnEvpnAdvertiseIPvfourUnicast{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 address family

`,
			Description: `IPv4 address family

`,
		},
	}
}
