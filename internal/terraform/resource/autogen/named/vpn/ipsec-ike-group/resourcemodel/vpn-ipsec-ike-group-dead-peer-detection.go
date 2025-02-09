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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (dead-peer-detection) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VpnIPsecIkeGroupDeadPeerDetection{}

// VpnIPsecIkeGroupDeadPeerDetection describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VpnIPsecIkeGroupDeadPeerDetection struct {
	// LeafNodes
	LeafVpnIPsecIkeGroupDeadPeerDetectionAction   types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafVpnIPsecIkeGroupDeadPeerDetectionInterval types.Number `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafVpnIPsecIkeGroupDeadPeerDetectionTimeout  types.Number `tfsdk:"timeout" vyos:"timeout,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecIkeGroupDeadPeerDetection) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"action":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (action) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Keep-alive failure action

    |  Format   |  Description                                                           |
    |-----------|------------------------------------------------------------------------|
    |  trap     |  Attempt to re-negotiate the connection when matching traffic is seen  |
    |  clear    |  Remove the connection immediately                                     |
    |  restart  |  Attempt to re-negotiate the connection immediately                    |
`,
			Description: `Keep-alive failure action

    |  Format   |  Description                                                           |
    |-----------|------------------------------------------------------------------------|
    |  trap     |  Attempt to re-negotiate the connection when matching traffic is seen  |
    |  clear    |  Remove the connection immediately                                     |
    |  restart  |  Attempt to re-negotiate the connection immediately                    |
`,

			// Default:          stringdefault.StaticString(`clear`),
			Computed: true,
		},

		"interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (interval) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Keep-alive interval

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  2-86400  |  Keep-alive interval in seconds  |
`,
			Description: `Keep-alive interval

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  2-86400  |  Keep-alive interval in seconds  |
`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"timeout":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (timeout) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Dead Peer Detection keep-alive timeout (IKEv1 only)

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  2-86400  |  Keep-alive timeout in seconds  |
`,
			Description: `Dead Peer Detection keep-alive timeout (IKEv1 only)

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  2-86400  |  Keep-alive timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`120`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
