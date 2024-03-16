package crud

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Read method to define the logic which refreshes the Terraform state for the resource.
func Read(ctx context.Context, r helpers.VyosResource, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Read Resource")
	tflog.Trace(ctx, "Fetching data model")
	stateModel := r.GetModel()

	// Read Terraform prior state data into the model
	tflog.Trace(ctx, "Fetching state data")
	diags := req.State.Get(ctx, stateModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Fetch live state from Vyos
	err := read(ctx, r.GetClient(), stateModel)
	if err != nil {
		resp.Diagnostics.AddError("API Read error", err.Error())
	}

	// Save updated data into Terraform state
	tflog.Info(ctx, "Saving state")
	tflog.Trace(ctx, "Setting state to", map[string]interface{}{"data": fmt.Sprintf("%#v", stateModel)})
	resp.Diagnostics.Append(resp.State.Set(ctx, stateModel)...)
}

// read populates resource model
// model must be a ptr
func read(ctx context.Context, client client.Client, model helpers.VyosTopResourceDataModel) error {
	tflog.Debug(ctx, "Fetching API data")
	response, err := client.Read(ctx, model.GetVyosPath())
	if err != nil {
		if err.Error() == "API ERROR: Configuration under specified path is empty\n" {
			tflog.Warn(ctx, "API Error: Empty resource", map[string]interface{}{"vyos-path": model.GetVyosPath(), "error": err, "response": response})
		} else {
			tflog.Error(ctx, "API Error", map[string]interface{}{"vyos-path": model.GetVyosPath(), "error": err, "response": response})
			return fmt.Errorf("unable to read %s, got error: %s", model.GetVyosPath(), err)
		}
	} else {
		if responseAssrt, ok := response.(map[string]any); ok {
			err := helpers.UnmarshalVyos(ctx, responseAssrt, model)
			if err != nil {
				return fmt.Errorf("vyos API response unmarshalling error: %#v", err)
			}
		} else {
			return fmt.Errorf("wrong API return type, expected map[string]any, got: %#v", response)
		}
	}

	return nil
}
