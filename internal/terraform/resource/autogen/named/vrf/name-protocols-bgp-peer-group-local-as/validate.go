/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedvrfnameprotocolsbgppeergrouplocalas code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgppeergrouplocalas

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (local-as) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfNameProtocolsBgpPeerGroupLocalAs{}
	_ resource.ResourceWithConfigure   = &vrfNameProtocolsBgpPeerGroupLocalAs{}
	_ resource.ResourceWithImportState = &vrfNameProtocolsBgpPeerGroupLocalAs{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpPeerGroupLocalAs{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpPeerGroupLocalAs{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpPeerGroupLocalAs{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpPeerGroupLocalAs{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpPeerGroupLocalAs{}
