package crud

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers/tools"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"
)

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func Create(ctx context.Context, r helpers.VyosResource, req resource.CreateRequest, resp *resource.CreateResponse) {
	tools.Debug(ctx, "New Resource")
	ctx = r.GetProviderConfig().CtxMutilatorRun(ctx)

	tools.Trace(ctx, "Fetching data model")
	planModel := r.GetModel()

	// Read Terraform plan data into the model
	tools.Trace(ctx, "Fetching plan data")
	resp.Diagnostics.Append(req.Plan.Get(ctx, planModel)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Setup timeout
	createTimeout, diags := planModel.GetTimeouts().Create(ctx, time.Duration(r.GetProviderConfig().Config.CrudDefaultTimeouts)*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, createTimeout)
	defer cancel()

	// Fetch resource from API
	err := create(ctx, r.GetProviderConfig(), r.GetClient(), planModel)
	if err != nil {
		resp.Diagnostics.Append(
			diag.NewErrorDiagnostic(
				"unable to create resource",
				err.Error(),
			))
		return
	}

	// Add ID to the resource model
	planModel.SetID(planModel.GetVyosPath())
	tools.Info(ctx, "setting newly created resource id", map[string]interface{}{"vyos-path": planModel.GetVyosPath()})

	// Save data to Terraform state
	tools.Debug(ctx, "resource created, saving state")
	helpers.UnknownToNull(ctx, planModel)
	resp.Diagnostics.Append(resp.State.Set(ctx, planModel)...)
}

// create populates resource model
// model must be a ptr
// this function is separated out to keep the terraform provider
// logic and API logic separate so we can test the API logic easier
func create(ctx context.Context, providerCfg data.ProviderData, c client.Client, planModel helpers.VyosTopResourceDataModel) error {
	// Check if nearest parent exists
	if !providerCfg.Config.CrudSkipCheckParentBeforeCreate && (len(planModel.GetVyosNamedParentPath()) > 0) {
		tools.Debug(ctx, fmt.Sprintf("checking for parent: '%s'", planModel.GetVyosNamedParentPath()))

		// Check timeout before retrying
		var lastErr error
		boMs := 500
		boExp := 1.04
		boMaxS := 10.0
		retryTotalDelay := time.Duration(0)
		retryCnt := 0
	OuterLoop:
		for {
			select {
			case <-ctx.Done():

				tools.Warn(ctx, "retry timeout reached")
				return lastErr
			default:
				parentExists, err := c.Has(ctx, planModel.GetVyosNamedParentPath())
				if err != nil {
					return fmt.Errorf("parent check error for: '%s': %w", planModel.GetVyosNamedParentPath(), err)
				}

				if parentExists {
					break OuterLoop
				} else {
					tools.Warn(ctx, "parent resource missing, retrying", map[string]interface{}{"resource": planModel.GetVyosPath()})
					lastErr = fmt.Errorf("parent resource missing: '%s': ", planModel.GetVyosNamedParentPath())
				}

				// No Deadline means we do not wish to retry
				if _, ok := ctx.Deadline(); !ok {

					tools.Warn(ctx, "no retry deadline configured, disabling retry.")
					if lastErr != nil {
						return lastErr
					}
					break OuterLoop
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

	// Check if resource already exists
	if !providerCfg.Config.CrudSkipExistingResourceCheck {
		tools.Debug(ctx, fmt.Sprintf("checking for existing resource: '%s'", planModel.GetVyosPath()))

		// Check timeout before retrying
		var lastErr error
		boMs := 500
		boExp := 1.04
		boMaxS := 10.0
		retryTotalDelay := time.Duration(0)
		retryCnt := 0
	EL:
		for {
			select {
			case <-ctx.Done():

				tools.Warn(ctx, "retry timeout reached")
				return lastErr
			default:
				resourceExists, err := c.Has(ctx, planModel.GetVyosPath())
				if err != nil {
					return fmt.Errorf("resource check error for: [%s]: %w", strings.Join(planModel.GetVyosPath(), " "), err)
				} else if !resourceExists {
					break EL
				} else {
					tools.Info(ctx, "resource already exists, retrying", map[string]interface{}{"resource": planModel.GetVyosPath()})
					lastErr = fmt.Errorf("resource already exists: [%s]", strings.Join(planModel.GetVyosPath(), " "))
				}

				// No Deadline means we do not wish to retry
				if _, ok := ctx.Deadline(); !ok {

					tools.Warn(ctx, "no retry deadline configured, disabling retry.")
					if lastErr != nil {
						return lastErr
					}
					break EL
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

	// Marshal resource model for vyos
	tools.Trace(ctx, "Marshalling plan for VyOS")
	vyosData, err := helpers.MarshalVyos(ctx, planModel)
	if err != nil {
		return fmt.Errorf("marshalling error: %w", err)
	}

	// Create vyos api ops
	tools.Trace(ctx, "Formatting vyos operations")
	vyosOps := helpers.GenerateVyosOps(ctx, planModel.GetVyosPath(), vyosData)
	tools.Trace(ctx, "Compiled vyos operations", map[string]interface{}{"vyosOps": vyosOps})

	// Stage changes
	tools.Debug(ctx, "staging vyos changes api calls")
	c.StageSet(ctx, vyosOps)

	// Commit changes
	tools.Info(ctx, "committing vyos changes api calls")
	response, err := c.CommitChanges(ctx)
	if err != nil {
		return fmt.Errorf("unable to create resource '%s' due to client error: %w", planModel.GetVyosPath(), err)
	}
	if response != nil {
		tools.Warn(ctx, "Got non-nil response from API", map[string]interface{}{"response": response})
	}

	return nil
}
