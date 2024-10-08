// Package globalservicehttpscertificates code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicehttpscertificates

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceHTTPSCertificates{}
	_ resource.ResourceWithConfigure = &serviceHTTPSCertificates{}
)

// var _ resource.ResourceWithConfigValidators = &serviceHTTPSCertificates{}
// var _ resource.ResourceWithModifyPlan = &serviceHTTPSCertificates{}
// var _ resource.ResourceWithUpgradeState = &serviceHTTPSCertificates{}
// var _ resource.ResourceWithValidateConfig = &serviceHTTPSCertificates{}
// var _ resource.ResourceWithImportState = &serviceHTTPSCertificates{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceHTTPSCertificates{}
