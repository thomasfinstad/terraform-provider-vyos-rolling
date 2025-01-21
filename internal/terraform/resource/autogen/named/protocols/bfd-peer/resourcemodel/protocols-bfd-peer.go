/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (peer) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsBfdPeer{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (peer) */
// ProtocolsBfdPeerSelfIdentifier used by TagNodes to keep the id info
type ProtocolsBfdPeerSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (peer) */

	ProtocolsBfdPeer types.String `tfsdk:"peer" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (bfd) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (protocols) */
}

// ProtocolsBfdPeer describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type ProtocolsBfdPeer struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *ProtocolsBfdPeerSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafProtocolsBfdPeerProfile    types.String `tfsdk:"profile" vyos:"profile,omitempty"`
	LeafProtocolsBfdPeerEchoMode   types.Bool   `tfsdk:"echo_mode" vyos:"echo-mode,omitempty"`
	LeafProtocolsBfdPeerMinimumTTL types.Number `tfsdk:"minimum_ttl" vyos:"minimum-ttl,omitempty"`
	LeafProtocolsBfdPeerPassive    types.Bool   `tfsdk:"passive" vyos:"passive,omitempty"`
	LeafProtocolsBfdPeerShutdown   types.Bool   `tfsdk:"shutdown" vyos:"shutdown,omitempty"`
	LeafProtocolsBfdPeerMultihop   types.Bool   `tfsdk:"multihop" vyos:"multihop,omitempty"`
	LeafProtocolsBfdPeerVrf        types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes

	// Nodes

	NodeProtocolsBfdPeerSource *ProtocolsBfdPeerSource `tfsdk:"source" vyos:"source,omitempty"`

	NodeProtocolsBfdPeerInterval *ProtocolsBfdPeerInterval `tfsdk:"interval" vyos:"interval,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsBfdPeer) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsBfdPeer) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsBfdPeer) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBfdPeer) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"peer",
		o.SelfIdentifier.ProtocolsBfdPeer.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsBfdPeer) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (bfd) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (protocols) */
		"protocols", // Node

		"bfd", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsBfdPeer) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (bfd) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (protocols) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBfdPeer) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"peer": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Configures BFD peer to listen and talk to

    |  Format  |  Description            |
    |----------|-------------------------|
    |  ipv4    |  BFD peer IPv4 address  |
    |  ipv6    |  BFD peer IPv6 address  |
`,
					Description: `Configures BFD peer to listen and talk to

    |  Format  |  Description            |
    |----------|-------------------------|
    |  ipv4    |  BFD peer IPv4 address  |
    |  ipv6    |  BFD peer IPv6 address  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in peer, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  peer, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (bfd) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (protocols) */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"profile":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (profile) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Use settings from BFD profile

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  BFD profile name  |
`,
			Description: `Use settings from BFD profile

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  BFD profile name  |
`,
		},

		"echo_mode":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (echo-mode) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enables the echo transmission mode

`,
			Description: `Enables the echo transmission mode

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"minimum_ttl":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (minimum-ttl) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Expect packets with at least this TTL

    |  Format  |  Description           |
    |----------|------------------------|
    |  1-254   |  Minimum TTL expected  |
`,
			Description: `Expect packets with at least this TTL

    |  Format  |  Description           |
    |----------|------------------------|
    |  1-254   |  Minimum TTL expected  |
`,
		},

		"passive":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (passive) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not attempt to start sessions

`,
			Description: `Do not attempt to start sessions

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"shutdown":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (shutdown) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable this peer

`,
			Description: `Disable this peer

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"multihop":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (multihop) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Allow this BFD peer to not be directly connected

`,
			Description: `Allow this BFD peer to not be directly connected

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"vrf":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (vrf) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
			Description: `VRF instance name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
		},

		// TagNodes

		// Nodes

		"source": schema.SingleNestedAttribute{
			Attributes: ProtocolsBfdPeerSource{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Bind listener to specified interface/address, mandatory for IPv6

`,
			Description: `Bind listener to specified interface/address, mandatory for IPv6

`,
		},

		"interval": schema.SingleNestedAttribute{
			Attributes: ProtocolsBfdPeerInterval{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Configure timer intervals

`,
			Description: `Configure timer intervals

`,
		},
	}
}
