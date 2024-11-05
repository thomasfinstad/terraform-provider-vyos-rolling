package tests

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	ntpResourceModel "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/global/service/ntp/resourcemodel"
)

// TestTfstateToNamedResourceMergedIntoGlobal
// Verify that what would have been a named resource that has to be merged
// in under a global resource results in a map of config parameters and works as expected
func TestTfstateToNamedResourceMergedIntoGlobal(t *testing.T) {
	ctx := context.Background()

	resp := &resource.CreateResponse{
		State: tfsdk.State{
			Raw: tftypes.NewValue(
				tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"allow_client": tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"address": tftypes.List{
									ElementType: tftypes.String,
								},
							},
						},
						"id":          tftypes.String,
						"interface":   tftypes.String,
						"leap_second": tftypes.String,
						"listen_address": tftypes.List{
							ElementType: tftypes.String,
						},
						"server": tftypes.Map{
							ElementType: tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"interleave": tftypes.Bool,
									"noselect":   tftypes.Bool,
									"nts":        tftypes.Bool,
									"pool":       tftypes.Bool,
									"prefer":     tftypes.Bool,
									"ptp":        tftypes.Bool,
								},
							},
						},
						"timeouts": tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"create": tftypes.String,
							},
						},
						"vrf": tftypes.String,
					},
				},
				map[string]tftypes.Value{
					"allow_client": tftypes.NewValue(
						tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"address": tftypes.List{
									ElementType: tftypes.String,
								},
							},
						},
						map[string]tftypes.Value{
							"address": tftypes.NewValue(
								tftypes.List{
									ElementType: tftypes.String,
								},
								[]tftypes.Value{
									tftypes.NewValue(tftypes.String, "1.2.3.4"),
								},
							),
						},
					),
					"id":             tftypes.NewValue(tftypes.String, nil),
					"interface":      tftypes.NewValue(tftypes.String, nil),
					"leap_second":    tftypes.NewValue(tftypes.String, nil),
					"listen_address": tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, nil),
					"server": tftypes.NewValue(
						tftypes.Map{
							ElementType: tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"interleave": tftypes.Bool,
									"noselect":   tftypes.Bool,
									"nts":        tftypes.Bool,
									"pool":       tftypes.Bool,
									"prefer":     tftypes.Bool,
									"ptp":        tftypes.Bool,
								},
							},
						},
						map[string]tftypes.Value{
							"no.pool.ntp.org": tftypes.NewValue(
								tftypes.Object{
									AttributeTypes: map[string]tftypes.Type{
										"interleave": tftypes.Bool,
										"noselect":   tftypes.Bool,
										"nts":        tftypes.Bool,
										"pool":       tftypes.Bool,
										"prefer":     tftypes.Bool,
										"ptp":        tftypes.Bool,
									},
								},
								map[string]tftypes.Value{
									"interleave": tftypes.NewValue(tftypes.Bool, false),
									"noselect":   tftypes.NewValue(tftypes.Bool, false),
									"nts":        tftypes.NewValue(tftypes.Bool, false),
									"pool":       tftypes.NewValue(tftypes.Bool, false),
									"prefer":     tftypes.NewValue(tftypes.Bool, true),
									"ptp":        tftypes.NewValue(tftypes.Bool, false),
								},
							),
						},
					),
					"timeouts": tftypes.NewValue(
						tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"create": tftypes.String,
							},
						},
						nil),
					"vrf": tftypes.NewValue(tftypes.String, nil),
				},
			),
			Schema: resourceschema.Schema{
				Attributes: ntpResourceModel.ServiceNtp{}.ResourceSchemaAttributes(ctx),
			},
		},
	}

	m := ntpResourceModel.ServiceNtp{}

	diags := resp.State.Get(ctx, &m)

	if diags.HasError() {
		t.Errorf("Failed with diags: %v", diags)
	}
}
