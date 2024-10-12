// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceSnmpListenAddress{}

// ServiceSnmpListenAddress describes the resource data model.
type ServiceSnmpListenAddress struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceSnmpListenAddressPort types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceSnmpListenAddress) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceSnmpListenAddress) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceSnmpListenAddress) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceSnmpListenAddress) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"listen-address",
		o.SelfIdentifier.Attributes()["listen_address"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceSnmpListenAddress) GetVyosParentPath() []string {
	return []string{
		"service",

		"snmp",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceSnmpListenAddress) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSnmpListenAddress) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"listen_address": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `IP address to listen for incoming SNMP requests

    |  Format  |  Description                                        |
    |----------|-----------------------------------------------------|
    |  ipv4    |  IPv4 address to listen for incoming SNMP requests  |
    |  ipv6    |  IPv6 address to listen for incoming SNMP requests  |
`,
					Description: `IP address to listen for incoming SNMP requests

    |  Format  |  Description                                        |
    |----------|-----------------------------------------------------|
    |  ipv4    |  IPv4 address to listen for incoming SNMP requests  |
    |  ipv6    |  IPv6 address to listen for incoming SNMP requests  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in listen_address, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  listen_address, value must match: ^[a-zA-Z0-9-_]*$",
							),
						),
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"port": schema.NumberAttribute{
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

			// Default:          stringdefault.StaticString(`161`),
			Computed: true,
		},

		// Nodes

	}
}
