// Package globalprotocolsigmpproxy code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsigmpproxy

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/global/protocols/igmp-proxy/resourcemodel"
)

// NewProtocolsIgmpProxy method to return the example resource reference
func NewProtocolsIgmpProxy() resource.Resource {
	return &protocolsIgmpProxy{
		model: &resourcemodel.ProtocolsIgmpProxy{},
	}
}

// protocolsIgmpProxy defines the resource implementation.
type protocolsIgmpProxy struct {
	client *client.Client
	model  *resourcemodel.ProtocolsIgmpProxy
}

// GetClient returns the vyos api client
func (r *protocolsIgmpProxy) GetClient() *client.Client {
	return r.client
}

func (r *protocolsIgmpProxy) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
