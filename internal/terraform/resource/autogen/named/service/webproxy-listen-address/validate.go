// Package namedservicewebproxylistenaddress code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicewebproxylistenaddress

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource              = &serviceWebproxyListenAddress{}
	_ resource.ResourceWithConfigure = &serviceWebproxyListenAddress{}
)

// var _ resource.ResourceWithConfigValidators = &serviceWebproxyListenAddress{}
// var _ resource.ResourceWithModifyPlan = &serviceWebproxyListenAddress{}
// var _ resource.ResourceWithUpgradeState = &serviceWebproxyListenAddress{}
// var _ resource.ResourceWithValidateConfig = &serviceWebproxyListenAddress{}
// var _ resource.ResourceWithImportState = &serviceWebproxyListenAddress{}

// Ensure we fully satisfy helper pkg interfaces
var _ helpers.VyosResource = &serviceWebproxyListenAddress{}