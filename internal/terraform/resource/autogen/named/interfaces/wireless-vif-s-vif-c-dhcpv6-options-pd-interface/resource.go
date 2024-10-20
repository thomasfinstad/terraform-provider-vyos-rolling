/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedinterfaceswirelessvifsvifcdhcpvsixoptionspdinterface code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedinterfaceswirelessvifsvifcdhcpvsixoptionspdinterface

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
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/interfaces/wireless-vif-s-vif-c-dhcpv6-options-pd-interface/resourcemodel"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-tagnode-based-full.gotmpl */
// NewInterfacesWirelessVifSVifCDhcpvsixOptionsPdInterface method to return the example resource reference
func NewInterfacesWirelessVifSVifCDhcpvsixOptionsPdInterface() resource.Resource {
	return &interfacesWirelessVifSVifCDhcpvsixOptionsPdInterface{
		model: &resourcemodel.InterfacesWirelessVifSVifCDhcpvsixOptionsPdInterface{},
	}
}

// interfacesWirelessVifSVifCDhcpvsixOptionsPdInterface defines the resource implementation.
type interfacesWirelessVifSVifCDhcpvsixOptionsPdInterface struct {
	providerData data.ProviderData
	model        *resourcemodel.InterfacesWirelessVifSVifCDhcpvsixOptionsPdInterface
}

// GetClient returns the vyos api client
func (r *interfacesWirelessVifSVifCDhcpvsixOptionsPdInterface) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *interfacesWirelessVifSVifCDhcpvsixOptionsPdInterface) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *interfacesWirelessVifSVifCDhcpvsixOptionsPdInterface) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *interfacesWirelessVifSVifCDhcpvsixOptionsPdInterface) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *interfacesWirelessVifSVifCDhcpvsixOptionsPdInterface) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
