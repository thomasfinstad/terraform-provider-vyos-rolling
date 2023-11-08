// Package namedvrfnameprotocolsospfareavirtuallink code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfareavirtuallink

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/vrf/name-protocols-ospf-area-virtual-link/resourcemodel"
)

// NewVrfNameProtocolsOspfAreaVirtualLink method to return the example resource reference
func NewVrfNameProtocolsOspfAreaVirtualLink() resource.Resource {
	return &vrfNameProtocolsOspfAreaVirtualLink{
		model: resourcemodel.VrfNameProtocolsOspfAreaVirtualLink{},
	}
}

// vrfNameProtocolsOspfAreaVirtualLink defines the resource implementation.
type vrfNameProtocolsOspfAreaVirtualLink struct {
	ResourceName string
	client       *client.Client
	model        resourcemodel.VrfNameProtocolsOspfAreaVirtualLink
}

// GetName returns resource name
func (r *vrfNameProtocolsOspfAreaVirtualLink) GetName() string {
	return r.ResourceName
}

func (r *vrfNameProtocolsOspfAreaVirtualLink) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
