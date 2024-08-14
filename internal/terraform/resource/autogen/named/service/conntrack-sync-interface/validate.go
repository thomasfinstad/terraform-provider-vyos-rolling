// Package namedserviceconntracksyncinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedserviceconntracksyncinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceConntrackSyncInterface{}
	_ resource.ResourceWithConfigure = &serviceConntrackSyncInterface{}
)

// var _ resource.ResourceWithConfigValidators = &serviceConntrackSyncInterface{}
// var _ resource.ResourceWithModifyPlan = &serviceConntrackSyncInterface{}
// var _ resource.ResourceWithUpgradeState = &serviceConntrackSyncInterface{}
// var _ resource.ResourceWithValidateConfig = &serviceConntrackSyncInterface{}
// var _ resource.ResourceWithImportState = &serviceConntrackSyncInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceConntrackSyncInterface{}