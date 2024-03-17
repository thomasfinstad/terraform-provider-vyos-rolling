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
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"
)

// TODO add "timeout" and "retry" functionality to required parent check
// TODO add "timeout" and "retry" functionality to existing resource check

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
func create(ctx context.Context, providerCfg data.ProviderData, c client.Client, planModel helpers.VyosTopResourceDataModel) error {
	// TODO consolidate parent check and self check API calls
	//  When running the API call for the parent check we should have enough data
	//  to also check if the resource already exists and can skip the extra
	//  API call

	// Check if nearest parent exists
	if !providerCfg.Config.CrudSkipCheckParentBeforeCreate && (len(planModel.GetVyosNamedParentPath()) > 0) {
		parentSemiPath := planModel.GetVyosNamedParentPath()[:len(planModel.GetVyosNamedParentPath())-2]
		parentIDComponenets := planModel.GetVyosNamedParentPath()[len(planModel.GetVyosNamedParentPath())-2:]

		tflog.Debug(ctx, fmt.Sprintf("checking for parent: '%s' under '%s'", parentIDComponenets, strings.Join(parentSemiPath, " ")))

		// TODO Look into get values API call to reduce overhead of multiple lookups
		//  Ref: https://vyos.dev/T6135
		//  This method returns a lot more than is strictly needed, but it saves us from checking if the
		//  parent actually exists but is empty after a perfectly targeted search and APINotFoundError.
		//  The response for creating a "firewall ipv4 name xYz rule 123" resource will look up
		//  the entirety of all firewall rules "firewall ipv4", this can easily return a disgustingly large
		//  blob, and is very wasteful.
		//  There is a chance that it would be better to not support empty resources at all.
		ret, err := c.Read(ctx, parentSemiPath)
		if err != nil {
			var apiNotFoundError *client.APINotFoundError
			if errors.As(err, &apiNotFoundError) {
				return fmt.Errorf(
					"missing parent config: '%s': %w",
					strings.Join(planModel.GetVyosNamedParentPath(), " "),
					err,
				)
			}
			return fmt.Errorf("API error: '%s': %w", parentSemiPath, err)
		}

		// Check for empty but existing parent resource
		tmpRet := ret
		for _, idComponent := range parentIDComponenets {
			var ok bool
			if tmpRet, ok = tmpRet.(map[string]any)[idComponent]; !ok {
				tflog.Debug(ctx, fmt.Sprintf("missing parent: '%s', API returned: '%v'", strings.Join(planModel.GetVyosNamedParentPath(), " "), ret))
				return fmt.Errorf("missing parent: '%s': ", strings.Join(append(parentSemiPath, parentIDComponenets...), " "))
			}
		}
	}

	// Check if resource already exists
	if !providerCfg.Config.CrudSkipExistingResourceCheck {
		tflog.Debug(ctx, fmt.Sprintf("checking for existing resource: '%s'", planModel.GetVyosPath()))

		_, err := c.Read(ctx, planModel.GetVyosPath())
		if err == nil {
			return fmt.Errorf("resource already exists: '%s'", strings.Join(planModel.GetVyosPath(), " "))
		}

		var apiNotFoundError *client.APINotFoundError
		if !errors.As(err, &apiNotFoundError) {
			return fmt.Errorf("API error: %w", err)
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
