// Package namedqospolicypriorityqueueclass code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package namedqospolicypriorityqueueclass

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r qosPolicyPriorityQueueClass) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	data := r.model

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		resp.Diagnostics.AddError("Json Marshalling Error", fmt.Sprintf("%s", err))
		return
	}

	var vyosData map[string]interface{}
	err = json.Unmarshal(jsonData, &vyosData)
	if err != nil {
		resp.Diagnostics.AddError("Json Unmarshalling Error", fmt.Sprintf("%s", err))
		return
	}

	// Create vyos api ops
	vyosOps := helpers.GenerateVyosOps(ctx, data.GetVyosPath(), vyosData)
	tflog.Error(ctx, "Compiled vyos operations", map[string]interface{}{"vyosOps": vyosOps})

	// Stage changes
	r.client.StageSet(ctx, vyosOps)

	// Commit changes
	response, err := r.client.CommitChanges(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create %s, got error: %s", r.ResourceName, err))
		return
	}
	if response != nil {
		tflog.Warn(ctx, "Got non-nil response from API", map[string]interface{}{"response": response})
	}

	// Save ID into the Terraform state.data.SelfIdentifier = types.NumberValue(data.SelfIdentifier.ValueBigFloat().String())

	// Save data to Terraform state
	tflog.Trace(ctx, "resource created")
	tflog.Error(ctx, "Setting state", map[string]interface{}{"data": fmt.Sprintf("%#v", data)})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read method to define the logic which refreshes the Terraform state for the resource.
func (r qosPolicyPriorityQueueClass) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	data := r.model

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	vyosPath := data.GetVyosPath()

	// Fetch live state from Vyos
	response, err := r.client.Read(ctx, vyosPath)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create %s, got error: %s", r.ResourceName, err))
		return
	}
	if response != nil {
		tflog.Warn(ctx, "Got non-nil response from API", map[string]interface{}{"response": response})
	}

	jsonStr, err := json.Marshal(response.(map[string]interface{}))
	err = json.Unmarshal(jsonStr, &r.model)
	if err != nil {
		resp.Diagnostics.AddError("Json Unmarshal Error", fmt.Sprintf("%s", err))
		return
	}

	// Save updated data into Terraform state
	tflog.Error(ctx, "Setting state", map[string]interface{}{"data": fmt.Sprintf("%#v", data)})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update method to define the logic which updates the resource and sets the updated Terraform state on success.
func (r qosPolicyPriorityQueueClass) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	data := r.model

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		resp.Diagnostics.AddError("Json Marshalling Error", fmt.Sprintf("%s", err))
		return
	}

	var vyosData map[string]interface{}
	err = json.Unmarshal(jsonData, &vyosData)
	if err != nil {
		resp.Diagnostics.AddError("Json Unmarshalling Error", fmt.Sprintf("%s", err))
		return
	}

	// Create vyos api ops
	vyosOps := helpers.GenerateVyosOps(ctx, data.GetVyosPath(), vyosData)
	tflog.Error(ctx, "Compiled vyos operations", map[string]interface{}{"vyosOps": vyosOps})

	// Stage changes
	r.client.StageSet(ctx, vyosOps)

	// Commit changes
	response, err := r.client.CommitChanges(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create %s, got error: %s", r.ResourceName, err))
		return
	}
	if response != nil {
		tflog.Warn(ctx, "Got non-nil response from API", map[string]interface{}{"response": response})
	}

	// Save ID into the Terraform state.data.SelfIdentifier = types.NumberValue(data.SelfIdentifier.ValueBigFloat().String())

	// Save data to Terraform state
	tflog.Trace(ctx, "resource created")
	tflog.Error(ctx, "Setting state", map[string]interface{}{"data": fmt.Sprintf("%#v", data)})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete method to define the logic which deletes the resource and removes the Terraform state on success.
func (r qosPolicyPriorityQueueClass) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	data := r.model

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete example, got error: %s", err))
	//     return
	// }
}
