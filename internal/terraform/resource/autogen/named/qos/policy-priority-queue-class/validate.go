/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedqospolicypriorityqueueclass code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicypriorityqueueclass

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (class) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &qosPolicyPriorityQueueClass{}
	_ resource.ResourceWithConfigure   = &qosPolicyPriorityQueueClass{}
	_ resource.ResourceWithImportState = &qosPolicyPriorityQueueClass{}
)

// var _ resource.ResourceWithConfigValidators = &qosPolicyPriorityQueueClass{}
// var _ resource.ResourceWithModifyPlan = &qosPolicyPriorityQueueClass{}
// var _ resource.ResourceWithUpgradeState = &qosPolicyPriorityQueueClass{}
// var _ resource.ResourceWithValidateConfig = &qosPolicyPriorityQueueClass{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosPolicyPriorityQueueClass{}
