// Package globalservicemonitoringtelegrafsplunkauthentication code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicemonitoringtelegrafsplunkauthentication

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/service/monitoring-telegraf-splunk-authentication/resourcemodel"
)

// NewServiceMonitoringTelegrafSplunkAuthentication method to return the example resource reference
func NewServiceMonitoringTelegrafSplunkAuthentication() resource.Resource {
	return &serviceMonitoringTelegrafSplunkAuthentication{
		model: &resourcemodel.ServiceMonitoringTelegrafSplunkAuthentication{},
	}
}

// serviceMonitoringTelegrafSplunkAuthentication defines the resource implementation.
type serviceMonitoringTelegrafSplunkAuthentication struct {
	providerData data.ProviderData
	model        *resourcemodel.ServiceMonitoringTelegrafSplunkAuthentication
}

// GetClient returns the vyos api client
func (r *serviceMonitoringTelegrafSplunkAuthentication) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *serviceMonitoringTelegrafSplunkAuthentication) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *serviceMonitoringTelegrafSplunkAuthentication) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *serviceMonitoringTelegrafSplunkAuthentication) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
