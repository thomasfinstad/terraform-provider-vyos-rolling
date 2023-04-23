package helpers

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// FromTerraformToVyos takes a resource data model and converts it to vyos api operation compatible list of lists of strings
func FromTerraformToVyos(ctx context.Context, rdm CustomResourceDataModel) [][]string {

	vyosPath, values := rdm.GetValues()

	return iron(ctx, vyosPath, values)
}

func iron(ctx context.Context, vyosPath []string, values map[string]attr.Value) [][]string {

	ret := [][]string{}

	for key, value := range values {

		key = strings.ReplaceAll(key, "_", "-")
		cVyosPath := append(vyosPath, key)

		switch value := value.(type) {
		case types.String:
			if v := value.ValueString(); v != "" {
				ret = append(
					ret,
					append(cVyosPath, v),
				)
			} else {
				tflog.Trace(ctx, "Value is empty", map[string]interface{}{"key": key})
			}

		case types.Map:
			ret = append(
				ret,
				iron(ctx, cVyosPath, value.Elements())...,
			)
		case types.Object:
			ret = append(
				ret,
				iron(ctx, cVyosPath, value.Attributes())...,
			)
		default:
			tflog.Error(ctx, "No handling of terraform value type", map[string]interface{}{"type": fmt.Sprintf("%T", value)})
		}
	}

	return ret
}
