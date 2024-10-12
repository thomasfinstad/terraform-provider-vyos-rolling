// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceDNSForwardingAuthoritativeDomainRecordsSpf{}

// ServiceDNSForwardingAuthoritativeDomainRecordsSpf describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecordsSpf struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceDNSForwardingAuthoritativeDomainRecordsSpfValue   types.String `tfsdk:"value" vyos:"value,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsSpfTTL     types.Number `tfsdk:"ttl" vyos:"ttl,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsSpfDisable types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsSpf) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsSpf) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsSpf) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsSpf) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"spf",
		o.SelfIdentifier.Attributes()["spf"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsSpf) GetVyosParentPath() []string {
	return []string{
		"service",

		"dns",

		"forwarding",

		"authoritative-domain",

		o.SelfIdentifier.Attributes()["authoritative_domain"].(types.String).ValueString(),

		"records",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsSpf) GetVyosNamedParentPath() []string {
	return []string{
		"service",

		"dns",

		"forwarding",

		"authoritative-domain",

		o.SelfIdentifier.Attributes()["authoritative_domain"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsSpf) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"spf": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `SPF record

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  txt     |  A DNS name relative to the root record  |
    |  @       |  Root record                             |
`,
					Description: `SPF record

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  txt     |  A DNS name relative to the root record  |
    |  @       |  Root record                             |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in spf, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  spf, value must match: ^[a-zA-Z0-9-_]*$",
							),
						),
					},
				},

				"authoritative_domain": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Domain to host authoritative records for

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  txt     |  An absolute DNS domain name  |
`,
					Description: `Domain to host authoritative records for

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  txt     |  An absolute DNS domain name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in authoritative_domain, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  authoritative_domain, value must match: ^[a-zA-Z0-9-_]*$",
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

		"value": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Record contents

    |  Format  |  Description      |
    |----------|-------------------|
    |  txt     |  Record contents  |
`,
			Description: `Record contents

    |  Format  |  Description      |
    |----------|-------------------|
    |  txt     |  Record contents  |
`,
		},

		"ttl": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Time-to-live (TTL)

    |  Format        |  Description     |
    |----------------|------------------|
    |  0-2147483647  |  TTL in seconds  |
`,
			Description: `Time-to-live (TTL)

    |  Format        |  Description     |
    |----------------|------------------|
    |  0-2147483647  |  TTL in seconds  |
`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
