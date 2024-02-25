package crud

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Delete method to define the logic which deletes the resource and removes the Terraform state on success.
func Delete(ctx context.Context, r helpers.VyosResource, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete Resource")
	tflog.Trace(ctx, "Fetching data model")
	stateModel := r.GetModel()

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, stateModel)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete resource
	vyosOps := [][]string{stateModel.GetVyosPath()}
	tflog.Error(ctx, "Compiling vyos operations", map[string]interface{}{"vyosOps": vyosOps})

	r.GetClient().StageDelete(ctx, vyosOps)

	// Commit changes to api
	response, err := r.GetClient().CommitChanges(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create %s, got error: %s", stateModel.GetVyosPath(), err))
		return
	}
	if response != nil {
		tflog.Error(ctx, "Got non-nil response from API", map[string]interface{}{"response": response})
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Got non-nil response from API: %s", response))
	}

	// Save data to Terraform state
	tflog.Info(ctx, "resource deleted")
}
