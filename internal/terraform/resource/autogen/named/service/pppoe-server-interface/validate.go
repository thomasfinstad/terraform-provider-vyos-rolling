// Package namedservicepppoeserverinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicepppoeserverinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &servicePppoeServerInterface{}
	_ resource.ResourceWithConfigure = &servicePppoeServerInterface{}
)

// var _ resource.ResourceWithConfigValidators = &servicePppoeServerInterface{}
// var _ resource.ResourceWithModifyPlan = &servicePppoeServerInterface{}
// var _ resource.ResourceWithUpgradeState = &servicePppoeServerInterface{}
// var _ resource.ResourceWithValidateConfig = &servicePppoeServerInterface{}
// var _ resource.ResourceWithImportState = &servicePppoeServerInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &servicePppoeServerInterface{}
