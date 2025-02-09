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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (interface) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ServiceConntrackSyncInterface{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (interface) */
// ServiceConntrackSyncInterfaceSelfIdentifier used by TagNodes to keep the id info
type ServiceConntrackSyncInterfaceSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (interface) */

	ServiceConntrackSyncInterface types.String `tfsdk:"interface" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (conntrack-sync) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (service) */
}

// ServiceConntrackSyncInterface describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ServiceConntrackSyncInterface struct {
	// LeafNodes
	LeafServiceConntrackSyncInterfacePeer types.String `tfsdk:"peer" vyos:"peer,omitempty"`
	LeafServiceConntrackSyncInterfacePort types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceConntrackSyncInterface) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"peer":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (peer) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address of the peer to send the UDP conntrack info too. This disable multicast.

    |  Format  |  Description                                    |
    |----------|-------------------------------------------------|
    |  ipv4    |  IP address to listen for incoming connections  |
`,
			Description: `IP address of the peer to send the UDP conntrack info too. This disable multicast.

    |  Format  |  Description                                    |
    |----------|-------------------------------------------------|
    |  ipv4    |  IP address to listen for incoming connections  |
`,
		},

		"port":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (port) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
			Description: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
		},

		// TagNodes

		// Nodes

	}
}
