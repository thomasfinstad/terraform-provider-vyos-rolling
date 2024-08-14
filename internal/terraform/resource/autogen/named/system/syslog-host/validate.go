// Package namedsystemsysloghost code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedsystemsysloghost

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &systemSyslogHost{}
	_ resource.ResourceWithConfigure = &systemSyslogHost{}
)

// var _ resource.ResourceWithConfigValidators = &systemSyslogHost{}
// var _ resource.ResourceWithModifyPlan = &systemSyslogHost{}
// var _ resource.ResourceWithUpgradeState = &systemSyslogHost{}
// var _ resource.ResourceWithValidateConfig = &systemSyslogHost{}
// var _ resource.ResourceWithImportState = &systemSyslogHost{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &systemSyslogHost{}