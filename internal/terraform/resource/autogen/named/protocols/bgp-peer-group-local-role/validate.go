// Package namedprotocolsbgppeergrouplocalrole code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgppeergrouplocalrole

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &protocolsBgpPeerGroupLocalRole{}
	_ resource.ResourceWithConfigure   = &protocolsBgpPeerGroupLocalRole{}
	_ resource.ResourceWithImportState = &protocolsBgpPeerGroupLocalRole{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpPeerGroupLocalRole{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpPeerGroupLocalRole{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpPeerGroupLocalRole{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpPeerGroupLocalRole{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpPeerGroupLocalRole{}
