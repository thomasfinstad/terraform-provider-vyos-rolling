// Package namedservicentpserver code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicentpserver

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceNtpServer{}
	_ resource.ResourceWithConfigure = &serviceNtpServer{}
)

// var _ resource.ResourceWithConfigValidators = &serviceNtpServer{}
// var _ resource.ResourceWithModifyPlan = &serviceNtpServer{}
// var _ resource.ResourceWithUpgradeState = &serviceNtpServer{}
// var _ resource.ResourceWithValidateConfig = &serviceNtpServer{}
// var _ resource.ResourceWithImportState = &serviceNtpServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceNtpServer{}