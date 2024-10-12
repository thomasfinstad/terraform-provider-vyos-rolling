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

var _ helpers.VyosTopResourceDataModel = &ServiceDNSForwardingDomain{}

// ServiceDNSForwardingDomain describes the resource data model.
type ServiceDNSForwardingDomain struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceDNSForwardingDomainAddnta           types.Bool `tfsdk:"addnta" vyos:"addnta,omitempty"`
	LeafServiceDNSForwardingDomainRecursionDesired types.Bool `tfsdk:"recursion_desired" vyos:"recursion-desired,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagServiceDNSForwardingDomainNameServer bool `tfsdk:"-" vyos:"name-server,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceDNSForwardingDomain) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceDNSForwardingDomain) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceDNSForwardingDomain) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDNSForwardingDomain) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"domain",
		o.SelfIdentifier.Attributes()["domain"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceDNSForwardingDomain) GetVyosParentPath() []string {
	return []string{
		"service",

		"dns",

		"forwarding",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceDNSForwardingDomain) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDNSForwardingDomain) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"domain": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Domain to forward to a custom DNS server

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  txt     |  An absolute DNS domain name  |
`,
					Description: `Domain to forward to a custom DNS server

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
									"double underscores in domain, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  domain, value must match: ^[a-zA-Z0-9-_]*$",
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

		"addnta": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Add NTA (negative trust anchor) for this domain (must be set if the domain does not support DNSSEC)

`,
			Description: `Add NTA (negative trust anchor) for this domain (must be set if the domain does not support DNSSEC)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"recursion_desired": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Set the "recursion desired" bit in requests to the upstream nameserver

`,
			Description: `Set the "recursion desired" bit in requests to the upstream nameserver

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
