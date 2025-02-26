/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (tcp) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &FirewallGlobalOptionsTimeoutTCP{}

// FirewallGlobalOptionsTimeoutTCP describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type FirewallGlobalOptionsTimeoutTCP struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafFirewallGlobalOptionsTimeoutTCPCloseWait   types.Number `tfsdk:"close_wait" vyos:"close-wait,omitempty"`
	LeafFirewallGlobalOptionsTimeoutTCPClose       types.Number `tfsdk:"close" vyos:"close,omitempty"`
	LeafFirewallGlobalOptionsTimeoutTCPEstablished types.Number `tfsdk:"established" vyos:"established,omitempty"`
	LeafFirewallGlobalOptionsTimeoutTCPFinWait     types.Number `tfsdk:"fin_wait" vyos:"fin-wait,omitempty"`
	LeafFirewallGlobalOptionsTimeoutTCPLastAck     types.Number `tfsdk:"last_ack" vyos:"last-ack,omitempty"`
	LeafFirewallGlobalOptionsTimeoutTCPSynRecv     types.Number `tfsdk:"syn_recv" vyos:"syn-recv,omitempty"`
	LeafFirewallGlobalOptionsTimeoutTCPSynSent     types.Number `tfsdk:"syn_sent" vyos:"syn-sent,omitempty"`
	LeafFirewallGlobalOptionsTimeoutTCPTimeWait    types.Number `tfsdk:"time_wait" vyos:"time-wait,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *FirewallGlobalOptionsTimeoutTCP) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *FirewallGlobalOptionsTimeoutTCP) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *FirewallGlobalOptionsTimeoutTCP) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallGlobalOptionsTimeoutTCP) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"tcp",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *FirewallGlobalOptionsTimeoutTCP) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (timeout) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (global-options) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (firewall) */
		"firewall", // Node

		"global-options", // Node

		"timeout", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *FirewallGlobalOptionsTimeoutTCP) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (timeout) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (global-options) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (firewall) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallGlobalOptionsTimeoutTCP) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"close_wait":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (close-wait) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP CLOSE-WAIT timeout in seconds

    |  Format      |  Description                        |
    |--------------|-------------------------------------|
    |  1-21474836  |  TCP CLOSE-WAIT timeout in seconds  |
`,
			Description: `TCP CLOSE-WAIT timeout in seconds

    |  Format      |  Description                        |
    |--------------|-------------------------------------|
    |  1-21474836  |  TCP CLOSE-WAIT timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		"close":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (close) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP CLOSE timeout in seconds

    |  Format      |  Description                   |
    |--------------|--------------------------------|
    |  1-21474836  |  TCP CLOSE timeout in seconds  |
`,
			Description: `TCP CLOSE timeout in seconds

    |  Format      |  Description                   |
    |--------------|--------------------------------|
    |  1-21474836  |  TCP CLOSE timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"established":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (established) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP ESTABLISHED timeout in seconds

    |  Format      |  Description                         |
    |--------------|--------------------------------------|
    |  1-21474836  |  TCP ESTABLISHED timeout in seconds  |
`,
			Description: `TCP ESTABLISHED timeout in seconds

    |  Format      |  Description                         |
    |--------------|--------------------------------------|
    |  1-21474836  |  TCP ESTABLISHED timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`432000`),
			Computed: true,
		},

		"fin_wait":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (fin-wait) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP FIN-WAIT timeout in seconds

    |  Format      |  Description                      |
    |--------------|-----------------------------------|
    |  1-21474836  |  TCP FIN-WAIT timeout in seconds  |
`,
			Description: `TCP FIN-WAIT timeout in seconds

    |  Format      |  Description                      |
    |--------------|-----------------------------------|
    |  1-21474836  |  TCP FIN-WAIT timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`120`),
			Computed: true,
		},

		"last_ack":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (last-ack) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP LAST-ACK timeout in seconds

    |  Format      |  Description                      |
    |--------------|-----------------------------------|
    |  1-21474836  |  TCP LAST-ACK timeout in seconds  |
`,
			Description: `TCP LAST-ACK timeout in seconds

    |  Format      |  Description                      |
    |--------------|-----------------------------------|
    |  1-21474836  |  TCP LAST-ACK timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"syn_recv":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (syn-recv) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP SYN-RECEIVED timeout in seconds

    |  Format      |  Description                          |
    |--------------|---------------------------------------|
    |  1-21474836  |  TCP SYN-RECEIVED timeout in seconds  |
`,
			Description: `TCP SYN-RECEIVED timeout in seconds

    |  Format      |  Description                          |
    |--------------|---------------------------------------|
    |  1-21474836  |  TCP SYN-RECEIVED timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		"syn_sent":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (syn-sent) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP SYN-SENT timeout in seconds

    |  Format      |  Description                      |
    |--------------|-----------------------------------|
    |  1-21474836  |  TCP SYN-SENT timeout in seconds  |
`,
			Description: `TCP SYN-SENT timeout in seconds

    |  Format      |  Description                      |
    |--------------|-----------------------------------|
    |  1-21474836  |  TCP SYN-SENT timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`120`),
			Computed: true,
		},

		"time_wait":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (time-wait) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TCP TIME-WAIT timeout in seconds

    |  Format      |  Description                       |
    |--------------|------------------------------------|
    |  1-21474836  |  TCP TIME-WAIT timeout in seconds  |
`,
			Description: `TCP TIME-WAIT timeout in seconds

    |  Format      |  Description                       |
    |--------------|------------------------------------|
    |  1-21474836  |  TCP TIME-WAIT timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`120`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
