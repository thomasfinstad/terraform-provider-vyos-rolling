// Package globalservicesaltminion code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicesaltminion

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/service/salt-minion/resourcemodel"
)

// NewServiceSaltMinion method to return the example resource reference
func NewServiceSaltMinion() resource.Resource {
	return &serviceSaltMinion{
		model: &resourcemodel.ServiceSaltMinion{},
	}
}

// serviceSaltMinion defines the resource implementation.
type serviceSaltMinion struct {
	providerData data.ProviderData
	model        *resourcemodel.ServiceSaltMinion
}

// GetClient returns the vyos api client
func (r *serviceSaltMinion) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *serviceSaltMinion) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *serviceSaltMinion) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *serviceSaltMinion) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
