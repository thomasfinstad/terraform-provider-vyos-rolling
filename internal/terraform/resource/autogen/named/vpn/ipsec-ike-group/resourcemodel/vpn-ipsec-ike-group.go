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

var _ helpers.VyosTopResourceDataModel = &VpnIPsecIkeGroup{}

// VpnIPsecIkeGroup describes the resource data model.
type VpnIPsecIkeGroup struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVpnIPsecIkeGroupCloseAction   types.String `tfsdk:"close_action" vyos:"close-action,omitempty"`
	LeafVpnIPsecIkeGroupIkevtwoReauth types.Bool   `tfsdk:"ikev2_reauth" vyos:"ikev2-reauth,omitempty"`
	LeafVpnIPsecIkeGroupKeyExchange   types.String `tfsdk:"key_exchange" vyos:"key-exchange,omitempty"`
	LeafVpnIPsecIkeGroupLifetime      types.Number `tfsdk:"lifetime" vyos:"lifetime,omitempty"`
	LeafVpnIPsecIkeGroupDisableMobike types.Bool   `tfsdk:"disable_mobike" vyos:"disable-mobike,omitempty"`
	LeafVpnIPsecIkeGroupMode          types.String `tfsdk:"mode" vyos:"mode,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagVpnIPsecIkeGroupProposal bool `tfsdk:"-" vyos:"proposal,child"`

	// Nodes
	NodeVpnIPsecIkeGroupDeadPeerDetection *VpnIPsecIkeGroupDeadPeerDetection `tfsdk:"dead_peer_detection" vyos:"dead-peer-detection,omitempty"`
}

// SetID configures the resource ID
func (o *VpnIPsecIkeGroup) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnIPsecIkeGroup) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnIPsecIkeGroup) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnIPsecIkeGroup) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"ike-group",
		o.SelfIdentifier.Attributes()["ike_group"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnIPsecIkeGroup) GetVyosParentPath() []string {
	return []string{
		"vpn",

		"ipsec",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VpnIPsecIkeGroup) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecIkeGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"ike_group": schema.StringAttribute{
						Required: true,
						MarkdownDescription: `Internet Key Exchange (IKE) group name

`,
						Description: `Internet Key Exchange (IKE) group name

`,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						}, Validators: []validator.String{
							stringvalidator.All(
								helpers.StringNot(
									stringvalidator.RegexMatches(
										regexp.MustCompile(`^.*__.*$`),
										"double underscores in ike_group, conflicts with the internal resource id",
									),
								),
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  ike_group, value must match: ^[a-zA-Z0-9-_]*$",
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

		"close_action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Action to take if a child SA is unexpectedly closed

    |  Format  |  Description                                            |
    |----------|---------------------------------------------------------|
    |  none    |  Do nothing                                             |
    |  trap    |  Attempt to re-negotiate when matching traffic is seen  |
    |  start   |  Attempt to re-negotiate the connection immediately     |
`,
			Description: `Action to take if a child SA is unexpectedly closed

    |  Format  |  Description                                            |
    |----------|---------------------------------------------------------|
    |  none    |  Do nothing                                             |
    |  trap    |  Attempt to re-negotiate when matching traffic is seen  |
    |  start   |  Attempt to re-negotiate the connection immediately     |
`,

			// Default:          stringdefault.StaticString(`none`),
			Computed: true,
		},

		"ikev2_reauth": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Re-authentication of the remote peer during an IKE re-key (IKEv2 only)

`,
			Description: `Re-authentication of the remote peer during an IKE re-key (IKEv2 only)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"key_exchange": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IKE version

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  ikev1   |  Use IKEv1 for key exchange  |
    |  ikev2   |  Use IKEv2 for key exchange  |
`,
			Description: `IKE version

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  ikev1   |  Use IKEv1 for key exchange  |
    |  ikev2   |  Use IKEv2 for key exchange  |
`,
		},

		"lifetime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `IKE lifetime

    |  Format   |  Description              |
    |-----------|---------------------------|
    |  0-86400  |  IKE lifetime in seconds  |
`,
			Description: `IKE lifetime

    |  Format   |  Description              |
    |-----------|---------------------------|
    |  0-86400  |  IKE lifetime in seconds  |
`,

			// Default:          stringdefault.StaticString(`28800`),
			Computed: true,
		},

		"disable_mobike": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable MOBIKE Support (IKEv2 only)

`,
			Description: `Disable MOBIKE Support (IKEv2 only)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IKEv1 phase 1 mode

    |  Format      |  Description                                          |
    |--------------|-------------------------------------------------------|
    |  main        |  Use the main mode (recommended)                      |
    |  aggressive  |  Use the aggressive mode (insecure, not recommended)  |
`,
			Description: `IKEv1 phase 1 mode

    |  Format      |  Description                                          |
    |--------------|-------------------------------------------------------|
    |  main        |  Use the main mode (recommended)                      |
    |  aggressive  |  Use the aggressive mode (insecure, not recommended)  |
`,

			// Default:          stringdefault.StaticString(`main`),
			Computed: true,
		},

		// Nodes

		"dead_peer_detection": schema.SingleNestedAttribute{
			Attributes: VpnIPsecIkeGroupDeadPeerDetection{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Dead Peer Detection (DPD)

`,
			Description: `Dead Peer Detection (DPD)

`,
		},
	}
}