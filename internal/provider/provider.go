package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/supabase/cli/pkg/api"
)

// Ensure SupabaseProvider satisfies various provider interfaces.
var _ provider.Provider = &SupabaseProvider{}

// SupabaseProvider defines the provider implementation.
type SupabaseProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// SupabaseProviderModel describes the provider data model.
type SupabaseProviderModel struct {
	AccessToken types.String `tfsdk:"access_token"`
	Endpoint    types.String `tfsdk:"endpoint"`
}

func (p *SupabaseProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "supabase"
	resp.Version = p.version
}

func (p *SupabaseProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"access_token": schema.StringAttribute{
				MarkdownDescription: "Management API access token.",
				Optional:            true,
			},
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Management API endpoint.",
				Optional:            true,
			},
		},
	}
}

func (p *SupabaseProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data SupabaseProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var access_token = os.Getenv("SUPABASE_TOKEN")
	if access_token == "" {
		access_token = data.AccessToken.ValueString()
	}

	var endpoint = "api.supabase.com"
	if data.Endpoint.IsNull() {
		endpoint = data.Endpoint.ValueString()
	}

	// Example client configuration for data sources and resources
	client, _ := api.NewClientWithResponses(endpoint)
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *SupabaseProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewExampleResource,
	}
}

func (p *SupabaseProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewExampleDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &SupabaseProvider{
			version: version,
		}
	}
}
