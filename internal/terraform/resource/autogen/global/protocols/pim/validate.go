/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolspim code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolspim

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (pim) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsPim{}
	_ resource.ResourceWithConfigure   = &protocolsPim{}
	_ resource.ResourceWithImportState = &protocolsPim{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsPim{}
// var _ resource.ResourceWithModifyPlan = &protocolsPim{}
// var _ resource.ResourceWithUpgradeState = &protocolsPim{}
// var _ resource.ResourceWithValidateConfig = &protocolsPim{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsPim{}
