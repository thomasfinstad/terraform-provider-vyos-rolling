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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (proposal) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VpnIPsecEspGroupProposal{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (proposal) */
// VpnIPsecEspGroupProposalSelfIdentifier used by TagNodes to keep the id info
type VpnIPsecEspGroupProposalSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (proposal) */

	VpnIPsecEspGroupProposal types.Number `tfsdk:"proposal" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (esp-group) */

	VpnIPsecEspGroup types.String `tfsdk:"esp_group" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (ipsec) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (vpn) */
}

// VpnIPsecEspGroupProposal describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type VpnIPsecEspGroupProposal struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *VpnIPsecEspGroupProposalSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafVpnIPsecEspGroupProposalEncryption types.String `tfsdk:"encryption" vyos:"encryption,omitempty"`
	LeafVpnIPsecEspGroupProposalHash       types.String `tfsdk:"hash" vyos:"hash,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *VpnIPsecEspGroupProposal) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnIPsecEspGroupProposal) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnIPsecEspGroupProposal) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnIPsecEspGroupProposal) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"proposal",
		o.SelfIdentifier.VpnIPsecEspGroupProposal.ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnIPsecEspGroupProposal) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (esp-group) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (ipsec) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (vpn) */
		"vpn", // Node

		"ipsec", // Node

		"esp-group",
		o.SelfIdentifier.VpnIPsecEspGroup.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VpnIPsecEspGroupProposal) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (esp-group) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (esp-group) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (ipsec) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (vpn) */
		"vpn", // Node

		"ipsec", // Node

		"esp-group",
		o.SelfIdentifier.VpnIPsecEspGroup.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecEspGroupProposal) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"proposal": schema.NumberAttribute{
					Required: true,
					MarkdownDescription: `ESP group proposal

    |  Format   |  Description                |
    |-----------|-----------------------------|
    |  1-65535  |  ESP group proposal number  |
`,
					Description: `ESP group proposal

    |  Format   |  Description                |
    |-----------|-----------------------------|
    |  1-65535  |  ESP group proposal number  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (esp-group) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (ipsec) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (vpn) */

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
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  esp_group, value must match: ^[.:a-zA-Z0-9-_/]+$",
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

		"encryption":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (encryption) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Encryption algorithm

    |  Format             |  Description                                 |
    |---------------------|----------------------------------------------|
    |  null               |  Null encryption                             |
    |  aes128             |  128 bit AES-CBC                             |
    |  aes192             |  192 bit AES-CBC                             |
    |  aes256             |  256 bit AES-CBC                             |
    |  aes128ctr          |  128 bit AES-COUNTER                         |
    |  aes192ctr          |  192 bit AES-COUNTER                         |
    |  aes256ctr          |  256 bit AES-COUNTER                         |
    |  aes128ccm64        |  128 bit AES-CCM with 64 bit ICV             |
    |  aes192ccm64        |  192 bit AES-CCM with 64 bit ICV             |
    |  aes256ccm64        |  256 bit AES-CCM with 64 bit ICV             |
    |  aes128ccm96        |  128 bit AES-CCM with 96 bit ICV             |
    |  aes192ccm96        |  192 bit AES-CCM with 96 bit ICV             |
    |  aes256ccm96        |  256 bit AES-CCM with 96 bit ICV             |
    |  aes128ccm128       |  128 bit AES-CCM with 128 bit ICV            |
    |  aes192ccm128       |  192 bit AES-CCM with 128 bit IC             |
    |  aes256ccm128       |  256 bit AES-CCM with 128 bit ICV            |
    |  aes128gcm64        |  128 bit AES-GCM with 64 bit ICV             |
    |  aes192gcm64        |  192 bit AES-GCM with 64 bit ICV             |
    |  aes256gcm64        |  256 bit AES-GCM with 64 bit ICV             |
    |  aes128gcm96        |  128 bit AES-GCM with 96 bit ICV             |
    |  aes192gcm96        |  192 bit AES-GCM with 96 bit ICV             |
    |  aes256gcm96        |  256 bit AES-GCM with 96 bit ICV             |
    |  aes128gcm128       |  128 bit AES-GCM with 128 bit ICV            |
    |  aes192gcm128       |  192 bit AES-GCM with 128 bit ICV            |
    |  aes256gcm128       |  256 bit AES-GCM with 128 bit ICV            |
    |  aes128gmac         |  Null encryption with 128 bit AES-GMAC       |
    |  aes192gmac         |  Null encryption with 192 bit AES-GMAC       |
    |  aes256gmac         |  Null encryption with 256 bit AES-GMAC       |
    |  3des               |  168 bit 3DES-EDE-CBC                        |
    |  blowfish128        |  128 bit Blowfish-CBC                        |
    |  blowfish192        |  192 bit Blowfish-CBC                        |
    |  blowfish256        |  256 bit Blowfish-CBC                        |
    |  camellia128        |  128 bit Camellia-CBC                        |
    |  camellia192        |  192 bit Camellia-CBC                        |
    |  camellia256        |  256 bit Camellia-CBC                        |
    |  camellia128ctr     |  128 bit Camellia-COUNTER                    |
    |  camellia192ctr     |  192 bit Camellia-COUNTER                    |
    |  camellia256ctr     |  256 bit Camellia-COUNTER                    |
    |  camellia128ccm64   |  128 bit Camellia-CCM with 64 bit ICV        |
    |  camellia192ccm64   |  192 bit Camellia-CCM with 64 bit ICV        |
    |  camellia256ccm64   |  256 bit Camellia-CCM with 64 bit ICV        |
    |  camellia128ccm96   |  128 bit Camellia-CCM with 96 bit ICV        |
    |  camellia192ccm96   |  192 bit Camellia-CCM with 96 bit ICV        |
    |  camellia256ccm96   |  256 bit Camellia-CCM with 96 bit ICV        |
    |  camellia128ccm128  |  128 bit Camellia-CCM with 128 bit ICV       |
    |  camellia192ccm128  |  192 bit Camellia-CCM with 128 bit ICV       |
    |  camellia256ccm128  |  256 bit Camellia-CCM with 128 bit ICV       |
    |  serpent128         |  128 bit Serpent-CBC                         |
    |  serpent192         |  192 bit Serpent-CBC                         |
    |  serpent256         |  256 bit Serpent-CBC                         |
    |  twofish128         |  128 bit Twofish-CBC                         |
    |  twofish192         |  192 bit Twofish-CBC                         |
    |  twofish256         |  256 bit Twofish-CBC                         |
    |  cast128            |  128 bit CAST-CBC                            |
    |  chacha20poly1305   |  256 bit ChaCha20/Poly1305 with 128 bit ICV  |
`,
			Description: `Encryption algorithm

    |  Format             |  Description                                 |
    |---------------------|----------------------------------------------|
    |  null               |  Null encryption                             |
    |  aes128             |  128 bit AES-CBC                             |
    |  aes192             |  192 bit AES-CBC                             |
    |  aes256             |  256 bit AES-CBC                             |
    |  aes128ctr          |  128 bit AES-COUNTER                         |
    |  aes192ctr          |  192 bit AES-COUNTER                         |
    |  aes256ctr          |  256 bit AES-COUNTER                         |
    |  aes128ccm64        |  128 bit AES-CCM with 64 bit ICV             |
    |  aes192ccm64        |  192 bit AES-CCM with 64 bit ICV             |
    |  aes256ccm64        |  256 bit AES-CCM with 64 bit ICV             |
    |  aes128ccm96        |  128 bit AES-CCM with 96 bit ICV             |
    |  aes192ccm96        |  192 bit AES-CCM with 96 bit ICV             |
    |  aes256ccm96        |  256 bit AES-CCM with 96 bit ICV             |
    |  aes128ccm128       |  128 bit AES-CCM with 128 bit ICV            |
    |  aes192ccm128       |  192 bit AES-CCM with 128 bit IC             |
    |  aes256ccm128       |  256 bit AES-CCM with 128 bit ICV            |
    |  aes128gcm64        |  128 bit AES-GCM with 64 bit ICV             |
    |  aes192gcm64        |  192 bit AES-GCM with 64 bit ICV             |
    |  aes256gcm64        |  256 bit AES-GCM with 64 bit ICV             |
    |  aes128gcm96        |  128 bit AES-GCM with 96 bit ICV             |
    |  aes192gcm96        |  192 bit AES-GCM with 96 bit ICV             |
    |  aes256gcm96        |  256 bit AES-GCM with 96 bit ICV             |
    |  aes128gcm128       |  128 bit AES-GCM with 128 bit ICV            |
    |  aes192gcm128       |  192 bit AES-GCM with 128 bit ICV            |
    |  aes256gcm128       |  256 bit AES-GCM with 128 bit ICV            |
    |  aes128gmac         |  Null encryption with 128 bit AES-GMAC       |
    |  aes192gmac         |  Null encryption with 192 bit AES-GMAC       |
    |  aes256gmac         |  Null encryption with 256 bit AES-GMAC       |
    |  3des               |  168 bit 3DES-EDE-CBC                        |
    |  blowfish128        |  128 bit Blowfish-CBC                        |
    |  blowfish192        |  192 bit Blowfish-CBC                        |
    |  blowfish256        |  256 bit Blowfish-CBC                        |
    |  camellia128        |  128 bit Camellia-CBC                        |
    |  camellia192        |  192 bit Camellia-CBC                        |
    |  camellia256        |  256 bit Camellia-CBC                        |
    |  camellia128ctr     |  128 bit Camellia-COUNTER                    |
    |  camellia192ctr     |  192 bit Camellia-COUNTER                    |
    |  camellia256ctr     |  256 bit Camellia-COUNTER                    |
    |  camellia128ccm64   |  128 bit Camellia-CCM with 64 bit ICV        |
    |  camellia192ccm64   |  192 bit Camellia-CCM with 64 bit ICV        |
    |  camellia256ccm64   |  256 bit Camellia-CCM with 64 bit ICV        |
    |  camellia128ccm96   |  128 bit Camellia-CCM with 96 bit ICV        |
    |  camellia192ccm96   |  192 bit Camellia-CCM with 96 bit ICV        |
    |  camellia256ccm96   |  256 bit Camellia-CCM with 96 bit ICV        |
    |  camellia128ccm128  |  128 bit Camellia-CCM with 128 bit ICV       |
    |  camellia192ccm128  |  192 bit Camellia-CCM with 128 bit ICV       |
    |  camellia256ccm128  |  256 bit Camellia-CCM with 128 bit ICV       |
    |  serpent128         |  128 bit Serpent-CBC                         |
    |  serpent192         |  192 bit Serpent-CBC                         |
    |  serpent256         |  256 bit Serpent-CBC                         |
    |  twofish128         |  128 bit Twofish-CBC                         |
    |  twofish192         |  192 bit Twofish-CBC                         |
    |  twofish256         |  256 bit Twofish-CBC                         |
    |  cast128            |  128 bit CAST-CBC                            |
    |  chacha20poly1305   |  256 bit ChaCha20/Poly1305 with 128 bit ICV  |
`,

			// Default:          stringdefault.StaticString(`aes128`),
			Computed: true,
		},

		"hash":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (hash) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Hash algorithm

    |  Format      |  Description        |
    |--------------|---------------------|
    |  md5         |  MD5 HMAC           |
    |  md5_128     |  MD5_128 HMAC       |
    |  sha1        |  SHA1 HMAC          |
    |  sha1_160    |  SHA1_160 HMAC      |
    |  sha256      |  SHA2_256_128 HMAC  |
    |  sha256_96   |  SHA2_256_96 HMAC   |
    |  sha384      |  SHA2_384_192 HMAC  |
    |  sha512      |  SHA2_512_256 HMAC  |
    |  aesxcbc     |  AES XCBC           |
    |  aescmac     |  AES CMAC           |
    |  aes128gmac  |  128-bit AES-GMAC   |
    |  aes192gmac  |  192-bit AES-GMAC   |
    |  aes256gmac  |  256-bit AES-GMAC   |
`,
			Description: `Hash algorithm

    |  Format      |  Description        |
    |--------------|---------------------|
    |  md5         |  MD5 HMAC           |
    |  md5_128     |  MD5_128 HMAC       |
    |  sha1        |  SHA1 HMAC          |
    |  sha1_160    |  SHA1_160 HMAC      |
    |  sha256      |  SHA2_256_128 HMAC  |
    |  sha256_96   |  SHA2_256_96 HMAC   |
    |  sha384      |  SHA2_384_192 HMAC  |
    |  sha512      |  SHA2_512_256 HMAC  |
    |  aesxcbc     |  AES XCBC           |
    |  aescmac     |  AES CMAC           |
    |  aes128gmac  |  128-bit AES-GMAC   |
    |  aes192gmac  |  192-bit AES-GMAC   |
    |  aes256gmac  |  256-bit AES-GMAC   |
`,

			// Default:          stringdefault.StaticString(`sha1`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
