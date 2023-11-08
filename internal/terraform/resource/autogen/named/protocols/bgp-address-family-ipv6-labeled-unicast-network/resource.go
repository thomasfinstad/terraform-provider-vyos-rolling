// Package namedprotocolsbgpaddressfamilyipvsixlabeledunicastnetwork code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsbgpaddressfamilyipvsixlabeledunicastnetwork

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/protocols/bgp-address-family-ipv6-labeled-unicast-network/resourcemodel"
)

// NewProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork method to return the example resource reference
func NewProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork() resource.Resource {
	return &protocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork{
		model: resourcemodel.ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork{},
	}
}

// protocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork defines the resource implementation.
type protocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork struct {
	ResourceName string
	client       *client.Client
	model        resourcemodel.ProtocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork
}

// GetName returns resource name
func (r *protocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) GetName() string {
	return r.ResourceName
}

func (r *protocolsBgpAddressFamilyIPvsixLabeledUnicastNetwork) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
