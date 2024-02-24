// Package namedvrfnameprotocolsospfvthreeinterface code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedvrfnameprotocolsospfvthreeinterface

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *vrfNameProtocolsOspfvthreeInterface) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "New Resource")
	tflog.Trace(ctx, "Fetching data model")
	planModel := r.model

	// Read Terraform plan data into the model
	tflog.Trace(ctx, "Fetching plan data")
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planModel)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Marshal resource model for vyos
	tflog.Trace(ctx, "Marshalling plan for VyOS")
	vyosData, err := helpers.MarshalVyos(ctx, planModel)
	if err != nil {
		resp.Diagnostics.AddError("API Marshalling error", fmt.Sprintf("%s", err))
		return
	}

	// Create vyos api ops
	tflog.Trace(ctx, "Formatting vyos operations")
	vyosOps := helpers.GenerateVyosOps(ctx, planModel.GetVyosPath(), vyosData)
	tflog.Trace(ctx, "Compiled vyos operations", map[string]interface{}{"vyosOps": vyosOps})

	// Stage changes
	tflog.Debug(ctx, "staging vyos changes api calls")
	r.GetClient().StageSet(ctx, vyosOps)

	// Commit changes
	tflog.Info(ctx, "committing vyos changes api calls")
	response, err := r.GetClient().CommitChanges(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create %s, got error: %s", planModel.GetVyosPath(), err))
		return
	}
	if response != nil {
		tflog.Warn(ctx, "Got non-nil response from API", map[string]interface{}{"response": response})
	}

	// Add ID to the resource model
	planModel.SetID(planModel.GetVyosPath())
	tflog.Info(ctx, "setting newly created resource id", map[string]interface{}{"vyos-path": planModel.GetVyosPath()})

	// Save data to Terraform state
	tflog.Trace(ctx, "resource created")
	tflog.Error(ctx, "Setting state", map[string]interface{}{"data": fmt.Sprintf("%#v", planModel)})
	resp.Diagnostics.Append(resp.State.Set(ctx, &planModel)...)
}

// Read method to define the logic which refreshes the Terraform state for the resource.
func (r *vrfNameProtocolsOspfvthreeInterface) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Read Resource")
	tflog.Trace(ctx, "Fetching data model")
	stateModel := r.model

	// Read Terraform prior state data into the model
	tflog.Trace(ctx, "Fetching state data")
	resp.Diagnostics.Append(req.State.Get(ctx, &stateModel)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Fetch live state from Vyos
	tflog.Debug(ctx, "Fetching API data")
	response, err := r.GetClient().Read(ctx, stateModel.GetVyosPath())
	if err != nil {
		if err.Error() == "API ERROR: Configuration under specified path is empty\n" {
			tflog.Warn(ctx, "API Error: Empty resource", map[string]interface{}{"vyos-path": stateModel.GetVyosPath(), "error": err, "response": response})
		} else {
			tflog.Error(ctx, "API Error", map[string]interface{}{"vyos-path": stateModel.GetVyosPath(), "error": err, "response": response})
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read %s, got error: %s", stateModel.GetVyosPath(), err))
			return
		}
	} else {
		// Populate resource model
		if responseAssrt, ok := response.(map[string]any); ok {
			err := helpers.UnmarshalVyos(ctx, responseAssrt, stateModel)
			if err != nil {
				resp.Diagnostics.AddError("Vyos API response unmarshalling error", fmt.Sprintf("error=%#v", err))
				return
			}
		} else {
			resp.Diagnostics.AddError("Wrong API return type, expected map[string]any.", fmt.Sprintf("response=%#v", response))
			return
		}
	}

	// Save updated data into Terraform state
	tflog.Info(ctx, "Saving state")
	tflog.Trace(ctx, "Setting state to", map[string]interface{}{"data": fmt.Sprintf("%#v", stateModel)})
	resp.Diagnostics.Append(resp.State.Set(ctx, &stateModel)...)
}

// Update method to define the logic which updates the resource and sets the updated Terraform state on success.
func (r *vrfNameProtocolsOspfvthreeInterface) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update Resource")
	tflog.Trace(ctx, "Fetching data model")
	planModel := r.model
	stateModel := r.model

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planModel)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateModel)...)
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
	resp.Diagnostics.Append(resp.State.Set(ctx, &planModel)...)
}

// Delete method to define the logic which deletes the resource and removes the Terraform state on success.
func (r *vrfNameProtocolsOspfvthreeInterface) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete Resource")
	tflog.Trace(ctx, "Fetching data model")
	stateModel := r.model

	// Read Terraform plan data into the model
	// resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateModel)...)
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
