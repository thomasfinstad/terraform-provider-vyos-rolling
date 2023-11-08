// Package namedprotocolsospfarearange code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsospfarearange

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// GetClient returns the vyos api client
func (r *protocolsOspfAreaRange) GetClient() *client.Client {
	return r.client
}

// GetModel returns the resource model as defined by the interface in helpers
func (r *protocolsOspfAreaRange) GetModel() helpers.VyosTopResourceDataModel {
	return &r.model
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *protocolsOspfAreaRange) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	helpers.Create(ctx, r, req, resp)
}

// Read method to define the logic which refreshes the Terraform state for the resource.
func (r *protocolsOspfAreaRange) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	helpers.Read(ctx, r, req, resp)
}

// Update method to define the logic which updates the resource and sets the updated Terraform state on success.
func (r *protocolsOspfAreaRange) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	helpers.Update(ctx, r, req, resp)
}

// Delete method to define the logic which deletes the resource and removes the Terraform state on success.
func (r *protocolsOspfAreaRange) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	helpers.Delete(ctx, r, req, resp)
}
