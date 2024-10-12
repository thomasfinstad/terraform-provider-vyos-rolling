// Package namedprotocolsbgppeergrouplocalas code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgppeergrouplocalas

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &protocolsBgpPeerGroupLocalAs{}
	_ resource.ResourceWithConfigure = &protocolsBgpPeerGroupLocalAs{}
)

// var _ resource.ResourceWithConfigValidators = &protocolsBgpPeerGroupLocalAs{}
// var _ resource.ResourceWithModifyPlan = &protocolsBgpPeerGroupLocalAs{}
// var _ resource.ResourceWithUpgradeState = &protocolsBgpPeerGroupLocalAs{}
// var _ resource.ResourceWithValidateConfig = &protocolsBgpPeerGroupLocalAs{}
// var _ resource.ResourceWithImportState = &protocolsBgpPeerGroupLocalAs{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &protocolsBgpPeerGroupLocalAs{}
