// Package globalserviceidsddosprotection code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceidsddosprotection

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceIDsDdosProtection{}
	_ resource.ResourceWithConfigure = &serviceIDsDdosProtection{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIDsDdosProtection{}
// var _ resource.ResourceWithModifyPlan = &serviceIDsDdosProtection{}
// var _ resource.ResourceWithUpgradeState = &serviceIDsDdosProtection{}
// var _ resource.ResourceWithValidateConfig = &serviceIDsDdosProtection{}
// var _ resource.ResourceWithImportState = &serviceIDsDdosProtection{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIDsDdosProtection{}
