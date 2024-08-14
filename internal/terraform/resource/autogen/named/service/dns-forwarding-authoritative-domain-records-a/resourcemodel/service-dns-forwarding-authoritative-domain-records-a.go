// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
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

var _ helpers.VyosTopResourceDataModel = &ServiceDNSForwardingAuthoritativeDomainRecordsA{}

// ServiceDNSForwardingAuthoritativeDomainRecordsA describes the resource data model.
type ServiceDNSForwardingAuthoritativeDomainRecordsA struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceDNSForwardingAuthoritativeDomainRecordsAAddress types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsATTL     types.Number `tfsdk:"ttl" vyos:"ttl,omitempty"`
	LeafServiceDNSForwardingAuthoritativeDomainRecordsADisable types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsA) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsA) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsA) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsA) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"a",
		o.SelfIdentifier.Attributes()["a"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsA) GetVyosParentPath() []string {
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
func (o *ServiceDNSForwardingAuthoritativeDomainRecordsA) GetVyosNamedParentPath() []string {
	return []string{
		"service",

		"dns",

		"forwarding",

		"authoritative-domain",

		o.SelfIdentifier.Attributes()["authoritative_domain"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingAuthoritativeDomainRecordsA) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"a": schema.StringAttribute{
						Required: true,
						MarkdownDescription: `A record

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  txt     |  A DNS name relative to the root record  |
    |  @       |  Root record                             |
    |  any     |  Wildcard record (any subdomain)         |
`,
						Description: `A record

    |  Format  |  Description                             |
    |----------|------------------------------------------|
    |  txt     |  A DNS name relative to the root record  |
    |  @       |  Root record                             |
    |  any     |  Wildcard record (any subdomain)         |
`,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						}, Validators: []validator.String{
							stringvalidator.All(
								helpers.StringNot(
									stringvalidator.RegexMatches(
										regexp.MustCompile(`^.*__.*$`),
										"double underscores in a, conflicts with the internal resource id",
									),
								),
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  a, value must match: ^[a-zA-Z0-9-_]*$",
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
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IPv4 address

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
`,
			Description: `IPv4 address

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
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