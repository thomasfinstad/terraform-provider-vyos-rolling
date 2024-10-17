// Package globalservicetftpserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicetftpserver

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceTftpServer{}
	_ resource.ResourceWithConfigure   = &serviceTftpServer{}
	_ resource.ResourceWithImportState = &serviceTftpServer{}
)

// var _ resource.ResourceWithConfigValidators = &serviceTftpServer{}
// var _ resource.ResourceWithModifyPlan = &serviceTftpServer{}
// var _ resource.ResourceWithUpgradeState = &serviceTftpServer{}
// var _ resource.ResourceWithValidateConfig = &serviceTftpServer{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceTftpServer{}
