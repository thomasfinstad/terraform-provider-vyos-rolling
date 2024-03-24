package crud

import (
	"context"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/client/clienterrors"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	cruderrors "github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers/crud/cruderror"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers/tools"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"
)

// Delete method to define the logic which deletes the resource and removes the Terraform state on success.
func Delete(ctx context.Context, r helpers.VyosResource, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tools.Debug(ctx, "Delete Resource")
	ctx = r.GetProviderConfig().CtxMutilatorRun(ctx)

	tools.Trace(ctx, "Fetching data model")
	stateModel := r.GetModel()
	c := r.GetClient()

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, stateModel)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Setup timeout
	createTimeout, diags := stateModel.GetTimeouts().Delete(ctx, time.Duration(r.GetProviderConfig().Config.CrudDefaultTimeouts)*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, createTimeout)
	defer cancel()

	// Delete resource
	err := delete(ctx, r.GetProviderConfig(), c, stateModel)
	if err != nil {
		var rnfErr cruderrors.ResourceNotFoundError
		if !errors.As(err, &rnfErr) {
			resp.Diagnostics.Append(err)
			return
		}
	}

	// Save data to Terraform state
	resp.State.RemoveResource(ctx)
	tools.Info(ctx, "resource deleted")
}

// delete removes the resource
// model must be a ptr
// this function is separated out to keep the terraform provider
// logic and API logic separate so we can test the API logic easier
func delete(ctx context.Context, providerCfg data.ProviderData, c client.Client, stateModel helpers.VyosTopResourceDataModel) cruderrors.CrudError {

	if stateModel.IsGlobalResource() {
		// Handle global resources

		// Check if resource has children
		err := read(ctx, c, stateModel)
		if err != nil {
			var nfErr clienterrors.NotFoundError
			if !errors.As(err, &nfErr) {
				return cruderrors.WrapIntoResourceError(stateModel, err)
			}
		}
		_, hasChild := helpers.GetChild(ctx, stateModel)

		// global resources do not handle children the same way as named resources
		if hasChild {
			// global resources with children should remove their attributes only
			vyosData, err := helpers.MarshalVyos(ctx, stateModel)
			if err != nil {
				return cruderrors.WrapIntoResourceError(stateModel, err)
			}
			c.StageDelete(ctx, helpers.GenerateVyosOps(ctx, stateModel.GetVyosPath(), vyosData))
		} else {
			// global resources with no children should delete the entire resource
			c.StageDelete(ctx, helpers.GenerateVyosOps(ctx, stateModel.GetVyosPath(), nil))
		}
	} else {
		// Handle named resources

		if providerCfg.Config.CrudSkipCheckChildBeforeDelete {
			// If we do not care about child resources we can nuke the entire resource
			c.StageDelete(ctx, helpers.GenerateVyosOps(ctx, stateModel.GetVyosPath(), nil))
		} else {

			// Check timeout before retrying
			var lastErr cruderrors.CrudError
			boMs := 500
			boExp := 1.04
			boMaxS := 10.0
			retryTotalDelay := time.Duration(0)
			retryCnt := 0

		L:
			for {
				select {
				case <-ctx.Done():

					tools.Warn(ctx, "retry timeout reached")
					return cruderrors.WrapIntoResourceError(stateModel, lastErr)
				default:
					// Check if resource has children
					readErr := read(ctx, c, stateModel)
					if readErr != nil {
						var nfErr clienterrors.NotFoundError
						if !errors.As(readErr, &nfErr) {
							return cruderrors.WrapIntoResourceError(stateModel, readErr)
						}
					}
					child, hasChild := helpers.GetChild(ctx, stateModel)
					if !hasChild {
						// named resources with no children means we can delete the entire resource
						c.StageDelete(ctx, helpers.GenerateVyosOps(ctx, stateModel.GetVyosPath(), nil))
						break L
					}

					// named resources should not delete if there is a child
					tools.Warn(ctx, "child resource detected, retrying", map[string]interface{}{"child": child})
					// lastErr = fmt.Errorf("child resource detected: %s", child)
					lastErr = cruderrors.NewResourceError(stateModel, "child resource detected: %s", child)

					// No Deadline means we do not wish to retry
					if _, ok := ctx.Deadline(); !ok {

						tools.Warn(ctx, "no retry deadline configured, disabling retry.")
						return lastErr
					}

					// Wait a bit before allowing the next attempt
					boMs = int(float64(boMs) * boExp)
					backOff := min(
						time.Duration(boMs)*time.Millisecond,
						time.Duration(boMaxS)*time.Second,
					)

					tools.Info(ctx, "delaying before next retry", map[string]interface{}{"retryTotalDelay": retryTotalDelay, "retryCnt": retryCnt, "backOff": backOff})
					time.Sleep(backOff)

					retryTotalDelay += backOff
					retryCnt++
				}
			}
		}
	}

	// Stage and Commit changes to api
	response, err := c.CommitChanges(ctx)
	if err != nil {
		return cruderrors.WrapIntoResourceError(stateModel, err)
	}
	if response != nil {
		tools.Error(ctx, "got non-nil response from API", map[string]interface{}{"response": response})
		return cruderrors.NewResourceError(stateModel, "got non-nil response from API: %s", response)
	}

	return nil
}
