// Package globalprotocolsisisdefaultinformationoriginateipvfourlevelone code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsisisdefaultinformationoriginateipvfourlevelone

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/global/protocols/isis-default-information-originate-ipv4-level-1/resourcemodel"
)

// NewProtocolsIsisDefaultInformationOriginateIPvfourLevelOne method to return the example resource reference
func NewProtocolsIsisDefaultInformationOriginateIPvfourLevelOne() resource.Resource {
	return &protocolsIsisDefaultInformationOriginateIPvfourLevelOne{
		model: &resourcemodel.ProtocolsIsisDefaultInformationOriginateIPvfourLevelOne{},
	}
}

// protocolsIsisDefaultInformationOriginateIPvfourLevelOne defines the resource implementation.
type protocolsIsisDefaultInformationOriginateIPvfourLevelOne struct {
	client *client.Client
	model  *resourcemodel.ProtocolsIsisDefaultInformationOriginateIPvfourLevelOne
}

// GetClient returns the vyos api client
func (r *protocolsIsisDefaultInformationOriginateIPvfourLevelOne) GetClient() *client.Client {
	return r.client
}

func (r *protocolsIsisDefaultInformationOriginateIPvfourLevelOne) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
