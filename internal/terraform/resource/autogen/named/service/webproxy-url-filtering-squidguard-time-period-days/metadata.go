/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package namedservicewebproxyurlfilteringsquidguardtimeperioddays code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package namedservicewebproxyurlfilteringsquidguardtimeperioddays

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

/* tools/generate-terraform-resource-full/templates/resources/common/metadata.gotmpl */
// Metadata method to define the resource type name.
func (r serviceWebproxyURLFilteringSquIDguardTimePeriodDays) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_webproxy_url_filtering_squidguard_time_period_days"
}
