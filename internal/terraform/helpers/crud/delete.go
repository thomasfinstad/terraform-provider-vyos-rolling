package crud

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"
)

// Delete method to define the logic which deletes the resource and removes the Terraform state on success.
func Delete(ctx context.Context, r helpers.VyosResource, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete Resource")
	tflog.Trace(ctx, "Fetching data model")
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
		resp.Diagnostics.AddError(fmt.Sprintf("delete error for: '%s'", stateModel.GetVyosPath()), fmt.Errorf("got non-nil response from API: %w", err).Error())
		return
	}

	// Save data to Terraform state
	resp.State.RemoveResource(ctx)
	tflog.Info(ctx, "resource deleted")
}

// delete removes the resource
// model must be a ptr
// this function is seperated out to keep the terraform provider
// logic and API logic seperate so we can test the API logic easier
func delete(ctx context.Context, providerCfg data.ProviderData, c client.Client, stateModel helpers.VyosTopResourceDataModel) error {
	if providerCfg.Config.CrudSkipCheckChildBeforeDelete {
		// If we do not care about child resources we can nuke it all
		c.StageDelete(ctx, helpers.GenerateVyosOps(ctx, stateModel.GetVyosPath(), nil))
	} else {
		if stateModel.IsGlobalResource() {
			// Handle global resources

			// Check if resource has children
			err := read(ctx, c, stateModel)
			if err != nil {
				return fmt.Errorf("unable to find resource: %w", err)
			}
			_, hasChild := helpers.GetChild(ctx, stateModel)

			// global resources do not handle children the same way as named resources
			if hasChild {
				// global resources with children should remove their attributes only
				vyosData, err := helpers.MarshalVyos(ctx, stateModel)
				if err != nil {
					return fmt.Errorf("format delete operations: %w", err)
				}
				c.StageDelete(ctx, helpers.GenerateVyosOps(ctx, stateModel.GetVyosPath(), vyosData))
			} else {
				// global resources with no children should delete the entire resource
				c.StageDelete(ctx, helpers.GenerateVyosOps(ctx, stateModel.GetVyosPath(), nil))
			}
		} else {
			// Handle named resources

			// Check timeout before retrying
			var lastErr error
			boMs := 500
			boExp := 1.04
			boMaxS := 10.0
			retryTotalDelay := time.Duration(0)
			retryCnt := 0

		L:
			for {
				select {
				case <-ctx.Done():
					log.Println("retry timeout reached")
					tflog.Warn(ctx, "retry timeout reached")
					return lastErr
				default:
					// Check if resource has children
					err := read(ctx, c, stateModel)
					if err != nil {
						return fmt.Errorf("unable to find resource: %w", err)
					}
					child, hasChild := helpers.GetChild(ctx, stateModel)
					if !hasChild {
						// named resources with no children means we can delete the resource
						c.StageDelete(ctx, helpers.GenerateVyosOps(ctx, stateModel.GetVyosPath(), nil))
						break L
					}

					// named resources should not delete if there is a child
					tflog.Warn(ctx, "child resource detected, retrying", map[string]interface{}{"child": child})
					lastErr = fmt.Errorf("child resource detected: %s", child)

					// No Deadline means we do not wish to retry
					if _, ok := ctx.Deadline(); !ok {
						log.Println("no retry deadline configured, disabling retry.", retryTotalDelay)
						tflog.Warn(ctx, "no retry deadline configured, disabling retry.")
						if lastErr != nil {
							return lastErr
						}
						break L
					}

					// Wait a bit before allowing the next attempt
					boMs = int(float64(boMs) * boExp)
					backOff := min(
						time.Duration(boMs)*time.Millisecond,
						time.Duration(boMaxS)*time.Second,
					)
					log.Println("Total retry delay:", retryTotalDelay)
					log.Println("Total retry count:", retryCnt)
					log.Println("Retry delay..:", backOff)
					tflog.Info(ctx, "delaying before next retry", map[string]interface{}{"retryTotalDelay": retryTotalDelay, "retryCnt": retryCnt, "backOff": backOff})
					time.Sleep(backOff)
					log.Println("Retrying...")
					retryTotalDelay += backOff
					retryCnt++
				}
			}
		}
	}

	// Stage and Commit changes to api
	response, err := c.CommitChanges(ctx)
	if err != nil {
		return fmt.Errorf("unable to delete %s: %w", stateModel.GetVyosPath(), err)
	}
	if response != nil {
		tflog.Error(ctx, "got non-nil response from API", map[string]interface{}{"response": response})
		return fmt.Errorf("got non-nil response from API: %s", response)
	}

	return nil
}
