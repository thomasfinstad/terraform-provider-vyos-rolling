/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsisisfastreroutelfalocaltiebreakerdownstreamindex code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsisisfastreroutelfalocaltiebreakerdownstreamindex

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (index) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndex{}
	_ resource.ResourceWithConfigure   = &protocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndex{}
	_ resource.ResourceWithImportState = &protocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndex{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndex{}
// var _ resource.ResourceWithModifyPlan = &protocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndex{}
// var _ resource.ResourceWithUpgradeState = &protocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndex{}
// var _ resource.ResourceWithValidateConfig = &protocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndex{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsIsisFastRerouteLfaLocalTiebreakerDownstreamIndex{}
