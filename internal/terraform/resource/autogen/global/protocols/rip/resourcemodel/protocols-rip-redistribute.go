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

var _ helpers.VyosResourceDataModel = &ProtocolsRIPRedistribute{}

// ProtocolsRIPRedistribute describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ProtocolsRIPRedistribute struct {
	// LeafNodes

	// TagNodes

	// Nodes

	ExistsNodeProtocolsRIPRedistributeBgp bool `tfsdk:"-" vyos:"bgp,child"`

	ExistsNodeProtocolsRIPRedistributeConnected bool `tfsdk:"-" vyos:"connected,child"`

	ExistsNodeProtocolsRIPRedistributeIsis bool `tfsdk:"-" vyos:"isis,child"`

	ExistsNodeProtocolsRIPRedistributeKernel bool `tfsdk:"-" vyos:"kernel,child"`

	ExistsNodeProtocolsRIPRedistributeOspf bool `tfsdk:"-" vyos:"ospf,child"`

	ExistsNodeProtocolsRIPRedistributeStatic bool `tfsdk:"-" vyos:"static,child"`

	ExistsNodeProtocolsRIPRedistributeBabel bool `tfsdk:"-" vyos:"babel,child"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPRedistribute) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{

		// LeafNodes

		// TagNodes

		// Nodes

	}
}
