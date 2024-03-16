// Package namedqospolicynetworkemulator code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicynetworkemulator

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/qos/policy-network-emulator/resourcemodel"
)

// NewQosPolicyNetworkEmulator method to return the example resource reference
func NewQosPolicyNetworkEmulator() resource.Resource {
	return &qosPolicyNetworkEmulator{
		model: &resourcemodel.QosPolicyNetworkEmulator{},
	}
}

// qosPolicyNetworkEmulator defines the resource implementation.
type qosPolicyNetworkEmulator struct {
	providerData data.ProviderData
	model        *resourcemodel.QosPolicyNetworkEmulator
}

// GetClient returns the vyos api client
func (r *qosPolicyNetworkEmulator) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *qosPolicyNetworkEmulator) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

func (r *qosPolicyNetworkEmulator) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
