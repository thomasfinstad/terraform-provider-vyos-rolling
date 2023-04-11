package resourcefull

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &vpn_ipsec_esp_group{}

// var _ resource.ResourceWithImportState = &vpn_ipsec_esp_group{}

// vpn_ipsec_esp_group defines the resource implementation.
type vpn_ipsec_esp_group struct {
	client *http.Client
}

// vpn ipsec esp-groupModel describes the resource data model.
type vpn_ipsec_esp_groupModel struct {
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
	ID                    types.String `tfsdk:"id"`

	Proposal types.Map `tfsdk:proposal`

	Compression types.String `tfsdk:compression`

	Lifetime types.String `tfsdk:lifetime`

	Life_bytes types.String `tfsdk:life_bytes`

	Life_packets types.String `tfsdk:life_packets`

	Mode types.String `tfsdk:mode`

	Pfs types.String `tfsdk:pfs`
}

// Metadata method to define the resource type name.
func (r *vpn_ipsec_esp_group) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn ipsec esp-group"
}

// vpn_ipsec_esp_groupResource method to return the example resource reference
func vpn_ipsec_esp_groupResource() resource.Resource {
	return &vpn_ipsec_esp_group{}
}

// Schema method to define the schema for any resource configuration, plan, and state data.
func (r *vpn_ipsec_esp_group) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"proposal": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						// TODO handle non-string types
						"encryption": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Encryption algorithm

Format: null
Null encryption
Format: aes128
128 bit AES-CBC
Format: aes192
192 bit AES-CBC
Format: aes256
256 bit AES-CBC
Format: aes128ctr
128 bit AES-COUNTER
Format: aes192ctr
192 bit AES-COUNTER
Format: aes256ctr
256 bit AES-COUNTER
Format: aes128ccm64
128 bit AES-CCM with 64 bit ICV
Format: aes192ccm64
192 bit AES-CCM with 64 bit ICV
Format: aes256ccm64
256 bit AES-CCM with 64 bit ICV
Format: aes128ccm96
128 bit AES-CCM with 96 bit ICV
Format: aes192ccm96
192 bit AES-CCM with 96 bit ICV
Format: aes256ccm96
256 bit AES-CCM with 96 bit ICV
Format: aes128ccm128
128 bit AES-CCM with 128 bit ICV
Format: aes192ccm128
192 bit AES-CCM with 128 bit IC
Format: aes256ccm128
256 bit AES-CCM with 128 bit ICV
Format: aes128gcm64
128 bit AES-GCM with 64 bit ICV
Format: aes192gcm64
192 bit AES-GCM with 64 bit ICV
Format: aes256gcm64
256 bit AES-GCM with 64 bit ICV
Format: aes128gcm96
128 bit AES-GCM with 96 bit ICV
Format: aes192gcm96
192 bit AES-GCM with 96 bit ICV
Format: aes256gcm96
256 bit AES-GCM with 96 bit ICV
Format: aes128gcm128
128 bit AES-GCM with 128 bit ICV
Format: aes192gcm128
192 bit AES-GCM with 128 bit ICV
Format: aes256gcm128
256 bit AES-GCM with 128 bit ICV
Format: aes128gmac
Null encryption with 128 bit AES-GMAC
Format: aes192gmac
Null encryption with 192 bit AES-GMAC
Format: aes256gmac
Null encryption with 256 bit AES-GMAC
Format: 3des
168 bit 3DES-EDE-CBC
Format: blowfish128
128 bit Blowfish-CBC
Format: blowfish192
192 bit Blowfish-CBC
Format: blowfish256
256 bit Blowfish-CBC
Format: camellia128
128 bit Camellia-CBC
Format: camellia192
192 bit Camellia-CBC
Format: camellia256
256 bit Camellia-CBC
Format: camellia128ctr
128 bit Camellia-COUNTER
Format: camellia192ctr
192 bit Camellia-COUNTER
Format: camellia256ctr
256 bit Camellia-COUNTER
Format: camellia128ccm64
128 bit Camellia-CCM with 64 bit ICV
Format: camellia192ccm64
192 bit Camellia-CCM with 64 bit ICV
Format: camellia256ccm64
256 bit Camellia-CCM with 64 bit ICV
Format: camellia128ccm96
128 bit Camellia-CCM with 96 bit ICV
Format: camellia192ccm96
192 bit Camellia-CCM with 96 bit ICV
Format: camellia256ccm96
256 bit Camellia-CCM with 96 bit ICV
Format: camellia128ccm128
128 bit Camellia-CCM with 128 bit ICV
Format: camellia192ccm128
192 bit Camellia-CCM with 128 bit ICV
Format: camellia256ccm128
256 bit Camellia-CCM with 128 bit ICV
Format: serpent128
128 bit Serpent-CBC
Format: serpent192
192 bit Serpent-CBC
Format: serpent256
256 bit Serpent-CBC
Format: twofish128
128 bit Twofish-CBC
Format: twofish192
192 bit Twofish-CBC
Format: twofish256
256 bit Twofish-CBC
Format: cast128
128 bit CAST-CBC
Format: chacha20poly1305
256 bit ChaCha20/Poly1305 with 128 bit ICV
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

							Default:  stringdefault.StaticString(`aes128`),
							Computed: true,
						},

						// TODO handle non-string types
						"hash": schema.StringAttribute{
							// CustomType:          basetypes.StringTypable(nil),
							// Required:            false,
							Optional: true,
							// Sensitive:           false,
							// Description:         "",
							MarkdownDescription: `Hash algorithm

Format: md5
MD5 HMAC
Format: md5_128
MD5_128 HMAC
Format: sha1
SHA1 HMAC
Format: sha1_160
SHA1_160 HMAC
Format: sha256
SHA2_256_128 HMAC
Format: sha256_96
SHA2_256_96 HMAC
Format: sha384
SHA2_384_192 HMAC
Format: sha512
SHA2_512_256 HMAC
Format: aesxcbc
AES XCBC
Format: aescmac
AES CMAC
Format: aes128gmac
128-bit AES-GMAC
Format: aes192gmac
192-bit AES-GMAC
Format: aes256gmac
256-bit AES-GMAC
`,
							// DeprecationMessage:  "",
							// TODO Recreate some of vyos validators for use in leafnodes
							// Validators:          []validator.String(nil),
							// PlanModifiers:       []planmodifier.String(nil),

							Default:  stringdefault.StaticString(`sha1`),
							Computed: true,
						},

						// CustomType:    basetypes.ObjectTypable(nil),
						// Validators:    []validator.Object(nil),
						// PlanModifiers: []planmodifier.Object(nil),
					},
				},
				// CustomType:          basetypes.MapTypable(nil),
				// Required:            false,
				Optional: true,
				// Computed:            false,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `ESP group proposal

Format: u32:1-65535
ESP group proposal number
`,
				// DeprecationMessage:  "",
				// Validators:          []validator.Map(nil),
				// PlanModifiers:       []planmodifier.Map(nil),
				// TODO investigate if tagnode defaults can be handled
				// Default:             defaults.Map(nil),
			},

			// TODO handle non-string types
			"compression": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Enable ESP compression

`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"lifetime": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Security Association time to expire

Format: u32:30-86400
SA lifetime in seconds
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

				Default:  stringdefault.StaticString(`3600`),
				Computed: true,
			},

			// TODO handle non-string types
			"life_bytes": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Security Association byte count to expire

Format: u32:1024-26843545600000
SA life in bytes
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"life_packets": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `Security Association packet count to expire

Format: u32:1000-26843545600000
SA life in packets
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

			},

			// TODO handle non-string types
			"mode": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `ESP mode

Format: tunnel
Tunnel mode
Format: transport
Transport mode
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

				Default:  stringdefault.StaticString(`tunnel`),
				Computed: true,
			},

			// TODO handle non-string types
			"pfs": schema.StringAttribute{
				// CustomType:          basetypes.StringTypable(nil),
				// Required:            false,
				Optional: true,
				// Sensitive:           false,
				// Description:         "",
				MarkdownDescription: `ESP Perfect Forward Secrecy

Format: enable
Inherit Diffie-Hellman group from the IKE group
Format: dh-group1
Use Diffie-Hellman group 1 (modp768)
Format: dh-group2
Use Diffie-Hellman group 2 (modp1024)
Format: dh-group5
Use Diffie-Hellman group 5 (modp1536)
Format: dh-group14
Use Diffie-Hellman group 14 (modp2048)
Format: dh-group15
Use Diffie-Hellman group 15 (modp3072)
Format: dh-group16
Use Diffie-Hellman group 16 (modp4096)
Format: dh-group17
Use Diffie-Hellman group 17 (modp6144)
Format: dh-group18
Use Diffie-Hellman group 18 (modp8192)
Format: dh-group19
Use Diffie-Hellman group 19 (ecp256)
Format: dh-group20
Use Diffie-Hellman group 20 (ecp384)
Format: dh-group21
Use Diffie-Hellman group 21 (ecp521)
Format: dh-group22
Use Diffie-Hellman group 22 (modp1024s160)
Format: dh-group23
Use Diffie-Hellman group 23 (modp2048s224)
Format: dh-group24
Use Diffie-Hellman group 24 (modp2048s256)
Format: dh-group25
Use Diffie-Hellman group 25 (ecp192)
Format: dh-group26
Use Diffie-Hellman group 26 (ecp224)
Format: dh-group27
Use Diffie-Hellman group 27 (ecp224bp)
Format: dh-group28
Use Diffie-Hellman group 28 (ecp256bp)
Format: dh-group29
Use Diffie-Hellman group 29 (ecp384bp)
Format: dh-group30
Use Diffie-Hellman group 30 (ecp512bp)
Format: dh-group31
Use Diffie-Hellman group 31 (curve25519)
Format: dh-group32
Use Diffie-Hellman group 32 (curve448)
Format: disable
Disable PFS
`,
				// DeprecationMessage:  "",
				// TODO Recreate some of vyos validators for use in leafnodes
				// Validators:          []validator.String(nil),
				// PlanModifiers:       []planmodifier.String(nil),

				Default:  stringdefault.StaticString(`enable`),
				Computed: true,
			},
		},
	}
}

// Create method to define the logic which creates the resource and sets its initial Terraform state.
func (r *vpn_ipsec_esp_group) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *vpn_ipsec_esp_groupModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create example, got error: %s", err))
	//     return
	// }

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	data.ID = types.StringValue("example-id")

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read method to define the logic which refreshes the Terraform state for the resource.
func (r *vpn_ipsec_esp_group) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *vpn_ipsec_esp_groupModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update method to define the logic which updates the resource and sets the updated Terraform state on success.
func (r *vpn_ipsec_esp_group) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *vpn_ipsec_esp_groupModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete method to define the logic which deletes the resource and removes the Terraform state on success.
func (r *vpn_ipsec_esp_group) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *vpn_ipsec_esp_groupModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete example, got error: %s", err))
	//     return
	// }
}
