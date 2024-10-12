// Package globalservicetcpaccesscontroldeny code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicetcpaccesscontroldeny

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceTCPAccessControlDeny{}
	_ resource.ResourceWithConfigure = &serviceTCPAccessControlDeny{}
)

// var _ resource.ResourceWithConfigValidators = &serviceTCPAccessControlDeny{}
// var _ resource.ResourceWithModifyPlan = &serviceTCPAccessControlDeny{}
// var _ resource.ResourceWithUpgradeState = &serviceTCPAccessControlDeny{}
// var _ resource.ResourceWithValidateConfig = &serviceTCPAccessControlDeny{}
// var _ resource.ResourceWithImportState = &serviceTCPAccessControlDeny{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceTCPAccessControlDeny{}
