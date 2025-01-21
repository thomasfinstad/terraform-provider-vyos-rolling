/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (mld) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsPimsixInterfaceMld{}

// ProtocolsPimsixInterfaceMld describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type ProtocolsPimsixInterfaceMld struct {
	// LeafNodes
	LeafProtocolsPimsixInterfaceMldDisable                 types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafProtocolsPimsixInterfaceMldLastMemberQueryCount    types.Number `tfsdk:"last_member_query_count" vyos:"last-member-query-count,omitempty"`
	LeafProtocolsPimsixInterfaceMldLastMemberQueryInterval types.Number `tfsdk:"last_member_query_interval" vyos:"last-member-query-interval,omitempty"`
	LeafProtocolsPimsixInterfaceMldInterval                types.Number `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafProtocolsPimsixInterfaceMldMaxResponseTime         types.Number `tfsdk:"max_response_time" vyos:"max-response-time,omitempty"`
	LeafProtocolsPimsixInterfaceMldVersion                 types.String `tfsdk:"version" vyos:"version,omitempty"`

	// TagNodes

	ExistsTagProtocolsPimsixInterfaceMldJoin bool `tfsdk:"-" vyos:"join,child"`

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsPimsixInterfaceMld) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"disable":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (disable) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"last_member_query_count":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (last-member-query-count) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Last member query count

    |  Format  |  Description  |
    |----------|---------------|
    |  1-255   |  Count        |
`,
			Description: `Last member query count

    |  Format  |  Description  |
    |----------|---------------|
    |  1-255   |  Count        |
`,
		},

		"last_member_query_interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (last-member-query-interval) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Last member query interval

    |  Format       |  Description                                 |
    |---------------|----------------------------------------------|
    |  100-6553500  |  Last member query interval in milliseconds  |
`,
			Description: `Last member query interval

    |  Format       |  Description                                 |
    |---------------|----------------------------------------------|
    |  100-6553500  |  Last member query interval in milliseconds  |
`,
		},

		"interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (interval) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Query interval

    |  Format   |  Description                |
    |-----------|-----------------------------|
    |  1-65535  |  Query interval in seconds  |
`,
			Description: `Query interval

    |  Format   |  Description                |
    |-----------|-----------------------------|
    |  1-65535  |  Query interval in seconds  |
`,
		},

		"max_response_time":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (max-response-time) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Max query response time

    |  Format       |  Description                           |
    |---------------|----------------------------------------|
    |  100-6553500  |  Query response value in milliseconds  |
`,
			Description: `Max query response time

    |  Format       |  Description                           |
    |---------------|----------------------------------------|
    |  100-6553500  |  Query response value in milliseconds  |
`,
		},

		"version":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (version) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `MLD version

    |  Format  |  Description    |
    |----------|-----------------|
    |  1       |  MLD version 1  |
    |  2       |  MLD version 2  |
`,
			Description: `MLD version

    |  Format  |  Description    |
    |----------|-----------------|
    |  1       |  MLD version 1  |
    |  2       |  MLD version 2  |
`,

			// Default:          stringdefault.StaticString(`2`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
