/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsBgpPeerGroup{}

// ProtocolsBgpPeerGroup describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type ProtocolsBgpPeerGroup struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl */
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsBgpPeerGroupDescrIPtion                  types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafProtocolsBgpPeerGroupDisableCapabilityNegotiation types.Bool   `tfsdk:"disable_capability_negotiation" vyos:"disable-capability-negotiation,omitempty"`
	LeafProtocolsBgpPeerGroupDisableConnectedCheck        types.Bool   `tfsdk:"disable_connected_check" vyos:"disable-connected-check,omitempty"`
	LeafProtocolsBgpPeerGroupEbgpMultihop                 types.Number `tfsdk:"ebgp_multihop" vyos:"ebgp-multihop,omitempty"`
	LeafProtocolsBgpPeerGroupGracefulRestart              types.String `tfsdk:"graceful_restart" vyos:"graceful-restart,omitempty"`
	LeafProtocolsBgpPeerGroupOverrIDeCapability           types.Bool   `tfsdk:"override_capability" vyos:"override-capability,omitempty"`
	LeafProtocolsBgpPeerGroupPassive                      types.Bool   `tfsdk:"passive" vyos:"passive,omitempty"`
	LeafProtocolsBgpPeerGroupPassword                     types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafProtocolsBgpPeerGroupShutdown                     types.Bool   `tfsdk:"shutdown" vyos:"shutdown,omitempty"`
	LeafProtocolsBgpPeerGroupUpdateSource                 types.String `tfsdk:"update_source" vyos:"update-source,omitempty"`
	LeafProtocolsBgpPeerGroupRemoteAs                     types.String `tfsdk:"remote_as" vyos:"remote-as,omitempty"`
	LeafProtocolsBgpPeerGroupPort                         types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes

	ExistsTagProtocolsBgpPeerGroupLocalAs bool `tfsdk:"-" vyos:"local-as,child"`

	ExistsTagProtocolsBgpPeerGroupLocalRole bool `tfsdk:"-" vyos:"local-role,child"`

	// Nodes

	NodeProtocolsBgpPeerGroupAddressFamily *ProtocolsBgpPeerGroupAddressFamily `tfsdk:"address_family" vyos:"address-family,omitempty"`

	NodeProtocolsBgpPeerGroupBfd *ProtocolsBgpPeerGroupBfd `tfsdk:"bfd" vyos:"bfd,omitempty"`

	NodeProtocolsBgpPeerGroupCapability *ProtocolsBgpPeerGroupCapability `tfsdk:"capability" vyos:"capability,omitempty"`

	NodeProtocolsBgpPeerGroupPathAttribute *ProtocolsBgpPeerGroupPathAttribute `tfsdk:"path_attribute" vyos:"path-attribute,omitempty"`

	NodeProtocolsBgpPeerGroupTTLSecURIty *ProtocolsBgpPeerGroupTTLSecURIty `tfsdk:"ttl_security" vyos:"ttl-security,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsBgpPeerGroup) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsBgpPeerGroup) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsBgpPeerGroup) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpPeerGroup) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"peer-group",
		o.SelfIdentifier.Attributes()["peer_group"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsBgpPeerGroup) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"protocols", // Node

		"bgp", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsBgpPeerGroup) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroup) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"peer_group": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Name of peer-group

`,
					Description: `Name of peer-group

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in peer_group, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  peer_group, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"description":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
			Description: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
		},

		"disable_capability_negotiation":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable capability negotiation with this neighbor

`,
			Description: `Disable capability negotiation with this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"disable_connected_check":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Allow peerings between eBGP peer using loopback/dummy address

`,
			Description: `Allow peerings between eBGP peer using loopback/dummy address

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ebgp_multihop":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Allow this EBGP neighbor to not be on a directly connected network

    |  Format  |  Description     |
    |----------|------------------|
    |  1-255   |  Number of hops  |
`,
			Description: `Allow this EBGP neighbor to not be on a directly connected network

    |  Format  |  Description     |
    |----------|------------------|
    |  1-255   |  Number of hops  |
`,
		},

		"graceful_restart":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP graceful restart functionality

    |  Format          |  Description                                            |
    |------------------|---------------------------------------------------------|
    |  enable          |  Enable BGP graceful restart at peer level              |
    |  disable         |  Disable BGP graceful restart at peer level             |
    |  restart-helper  |  Enable BGP graceful restart helper only functionality  |
`,
			Description: `BGP graceful restart functionality

    |  Format          |  Description                                            |
    |------------------|---------------------------------------------------------|
    |  enable          |  Enable BGP graceful restart at peer level              |
    |  disable         |  Disable BGP graceful restart at peer level             |
    |  restart-helper  |  Enable BGP graceful restart helper only functionality  |
`,
		},

		"override_capability":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Ignore capability negotiation with specified neighbor

`,
			Description: `Ignore capability negotiation with specified neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"passive":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not initiate a session with this neighbor

`,
			Description: `Do not initiate a session with this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"password":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP MD5 password

`,
			Description: `BGP MD5 password

`,
		},

		"shutdown":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administratively shutdown this neighbor

`,
			Description: `Administratively shutdown this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"update_source":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source IP of routing updates

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  ipv4    |  IPv4 address of route source  |
    |  ipv6    |  IPv6 address of route source  |
    |  txt     |  Interface as route source     |
`,
			Description: `Source IP of routing updates

    |  Format  |  Description                   |
    |----------|--------------------------------|
    |  ipv4    |  IPv4 address of route source  |
    |  ipv6    |  IPv6 address of route source  |
    |  txt     |  Interface as route source     |
`,
		},

		"remote_as":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor BGP AS number

    |  Format        |  Description                         |
    |----------------|--------------------------------------|
    |  1-4294967294  |  Neighbor AS number                  |
    |  external      |  Any AS different from the local AS  |
    |  internal      |  Neighbor AS number                  |
`,
			Description: `Neighbor BGP AS number

    |  Format        |  Description                         |
    |----------------|--------------------------------------|
    |  1-4294967294  |  Neighbor AS number                  |
    |  external      |  Any AS different from the local AS  |
    |  internal      |  Neighbor AS number                  |
`,
		},

		"port":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
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
		},

		// TagNodes

		// Nodes

		"address_family": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupAddressFamily{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Address-family parameters

`,
			Description: `Address-family parameters

`,
		},

		"bfd": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupBfd{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD) support

`,
			Description: `Enable Bidirectional Forwarding Detection (BFD) support

`,
		},

		"capability": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupCapability{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Advertise capabilities to this peer-group

`,
			Description: `Advertise capabilities to this peer-group

`,
		},

		"path_attribute": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupPathAttribute{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Manipulate path attributes from incoming UPDATE messages

`,
			Description: `Manipulate path attributes from incoming UPDATE messages

`,
		},

		"ttl_security": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpPeerGroupTTLSecURIty{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Ttl security mechanism

`,
			Description: `Ttl security mechanism

`,
		},
	}
}
