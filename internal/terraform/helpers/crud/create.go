package crud

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// TODO add "timeout" and "retry" functionality to required parent check
// TODO add "timeout" and "retry" functionality to existing resource check

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func Create(ctx context.Context, r helpers.VyosResource, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "New Resource")
	tflog.Trace(ctx, "Fetching data model")
	planModel := r.GetModel()
	c := r.GetClient()
	providerCfg := r.GetProviderConfig()

	// Read Terraform plan data into the model
	tflog.Trace(ctx, "Fetching plan data")
	resp.Diagnostics.Append(req.Plan.Get(ctx, planModel)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if nearest parent exists
	if !providerCfg.Config.CrudSkipCheckParentBeforeCreate && (len(planModel.GetVyosParentPath()) > 0) {
		parentSemiPath := planModel.GetVyosParentPath()[:len(planModel.GetVyosParentPath())-2]
		parentIDComponenets := planModel.GetVyosParentPath()[len(planModel.GetVyosParentPath())-2:]

		tflog.Debug(ctx, fmt.Sprintf("checking for parent: '%s' under '%s'", parentIDComponenets, strings.Join(parentSemiPath, " ")))
		ret, err := c.Read(ctx, parentSemiPath)
		if err != nil {
			var apiNotFoundError *client.APINotFoundError
			if errors.As(err, &apiNotFoundError) {
				resp.Diagnostics.Append(
					diag.NewErrorDiagnostic(
						fmt.Sprintf("missing parent: '%s'", strings.Join(planModel.GetVyosParentPath(), " ")),
						fmt.Sprintf(`[%s] create the parent resource or configure the provider to ignore missing parents`, err.Error()),
					))
				return
			}

			resp.Diagnostics.Append(
				diag.NewErrorDiagnostic(
					"error reading API endpoint",
					fmt.Sprintf("unable to read '%s', got error: '%s'", planModel.GetVyosParentPath(), err),
				))
			return

		}

		tmpRet := ret
		for _, idComponent := range parentIDComponenets {
			var ok bool
			if tmpRet, ok = tmpRet.(map[string]any)[idComponent]; !ok {
				tflog.Info(ctx, fmt.Sprintf("missing parent: '%s', API returned: '%v'", strings.Join(planModel.GetVyosParentPath(), " "), ret))
				resp.Diagnostics.Append(
					diag.NewErrorDiagnostic(
						fmt.Sprintf("missing parent: '%s'", strings.Join(append(parentSemiPath, parentIDComponenets...), " ")),
						`create the parent  resource or configure the provider to ignore missing parents`,
					))
				return
			}
		}
	}

	// Check if resource already exists
	if !providerCfg.Config.CrudSkipExistingResourceCheck {
		_, err := c.Read(ctx, planModel.GetVyosPath())
		if err == nil {
			resp.Diagnostics.Append(
				diag.NewErrorDiagnostic(
					fmt.Sprintf("resource already exists: '%s'", strings.Join(planModel.GetVyosPath(), " ")),
					"remove the resource in vyos, import the resource, or configure the provider to overwrite existing resources",
				))
			return
		}

		var apiNotFoundError *client.APINotFoundError
		if !errors.As(err, &apiNotFoundError) {
			resp.Diagnostics.Append(
				diag.NewErrorDiagnostic(
					"error reading API endpoint",
					fmt.Sprintf("unable to read %s, got error: %s", planModel.GetVyosPath(), err),
				))
			return
		}
	}

	// Marshal resource model for vyos
	tflog.Trace(ctx, "Marshalling plan for VyOS")
	vyosData, err := helpers.MarshalVyos(ctx, planModel)
	if err != nil {
		resp.Diagnostics.AddError("API Marshalling error", fmt.Sprintf("%s", err))
		return
	}

	// Create vyos api ops
	tflog.Trace(ctx, "Formatting vyos operations")
	vyosOps := helpers.GenerateVyosOps(ctx, planModel.GetVyosPath(), vyosData)
	tflog.Trace(ctx, "Compiled vyos operations", map[string]interface{}{"vyosOps": vyosOps})

	// Stage changes
	tflog.Debug(ctx, "staging vyos changes api calls")
	c.StageSet(ctx, vyosOps)

	// Commit changes
	tflog.Info(ctx, "committing vyos changes api calls")
	response, err := c.CommitChanges(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create %s, got error: %s", planModel.GetVyosPath(), err))
		return
	}
	if response != nil {
		tflog.Warn(ctx, "Got non-nil response from API", map[string]interface{}{"response": response})
	}

	// Add ID to the resource model
	planModel.SetID(planModel.GetVyosPath())
	tflog.Info(ctx, "setting newly created resource id", map[string]interface{}{"vyos-path": planModel.GetVyosPath()})

	// Save data to Terraform state
	tflog.Trace(ctx, "resource created")
	tflog.Error(ctx, "Setting state", map[string]interface{}{"data": fmt.Sprintf("%#v", planModel)})
	resp.Diagnostics.Append(resp.State.Set(ctx, planModel)...)
}
