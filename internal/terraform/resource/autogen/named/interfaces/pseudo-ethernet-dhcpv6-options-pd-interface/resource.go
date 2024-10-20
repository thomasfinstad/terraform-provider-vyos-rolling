/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedinterfacespseudoethernetdhcpvsixoptionspdinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfacespseudoethernetdhcpvsixoptionspdinterface

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/interfaces/pseudo-ethernet-dhcpv6-options-pd-interface/resourcemodel"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-tagnode-based-full.gotmpl */
// NewInterfacesPseudoEthernetDhcpvsixOptionsPdInterface method to return the example resource reference
func NewInterfacesPseudoEthernetDhcpvsixOptionsPdInterface() resource.Resource {
	return &interfacesPseudoEthernetDhcpvsixOptionsPdInterface{
		model: &resourcemodel.InterfacesPseudoEthernetDhcpvsixOptionsPdInterface{},
	}
}

// interfacesPseudoEthernetDhcpvsixOptionsPdInterface defines the resource implementation.
type interfacesPseudoEthernetDhcpvsixOptionsPdInterface struct {
	providerData data.ProviderData
	model        *resourcemodel.InterfacesPseudoEthernetDhcpvsixOptionsPdInterface
}

// GetClient returns the vyos api client
func (r *interfacesPseudoEthernetDhcpvsixOptionsPdInterface) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *interfacesPseudoEthernetDhcpvsixOptionsPdInterface) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *interfacesPseudoEthernetDhcpvsixOptionsPdInterface) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *interfacesPseudoEthernetDhcpvsixOptionsPdInterface) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *interfacesPseudoEthernetDhcpvsixOptionsPdInterface) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
