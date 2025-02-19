/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (ipv4-unicast) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsBgpAddressFamilyIPvfourUnicast{}

// ProtocolsBgpAddressFamilyIPvfourUnicast describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ProtocolsBgpAddressFamilyIPvfourUnicast struct {
	// LeafNodes

	// TagNodes

	ExistsTagProtocolsBgpAddressFamilyIPvfourUnicastAggregateAddress bool `tfsdk:"-" vyos:"aggregate-address,child"`

	ExistsTagProtocolsBgpAddressFamilyIPvfourUnicastNetwork bool `tfsdk:"-" vyos:"network,child"`

	// Nodes

	ExistsNodeProtocolsBgpAddressFamilyIPvfourUnicastDistance bool `tfsdk:"-" vyos:"distance,child"`

	ExistsNodeProtocolsBgpAddressFamilyIPvfourUnicastExport bool `tfsdk:"-" vyos:"export,child"`

	ExistsNodeProtocolsBgpAddressFamilyIPvfourUnicastImport bool `tfsdk:"-" vyos:"import,child"`

	// Ignoring Node `ProtocolsBgpAddressFamilyIPvfourUnicastLabel`.

	ExistsNodeProtocolsBgpAddressFamilyIPvfourUnicastMaximumPaths bool `tfsdk:"-" vyos:"maximum-paths,child"`

	// Ignoring Node `ProtocolsBgpAddressFamilyIPvfourUnicastRd`.

	// Ignoring Node `ProtocolsBgpAddressFamilyIPvfourUnicastRouteMap`.

	// Ignoring Node `ProtocolsBgpAddressFamilyIPvfourUnicastRouteTarget`.

	// Ignoring Node `ProtocolsBgpAddressFamilyIPvfourUnicastNexthop`.

	// Ignoring Node `ProtocolsBgpAddressFamilyIPvfourUnicastRedistribute`.

	// Ignoring Node `ProtocolsBgpAddressFamilyIPvfourUnicastSID`.
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvfourUnicast) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{

		// LeafNodes

		// TagNodes

		// Nodes

	}
}
