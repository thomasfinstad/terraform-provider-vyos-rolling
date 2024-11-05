package tests

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflogtest"
	protocolsBgpNeighbor "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/protocols/bgp-neighbor/resourcemodel"
	serviceDhcpServerSharedNetworkName "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/service/dhcp-server-shared-network-name/resourcemodel"
)

// TestIpInNamedResourceIdentifier allow IPv6 in resource identifiers
//
// Regression test for bug: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/issues/222
func TestIpInNamedResourceIdentifier(t *testing.T) {
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	validators := protocolsBgpNeighbor.ProtocolsBgpNeighbor{}.ResourceSchemaAttributes(ctx)["identifier"].(schema.SingleNestedAttribute).Attributes["neighbor"].(schema.StringAttribute).Validators

	testValues := []string{"10.1.1.1", "2001:db8::1:0:0:1"}

	for _, v := range testValues {

		resp := make([]validator.StringResponse, len(validators))
		for i := range validators {
			validators[i].ValidateString(ctx,
				validator.StringRequest{
					Path:        path.Empty(),
					ConfigValue: basetypes.NewStringValue(v),
				},
				&resp[i],
			)
		}

		for _, r := range resp {
			for _, d := range r.Diagnostics {
				t.Logf("Validation failure: %s", d.Detail())
				t.Fail()
			}
		}
	}
}

// TestNetworkInNamedResourceIdentifier allow IPv6 in resource identifiers
// Regression test for bug: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/issues/224
func TestNetworkInNamedResourceIdentifier(t *testing.T) {
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	validators := serviceDhcpServerSharedNetworkName.ServiceDhcpServerSharedNetworkName{}.ResourceSchemaAttributes(ctx)["identifier"].(schema.SingleNestedAttribute).Attributes["shared_network_name"].(schema.StringAttribute).Validators

	testValues := []string{"192.168.128.0/24", "10.0.0.0/8"}

	for _, v := range testValues {

		resp := make([]validator.StringResponse, len(validators))
		for i := range validators {
			validators[i].ValidateString(ctx,
				validator.StringRequest{
					Path:        path.Empty(),
					ConfigValue: basetypes.NewStringValue(v),
				},
				&resp[i],
			)
		}

		for _, r := range resp {
			for _, d := range r.Diagnostics {
				t.Logf("Validation failure: %s", d.Detail())
				t.Fail()
			}
		}
	}
}
