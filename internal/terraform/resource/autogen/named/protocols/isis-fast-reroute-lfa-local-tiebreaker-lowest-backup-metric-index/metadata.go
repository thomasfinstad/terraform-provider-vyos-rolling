/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package namedprotocolsisisfastreroutelfalocaltiebreakerlowestbackupmetricindex code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedprotocolsisisfastreroutelfalocaltiebreakerlowestbackupmetricindex

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl #metadata (index) */
// Metadata method to define the resource type name.
func (r protocolsIsisFastRerouteLfaLocalTiebreakerLowestBackupMetricIndex) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protocols_isis_fast_reroute_lfa_local_tiebreaker_lowest_backup_metric_index"
}
