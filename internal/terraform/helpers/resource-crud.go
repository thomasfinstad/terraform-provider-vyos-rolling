package helpers

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func Create(ctx context.Context, r VyosResource, req resource.CreateRequest, resp *resource.CreateResponse) {
	data := r.GetModel()

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Marshal resource model for vyos
	vyosData, err := MarshalVyos(ctx, data)
	if err != nil {
		resp.Diagnostics.AddError("API Marshalling error", fmt.Sprintf("%s", err))
		return
	}

	// Create vyos api ops
	vyosOps := GenerateVyosOps(ctx, data.GetVyosPath(), vyosData)
	tflog.Error(ctx, "Compiled vyos operations", map[string]interface{}{"vyosOps": vyosOps})

	// Stage changes
	r.GetClient().StageSet(ctx, vyosOps)

	// Commit changes
	response, err := r.GetClient().CommitChanges(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create %s, got error: %s", r.GetName(), err))
		return
	}
	if response != nil {
		tflog.Warn(ctx, "Got non-nil response from API", map[string]interface{}{"response": response})
	}

	// Add ID to the resource model
	data.SetID(basetypes.NewStringValue(strings.Join(data.GetVyosPath(), "__")))

	// Save data to Terraform state
	tflog.Trace(ctx, "resource created")
	tflog.Error(ctx, "Setting state", map[string]interface{}{"data": fmt.Sprintf("%#v", data)})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read method to define the logic which refreshes the Terraform state for the resource.
func Read(ctx context.Context, r VyosResource, req resource.ReadRequest, resp *resource.ReadResponse) {
	data := r.GetModel()

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Use ID to generate vyospath (makes importing resources easier than using the function)
	if data.GetID().ValueString() == "" {
		resp.Diagnostics.AddError("Resource ID empty, can not read resource", fmt.Sprintf("%#v", data))
	}

	// Fetch live state from Vyos
	response, err := r.GetClient().Read(ctx, data.GetVyosPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read %s, got error: %s", r.GetName(), err))
		return
	}
	if response != nil {
		tflog.Warn(ctx, "Got non-nil response from API", map[string]interface{}{"response": response})
	}

	// Populate resource model
	if responseAssrt, ok := response.(map[string]any); ok {
		UnmarshalVyos(ctx, responseAssrt, data)
	} else {
		resp.Diagnostics.AddError("Wrong API return type, expected map[string]any.", fmt.Sprintf("response=%#v", response))
		return
	}

	// Save updated data into Terraform state
	tflog.Error(ctx, "Setting state", map[string]interface{}{"data": fmt.Sprintf("%#v", data)})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update method to define the logic which updates the resource and sets the updated Terraform state on success.
func Update(ctx context.Context, r VyosResource, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	data := r.GetModel()
	state := r.GetModel()

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing config
	vyosData, err := MarshalVyos(ctx, state)
	if err != nil {
		resp.Diagnostics.AddError("API Marshalling error", fmt.Sprintf("%s", err))
		return
	}

	vyosOps := GenerateVyosOps(ctx, state.GetVyosPath(), vyosData)
	tflog.Error(ctx, "Compiling vyos operations", map[string]interface{}{"vyosOps": vyosOps})

	r.GetClient().StageDelete(ctx, vyosOps)

	// Create new config
	vyosData, err = MarshalVyos(ctx, data)
	if err != nil {
		resp.Diagnostics.AddError("API Marshalling error", fmt.Sprintf("%s", err))
		return
	}

	vyosOps = GenerateVyosOps(ctx, data.GetVyosPath(), vyosData)
	tflog.Error(ctx, "Compiling vyos operations", map[string]interface{}{"vyosOps": vyosOps})

	r.GetClient().StageSet(ctx, vyosOps)

	// Commit changes to api
	response, err := r.GetClient().CommitChanges(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create %s, got error: %s", r.GetName(), err))
		return
	}
	if response != nil {
		tflog.Warn(ctx, "Got non-nil response from API", map[string]interface{}{"response": response})
	}

	// Save data to Terraform state
	tflog.Trace(ctx, "resource updated")
	tflog.Error(ctx, "Setting state", map[string]interface{}{"data": fmt.Sprintf("%#v", data)})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete method to define the logic which deletes the resource and removes the Terraform state on success.
func Delete(ctx context.Context, r VyosResource, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	state := r.GetModel()

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing config
	vyosData, err := MarshalVyos(ctx, state)
	if err != nil {
		resp.Diagnostics.AddError("API Marshalling error", fmt.Sprintf("%s", err))
		return
	}

	vyosOps := GenerateVyosOps(ctx, state.GetVyosPath(), vyosData)
	tflog.Error(ctx, "Compiling vyos operations", map[string]interface{}{"vyosOps": vyosOps})

	r.GetClient().StageDelete(ctx, vyosOps)

	// Commit changes to api
	response, err := r.GetClient().CommitChanges(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create %s, got error: %s", r.GetName(), err))
		return
	}
	if response != nil {
		tflog.Warn(ctx, "Got non-nil response from API", map[string]interface{}{"response": response})
	}

	// Save data to Terraform state
	blankState := r.GetModel()
	tflog.Trace(ctx, "resource updated")
	tflog.Error(ctx, "Setting state", map[string]interface{}{"data": fmt.Sprintf("%#v", blankState)})
	resp.Diagnostics.Append(resp.State.Set(ctx, &blankState)...)
}
