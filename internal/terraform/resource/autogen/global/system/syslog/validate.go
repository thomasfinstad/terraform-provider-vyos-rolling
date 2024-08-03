// Package globalsystemsyslog code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemsyslog

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemSyslog{}
	_ resource.ResourceWithConfigure = &systemSyslog{}
)

// var _ resource.ResourceWithConfigValidators = &systemSyslog{}
// var _ resource.ResourceWithModifyPlan = &systemSyslog{}
// var _ resource.ResourceWithUpgradeState = &systemSyslog{}
// var _ resource.ResourceWithValidateConfig = &systemSyslog{}
// var _ resource.ResourceWithImportState = &systemSyslog{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemSyslog{}
