// Package namedqostrafficmatchgroup code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedqostrafficmatchgroup

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &qosTrafficMatchGroup{}
	_ resource.ResourceWithConfigure   = &qosTrafficMatchGroup{}
	_ resource.ResourceWithImportState = &qosTrafficMatchGroup{}
)

// var _ resource.ResourceWithConfigValidators = &qosTrafficMatchGroup{}
// var _ resource.ResourceWithModifyPlan = &qosTrafficMatchGroup{}
// var _ resource.ResourceWithUpgradeState = &qosTrafficMatchGroup{}
// var _ resource.ResourceWithValidateConfig = &qosTrafficMatchGroup{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &qosTrafficMatchGroup{}
