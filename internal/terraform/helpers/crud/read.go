package crud

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client/clienterrors"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	cruderrors "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers/crud/cruderror"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers/tools"
)

// Read method to define the logic which refreshes the Terraform state for the resource.
func Read(ctx context.Context, r helpers.VyosResource, req resource.ReadRequest, resp *resource.ReadResponse) {
	tools.Debug(ctx, "Read Resource")
	ctx = r.GetProviderConfig().CtxMutilatorRun(ctx)

	tools.Trace(ctx, "Fetching data model")
	stateModel := r.GetModel()

	// Read Terraform prior state data into the model
	tools.Debug(ctx, "Fetching state data")
	diags := req.State.Get(ctx, stateModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tools.Trace(ctx, "Fetched state data", map[string]interface{}{"state-model": stateModel})

	// Setup timeout
	createTimeout, diags := stateModel.GetTimeouts().Read(ctx, time.Duration(r.GetProviderConfig().Config.CrudDefaultTimeouts)*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, createTimeout)
	defer cancel()

	// Fetch live state from Vyos
	err := read(ctx, r.GetClient(), stateModel)

	var rnfErr cruderrors.ResourceNotFoundError
	if errors.As(err, &rnfErr) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("unable to retrieve config", err.Error())
		return
	}

	// Save data to Terraform state
	tools.Debug(ctx, "resource created, saving state")
	helpers.UnknownToNull(ctx, stateModel)
	resp.Diagnostics.Append(resp.State.Set(ctx, stateModel)...)
}

// read populates resource model
// model must be a ptr
// this function is separated out to keep the terraform provider
// logic and API logic separate so we can test the API logic easier
func read(ctx context.Context, c client.Client, model helpers.VyosTopResourceDataModel) error {
	tools.Debug(ctx, "Fetching API data", map[string]interface{}{"vyos-path": model.GetVyosPath()})

	// Check if we exists, if so this means we are an empty resource
	ret, hasErr := c.Has(ctx, model.GetVyosPath())
	if hasErr != nil {
		return fmt.Errorf("[%s] resource lookup: %w", model.GetVyosNamedParentPath(), hasErr)
	}

	if !ret {
		return cruderrors.WrapIntoResourceNotFoundError(model, hasErr)
	}

	response, getErr := c.Get(ctx, model.GetVyosPath())
	// Error after successful client.Has call should mean empty resource
	var nfErr clienterrors.NotFoundError
	if errors.As(getErr, &nfErr) {
		err := helpers.UnmarshalVyos(ctx, make(map[string]any), model)
		if err != nil {
			return fmt.Errorf("empty unmarshal response: %w", err)
		}
		return nil
	}
	if getErr != nil {
		return getErr
	}

	// Populate full model config
	if responseAssrt, ok := response.(map[string]any); ok {
		err := helpers.UnmarshalVyos(ctx, responseAssrt, model)
		if err != nil {
			return fmt.Errorf("unmarshal response: %w", err)
		}
	} else {
		return fmt.Errorf("unknown api response: %#v", response)
	}

	return nil
}
