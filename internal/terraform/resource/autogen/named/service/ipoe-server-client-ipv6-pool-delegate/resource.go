/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedserviceipoeserverclientipvsixpooldelegate code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedserviceipoeserverclientipvsixpooldelegate

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
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/service/ipoe-server-client-ipv6-pool-delegate/resourcemodel"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-tagnode-based-full.gotmpl #resource-tagnode-based-full (delegate) */
// NewServiceIPoeServerClientIPvsixPoolDelegate method to return the example resource reference
func NewServiceIPoeServerClientIPvsixPoolDelegate() resource.Resource {
	return &serviceIPoeServerClientIPvsixPoolDelegate{
		model: &resourcemodel.ServiceIPoeServerClientIPvsixPoolDelegate{},
	}
}

// serviceIPoeServerClientIPvsixPoolDelegate defines the resource implementation.
type serviceIPoeServerClientIPvsixPoolDelegate struct {
	providerData data.ProviderData
	model        *resourcemodel.ServiceIPoeServerClientIPvsixPoolDelegate
}

// GetClient returns the vyos api client
func (r *serviceIPoeServerClientIPvsixPoolDelegate) GetClient() client.Client {
	return r.providerData.Client
}

// GetModel returns the resource model
func (r *serviceIPoeServerClientIPvsixPoolDelegate) GetModel() helpers.VyosTopResourceDataModel {
	return r.model
}

// GetProviderConfig returns global provider data config
func (r *serviceIPoeServerClientIPvsixPoolDelegate) GetProviderConfig() data.ProviderData {
	return r.providerData
}

func (r *serviceIPoeServerClientIPvsixPoolDelegate) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *serviceIPoeServerClientIPvsixPoolDelegate) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
