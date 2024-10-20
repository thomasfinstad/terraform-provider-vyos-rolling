/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedprotocolsnhrptunnel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsnhrptunnel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers/crud"
)

/* tools/generate-terraform-resource-full/templates/resources/common/crud.gotmpl */
// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *protocolsNhrpTunnel) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	crud.Create(ctx, r, req, resp)
}

// Read method to define the logic which refreshes the Terraform state for the resource.
func (r *protocolsNhrpTunnel) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	crud.Read(ctx, r, req, resp)
}

// Update method to define the logic which updates the resource and sets the updated Terraform state on success.
func (r *protocolsNhrpTunnel) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	crud.Update(ctx, r, req, resp)
}

// Delete method to define the logic which deletes the resource and removes the Terraform state on success.
func (r *protocolsNhrpTunnel) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	crud.Delete(ctx, r, req, resp)
}
