/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsstaticneighborproxynd code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsstaticneighborproxynd

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (nd) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsStaticNeighborProxyNd{}
	_ resource.ResourceWithConfigure   = &protocolsStaticNeighborProxyNd{}
	_ resource.ResourceWithImportState = &protocolsStaticNeighborProxyNd{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsStaticNeighborProxyNd{}
// var _ resource.ResourceWithModifyPlan = &protocolsStaticNeighborProxyNd{}
// var _ resource.ResourceWithUpgradeState = &protocolsStaticNeighborProxyNd{}
// var _ resource.ResourceWithValidateConfig = &protocolsStaticNeighborProxyNd{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsStaticNeighborProxyNd{}
