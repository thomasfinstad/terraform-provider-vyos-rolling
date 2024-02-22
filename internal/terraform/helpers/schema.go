package helpers

import (
	"context"
	"fmt"
	"slices"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GenerateVyosOps takes a resource data model and converts it to vyos api operation compatible list of lists of strings
func GenerateVyosOps(ctx context.Context, vyosPath []string, vyosData map[string]interface{}) [][]string {
	return iron(ctx, vyosPath, vyosData)
}

func iron(ctx context.Context, vyosPath []string, values map[string]interface{}) [][]string {

	ret := [][]string{}

	for key, value := range values {
		cVyosPath := append(vyosPath, key)

		switch value := value.(type) {

		// LeafNodes
		case string:
			tflog.Trace(ctx, "ironing string value", map[string]interface{}{"current-vyos-path": cVyosPath, "type": fmt.Sprintf("%T", value), "value": fmt.Sprintf("%#v", value)})
			val := append(cVyosPath, value)
			tflog.Trace(ctx, "appending to ret", map[string]interface{}{"ret": fmt.Sprintf("%#v", ret), "val": fmt.Sprintf("%#v", val)})
			ret = append(ret, val)
		// LeafNodes multi value
		case []string:
			tflog.Trace(ctx, "ironing slice of strings value", map[string]interface{}{"current-vyos-path": cVyosPath, "type": fmt.Sprintf("%T", value), "value": fmt.Sprintf("%#v", value)})
			for _, element := range value {
				val := append(cVyosPath, element)
				// Do not understand why I must clone here, but it seems the slice is a pointer, so it duplicates elements of the last one in the slice instead of adding all the uniqeue values otherwise
				val = slices.Clone(val)
				tflog.Trace(ctx, "appending to ret", map[string]interface{}{"ret": fmt.Sprintf("%#v", ret), "val": fmt.Sprintf("%#v", val)})
				ret = append(ret, val)
				tflog.Trace(ctx, "appended to ret", map[string]interface{}{"ret": fmt.Sprintf("%#v", ret), "val": fmt.Sprintf("%#v", val)})
			}

		// TagNodes and Nodes
		case map[string]interface{}:
			tflog.Trace(ctx, "ironing nested value", map[string]interface{}{"current-vyos-path": cVyosPath, "type": fmt.Sprintf("%T", value), "value": fmt.Sprintf("%#v", value)})
			tflog.Trace(ctx, "recursing for ret", map[string]interface{}{"cVyosPath": fmt.Sprintf("%#v", cVyosPath)})
			ret = append(
				ret,
				iron(ctx, cVyosPath, value)...,
			)

		// ERROR
		default:
			tflog.Error(ctx, "No handling of value type", map[string]interface{}{"type": fmt.Sprintf("%T", value), "key": key, "cVyosPath": cVyosPath})
			panic("unhandled type see last log entry")
		}
	}

	tflog.Trace(ctx, "ironing result", map[string]interface{}{"result": fmt.Sprintf("%#v", ret)})
	return ret
}
