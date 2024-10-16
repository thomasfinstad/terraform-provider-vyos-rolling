// Package namedpkiopenvpnsharedsecret code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedpkiopenvpnsharedsecret

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &pkiOpenvpnSharedSecret{}
	_ resource.ResourceWithConfigure   = &pkiOpenvpnSharedSecret{}
	_ resource.ResourceWithImportState = &pkiOpenvpnSharedSecret{}
)

// var _ resource.ResourceWithConfigValidators = &pkiOpenvpnSharedSecret{}
// var _ resource.ResourceWithModifyPlan = &pkiOpenvpnSharedSecret{}
// var _ resource.ResourceWithUpgradeState = &pkiOpenvpnSharedSecret{}
// var _ resource.ResourceWithValidateConfig = &pkiOpenvpnSharedSecret{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &pkiOpenvpnSharedSecret{}
