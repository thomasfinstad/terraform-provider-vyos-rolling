// Package globalserviceidsddosprotectionsflow code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceidsddosprotectionsflow

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceIDsDdosProtectionSflow{}
	_ resource.ResourceWithConfigure = &serviceIDsDdosProtectionSflow{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIDsDdosProtectionSflow{}
// var _ resource.ResourceWithModifyPlan = &serviceIDsDdosProtectionSflow{}
// var _ resource.ResourceWithUpgradeState = &serviceIDsDdosProtectionSflow{}
// var _ resource.ResourceWithValidateConfig = &serviceIDsDdosProtectionSflow{}
// var _ resource.ResourceWithImportState = &serviceIDsDdosProtectionSflow{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIDsDdosProtectionSflow{}
