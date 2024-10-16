// Package globalserviceipoeserverextendedscripts code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceipoeserverextendedscripts

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceIPoeServerExtendedScrIPts{}
	_ resource.ResourceWithConfigure   = &serviceIPoeServerExtendedScrIPts{}
	_ resource.ResourceWithImportState = &serviceIPoeServerExtendedScrIPts{}
)

// var _ resource.ResourceWithConfigValidators = &serviceIPoeServerExtendedScrIPts{}
// var _ resource.ResourceWithModifyPlan = &serviceIPoeServerExtendedScrIPts{}
// var _ resource.ResourceWithUpgradeState = &serviceIPoeServerExtendedScrIPts{}
// var _ resource.ResourceWithValidateConfig = &serviceIPoeServerExtendedScrIPts{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceIPoeServerExtendedScrIPts{}
