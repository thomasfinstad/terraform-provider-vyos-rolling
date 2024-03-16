package crud

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// TODO add retry and timeout
// TODO check provider config if we must check for parent before creating
// TODO check if required parent exists and fail after timeout if not
// TODO add timeout and retry to existing resource check

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func Create(ctx context.Context, r helpers.VyosResource, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "New Resource")
	tflog.Trace(ctx, "Fetching data model")
	planModel := r.GetModel()
	client := r.GetClient()
	providerCfg := r.GetProviderConfig()

	// Read Terraform plan data into the model
	tflog.Trace(ctx, "Fetching plan data")
	resp.Diagnostics.Append(req.Plan.Get(ctx, planModel)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if nearest parent exists

	// Check if resource already exists
	if !providerCfg.Config.CrudSkipExistingResourceCheck {
		_, err := client.Read(ctx, planModel.GetVyosPath())
		if err == nil {
			resp.Diagnostics.Append(
				diag.NewErrorDiagnostic(
					"resource already exists: '"+strings.Join(planModel.GetVyosPath(), " ")+"'",
					"remove the resource in vyos, import the resource, or configure the provider to overwrite existing resources",
				))
			return
		} else if err.Error() != "API ERROR: Configuration under specified path is empty\n" {
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
	client.StageSet(ctx, vyosOps)

	// Commit changes
	tflog.Info(ctx, "committing vyos changes api calls")
	response, err := client.CommitChanges(ctx)
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
