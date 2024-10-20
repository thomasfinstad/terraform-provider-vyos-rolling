/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedvrfnameprotocolsbgpaddressfamilyltwovpnevpnvni code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpaddressfamilyltwovpnevpnvni

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni{}
	_ resource.ResourceWithConfigure   = &vrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni{}
	_ resource.ResourceWithImportState = &vrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpAddressFamilyLtwovpnEvpnVni{}
