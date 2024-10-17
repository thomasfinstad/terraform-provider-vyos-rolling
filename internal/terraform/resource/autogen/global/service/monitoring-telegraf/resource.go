// Package globalservicemonitoringtelegraf code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicemonitoringtelegraf

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/service/monitoring-telegraf/resourcemodel"
)

// NewServiceMonitoringTelegraf method to return the example resource reference
func NewServiceMonitoringTelegraf() resource.Resource {
	return &serviceMonitoringTelegraf{
		model: &resourcemodel.ServiceMonitoringTelegraf{},
	}
}

// serviceMonitoringTelegraf defines the resource implementation.
type serviceMonitoringTelegraf struct {
	providerData data.ProviderData
	model        *resourcemodel.ServiceMonitoringTelegraf
}

// GetClient returns the vyos api client
func (r *serviceMonitoringTelegraf) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *serviceMonitoringTelegraf) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *serviceMonitoringTelegraf) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *serviceMonitoringTelegraf) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *serviceMonitoringTelegraf) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
