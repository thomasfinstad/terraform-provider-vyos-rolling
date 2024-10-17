// Package globalservicemonitoringtelegrafazuredataexplorerauthentication code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicemonitoringtelegrafazuredataexplorerauthentication

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/service/monitoring-telegraf-azure-data-explorer-authentication/resourcemodel"
)

// NewServiceMonitoringTelegrafAzureDataExplorerAuthentication method to return the example resource reference
func NewServiceMonitoringTelegrafAzureDataExplorerAuthentication() resource.Resource {
	return &serviceMonitoringTelegrafAzureDataExplorerAuthentication{
		model: &resourcemodel.ServiceMonitoringTelegrafAzureDataExplorerAuthentication{},
	}
}

// serviceMonitoringTelegrafAzureDataExplorerAuthentication defines the resource implementation.
type serviceMonitoringTelegrafAzureDataExplorerAuthentication struct {
	providerData data.ProviderData
	model        *resourcemodel.ServiceMonitoringTelegrafAzureDataExplorerAuthentication
}

// GetClient returns the vyos api client
func (r *serviceMonitoringTelegrafAzureDataExplorerAuthentication) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *serviceMonitoringTelegrafAzureDataExplorerAuthentication) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *serviceMonitoringTelegrafAzureDataExplorerAuthentication) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *serviceMonitoringTelegrafAzureDataExplorerAuthentication) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *serviceMonitoringTelegrafAzureDataExplorerAuthentication) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
