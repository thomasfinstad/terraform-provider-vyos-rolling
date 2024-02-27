// Package namedvrfnameprotocolsospfredistributetable code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfredistributetable

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/vrf/name-protocols-ospf-redistribute-table/resourcemodel"
)

// NewVrfNameProtocolsOspfRedistributeTable method to return the example resource reference
func NewVrfNameProtocolsOspfRedistributeTable() resource.Resource {
	return &vrfNameProtocolsOspfRedistributeTable{
		model: &resourcemodel.VrfNameProtocolsOspfRedistributeTable{},
	}
}

// vrfNameProtocolsOspfRedistributeTable defines the resource implementation.
type vrfNameProtocolsOspfRedistributeTable struct {
	client *client.Client
	model  *resourcemodel.VrfNameProtocolsOspfRedistributeTable
}

// GetClient returns the vyos api client
func (r *vrfNameProtocolsOspfRedistributeTable) GetClient() *client.Client {
	return r.client
}

// GetModel returns the resource model
func (r *vrfNameProtocolsOspfRedistributeTable) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

func (r *vrfNameProtocolsOspfRedistributeTable) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
