// Package namedvrfnameprotocolsospfvthreearea code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfvthreearea

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/vrf/name-protocols-ospfv3-area/resourcemodel"
)

// NewVrfNameProtocolsOspfvthreeArea method to return the example resource reference
func NewVrfNameProtocolsOspfvthreeArea() resource.Resource {
	return &vrfNameProtocolsOspfvthreeArea{
		model: &resourcemodel.VrfNameProtocolsOspfvthreeArea{},
	}
}

// vrfNameProtocolsOspfvthreeArea defines the resource implementation.
type vrfNameProtocolsOspfvthreeArea struct {
	client *client.Client
	model  *resourcemodel.VrfNameProtocolsOspfvthreeArea
}

// GetClient returns the vyos api client
func (r *vrfNameProtocolsOspfvthreeArea) GetClient() *client.Client {
	return r.client
}

// GetModel returns the resource model
func (r *vrfNameProtocolsOspfvthreeArea) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

func (r *vrfNameProtocolsOspfvthreeArea) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = &client
}
