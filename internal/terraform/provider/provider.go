package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	resourcefull "github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/resource/full"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var _ provider.Provider = &VyosProvider{}

// VyosProvider defines the provider implementation.
type VyosProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// VyosProviderModel describes the provider data model.
type VyosProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	Key      types.String `tfsdk:"api_key"`
}

// Metadata method to define the provider type name for inclusion in each data source and resource type name.
func (p *VyosProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "vyos"
	resp.Version = p.version
}

// Schema method to define the schema for provider-level configuration.
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
		},
	}
}

// Configure method to configure shared clients for data source and resource implementations.
func (p *VyosProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring vyos client")

	var data VyosProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources
	client := http.DefaultClient
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
		//NewExampleDataSource,
	}
}

// New method to return the provider generator function
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &VyosProvider{
			version: version,
		}
	}
}
