// Package namednatstaticrule code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namednatstaticrule

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/nat/static-rule/resourcemodel"
)

// NewNatStaticRule method to return the example resource reference
func NewNatStaticRule() resource.Resource {
	return &natStaticRule{
		model: &resourcemodel.NatStaticRule{},
	}
}

// natStaticRule defines the resource implementation.
type natStaticRule struct {
	providerData data.ProviderData
	model        *resourcemodel.NatStaticRule
}

// GetClient returns the vyos api client
func (r *natStaticRule) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *natStaticRule) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *natStaticRule) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *natStaticRule) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
