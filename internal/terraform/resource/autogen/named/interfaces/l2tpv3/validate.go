// Package namedinterfacesltwotpvthree code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacesltwotpvthree

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &interfacesLtwotpvthree{}
	_ resource.ResourceWithConfigure = &interfacesLtwotpvthree{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesLtwotpvthree{}
// var _ resource.ResourceWithModifyPlan = &interfacesLtwotpvthree{}
// var _ resource.ResourceWithUpgradeState = &interfacesLtwotpvthree{}
// var _ resource.ResourceWithValidateConfig = &interfacesLtwotpvthree{}
// var _ resource.ResourceWithImportState = &interfacesLtwotpvthree{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesLtwotpvthree{}
