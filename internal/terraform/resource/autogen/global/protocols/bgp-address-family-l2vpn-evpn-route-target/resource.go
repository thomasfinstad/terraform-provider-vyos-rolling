/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package globalprotocolsbgpaddressfamilyltwovpnevpnroutetarget code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package globalprotocolsbgpaddressfamilyltwovpnevpnroutetarget

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"

	// Extra Imports
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/protocols/bgp-address-family-l2vpn-evpn-route-target/resourcemodel"
)

/* tools/generate-terraform-resource-full/templates/resources/global/resource-node-based-full.gotmpl #resource-node-based-full (route-target) */
// NewProtocolsBgpAddressFamilyLtwovpnEvpnRouteTarget method to return the example resource reference
func NewProtocolsBgpAddressFamilyLtwovpnEvpnRouteTarget() resource.Resource {
	return &protocolsBgpAddressFamilyLtwovpnEvpnRouteTarget{
		model: &resourcemodel.ProtocolsBgpAddressFamilyLtwovpnEvpnRouteTarget{},
	}
}

// protocolsBgpAddressFamilyLtwovpnEvpnRouteTarget defines the resource implementation.
type protocolsBgpAddressFamilyLtwovpnEvpnRouteTarget struct {
	providerData data.ProviderData
	model        *resourcemodel.ProtocolsBgpAddressFamilyLtwovpnEvpnRouteTarget
}

// GetClient returns the vyos api client
func (r *protocolsBgpAddressFamilyLtwovpnEvpnRouteTarget) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *protocolsBgpAddressFamilyLtwovpnEvpnRouteTarget) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *protocolsBgpAddressFamilyLtwovpnEvpnRouteTarget) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *protocolsBgpAddressFamilyLtwovpnEvpnRouteTarget) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *protocolsBgpAddressFamilyLtwovpnEvpnRouteTarget) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
