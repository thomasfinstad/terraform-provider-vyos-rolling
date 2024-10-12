// Package namedservicepppoeserverauthenticationradiusserver code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicepppoeserverauthenticationradiusserver

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/service/pppoe-server-authentication-radius-server/resourcemodel"
)

// NewServicePppoeServerAuthenticationRadiusServer method to return the example resource reference
func NewServicePppoeServerAuthenticationRadiusServer() resource.Resource {
	return &servicePppoeServerAuthenticationRadiusServer{
		model: &resourcemodel.ServicePppoeServerAuthenticationRadiusServer{},
	}
}

// servicePppoeServerAuthenticationRadiusServer defines the resource implementation.
type servicePppoeServerAuthenticationRadiusServer struct {
	providerData data.ProviderData
	model        *resourcemodel.ServicePppoeServerAuthenticationRadiusServer
}

// GetClient returns the vyos api client
func (r *servicePppoeServerAuthenticationRadiusServer) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *servicePppoeServerAuthenticationRadiusServer) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *servicePppoeServerAuthenticationRadiusServer) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *servicePppoeServerAuthenticationRadiusServer) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
