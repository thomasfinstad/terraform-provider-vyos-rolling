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

var _ helpers.VyosTopResourceDataModel = &VpnIPsecEspGroup{}

// VpnIPsecEspGroup describes the resource data model.
type VpnIPsecEspGroup struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVpnIPsecEspGroupCompression  types.Bool   `tfsdk:"compression" vyos:"compression,omitempty"`
	LeafVpnIPsecEspGroupLifetime     types.Number `tfsdk:"lifetime" vyos:"lifetime,omitempty"`
	LeafVpnIPsecEspGroupLifeBytes    types.Number `tfsdk:"life_bytes" vyos:"life-bytes,omitempty"`
	LeafVpnIPsecEspGroupLifePackets  types.Number `tfsdk:"life_packets" vyos:"life-packets,omitempty"`
	LeafVpnIPsecEspGroupDisableRekey types.Bool   `tfsdk:"disable_rekey" vyos:"disable-rekey,omitempty"`
	LeafVpnIPsecEspGroupMode         types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafVpnIPsecEspGroupPfs          types.String `tfsdk:"pfs" vyos:"pfs,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagVpnIPsecEspGroupProposal bool `tfsdk:"-" vyos:"proposal,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *VpnIPsecEspGroup) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnIPsecEspGroup) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnIPsecEspGroup) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnIPsecEspGroup) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"esp-group",
		o.SelfIdentifier.Attributes()["esp_group"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnIPsecEspGroup) GetVyosParentPath() []string {
	return []string{
		"vpn",

		"ipsec",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VpnIPsecEspGroup) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecEspGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"esp_group": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Encapsulating Security Payload (ESP) group name

`,
					Description: `Encapsulating Security Payload (ESP) group name

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in esp_group, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  esp_group, value must match: ^[a-zA-Z0-9-_]*$",
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

		"compression": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable ESP compression

`,
			Description: `Enable ESP compression

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"lifetime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Security Association time to expire

    |  Format    |  Description             |
    |------------|--------------------------|
    |  30-86400  |  SA lifetime in seconds  |
`,
			Description: `Security Association time to expire

    |  Format    |  Description             |
    |------------|--------------------------|
    |  30-86400  |  SA lifetime in seconds  |
`,

			// Default:          stringdefault.StaticString(`3600`),
			Computed: true,
		},

		"life_bytes": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Security Association byte count to expire

    |  Format               |  Description       |
    |-----------------------|--------------------|
    |  1024-26843545600000  |  SA life in bytes  |
`,
			Description: `Security Association byte count to expire

    |  Format               |  Description       |
    |-----------------------|--------------------|
    |  1024-26843545600000  |  SA life in bytes  |
`,
		},

		"life_packets": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Security Association packet count to expire

    |  Format               |  Description         |
    |-----------------------|----------------------|
    |  1000-26843545600000  |  SA life in packets  |
`,
			Description: `Security Association packet count to expire

    |  Format               |  Description         |
    |-----------------------|----------------------|
    |  1000-26843545600000  |  SA life in packets  |
`,
		},

		"disable_rekey": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not locally initiate a re-key of the SA, remote peer must re-key before expiration

`,
			Description: `Do not locally initiate a re-key of the SA, remote peer must re-key before expiration

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ESP mode

    |  Format     |  Description     |
    |-------------|------------------|
    |  tunnel     |  Tunnel mode     |
    |  transport  |  Transport mode  |
`,
			Description: `ESP mode

    |  Format     |  Description     |
    |-------------|------------------|
    |  tunnel     |  Tunnel mode     |
    |  transport  |  Transport mode  |
`,

			// Default:          stringdefault.StaticString(`tunnel`),
			Computed: true,
		},

		"pfs": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ESP Perfect Forward Secrecy

    |  Format      |  Description                                      |
    |--------------|---------------------------------------------------|
    |  enable      |  Inherit Diffie-Hellman group from the IKE group  |
    |  dh-group1   |  Use Diffie-Hellman group 1 (modp768)             |
    |  dh-group2   |  Use Diffie-Hellman group 2 (modp1024)            |
    |  dh-group5   |  Use Diffie-Hellman group 5 (modp1536)            |
    |  dh-group14  |  Use Diffie-Hellman group 14 (modp2048)           |
    |  dh-group15  |  Use Diffie-Hellman group 15 (modp3072)           |
    |  dh-group16  |  Use Diffie-Hellman group 16 (modp4096)           |
    |  dh-group17  |  Use Diffie-Hellman group 17 (modp6144)           |
    |  dh-group18  |  Use Diffie-Hellman group 18 (modp8192)           |
    |  dh-group19  |  Use Diffie-Hellman group 19 (ecp256)             |
    |  dh-group20  |  Use Diffie-Hellman group 20 (ecp384)             |
    |  dh-group21  |  Use Diffie-Hellman group 21 (ecp521)             |
    |  dh-group22  |  Use Diffie-Hellman group 22 (modp1024s160)       |
    |  dh-group23  |  Use Diffie-Hellman group 23 (modp2048s224)       |
    |  dh-group24  |  Use Diffie-Hellman group 24 (modp2048s256)       |
    |  dh-group25  |  Use Diffie-Hellman group 25 (ecp192)             |
    |  dh-group26  |  Use Diffie-Hellman group 26 (ecp224)             |
    |  dh-group27  |  Use Diffie-Hellman group 27 (ecp224bp)           |
    |  dh-group28  |  Use Diffie-Hellman group 28 (ecp256bp)           |
    |  dh-group29  |  Use Diffie-Hellman group 29 (ecp384bp)           |
    |  dh-group30  |  Use Diffie-Hellman group 30 (ecp512bp)           |
    |  dh-group31  |  Use Diffie-Hellman group 31 (curve25519)         |
    |  dh-group32  |  Use Diffie-Hellman group 32 (curve448)           |
    |  disable     |  Disable PFS                                      |
`,
			Description: `ESP Perfect Forward Secrecy

    |  Format      |  Description                                      |
    |--------------|---------------------------------------------------|
    |  enable      |  Inherit Diffie-Hellman group from the IKE group  |
    |  dh-group1   |  Use Diffie-Hellman group 1 (modp768)             |
    |  dh-group2   |  Use Diffie-Hellman group 2 (modp1024)            |
    |  dh-group5   |  Use Diffie-Hellman group 5 (modp1536)            |
    |  dh-group14  |  Use Diffie-Hellman group 14 (modp2048)           |
    |  dh-group15  |  Use Diffie-Hellman group 15 (modp3072)           |
    |  dh-group16  |  Use Diffie-Hellman group 16 (modp4096)           |
    |  dh-group17  |  Use Diffie-Hellman group 17 (modp6144)           |
    |  dh-group18  |  Use Diffie-Hellman group 18 (modp8192)           |
    |  dh-group19  |  Use Diffie-Hellman group 19 (ecp256)             |
    |  dh-group20  |  Use Diffie-Hellman group 20 (ecp384)             |
    |  dh-group21  |  Use Diffie-Hellman group 21 (ecp521)             |
    |  dh-group22  |  Use Diffie-Hellman group 22 (modp1024s160)       |
    |  dh-group23  |  Use Diffie-Hellman group 23 (modp2048s224)       |
    |  dh-group24  |  Use Diffie-Hellman group 24 (modp2048s256)       |
    |  dh-group25  |  Use Diffie-Hellman group 25 (ecp192)             |
    |  dh-group26  |  Use Diffie-Hellman group 26 (ecp224)             |
    |  dh-group27  |  Use Diffie-Hellman group 27 (ecp224bp)           |
    |  dh-group28  |  Use Diffie-Hellman group 28 (ecp256bp)           |
    |  dh-group29  |  Use Diffie-Hellman group 29 (ecp384bp)           |
    |  dh-group30  |  Use Diffie-Hellman group 30 (ecp512bp)           |
    |  dh-group31  |  Use Diffie-Hellman group 31 (curve25519)         |
    |  dh-group32  |  Use Diffie-Hellman group 32 (curve448)           |
    |  disable     |  Disable PFS                                      |
`,

			// Default:          stringdefault.StaticString(`enable`),
			Computed: true,
		},

		// Nodes

	}
}
