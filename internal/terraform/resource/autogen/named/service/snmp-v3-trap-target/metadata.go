// Package namedservicesnmpvthreetraptarget code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicesnmpvthreetraptarget

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r serviceSnmpVthreeTrapTarget) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_snmp_v3_trap_target"
}
