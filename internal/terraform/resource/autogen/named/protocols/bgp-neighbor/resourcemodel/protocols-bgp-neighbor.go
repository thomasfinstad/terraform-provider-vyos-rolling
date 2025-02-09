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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (neighbor) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsBgpNeighbor{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (neighbor) */
// ProtocolsBgpNeighborSelfIdentifier used by TagNodes to keep the id info
type ProtocolsBgpNeighborSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (neighbor) */

	ProtocolsBgpNeighbor types.String `tfsdk:"neighbor" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (bgp) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (protocols) */
}

// ProtocolsBgpNeighbor describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type ProtocolsBgpNeighbor struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *ProtocolsBgpNeighborSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafProtocolsBgpNeighborAdvertisementInterval        types.Number `tfsdk:"advertisement_interval" vyos:"advertisement-interval,omitempty"`
	LeafProtocolsBgpNeighborDescrIPtion                  types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafProtocolsBgpNeighborDisableCapabilityNegotiation types.Bool   `tfsdk:"disable_capability_negotiation" vyos:"disable-capability-negotiation,omitempty"`
	LeafProtocolsBgpNeighborDisableConnectedCheck        types.Bool   `tfsdk:"disable_connected_check" vyos:"disable-connected-check,omitempty"`
	LeafProtocolsBgpNeighborEbgpMultihop                 types.Number `tfsdk:"ebgp_multihop" vyos:"ebgp-multihop,omitempty"`
	LeafProtocolsBgpNeighborGracefulRestart              types.String `tfsdk:"graceful_restart" vyos:"graceful-restart,omitempty"`
	LeafProtocolsBgpNeighborOverrIDeCapability           types.Bool   `tfsdk:"override_capability" vyos:"override-capability,omitempty"`
	LeafProtocolsBgpNeighborPassive                      types.Bool   `tfsdk:"passive" vyos:"passive,omitempty"`
	LeafProtocolsBgpNeighborPassword                     types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafProtocolsBgpNeighborPeerGroup                    types.String `tfsdk:"peer_group" vyos:"peer-group,omitempty"`
	LeafProtocolsBgpNeighborRemoteAs                     types.String `tfsdk:"remote_as" vyos:"remote-as,omitempty"`
	LeafProtocolsBgpNeighborShutdown                     types.Bool   `tfsdk:"shutdown" vyos:"shutdown,omitempty"`
	LeafProtocolsBgpNeighborSolo                         types.Bool   `tfsdk:"solo" vyos:"solo,omitempty"`
	LeafProtocolsBgpNeighborEnforceFirstAs               types.Bool   `tfsdk:"enforce_first_as" vyos:"enforce-first-as,omitempty"`
	LeafProtocolsBgpNeighborStrictCapabilityMatch        types.Bool   `tfsdk:"strict_capability_match" vyos:"strict-capability-match,omitempty"`
	LeafProtocolsBgpNeighborUpdateSource                 types.String `tfsdk:"update_source" vyos:"update-source,omitempty"`
	LeafProtocolsBgpNeighborPort                         types.Number `tfsdk:"port" vyos:"port,omitempty"`

	// TagNodes

	ExistsTagProtocolsBgpNeighborLocalAs bool `tfsdk:"-" vyos:"local-as,child"`

	ExistsTagProtocolsBgpNeighborLocalRole bool `tfsdk:"-" vyos:"local-role,child"`

	// Nodes

	NodeProtocolsBgpNeighborAddressFamily *ProtocolsBgpNeighborAddressFamily `tfsdk:"address_family" vyos:"address-family,omitempty"`

	NodeProtocolsBgpNeighborBfd *ProtocolsBgpNeighborBfd `tfsdk:"bfd" vyos:"bfd,omitempty"`

	NodeProtocolsBgpNeighborCapability *ProtocolsBgpNeighborCapability `tfsdk:"capability" vyos:"capability,omitempty"`

	NodeProtocolsBgpNeighborInterface *ProtocolsBgpNeighborInterface `tfsdk:"interface" vyos:"interface,omitempty"`

	NodeProtocolsBgpNeighborPathAttribute *ProtocolsBgpNeighborPathAttribute `tfsdk:"path_attribute" vyos:"path-attribute,omitempty"`

	NodeProtocolsBgpNeighborTimers *ProtocolsBgpNeighborTimers `tfsdk:"timers" vyos:"timers,omitempty"`

	NodeProtocolsBgpNeighborTTLSecURIty *ProtocolsBgpNeighborTTLSecURIty `tfsdk:"ttl_security" vyos:"ttl-security,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsBgpNeighbor) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsBgpNeighbor) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsBgpNeighbor) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpNeighbor) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"neighbor",
		o.SelfIdentifier.ProtocolsBgpNeighbor.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsBgpNeighbor) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (bgp) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (protocols) */
		"protocols", // Node

		"bgp", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsBgpNeighbor) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (bgp) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (protocols) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpNeighbor) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"neighbor": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `BGP neighbor

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  ipv4    |  BGP neighbor IP address    |
    |  ipv6    |  BGP neighbor IPv6 address  |
    |  txt     |  Interface name             |
`,
					Description: `BGP neighbor

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  ipv4    |  BGP neighbor IP address    |
    |  ipv6    |  BGP neighbor IPv6 address  |
    |  txt     |  Interface name             |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in neighbor, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  neighbor, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (bgp) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (protocols) */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"advertisement_interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (advertisement-interval) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum interval for sending routing updates

    |  Format  |  Description                        |
    |----------|-------------------------------------|
    |  0-600   |  Advertisement interval in seconds  |
`,
			Description: `Minimum interval for sending routing updates

    |  Format  |  Description                        |
    |----------|-------------------------------------|
    |  0-600   |  Advertisement interval in seconds  |
`,
		},

		"description":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (description) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (disable-capability-negotiation) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (disable-connected-check) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ebgp-multihop) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (graceful-restart) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (override-capability) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (passive) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (password) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `BGP MD5 password

`,
			Description: `BGP MD5 password

`,
		},

		"peer_group":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (peer-group) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Peer group for this peer

    |  Format  |  Description      |
    |----------|-------------------|
    |  txt     |  Peer-group name  |
`,
			Description: `Peer group for this peer

    |  Format  |  Description      |
    |----------|-------------------|
    |  txt     |  Peer-group name  |
`,
		},

		"remote_as":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (remote-as) */
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

		"shutdown":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (shutdown) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administratively shutdown this neighbor

`,
			Description: `Administratively shutdown this neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"solo":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (solo) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not send back prefixes learned from the neighbor

`,
			Description: `Do not send back prefixes learned from the neighbor

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"enforce_first_as":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (enforce-first-as) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Ensure the first AS in the AS path matches the peer AS

`,
			Description: `Ensure the first AS in the AS path matches the peer AS

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"strict_capability_match":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (strict-capability-match) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable strict capability negotiation

`,
			Description: `Enable strict capability negotiation

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"update_source":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (update-source) */
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

		"port":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (port) */
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
			Attributes: ProtocolsBgpNeighborAddressFamily{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Address-family parameters

`,
			Description: `Address-family parameters

`,
		},

		"bfd": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborBfd{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD) support

`,
			Description: `Enable Bidirectional Forwarding Detection (BFD) support

`,
		},

		"capability": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborCapability{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Advertise capabilities to this peer-group

`,
			Description: `Advertise capabilities to this peer-group

`,
		},

		"interface": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborInterface{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Interface parameters

`,
			Description: `Interface parameters

`,
		},

		"path_attribute": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborPathAttribute{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Manipulate path attributes from incoming UPDATE messages

`,
			Description: `Manipulate path attributes from incoming UPDATE messages

`,
		},

		"timers": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborTimers{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Neighbor timers

`,
			Description: `Neighbor timers

`,
		},

		"ttl_security": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpNeighborTTLSecURIty{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Ttl security mechanism

`,
			Description: `Ttl security mechanism

`,
		},
	}
}
