package crud

import (
	"context"
	"errors"
	"fmt"
	"strings"

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
		var apiNotFoundError *client.APINotFoundError
		if errors.As(err, &apiNotFoundError) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("API Read error", err.Error())
	}

	// Save updated data into Terraform state
	tflog.Info(ctx, "Saving state")
	tflog.Trace(ctx, "Setting state to", map[string]interface{}{"data": fmt.Sprintf("%#v", stateModel)})
	resp.Diagnostics.Append(resp.State.Set(ctx, stateModel)...)
}

// read populates resource model
// model must be a ptr
func read(ctx context.Context, c client.Client, model helpers.VyosTopResourceDataModel) error {
	tflog.Debug(ctx, "Fetching API data")
	response, err := c.Read(ctx, model.GetVyosPath())
	if err != nil {
		tflog.Warn(ctx, "API Error:", map[string]interface{}{"vyos-path": model.GetVyosPath(), "error": err, "response": response})

		// If we could not find ourselfs, check the parent for an empty resource config
		var apiNotFoundError *client.APINotFoundError
		if errors.As(err, &apiNotFoundError) && len(model.GetVyosParentPath()) > 0 {
			parentPath := model.GetVyosParentPath()
			selfIDComponenets := model.GetVyosPath()[len(model.GetVyosParentPath()):]

			tflog.Debug(ctx, fmt.Sprintf("checking for parent: '%s' under '%s'", selfIDComponenets, strings.Join(parentPath, " ")))
			ret, perr := c.Read(ctx, parentPath)
			if perr != nil {
				return err
			}

			// See if we can walk into parent and find our selvs step by step, owtherwise return original error
			tmpRet := ret
			for _, idComponent := range selfIDComponenets {
				var ok bool
				if tmpRet, ok = tmpRet.(map[string]any)[idComponent]; !ok {
					tflog.Debug(ctx, fmt.Sprintf("missing parent: '%s', API returned: '%v'", strings.Join(model.GetVyosNamedParentPath(), " "), ret))
					return err
				}
			}

			// Reaching this point means we found ourselvs as an empty resource
			err := helpers.UnmarshalVyos(ctx, make(map[string]any), model)
			if err != nil {
				return fmt.Errorf("vyos API response unmarshalling error: %w", err)
			}
			return nil
		}

		return err

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
