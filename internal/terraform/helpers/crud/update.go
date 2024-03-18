package crud

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// TODO wrap update() call in "timeout" and "retry" functionality

// Update method to define the logic which updates the resource and sets the updated Terraform state on success.
func Update(ctx context.Context, r helpers.VyosResource, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update Resource")
	tflog.Trace(ctx, "Fetching data model")
	// Current State as resource model
	stateModel := r.GetModel()

	// Plan as resource model
	//  ! since GetModel returns a ptr we must create a new struct with a ptr to not mix up state and plan
	//		This is only a problem here in Update() as other CRUD functions only need to keep track of 1 state/model
	modelType := reflect.TypeOf(stateModel)
	pointerType := modelType.Elem()
	reflectedValue := reflect.New(pointerType)
	planModel := reflectedValue.Interface().(helpers.VyosTopResourceDataModel)

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, stateModel)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, planModel)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := update(ctx, r.GetClient(), stateModel, planModel)
	if err != nil {
		resp.Diagnostics.AddError("API Config error", err.Error())
		return
	}

	// Save data to Terraform state
	tflog.Trace(ctx, "resource updated")
	tflog.Error(ctx, "Setting state", map[string]interface{}{"data": fmt.Sprintf("%#v", planModel)})
	resp.Diagnostics.Append(resp.State.Set(ctx, planModel)...)
}

func update(ctx context.Context, client client.Client, stateModel, planModel helpers.VyosTopResourceDataModel) error {
	// Delete existing config
	stateVyosData, err := helpers.MarshalVyos(ctx, stateModel)
	if err != nil {
		return fmt.Errorf("API marshalling error: %s", err)
	}

	vyosOpsState := helpers.GenerateVyosOps(ctx, stateModel.GetVyosPath(), stateVyosData)
	tflog.Trace(ctx, "Compiling vyos state operations", map[string]interface{}{"vyosOpsState": vyosOpsState})

	client.StageDelete(ctx, vyosOpsState)

	// Create new config
	planVyosData, err := helpers.MarshalVyos(ctx, planModel)
	if err != nil {
		return fmt.Errorf("API marshalling error: %s", err)
	}

	vyosOpsPlan := helpers.GenerateVyosOps(ctx, planModel.GetVyosPath(), planVyosData)
	tflog.Error(ctx, "Compiling vyos plan operations", map[string]interface{}{"vyosOpsPlan": vyosOpsPlan})

	client.StageSet(ctx, vyosOpsPlan)

	// Commit changes to api
	response, err := client.CommitChanges(ctx)
	if err != nil {
		return fmt.Errorf("client error: unable to create '%s', got error: %s", planModel.GetVyosPath(), err)
	}
	if response != nil {
		tflog.Warn(ctx, "Got non-nil response from API", map[string]interface{}{"response": response})
	}

	// Add ID to the resource model as plan fetching do not include it
	planModel.SetID(planModel.GetVyosPath())

	return nil
}
