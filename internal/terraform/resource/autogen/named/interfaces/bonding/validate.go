/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedinterfacesbonding code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesbonding

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (bonding) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesBonding{}
	_ resource.ResourceWithConfigure   = &interfacesBonding{}
	_ resource.ResourceWithImportState = &interfacesBonding{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesBonding{}
// var _ resource.ResourceWithModifyPlan = &interfacesBonding{}
// var _ resource.ResourceWithUpgradeState = &interfacesBonding{}
// var _ resource.ResourceWithValidateConfig = &interfacesBonding{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesBonding{}
