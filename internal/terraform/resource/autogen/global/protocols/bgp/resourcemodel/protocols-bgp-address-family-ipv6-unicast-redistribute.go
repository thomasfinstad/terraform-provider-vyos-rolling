/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (redistribute) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsBgpAddressFamilyIPvsixUnicastRedistribute{}

// ProtocolsBgpAddressFamilyIPvsixUnicastRedistribute describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ProtocolsBgpAddressFamilyIPvsixUnicastRedistribute struct {
	// LeafNodes

	// TagNodes

	ExistsTagProtocolsBgpAddressFamilyIPvsixUnicastRedistributeTable bool `tfsdk:"-" vyos:"table,child"`

	// Nodes

	ExistsNodeProtocolsBgpAddressFamilyIPvsixUnicastRedistributeBabel bool `tfsdk:"-" vyos:"babel,child"`

	ExistsNodeProtocolsBgpAddressFamilyIPvsixUnicastRedistributeConnected bool `tfsdk:"-" vyos:"connected,child"`

	ExistsNodeProtocolsBgpAddressFamilyIPvsixUnicastRedistributeIsis bool `tfsdk:"-" vyos:"isis,child"`

	ExistsNodeProtocolsBgpAddressFamilyIPvsixUnicastRedistributeKernel bool `tfsdk:"-" vyos:"kernel,child"`

	ExistsNodeProtocolsBgpAddressFamilyIPvsixUnicastRedistributeStatic bool `tfsdk:"-" vyos:"static,child"`

	ExistsNodeProtocolsBgpAddressFamilyIPvsixUnicastRedistributeOspfvthree bool `tfsdk:"-" vyos:"ospfv3,child"`

	ExistsNodeProtocolsBgpAddressFamilyIPvsixUnicastRedistributeRIPng bool `tfsdk:"-" vyos:"ripng,child"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvsixUnicastRedistribute) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{

		// LeafNodes

		// TagNodes

		// Nodes

	}
}
