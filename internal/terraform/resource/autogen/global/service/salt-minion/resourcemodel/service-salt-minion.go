/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceSaltMinion{}

// ServiceSaltMinion describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ServiceSaltMinion struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceSaltMinionHash            types.String `tfsdk:"hash" vyos:"hash,omitempty"`
	LeafServiceSaltMinionMaster          types.List   `tfsdk:"master" vyos:"master,omitempty"`
	LeafServiceSaltMinionID              types.String `tfsdk:"id_param" vyos:"id,omitempty"`
	LeafServiceSaltMinionInterval        types.Number `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafServiceSaltMinionMasterKey       types.String `tfsdk:"master_key" vyos:"master-key,omitempty"`
	LeafServiceSaltMinionSourceInterface types.String `tfsdk:"source_interface" vyos:"source-interface,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceSaltMinion) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceSaltMinion) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceSaltMinion) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceSaltMinion) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"salt-minion",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceSaltMinion) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"service", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceSaltMinion) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSaltMinion) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"hash":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Hash used when discovering file on master server (default: sha256)

`,
			Description: `Hash used when discovering file on master server (default: sha256)

`,

			// Default:          stringdefault.StaticString(`sha256`),
			Computed: true,
		},

		"master":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Hostname or IP address of the Salt master server

    |  Format    |  Description               |
    |------------|----------------------------|
    |  ipv4      |  Salt server IPv4 address  |
    |  ipv6      |  Salt server IPv6 address  |
    |  hostname  |  Salt server FQDN address  |
`,
			Description: `Hostname or IP address of the Salt master server

    |  Format    |  Description               |
    |------------|----------------------------|
    |  ipv4      |  Salt server IPv4 address  |
    |  ipv6      |  Salt server IPv6 address  |
    |  hostname  |  Salt server FQDN address  |
`,
		},

		"id_param":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Explicitly declare ID for this minion to use (default: hostname)

`,
			Description: `Explicitly declare ID for this minion to use (default: hostname)

`,
		},

		"interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval in minutes between updates (default: 60)

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  1-1440  |  Update interval in minutes  |
`,
			Description: `Interval in minutes between updates (default: 60)

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  1-1440  |  Update interval in minutes  |
`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		"master_key":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `URL with signature of master for auth reply verification

`,
			Description: `URL with signature of master for auth reply verification

`,
		},

		"source_interface":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface used to establish connection

    |  Format     |  Description     |
    |-------------|------------------|
    |  interface  |  Interface name  |
`,
			Description: `Interface used to establish connection

    |  Format     |  Description     |
    |-------------|------------------|
    |  interface  |  Interface name  |
`,
		},

		// TagNodes

		// Nodes

	}
}
