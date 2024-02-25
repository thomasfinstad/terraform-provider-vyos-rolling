package crud

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Read method to define the logic which refreshes the Terraform state for the resource.
func Read(ctx context.Context, r helpers.VyosResource, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Read Resource")
	tflog.Trace(ctx, "Fetching data model")
	stateModel := r.GetModel()

	// Read Terraform prior state data into the model
	tflog.Trace(ctx, "Fetching state data")
	resp.Diagnostics.Append(req.State.Get(ctx, stateModel)...)
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
	resp.Diagnostics.Append(resp.State.Set(ctx, stateModel)...)
}
