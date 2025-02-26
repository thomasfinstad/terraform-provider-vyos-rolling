/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedloadbalancingwanruleinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedloadbalancingwanruleinterface

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (interface) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &loadBalancingWanRuleInterface{}
	_ resource.ResourceWithConfigure   = &loadBalancingWanRuleInterface{}
	_ resource.ResourceWithImportState = &loadBalancingWanRuleInterface{}
)

// var _ resource.ResourceWithConfigValidators = &loadBalancingWanRuleInterface{}
// var _ resource.ResourceWithModifyPlan = &loadBalancingWanRuleInterface{}
// var _ resource.ResourceWithUpgradeState = &loadBalancingWanRuleInterface{}
// var _ resource.ResourceWithValidateConfig = &loadBalancingWanRuleInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &loadBalancingWanRuleInterface{}
