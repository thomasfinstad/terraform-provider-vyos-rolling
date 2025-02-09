/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalservicemonitoringzabbixagent code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicemonitoringzabbixagent

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl #validate (zabbix-agent) */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &serviceMonitoringZabbixAgent{}
	_ resource.ResourceWithConfigure   = &serviceMonitoringZabbixAgent{}
	_ resource.ResourceWithImportState = &serviceMonitoringZabbixAgent{}
)

// var _ resource.ResourceWithConfigValidators = &serviceMonitoringZabbixAgent{}
// var _ resource.ResourceWithModifyPlan = &serviceMonitoringZabbixAgent{}
// var _ resource.ResourceWithUpgradeState = &serviceMonitoringZabbixAgent{}
// var _ resource.ResourceWithValidateConfig = &serviceMonitoringZabbixAgent{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceMonitoringZabbixAgent{}
