package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	resourcefull "github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/autogen/named"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var _ provider.Provider = &VyosProvider{}

// VyosProvider defines the provider implementation.
type VyosProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	Version string
}

// VyosProviderModel Contains all global configurations for the provider
type VyosProviderModel struct {
	Endpoint    types.String `tfsdk:"endpoint"`
	Key         types.String `tfsdk:"api_key"`
	Certificate types.Object `tfsdk:"certificate"`
}

// Metadata method to define the provider type name for inclusion in each data source and resource type name.
func (p *VyosProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "vyos"
	resp.Version = p.Version
}

// Schema method to define the schema for provider-level configuration.
// TODO Validate cert and self signed cert
func (p *VyosProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "VyOS API Endpoint",
				Required:            true,
			},
			"api_key": schema.StringAttribute{
				MarkdownDescription: "VyOS API Key",
				Required:            true,
			},
			"certificate": schema.SingleNestedAttribute{

				Optional: true,
				Attributes: map[string]schema.Attribute{
					"disable_verify": schema.BoolAttribute{
						MarkdownDescription: "Disable remote certificate verification, useful for selfsigned certs.",
						Optional:            true,
					},
				},
			},
		},
	}
}

// Configure method to configure shared clients for data source and resource implementations.
func (p *VyosProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring vyos client")

	// TODO implement env var provider settings overrides
	// Check environment variables
	// apiToken := os.Getenv("EXAMPLECLOUD_API_TOKEN")
	// endpoint := os.Getenv("EXAMPLECLOUD_ENDPOINT")

	var providerModel VyosProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &providerModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	endpoint := providerModel.Endpoint.ValueString()
	apiKey := providerModel.Key.ValueString()

	// Configuration values validation
	if endpoint == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"VyOS Endpoint",
			"A valid http(s) endpoint is required",
		)
	}

	if apiKey == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"VyOS API Key",
			"API Key required to connect",
		)
	}

	disableVerifyAttr := providerModel.Certificate.Attributes()["disable_verify"]

	var disableVerify bool
	if disableVerifyAttr != nil {
		disableVerify = disableVerifyAttr.(types.Bool).ValueBool()
	} else {
		disableVerify = false
	}

	// Client configuration for data sources and resources
	client := client.NewClient(ctx, endpoint, apiKey, "TODO: add useragent with provider version", disableVerify)
	resp.DataSourceData = client
	resp.ResourceData = client
}

// Resources method to define the provider's resources.
func (p *VyosProvider) Resources(ctx context.Context) []func() resource.Resource {
	return resourcefull.GetResources()
}

// DataSources method to define the provider's data sources.
func (p *VyosProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		//NewDataSource,
	}
}

// New method to return the provider generator function
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &VyosProvider{
			Version: version,
		}
	}
}
