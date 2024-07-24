// Package namedvrfnameprotocolsbgpinterface code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsbgpinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &vrfNameProtocolsBgpInterface{}
	_ resource.ResourceWithConfigure = &vrfNameProtocolsBgpInterface{}
)

// var _ resource.ResourceWithConfigValidators = &vrfNameProtocolsBgpInterface{}
// var _ resource.ResourceWithModifyPlan = &vrfNameProtocolsBgpInterface{}
// var _ resource.ResourceWithUpgradeState = &vrfNameProtocolsBgpInterface{}
// var _ resource.ResourceWithValidateConfig = &vrfNameProtocolsBgpInterface{}
// var _ resource.ResourceWithImportState = &vrfNameProtocolsBgpInterface{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &vrfNameProtocolsBgpInterface{}
