package helpers

import (
	"context"
	"fmt"

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

		tflog.Error(ctx, "Value type", map[string]interface{}{"type": fmt.Sprintf("%T", value)})

		switch value := value.(type) {

		// LeafNodes
		case string:
			ret = append(
				ret,
				append(cVyosPath, value),
			)
		// TagNodes and Nodes
		case map[string]interface{}:
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

	return ret
}
