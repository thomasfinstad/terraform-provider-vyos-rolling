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

var _ helpers.VyosTopResourceDataModel = &ServiceIDsDdosProtection{}

// ServiceIDsDdosProtection describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ServiceIDsDdosProtection struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceIDsDdosProtectionAlertScrIPt     types.String `tfsdk:"alert_script" vyos:"alert-script,omitempty"`
	LeafServiceIDsDdosProtectionBanTime         types.Number `tfsdk:"ban_time" vyos:"ban-time,omitempty"`
	LeafServiceIDsDdosProtectionDirection       types.List   `tfsdk:"direction" vyos:"direction,omitempty"`
	LeafServiceIDsDdosProtectionExcludedNetwork types.List   `tfsdk:"excluded_network" vyos:"excluded-network,omitempty"`
	LeafServiceIDsDdosProtectionListenInterface types.List   `tfsdk:"listen_interface" vyos:"listen-interface,omitempty"`
	LeafServiceIDsDdosProtectionMode            types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafServiceIDsDdosProtectionNetwork         types.List   `tfsdk:"network" vyos:"network,omitempty"`

	// TagNodes

	// Nodes

	ExistsNodeServiceIDsDdosProtectionSflow bool `tfsdk:"-" vyos:"sflow,child"`

	NodeServiceIDsDdosProtectionThreshold *ServiceIDsDdosProtectionThreshold `tfsdk:"threshold" vyos:"threshold,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceIDsDdosProtection) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceIDsDdosProtection) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceIDsDdosProtection) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceIDsDdosProtection) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"ddos-protection",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceIDsDdosProtection) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"service", // Node

		"ids", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceIDsDdosProtection) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIDsDdosProtection) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"alert_script":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Path to fastnetmon alert script

`,
			Description: `Path to fastnetmon alert script

`,
		},

		"ban_time":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `How long we should keep an IP in blocked state

    |  Format        |  Description      |
    |----------------|-------------------|
    |  1-4294967294  |  Time in seconds  |
`,
			Description: `How long we should keep an IP in blocked state

    |  Format        |  Description      |
    |----------------|-------------------|
    |  1-4294967294  |  Time in seconds  |
`,

			// Default:          stringdefault.StaticString(`1900`),
			Computed: true,
		},

		"direction":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Direction for processing traffic

`,
			Description: `Direction for processing traffic

`,
		},

		"excluded_network":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Specify IPv4 and IPv6 networks which are going to be excluded from protection

    |  Format   |  Description                 |
    |-----------|------------------------------|
    |  ipv4net  |  IPv4 prefix(es) to exclude  |
    |  ipv6net  |  IPv6 prefix(es) to exclude  |
`,
			Description: `Specify IPv4 and IPv6 networks which are going to be excluded from protection

    |  Format   |  Description                 |
    |-----------|------------------------------|
    |  ipv4net  |  IPv4 prefix(es) to exclude  |
    |  ipv6net  |  IPv6 prefix(es) to exclude  |
`,
		},

		"listen_interface":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Listen interface for mirroring traffic

`,
			Description: `Listen interface for mirroring traffic

`,
		},

		"mode":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Traffic capture mode

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  mirror  |  Listen to mirrored traffic  |
    |  sflow   |  Capture sFlow flows         |
`,
			Description: `Traffic capture mode

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  mirror  |  Listen to mirrored traffic  |
    |  sflow   |  Capture sFlow flows         |
`,
		},

		"network":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Specify IPv4 and IPv6 networks which belong to you

    |  Format   |  Description           |
    |-----------|------------------------|
    |  ipv4net  |  Your IPv4 prefix(es)  |
    |  ipv6net  |  Your IPv6 prefix(es)  |
`,
			Description: `Specify IPv4 and IPv6 networks which belong to you

    |  Format   |  Description           |
    |-----------|------------------------|
    |  ipv4net  |  Your IPv4 prefix(es)  |
    |  ipv6net  |  Your IPv6 prefix(es)  |
`,
		},

		// TagNodes

		// Nodes

		"threshold": schema.SingleNestedAttribute{
			Attributes: ServiceIDsDdosProtectionThreshold{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Attack limits thresholds

`,
			Description: `Attack limits thresholds

`,
		},
	}
}
