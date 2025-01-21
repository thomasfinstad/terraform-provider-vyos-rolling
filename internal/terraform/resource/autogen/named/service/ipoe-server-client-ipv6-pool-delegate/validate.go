/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedserviceipoeserverclientipvsixpooldelegate code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedserviceipoeserverclientipvsixpooldelegate

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (delegate) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceIPoeServerClientIPvsixPoolDelegate{}
	_ resource.ResourceWithConfigure   = &serviceIPoeServerClientIPvsixPoolDelegate{}
	_ resource.ResourceWithImportState = &serviceIPoeServerClientIPvsixPoolDelegate{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIPoeServerClientIPvsixPoolDelegate{}
// var _ resource.ResourceWithModifyPlan = &serviceIPoeServerClientIPvsixPoolDelegate{}
// var _ resource.ResourceWithUpgradeState = &serviceIPoeServerClientIPvsixPoolDelegate{}
// var _ resource.ResourceWithValidateConfig = &serviceIPoeServerClientIPvsixPoolDelegate{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIPoeServerClientIPvsixPoolDelegate{}
