/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedvpnsstpclientipvsixpooldelegate code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvpnsstpclientipvsixpooldelegate

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
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/vpn/sstp-client-ipv6-pool-delegate/resourcemodel"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-tagnode-based-full.gotmpl #resource-tagnode-based-full (delegate) */
// NewVpnSstpClientIPvsixPoolDelegate method to return the example resource reference
func NewVpnSstpClientIPvsixPoolDelegate() resource.Resource {
	return &vpnSstpClientIPvsixPoolDelegate{
		model: &resourcemodel.VpnSstpClientIPvsixPoolDelegate{},
	}
}

// vpnSstpClientIPvsixPoolDelegate defines the resource implementation.
type vpnSstpClientIPvsixPoolDelegate struct {
	providerData data.ProviderData
	model        *resourcemodel.VpnSstpClientIPvsixPoolDelegate
}

// GetClient returns the vyos api client
func (r *vpnSstpClientIPvsixPoolDelegate) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *vpnSstpClientIPvsixPoolDelegate) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *vpnSstpClientIPvsixPoolDelegate) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *vpnSstpClientIPvsixPoolDelegate) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *vpnSstpClientIPvsixPoolDelegate) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
