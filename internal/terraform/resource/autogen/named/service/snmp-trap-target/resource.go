// Package namedservicesnmptraptarget code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicesnmptraptarget

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/service/snmp-trap-target/resourcemodel"
)

// NewServiceSnmpTrapTarget method to return the example resource reference
func NewServiceSnmpTrapTarget() resource.Resource {
	return &serviceSnmpTrapTarget{
		model: &resourcemodel.ServiceSnmpTrapTarget{},
	}
}

// serviceSnmpTrapTarget defines the resource implementation.
type serviceSnmpTrapTarget struct {
	providerData data.ProviderData
	model        *resourcemodel.ServiceSnmpTrapTarget
}

// GetClient returns the vyos api client
func (r *serviceSnmpTrapTarget) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *serviceSnmpTrapTarget) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *serviceSnmpTrapTarget) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *serviceSnmpTrapTarget) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *serviceSnmpTrapTarget) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
