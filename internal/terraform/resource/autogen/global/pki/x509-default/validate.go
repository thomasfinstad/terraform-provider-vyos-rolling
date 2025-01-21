/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalpkixfivezeroninedefault code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalpkixfivezeroninedefault

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (default) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &pkiXfivezeronineDefault{}
	_ resource.ResourceWithConfigure   = &pkiXfivezeronineDefault{}
	_ resource.ResourceWithImportState = &pkiXfivezeronineDefault{}
)

// var _ resource.ResourceWithConfigValidators = &pkiXfivezeronineDefault{}
// var _ resource.ResourceWithModifyPlan = &pkiXfivezeronineDefault{}
// var _ resource.ResourceWithUpgradeState = &pkiXfivezeronineDefault{}
// var _ resource.ResourceWithValidateConfig = &pkiXfivezeronineDefault{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &pkiXfivezeronineDefault{}
