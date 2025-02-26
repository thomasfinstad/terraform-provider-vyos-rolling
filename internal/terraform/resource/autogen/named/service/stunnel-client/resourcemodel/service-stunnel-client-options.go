/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (options) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceStunnelClientOptions{}

// ServiceStunnelClientOptions describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ServiceStunnelClientOptions struct {
	// LeafNodes
	LeafServiceStunnelClientOptionsAuthentication types.String `tfsdk:"authentication" vyos:"authentication,omitempty"`
	LeafServiceStunnelClientOptionsDomain         types.String `tfsdk:"domain" vyos:"domain,omitempty"`
	LeafServiceStunnelClientOptionsPassword       types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafServiceStunnelClientOptionsUsername       types.String `tfsdk:"username" vyos:"username,omitempty"`

	// TagNodes

	// Nodes

	NodeServiceStunnelClientOptionsHost *ServiceStunnelClientOptionsHost `tfsdk:"host" vyos:"host,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceStunnelClientOptions) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"authentication":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (authentication) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Authentication type for the protocol negotiations

    |  Format  |  Description                                                |
    |----------|-------------------------------------------------------------|
    |  basic   |  The default 'connect' authentication type                  |
    |  ntlm    |  Supported authentication types for the 'connect' protocol  |
    |  plain   |  The default 'smtp' authentication type                     |
    |  login   |  Supported authentication types for the 'smtp' protocol     |
`,
			Description: `Authentication type for the protocol negotiations

    |  Format  |  Description                                                |
    |----------|-------------------------------------------------------------|
    |  basic   |  The default 'connect' authentication type                  |
    |  ntlm    |  Supported authentication types for the 'connect' protocol  |
    |  plain   |  The default 'smtp' authentication type                     |
    |  login   |  Supported authentication types for the 'smtp' protocol     |
`,
		},

		"domain":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (domain) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Domain for the 'connect' protocol.

    |  Format  |  Description  |
    |----------|---------------|
    |  domain  |  domain       |
`,
			Description: `Domain for the 'connect' protocol.

    |  Format  |  Description  |
    |----------|---------------|
    |  domain  |  domain       |
`,
		},

		"password":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (password) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Password for the protocol negotiations

    |  Format  |  Description              |
    |----------|---------------------------|
    |  txt     |  Authentication password  |
`,
			Description: `Password for the protocol negotiations

    |  Format  |  Description              |
    |----------|---------------------------|
    |  txt     |  Authentication password  |
`,
		},

		"username":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (username) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Username for the protocol negotiations

    |  Format  |  Description              |
    |----------|---------------------------|
    |  txt     |  Authentication username  |
`,
			Description: `Username for the protocol negotiations

    |  Format  |  Description              |
    |----------|---------------------------|
    |  txt     |  Authentication username  |
`,
		},

		// TagNodes

		// Nodes

		"host": schema.SingleNestedAttribute{
			Attributes: ServiceStunnelClientOptionsHost{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Destination address for the 'connect' protocol

`,
			Description: `Destination address for the 'connect' protocol

`,
		},
	}
}
