// Package globalservicewebproxyurlfiltering code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicewebproxyurlfiltering

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r serviceWebproxyURLFiltering) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_webproxy_url_filtering"
}
