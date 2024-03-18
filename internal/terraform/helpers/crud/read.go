package crud

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// TODO wrap read() call in "timeout" and "retry" functionality

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
		var apiNotFoundError *client.APINotFoundError
		if errors.As(err, &apiNotFoundError) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("read error", err.Error())
	}

	// Save updated data into Terraform state
	tflog.Info(ctx, "Saving state")
	tflog.Trace(ctx, "Setting state to", map[string]interface{}{"data": fmt.Sprintf("%#v", stateModel)})
	resp.Diagnostics.Append(resp.State.Set(ctx, stateModel)...)
}

// read populates resource model
// model must be a ptr
// this function is seperated out to keep the terraform provider
// logic and API logic seperate so we can test the API logic easier
func read(ctx context.Context, c client.Client, model helpers.VyosTopResourceDataModel) error {
	tflog.Debug(ctx, "Fetching API data")
	response, err := c.Read(ctx, model.GetVyosPath())
	if err != nil {
		tflog.Warn(ctx, "API Error:", map[string]interface{}{"vyos-path": model.GetVyosPath(), "error": err, "response": response})

		// Check if we exists, if so this means we are an empty resource
		ret, hasErr := c.Has(ctx, model.GetVyosPath())
		if hasErr != nil {
			return fmt.Errorf("resource check error for: '%s': %w", model.GetVyosNamedParentPath(), hasErr)
		}

		if !ret {
			return hasErr
		}

		// Marshal up as empty
		err := helpers.UnmarshalVyos(ctx, make(map[string]any), model)
		if err != nil {
			return fmt.Errorf("vyos API response unmarshalling error: %w", err)
		}
		return nil

	}

	if responseAssrt, ok := response.(map[string]any); ok {
		err := helpers.UnmarshalVyos(ctx, responseAssrt, model)
		if err != nil {
			return fmt.Errorf("vyos API response unmarshalling error: %w", err)
		}
	} else {
		return fmt.Errorf("wrong API return type, expected map[string]any, got: %#v", response)
	}

	return nil
}
