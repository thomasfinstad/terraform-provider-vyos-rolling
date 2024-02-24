// Package globalsystemoptiontcpclient code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalsystemoptiontcpclient

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/global/system/option-ssh-client/resourcemodel"
)

// NewSystemOptionTCPClient method to return the example resource reference
func NewSystemOptionTCPClient() resource.Resource {
	return &systemOptionTCPClient{
		model: &resourcemodel.SystemOptionTCPClient{},
	}
}

// systemOptionTCPClient defines the resource implementation.
type systemOptionTCPClient struct {
	client *client.Client
	model  *resourcemodel.SystemOptionTCPClient
}

// GetClient returns the vyos api client
func (r *systemOptionTCPClient) GetClient() *client.Client {
	return r.client
}

func (r *systemOptionTCPClient) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}
