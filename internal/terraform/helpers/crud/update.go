package crud

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Update method to define the logic which updates the resource and sets the updated Terraform state on success.
func Update(ctx context.Context, r helpers.VyosResource, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update Resource")
	tflog.Trace(ctx, "Fetching data model")
	planModel := r.GetModel()
	stateModel := r.GetModel()

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, planModel)...)
	resp.Diagnostics.Append(req.State.Get(ctx, stateModel)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing config
	stateVyosData, err := helpers.MarshalVyos(ctx, stateModel)
	if err != nil {
		resp.Diagnostics.AddError("API Marshalling error", fmt.Sprintf("%s", err))
		return
	}

	vyosOpsState := helpers.GenerateVyosOps(ctx, stateModel.GetVyosPath(), stateVyosData)
	tflog.Error(ctx, "Compiling vyos state operations", map[string]interface{}{"vyosOpsState": vyosOpsState})

	r.GetClient().StageDelete(ctx, vyosOpsState)

	// Create new config
	planVyosData, err := helpers.MarshalVyos(ctx, planModel)
	if err != nil {
		resp.Diagnostics.AddError("API Marshalling error", fmt.Sprintf("%s", err))
		return
	}

	vyosOpsPlan := helpers.GenerateVyosOps(ctx, planModel.GetVyosPath(), planVyosData)
	tflog.Error(ctx, "Compiling vyos plan operations", map[string]interface{}{"vyosOpsPlan": vyosOpsPlan})

	r.GetClient().StageSet(ctx, vyosOpsPlan)

	// Commit changes to api
	response, err := r.GetClient().CommitChanges(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create %s, got error: %s", planModel.GetVyosPath(), err))
		return
	}
	if response != nil {
		tflog.Warn(ctx, "Got non-nil response from API", map[string]interface{}{"response": response})
	}

	// Add ID to the resource model as plan fetching do not include it
	planModel.SetID(planModel.GetVyosPath())

	// Save data to Terraform state
	tflog.Trace(ctx, "resource updated")
	tflog.Error(ctx, "Setting state", map[string]interface{}{"data": fmt.Sprintf("%#v", planModel)})
	resp.Diagnostics.Append(resp.State.Set(ctx, planModel)...)
}
