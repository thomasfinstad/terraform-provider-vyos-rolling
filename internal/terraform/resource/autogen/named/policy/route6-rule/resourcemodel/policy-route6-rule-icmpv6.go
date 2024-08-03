// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &PolicyRoutesixRuleIcmpvsix{}

// PolicyRoutesixRuleIcmpvsix describes the resource data model.
type PolicyRoutesixRuleIcmpvsix struct {
	// LeafNodes
	LeafPolicyRoutesixRuleIcmpvsixType types.String `tfsdk:"type" vyos:"type,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyRoutesixRuleIcmpvsix) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `ICMP type-name

    |  Format                      |  Description          |
    |------------------------------|-----------------------|
    |  any                         |  Any ICMP type/code   |
    |  echo-reply                  |  ICMP type/code name  |
    |  pong                        |  ICMP type/code name  |
    |  destination-unreachable     |  ICMP type/code name  |
    |  network-unreachable         |  ICMP type/code name  |
    |  host-unreachable            |  ICMP type/code name  |
    |  protocol-unreachable        |  ICMP type/code name  |
    |  port-unreachable            |  ICMP type/code name  |
    |  fragmentation-needed        |  ICMP type/code name  |
    |  source-route-failed         |  ICMP type/code name  |
    |  network-unknown             |  ICMP type/code name  |
    |  host-unknown                |  ICMP type/code name  |
    |  network-prohibited          |  ICMP type/code name  |
    |  host-prohibited             |  ICMP type/code name  |
    |  TOS-network-unreachable     |  ICMP type/code name  |
    |  TOS-host-unreachable        |  ICMP type/code name  |
    |  communication-prohibited    |  ICMP type/code name  |
    |  host-precedence-violation   |  ICMP type/code name  |
    |  precedence-cutoff           |  ICMP type/code name  |
    |  source-quench               |  ICMP type/code name  |
    |  redirect                    |  ICMP type/code name  |
    |  network-redirect            |  ICMP type/code name  |
    |  host-redirect               |  ICMP type/code name  |
    |  TOS-network-redirect        |  ICMP type/code name  |
    |  TOS host-redirect           |  ICMP type/code name  |
    |  echo-request                |  ICMP type/code name  |
    |  ping                        |  ICMP type/code name  |
    |  router-advertisement        |  ICMP type/code name  |
    |  router-solicitation         |  ICMP type/code name  |
    |  time-exceeded               |  ICMP type/code name  |
    |  ttl-exceeded                |  ICMP type/code name  |
    |  ttl-zero-during-transit     |  ICMP type/code name  |
    |  ttl-zero-during-reassembly  |  ICMP type/code name  |
    |  parameter-problem           |  ICMP type/code name  |
    |  ip-header-bad               |  ICMP type/code name  |
    |  required-option-missing     |  ICMP type/code name  |
    |  timestamp-request           |  ICMP type/code name  |
    |  timestamp-reply             |  ICMP type/code name  |
    |  address-mask-request        |  ICMP type/code name  |
    |  address-mask-reply          |  ICMP type/code name  |
    |  packet-too-big              |  ICMP type/code name  |
`,
			Description: `ICMP type-name

    |  Format                      |  Description          |
    |------------------------------|-----------------------|
    |  any                         |  Any ICMP type/code   |
    |  echo-reply                  |  ICMP type/code name  |
    |  pong                        |  ICMP type/code name  |
    |  destination-unreachable     |  ICMP type/code name  |
    |  network-unreachable         |  ICMP type/code name  |
    |  host-unreachable            |  ICMP type/code name  |
    |  protocol-unreachable        |  ICMP type/code name  |
    |  port-unreachable            |  ICMP type/code name  |
    |  fragmentation-needed        |  ICMP type/code name  |
    |  source-route-failed         |  ICMP type/code name  |
    |  network-unknown             |  ICMP type/code name  |
    |  host-unknown                |  ICMP type/code name  |
    |  network-prohibited          |  ICMP type/code name  |
    |  host-prohibited             |  ICMP type/code name  |
    |  TOS-network-unreachable     |  ICMP type/code name  |
    |  TOS-host-unreachable        |  ICMP type/code name  |
    |  communication-prohibited    |  ICMP type/code name  |
    |  host-precedence-violation   |  ICMP type/code name  |
    |  precedence-cutoff           |  ICMP type/code name  |
    |  source-quench               |  ICMP type/code name  |
    |  redirect                    |  ICMP type/code name  |
    |  network-redirect            |  ICMP type/code name  |
    |  host-redirect               |  ICMP type/code name  |
    |  TOS-network-redirect        |  ICMP type/code name  |
    |  TOS host-redirect           |  ICMP type/code name  |
    |  echo-request                |  ICMP type/code name  |
    |  ping                        |  ICMP type/code name  |
    |  router-advertisement        |  ICMP type/code name  |
    |  router-solicitation         |  ICMP type/code name  |
    |  time-exceeded               |  ICMP type/code name  |
    |  ttl-exceeded                |  ICMP type/code name  |
    |  ttl-zero-during-transit     |  ICMP type/code name  |
    |  ttl-zero-during-reassembly  |  ICMP type/code name  |
    |  parameter-problem           |  ICMP type/code name  |
    |  ip-header-bad               |  ICMP type/code name  |
    |  required-option-missing     |  ICMP type/code name  |
    |  timestamp-request           |  ICMP type/code name  |
    |  timestamp-reply             |  ICMP type/code name  |
    |  address-mask-request        |  ICMP type/code name  |
    |  address-mask-reply          |  ICMP type/code name  |
    |  packet-too-big              |  ICMP type/code name  |
`,
		},

		// Nodes

	}
}
