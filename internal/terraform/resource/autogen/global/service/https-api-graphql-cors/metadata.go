// Package globalservicehttpsapigraphqlcors code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package globalservicehttpsapigraphqlcors

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Metadata method to define the resource type name.
func (r serviceHTTPSAPIGraphqlCors) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_https_api_graphql_cors"
}
