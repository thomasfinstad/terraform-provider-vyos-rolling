// Package namedserviceeventhandlerevent code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedserviceeventhandlerevent

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named/service/event-handler-event/resourcemodel"
)

// NewServiceEventHandlerEvent method to return the example resource reference
func NewServiceEventHandlerEvent() resource.Resource {
	return &serviceEventHandlerEvent{
		model: resourcemodel.ServiceEventHandlerEvent{},
	}
}

// serviceEventHandlerEvent defines the resource implementation.
type serviceEventHandlerEvent struct {
	ResourceName string
	client       *client.Client
	model        resourcemodel.ServiceEventHandlerEvent
}

func (r *serviceEventHandlerEvent) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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