/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namednatdestinationruleloadbalancebackend code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namednatdestinationruleloadbalancebackend

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (backend) */
// Metadata method to define the resource type name.
func (r natDestinationRuleLoadBalanceBackend) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nat_destination_rule_load_balance_backend"
}
