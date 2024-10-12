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

var _ helpers.VyosTopResourceDataModel = &ProtocolsOpenfabricDomain{}

// ProtocolsOpenfabricDomain describes the resource data model.
type ProtocolsOpenfabricDomain struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsOpenfabricDomainLogAdjacencyChanges types.Bool   `tfsdk:"log_adjacency_changes" vyos:"log-adjacency-changes,omitempty"`
	LeafProtocolsOpenfabricDomainSetOverloadBit      types.Bool   `tfsdk:"set_overload_bit" vyos:"set-overload-bit,omitempty"`
	LeafProtocolsOpenfabricDomainPurgeOriginator     types.Bool   `tfsdk:"purge_originator" vyos:"purge-originator,omitempty"`
	LeafProtocolsOpenfabricDomainFabricTier          types.Number `tfsdk:"fabric_tier" vyos:"fabric-tier,omitempty"`
	LeafProtocolsOpenfabricDomainLspGenInterval      types.Number `tfsdk:"lsp_gen_interval" vyos:"lsp-gen-interval,omitempty"`
	LeafProtocolsOpenfabricDomainLspRefreshInterval  types.Number `tfsdk:"lsp_refresh_interval" vyos:"lsp-refresh-interval,omitempty"`
	LeafProtocolsOpenfabricDomainMaxLspLifetime      types.Number `tfsdk:"max_lsp_lifetime" vyos:"max-lsp-lifetime,omitempty"`
	LeafProtocolsOpenfabricDomainSpfInterval         types.Number `tfsdk:"spf_interval" vyos:"spf-interval,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagProtocolsOpenfabricDomainInterface bool `tfsdk:"-" vyos:"interface,child"`

	// Nodes
	NodeProtocolsOpenfabricDomainDomainPassword *ProtocolsOpenfabricDomainDomainPassword `tfsdk:"domain_password" vyos:"domain-password,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsOpenfabricDomain) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsOpenfabricDomain) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsOpenfabricDomain) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOpenfabricDomain) GetVyosPath() []string {
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
func (o *ProtocolsOpenfabricDomain) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"openfabric",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsOpenfabricDomain) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOpenfabricDomain) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
					MarkdownDescription: `OpenFabric process name

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Domain name  |
`,
					Description: `OpenFabric process name

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Domain name  |
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

		"log_adjacency_changes": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log changes in adjacency state

`,
			Description: `Log changes in adjacency state

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"set_overload_bit": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Overload bit to avoid any transit traffic

`,
			Description: `Overload bit to avoid any transit traffic

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"purge_originator": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `RFC 6232 purge originator identification

`,
			Description: `RFC 6232 purge originator identification

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"fabric_tier": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Static tier number to advertise as location in the fabric

    |  Format  |  Description         |
    |----------|----------------------|
    |  0-14    |  Static tier number  |
`,
			Description: `Static tier number to advertise as location in the fabric

    |  Format  |  Description         |
    |----------|----------------------|
    |  0-14    |  Static tier number  |
`,
		},

		"lsp_gen_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum interval between regenerating same link-state packet (LSP)

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  1-120   |  Minimum interval in seconds  |
`,
			Description: `Minimum interval between regenerating same link-state packet (LSP)

    |  Format  |  Description                  |
    |----------|-------------------------------|
    |  1-120   |  Minimum interval in seconds  |
`,
		},

		"lsp_refresh_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Link-state packet (LSP) refresh interval

    |  Format   |  Description                      |
    |-----------|-----------------------------------|
    |  1-65235  |  LSP refresh interval in seconds  |
`,
			Description: `Link-state packet (LSP) refresh interval

    |  Format   |  Description                      |
    |-----------|-----------------------------------|
    |  1-65235  |  LSP refresh interval in seconds  |
`,
		},

		"max_lsp_lifetime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum link-state packet lifetime

    |  Format     |  Description                      |
    |-------------|-----------------------------------|
    |  360-65535  |  Maximum LSP lifetime in seconds  |
`,
			Description: `Maximum link-state packet lifetime

    |  Format     |  Description                      |
    |-------------|-----------------------------------|
    |  360-65535  |  Maximum LSP lifetime in seconds  |
`,
		},

		"spf_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum interval between SPF calculations

    |  Format  |  Description          |
    |----------|-----------------------|
    |  1-120   |  Interval in seconds  |
`,
			Description: `Minimum interval between SPF calculations

    |  Format  |  Description          |
    |----------|-----------------------|
    |  1-120   |  Interval in seconds  |
`,
		},

		// Nodes

		"domain_password": schema.SingleNestedAttribute{
			Attributes: ProtocolsOpenfabricDomainDomainPassword{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Authentication password for a routing domain

`,
			Description: `Authentication password for a routing domain

`,
		},
	}
}
