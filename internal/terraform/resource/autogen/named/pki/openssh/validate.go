// Package namedpkiopentcp code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpkiopentcp

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &pkiOpenTCP{}
	_ resource.ResourceWithConfigure = &pkiOpenTCP{}
)

// var _ resource.ResourceWithConfigValidators = &pkiOpenTCP{}
// var _ resource.ResourceWithModifyPlan = &pkiOpenTCP{}
// var _ resource.ResourceWithUpgradeState = &pkiOpenTCP{}
// var _ resource.ResourceWithValidateConfig = &pkiOpenTCP{}
// var _ resource.ResourceWithImportState = &pkiOpenTCP{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &pkiOpenTCP{}
