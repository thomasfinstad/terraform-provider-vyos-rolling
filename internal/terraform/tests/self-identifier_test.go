package tests

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	modelIfaceEth "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/interfaces/ethernet/resourcemodel"
)

/*
resource "vyos_interfaces_ethernet" "trunk" {
	identifier = {
		ethernet = local.default_eth
	}
	description = "Main trunk with all vlans"
	duplex = "auto"
	speed = "auto"
	hw_id = "06:00:00:00:00:${local.vyos_instance_nr}"
}
*/

/*
Stack trace from the terraform-provider-vyos-rolling_v13.0.202411271 plugin:

panic: interface conversion: attr.Value is nil, not basetypes.StringValue

goroutine 197 [running]:
github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/interfaces/ethernet/resourcemodel.(*InterfacesEthernet).GetVyosPath(0xc0006c0820)
        github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/interfaces/ethernet/resourcemodel/interfaces-ethernet.go:106 +0x29f
github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers/crud.read({0x1b5b2a8, 0xc000c59ea0}, {{{0x1b48fc0, 0xc00042a300}, 0x0, {0x0, 0x0}, 0x0}, {0x18ffc8a, 0x29}, ...}, ...)
        github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers/crud/read.go:67 +0x105
github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers/crud.Read({0x1b5b200, 0xc000635a10}, {0x1b55138, 0xc00003c3c0}, {{{{0x1b83ff8, 0xc0012f07e0}, {0x15bfae0, 0xc0012f01b0}}, {0x1b85cc8, 0xc0000194a0}}, ...}, ...)
        github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers/crud/read.go:44 +0x4d8
github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/interfaces/ethernet.(*interfacesEthernet).Read(0x0?, {0x1b5b200?, 0xc000635a10?}, {{{{0x1b83ff8, 0xc0012f07e0}, {0x15bfae0, 0xc0012f01b0}}, {0x1b85cc8, 0xc0000194a0}}, 0xc0004680d8, ...}, ...)
        github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/resource/autogen/named/interfaces/ethernet/crud.go:22 +0x6d
github.com/hashicorp/terraform-plugin-framework/internal/fwserver.(*Server).ReadResource(0xc00017d080, {0x1b5b200, 0xc000635a10}, 0xc000635b30, 0xc0005cb648)
        github.com/hashicorp/terraform-plugin-framework@v1.3.5/internal/fwserver/server_readresource.go:101 +0x62e
github.com/hashicorp/terraform-plugin-framework/internal/proto6server.(*Server).ReadResource(0xc00017d080, {0x1b5b200?, 0xc000635860?}, 0xc000406d00)
        github.com/hashicorp/terraform-plugin-framework@v1.3.5/internal/proto6server/server_readresource.go:55 +0x38e
github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server.(*server).ReadResource(0xc0001ec820, {0x1b5b200?, 0xc000634810?}, 0xc00168ea20)
        github.com/hashicorp/terraform-plugin-go@v0.18.0/tfprotov6/tf6server/server.go:749 +0x46f
github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6._Provider_ReadResource_Handler({0x18829a0, 0xc0001ec820}, {0x1b5b200, 0xc000634810}, 0xc000a976c0, 0x0)
        github.com/hashicorp/terraform-plugin-go@v0.18.0/tfprotov6/internal/tfplugin6/tfplugin6_grpc.pb.go:386 +0x1a6
google.golang.org/grpc.(*Server).processUnaryRPC(0xc0001e61e0, {0x1b84aa0, 0xc0004be000}, 0xc00037cfc0, 0xc0003261e0, 0x2a15e70, 0x0)
        google.golang.org/grpc@v1.56.1/server.go:1337 +0xdb8
google.golang.org/grpc.(*Server).handleStream(0xc0001e61e0, {0x1b84aa0, 0xc0004be000}, 0xc00037cfc0, 0x0)
        google.golang.org/grpc@v1.56.1/server.go:1714 +0x9da
google.golang.org/grpc.(*Server).serveStreams.func1.1()
        google.golang.org/grpc@v1.56.1/server.go:959 +0x87
created by google.golang.org/grpc.(*Server).serveStreams.func1 in goroutine 11
        google.golang.org/grpc@v1.56.1/server.go:957 +0x154

Error: The terraform-provider-vyos-rolling_v13.0.202411271 plugin crashed!
*/

// TestSelfIdentifier tests that that an empty resource marshals to the correct empty representation
func TestSelfIdentifier(t *testing.T) {

	want := []string{"interfaces", "ethernet", "dummy123"}

	model := &modelIfaceEth.InterfacesEthernet{
		SelfIdentifier: &modelIfaceEth.InterfacesEthernetSelfIdentifier{
			InterfacesEthernet: basetypes.NewStringValue("dummy123"),
		},
	}

	got := model.GetVyosPath()

	diff := deep.Equal(got, want)
	if diff != nil {
		t.Errorf("compare failed: %v", diff)
	}
}
