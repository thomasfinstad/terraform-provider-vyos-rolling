// Package namedpkikeypair code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedpkikeypair

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &pkiKeyPair{}
	_ resource.ResourceWithConfigure = &pkiKeyPair{}
)

// var _ resource.ResourceWithConfigValidators = &pkiKeyPair{}
// var _ resource.ResourceWithModifyPlan = &pkiKeyPair{}
// var _ resource.ResourceWithUpgradeState = &pkiKeyPair{}
// var _ resource.ResourceWithValidateConfig = &pkiKeyPair{}
// var _ resource.ResourceWithImportState = &pkiKeyPair{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &pkiKeyPair{}
