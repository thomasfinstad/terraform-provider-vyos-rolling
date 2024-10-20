/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedserviceipoeserverinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedserviceipoeserverinterface

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceIPoeServerInterface{}
	_ resource.ResourceWithConfigure   = &serviceIPoeServerInterface{}
	_ resource.ResourceWithImportState = &serviceIPoeServerInterface{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIPoeServerInterface{}
// var _ resource.ResourceWithModifyPlan = &serviceIPoeServerInterface{}
// var _ resource.ResourceWithUpgradeState = &serviceIPoeServerInterface{}
// var _ resource.ResourceWithValidateConfig = &serviceIPoeServerInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIPoeServerInterface{}
