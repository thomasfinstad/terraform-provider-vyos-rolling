/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package globalfirewallglobaloptions code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalfirewallglobaloptions

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallGlobalOptions{}
	_ resource.ResourceWithConfigure   = &firewallGlobalOptions{}
	_ resource.ResourceWithImportState = &firewallGlobalOptions{}
)

// var _ resource.ResourceWithConfigValidators = &firewallGlobalOptions{}
// var _ resource.ResourceWithModifyPlan = &firewallGlobalOptions{}
// var _ resource.ResourceWithUpgradeState = &firewallGlobalOptions{}
// var _ resource.ResourceWithValidateConfig = &firewallGlobalOptions{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &firewallGlobalOptions{}
