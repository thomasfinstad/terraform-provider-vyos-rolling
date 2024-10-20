/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedvrfnameprotocolsisissegmentroutingprefix code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsisissegmentroutingprefix

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
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/vrf/name-protocols-isis-segment-routing-prefix/resourcemodel"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-tagnode-based-full.gotmpl */
// NewVrfNameProtocolsIsisSegmentRoutingPrefix method to return the example resource reference
func NewVrfNameProtocolsIsisSegmentRoutingPrefix() resource.Resource {
	return &vrfNameProtocolsIsisSegmentRoutingPrefix{
		model: &resourcemodel.VrfNameProtocolsIsisSegmentRoutingPrefix{},
	}
}

// vrfNameProtocolsIsisSegmentRoutingPrefix defines the resource implementation.
type vrfNameProtocolsIsisSegmentRoutingPrefix struct {
	providerData data.ProviderData
	model        *resourcemodel.VrfNameProtocolsIsisSegmentRoutingPrefix
}

// GetClient returns the vyos api client
func (r *vrfNameProtocolsIsisSegmentRoutingPrefix) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *vrfNameProtocolsIsisSegmentRoutingPrefix) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *vrfNameProtocolsIsisSegmentRoutingPrefix) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *vrfNameProtocolsIsisSegmentRoutingPrefix) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *vrfNameProtocolsIsisSegmentRoutingPrefix) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
