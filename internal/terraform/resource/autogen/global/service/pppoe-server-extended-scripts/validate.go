// Package globalservicepppoeserverextendedscripts code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicepppoeserverextendedscripts

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &servicePppoeServerExtendedScrIPts{}
	_ resource.ResourceWithConfigure = &servicePppoeServerExtendedScrIPts{}
)

// var _ resource.ResourceWithConfigValidators = &servicePppoeServerExtendedScrIPts{}
// var _ resource.ResourceWithModifyPlan = &servicePppoeServerExtendedScrIPts{}
// var _ resource.ResourceWithUpgradeState = &servicePppoeServerExtendedScrIPts{}
// var _ resource.ResourceWithValidateConfig = &servicePppoeServerExtendedScrIPts{}
// var _ resource.ResourceWithImportState = &servicePppoeServerExtendedScrIPts{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &servicePppoeServerExtendedScrIPts{}