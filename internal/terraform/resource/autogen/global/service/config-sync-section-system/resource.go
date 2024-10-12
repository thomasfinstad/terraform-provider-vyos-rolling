// Package globalserviceconfigsyncsectionsystem code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceconfigsyncsectionsystem

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/service/config-sync-section-system/resourcemodel"
)

// NewServiceConfigSyncSectionSystem method to return the example resource reference
func NewServiceConfigSyncSectionSystem() resource.Resource {
	return &serviceConfigSyncSectionSystem{
		model: &resourcemodel.ServiceConfigSyncSectionSystem{},
	}
}

// serviceConfigSyncSectionSystem defines the resource implementation.
type serviceConfigSyncSectionSystem struct {
	providerData data.ProviderData
	model        *resourcemodel.ServiceConfigSyncSectionSystem
}

// GetClient returns the vyos api client
func (r *serviceConfigSyncSectionSystem) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *serviceConfigSyncSectionSystem) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *serviceConfigSyncSectionSystem) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *serviceConfigSyncSectionSystem) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(data.ProviderData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.providerData = data
}
