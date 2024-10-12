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

var _ helpers.VyosTopResourceDataModel = &VpnLtwotpRemoteAccessAuthenticationLocalUsersUsername{}

// VpnLtwotpRemoteAccessAuthenticationLocalUsersUsername describes the resource data model.
type VpnLtwotpRemoteAccessAuthenticationLocalUsersUsername struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVpnLtwotpRemoteAccessAuthenticationLocalUsersUsernameDisable  types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafVpnLtwotpRemoteAccessAuthenticationLocalUsersUsernamePassword types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafVpnLtwotpRemoteAccessAuthenticationLocalUsersUsernameStaticIP types.String `tfsdk:"static_ip" vyos:"static-ip,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeVpnLtwotpRemoteAccessAuthenticationLocalUsersUsernameRateLimit *VpnLtwotpRemoteAccessAuthenticationLocalUsersUsernameRateLimit `tfsdk:"rate_limit" vyos:"rate-limit,omitempty"`
}

// SetID configures the resource ID
func (o *VpnLtwotpRemoteAccessAuthenticationLocalUsersUsername) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnLtwotpRemoteAccessAuthenticationLocalUsersUsername) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnLtwotpRemoteAccessAuthenticationLocalUsersUsername) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnLtwotpRemoteAccessAuthenticationLocalUsersUsername) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"username",
		o.SelfIdentifier.Attributes()["username"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnLtwotpRemoteAccessAuthenticationLocalUsersUsername) GetVyosParentPath() []string {
	return []string{
		"vpn",

		"l2tp",

		"remote-access",

		"authentication",

		"local-users",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VpnLtwotpRemoteAccessAuthenticationLocalUsersUsername) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnLtwotpRemoteAccessAuthenticationLocalUsersUsername) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"username": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `User name for authentication

`,
					Description: `User name for authentication

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in username, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  username, value must match: ^[a-zA-Z0-9-_]*$",
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

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Password for authentication

`,
			Description: `Password for authentication

`,
		},

		"static_ip": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Static client IP address

`,
			Description: `Static client IP address

`,

			// Default:          stringdefault.StaticString(`&`),
			Computed: true,
		},

		// Nodes

		"rate_limit": schema.SingleNestedAttribute{
			Attributes: VpnLtwotpRemoteAccessAuthenticationLocalUsersUsernameRateLimit{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Upload/Download speed limits

`,
			Description: `Upload/Download speed limits

`,
		},
	}
}
