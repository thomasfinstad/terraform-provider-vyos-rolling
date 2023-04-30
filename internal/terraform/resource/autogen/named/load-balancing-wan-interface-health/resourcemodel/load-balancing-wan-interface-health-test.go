// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// LoadBalancingWanInterfaceHealthTest describes the resource data model.
type LoadBalancingWanInterfaceHealthTest struct {
	// LeafNodes
	LoadBalancingWanInterfaceHealthTestRespTime   customtypes.CustomStringValue `tfsdk:"resp_time" json:"resp-time,omitempty"`
	LoadBalancingWanInterfaceHealthTestTarget     customtypes.CustomStringValue `tfsdk:"target" json:"target,omitempty"`
	LoadBalancingWanInterfaceHealthTestTestScrIPt customtypes.CustomStringValue `tfsdk:"test_script" json:"test-script,omitempty"`
	LoadBalancingWanInterfaceHealthTestTTLLimit   customtypes.CustomStringValue `tfsdk:"ttl_limit" json:"ttl-limit,omitempty"`
	LoadBalancingWanInterfaceHealthTestType       customtypes.CustomStringValue `tfsdk:"type" json:"type,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o LoadBalancingWanInterfaceHealthTest) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"resp_time": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Ping response time (seconds)

|  Format  |  Description  |
|----------|---------------|
|  u32:1-30  |  Response time (seconds)  |

`,
		},

		"target": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Health target address

|  Format  |  Description  |
|----------|---------------|
|  ipv4  |  Health target address  |

`,
		},

		"test_script": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Path to user-defined script

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Script in /config/scripts  |

`,
		},

		"ttl_limit": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `TTL limit (hop count)

|  Format  |  Description  |
|----------|---------------|
|  u32:1-254  |  Number of hops  |

`,
		},

		"type": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `WLB test type

|  Format  |  Description  |
|----------|---------------|
|  ping  |  Test with ICMP echo response  |
|  ttl  |  Test with UDP TTL expired response  |
|  user-defined  |  User-defined test script  |

`,
		},

		// TagNodes

		// Nodes

	}
}
