/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedqostrafficmatchgroupmatch code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedqostrafficmatchgroupmatch

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (match) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &qosTrafficMatchGroupMatch{}
	_ resource.ResourceWithConfigure   = &qosTrafficMatchGroupMatch{}
	_ resource.ResourceWithImportState = &qosTrafficMatchGroupMatch{}
)

// var _ resource.ResourceWithConfigValidators = &qosTrafficMatchGroupMatch{}
// var _ resource.ResourceWithModifyPlan = &qosTrafficMatchGroupMatch{}
// var _ resource.ResourceWithUpgradeState = &qosTrafficMatchGroupMatch{}
// var _ resource.ResourceWithValidateConfig = &qosTrafficMatchGroupMatch{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosTrafficMatchGroupMatch{}
