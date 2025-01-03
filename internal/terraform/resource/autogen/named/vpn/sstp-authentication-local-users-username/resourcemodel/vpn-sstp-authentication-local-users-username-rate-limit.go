/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VpnSstpAuthenticationLocalUsersUsernameRateLimit{}

// VpnSstpAuthenticationLocalUsersUsernameRateLimit describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VpnSstpAuthenticationLocalUsersUsernameRateLimit struct {
	// LeafNodes
	LeafVpnSstpAuthenticationLocalUsersUsernameRateLimitUpload   types.String `tfsdk:"upload" vyos:"upload,omitempty"`
	LeafVpnSstpAuthenticationLocalUsersUsernameRateLimitDownload types.String `tfsdk:"download" vyos:"download,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnSstpAuthenticationLocalUsersUsernameRateLimit) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"upload":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Upload bandwidth limit in kbits/sec

`,
			Description: `Upload bandwidth limit in kbits/sec

`,
		},

		"download":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Download bandwidth limit in kbits/sec

`,
			Description: `Download bandwidth limit in kbits/sec

`,
		},

		// TagNodes

		// Nodes

	}
}
