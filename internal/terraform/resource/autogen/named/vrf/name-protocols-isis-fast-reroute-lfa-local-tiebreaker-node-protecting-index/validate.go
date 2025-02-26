/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedvrfnameprotocolsisisfastreroutelfalocaltiebreakernodeprotectingindex code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsisisfastreroutelfalocaltiebreakernodeprotectingindex

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (index) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex{}
	_ resource.ResourceWithConfigure   = &vrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex{}
	_ resource.ResourceWithImportState = &vrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtectingIndex{}
