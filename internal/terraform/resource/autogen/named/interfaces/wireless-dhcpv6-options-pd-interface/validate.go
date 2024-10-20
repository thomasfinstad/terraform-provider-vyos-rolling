/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedinterfaceswirelessdhcpvsixoptionspdinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswirelessdhcpvsixoptionspdinterface

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/validate.gotmpl */
// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &interfacesWirelessDhcpvsixOptionsPdInterface{}
	_ resource.ResourceWithConfigure   = &interfacesWirelessDhcpvsixOptionsPdInterface{}
	_ resource.ResourceWithImportState = &interfacesWirelessDhcpvsixOptionsPdInterface{}
)

// var _ resource.ResourceWithConfigValidators = &interfacesWirelessDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithModifyPlan = &interfacesWirelessDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithUpgradeState = &interfacesWirelessDhcpvsixOptionsPdInterface{}
// var _ resource.ResourceWithValidateConfig = &interfacesWirelessDhcpvsixOptionsPdInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &interfacesWirelessDhcpvsixOptionsPdInterface{}
