// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// FirewallGlobalOptionsStatePolicyEstablished describes the resource data model.
type FirewallGlobalOptionsStatePolicyEstablished struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafFirewallGlobalOptionsStatePolicyEstablishedAction   types.String `tfsdk:"action" vyos:"action,omitempty"`
	LeafFirewallGlobalOptionsStatePolicyEstablishedLog      types.Bool   `tfsdk:"log" vyos:"log,omitempty"`
	LeafFirewallGlobalOptionsStatePolicyEstablishedLogLevel types.String `tfsdk:"log_level" vyos:"log-level,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *FirewallGlobalOptionsStatePolicyEstablished) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallGlobalOptionsStatePolicyEstablished) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"firewall",

		"global-options",

		"state-policy",

		"established",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallGlobalOptionsStatePolicyEstablished) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Action for packets

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  accept  &emsp; |  Action to accept  |
    |  drop  &emsp; |  Action to drop  |
    |  reject  &emsp; |  Action to reject  |

`,
		},

		"log": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log packets hitting this rule

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"log_level": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set log-level. Log must be enable.

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  emerg  &emsp; |  Emerg log level  |
    |  alert  &emsp; |  Alert log level  |
    |  crit  &emsp; |  Critical log level  |
    |  err  &emsp; |  Error log level  |
    |  warn  &emsp; |  Warning log level  |
    |  notice  &emsp; |  Notice log level  |
    |  info  &emsp; |  Info log level  |
    |  debug  &emsp; |  Debug log level  |

`,
		},
	}
}