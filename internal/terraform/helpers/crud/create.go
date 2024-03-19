package crud

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"
)

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func Create(ctx context.Context, r helpers.VyosResource, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "New Resource")
	tflog.Trace(ctx, "Fetching data model")
	planModel := r.GetModel()

	// Read Terraform plan data into the model
	tflog.Trace(ctx, "Fetching plan data")
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
	tflog.Info(ctx, "setting newly created resource id", map[string]interface{}{"vyos-path": planModel.GetVyosPath()})

	// Save data to Terraform state
	tflog.Trace(ctx, "resource created")
	tflog.Error(ctx, "Setting state", map[string]interface{}{"data": fmt.Sprintf("%#v", planModel)})
	resp.Diagnostics.Append(resp.State.Set(ctx, planModel)...)
}

// create populates resource model
// model must be a ptr
// this function is seperated out to keep the terraform provider
// logic and API logic seperate so we can test the API logic easier
// TODO add retry support to create()
func create(ctx context.Context, providerCfg data.ProviderData, c client.Client, planModel helpers.VyosTopResourceDataModel) error {
	// Check if nearest parent exists
	if !providerCfg.Config.CrudSkipCheckParentBeforeCreate && (len(planModel.GetVyosNamedParentPath()) > 0) {
		tflog.Debug(ctx, fmt.Sprintf("checking for parent: '%s'", planModel.GetVyosNamedParentPath()))

		ret, err := c.Has(ctx, planModel.GetVyosNamedParentPath())
		if err != nil {
			return fmt.Errorf("parent check error for: '%s': %w", planModel.GetVyosNamedParentPath(), err)
		}

		if !ret {
			return fmt.Errorf("missing parent: '%s': ", planModel.GetVyosNamedParentPath())
		}
	}

	// Check if resource already exists
	if !providerCfg.Config.CrudSkipExistingResourceCheck {
		tflog.Debug(ctx, fmt.Sprintf("checking for existing resource: '%s'", planModel.GetVyosPath()))

		ret, err := c.Has(ctx, planModel.GetVyosPath())
		if err != nil {
			return fmt.Errorf("resource check error for: '%s': %w", planModel.GetVyosNamedParentPath(), err)
		}

		if ret {
			return fmt.Errorf("resource already exists: '%s'", strings.Join(planModel.GetVyosPath(), " "))
		}
	}

	// Marshal resource model for vyos
	tflog.Trace(ctx, "Marshalling plan for VyOS")
	vyosData, err := helpers.MarshalVyos(ctx, planModel)
	if err != nil {
		return fmt.Errorf("marshalling error: %w", err)
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
		return fmt.Errorf("unable to create resource '%s' due to client error: %w", planModel.GetVyosPath(), err)
	}
	if response != nil {
		tflog.Warn(ctx, "Got non-nil response from API", map[string]interface{}{"response": response})
	}

	return nil
}
