/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalprotocolspimregisteracceptlist code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolspimregisteracceptlist

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsPimRegisterAcceptList{}
	_ resource.ResourceWithConfigure   = &protocolsPimRegisterAcceptList{}
	_ resource.ResourceWithImportState = &protocolsPimRegisterAcceptList{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsPimRegisterAcceptList{}
// var _ resource.ResourceWithModifyPlan = &protocolsPimRegisterAcceptList{}
// var _ resource.ResourceWithUpgradeState = &protocolsPimRegisterAcceptList{}
// var _ resource.ResourceWithValidateConfig = &protocolsPimRegisterAcceptList{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsPimRegisterAcceptList{}
