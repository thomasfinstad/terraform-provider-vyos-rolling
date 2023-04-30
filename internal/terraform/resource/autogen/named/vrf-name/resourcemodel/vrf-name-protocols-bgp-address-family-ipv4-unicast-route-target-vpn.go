// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn struct {
	// LeafNodes
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpnBoth   customtypes.CustomStringValue `tfsdk:"both" json:"both,omitempty"`
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpnImport customtypes.CustomStringValue `tfsdk:"import" json:"import,omitempty"`
	VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpnExport customtypes.CustomStringValue `tfsdk:"export" json:"export,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourUnicastRouteTargetVpn) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"both": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Route Target both import and export

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |

`,
		},

		"import": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Route Target import

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |

`,
		},

		"export": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Route Target export

|  Format  |  Description  |
|----------|---------------|
|  txt  |  Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |

`,
		},

		// TagNodes

		// Nodes

	}
}
