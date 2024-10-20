/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedvpnipsecikegroupproposal code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnipsecikegroupproposal

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vpnIPsecIkeGroupProposal{}
	_ resource.ResourceWithConfigure   = &vpnIPsecIkeGroupProposal{}
	_ resource.ResourceWithImportState = &vpnIPsecIkeGroupProposal{}
)

// var _ resource.ResourceWithConfigValidators = &vpnIPsecIkeGroupProposal{}
// var _ resource.ResourceWithModifyPlan = &vpnIPsecIkeGroupProposal{}
// var _ resource.ResourceWithUpgradeState = &vpnIPsecIkeGroupProposal{}
// var _ resource.ResourceWithValidateConfig = &vpnIPsecIkeGroupProposal{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vpnIPsecIkeGroupProposal{}
