/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolspiminterfaceigmpjoin code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolspiminterfaceigmpjoin

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (join) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsPimInterfaceIgmpJoin{}
	_ resource.ResourceWithConfigure   = &protocolsPimInterfaceIgmpJoin{}
	_ resource.ResourceWithImportState = &protocolsPimInterfaceIgmpJoin{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsPimInterfaceIgmpJoin{}
// var _ resource.ResourceWithModifyPlan = &protocolsPimInterfaceIgmpJoin{}
// var _ resource.ResourceWithUpgradeState = &protocolsPimInterfaceIgmpJoin{}
// var _ resource.ResourceWithValidateConfig = &protocolsPimInterfaceIgmpJoin{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsPimInterfaceIgmpJoin{}
