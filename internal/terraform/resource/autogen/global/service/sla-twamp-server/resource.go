// Package globalserviceslatwampserver code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalserviceslatwampserver

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/global/service/sla-twamp-server/resourcemodel"
)

// NewServiceSLATwampServer method to return the example resource reference
func NewServiceSLATwampServer() resource.Resource {
	return &serviceSLATwampServer{
		model: &resourcemodel.ServiceSLATwampServer{},
	}
}

// serviceSLATwampServer defines the resource implementation.
type serviceSLATwampServer struct {
	client *client.Client
	model  *resourcemodel.ServiceSLATwampServer
}

// GetClient returns the vyos api client
func (r *serviceSLATwampServer) GetClient() *client.Client {
	return r.client
}

func (r *serviceSLATwampServer) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
