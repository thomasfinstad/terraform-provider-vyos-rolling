package helpers

import (
	"context"
	"fmt"
	"slices"
	"strconv"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers/tools"
	"golang.org/x/exp/maps"
)

// GenerateVyosOps takes a resource data model and converts it to vyos api operation compatible list of lists of strings
//
// Input:
//
//	vyosPath: []string{"firewall", "ipv4", "forward", "filter"}
//	vyosData: map[string]interface {}{"default-action":"reject", "default-log":map[string]interface {}{}}
//
// Output:
//
//	[][]string{[]string{"firewall", "ipv4", "forward", "filter", "default-log"}, []string{"firewall", "ipv4", "forward", "filter", "default-action", "reject"}}
func GenerateVyosOps(ctx context.Context, vyosPath []string, vyosData map[string]interface{}) [][]string {
	tools.Trace(ctx, "GenerateVyosOps Input", map[string]interface{}{"vyosPath": vyosPath, "vyosData": vyosData})
	tools.Debug(ctx, "GenerateVyosOps Input", map[string]interface{}{"vyosPath": vyosPath, "vyosData": vyosData})
	tools.Info(ctx, "GenerateVyosOps Input", map[string]interface{}{"vyosPath": vyosPath, "vyosData": vyosData})
	tools.Warn(ctx, "GenerateVyosOps Input", map[string]interface{}{"vyosPath": vyosPath, "vyosData": vyosData})
	tools.Error(ctx, "GenerateVyosOps Input", map[string]interface{}{"vyosPath": vyosPath, "vyosData": vyosData})
	vyosOps := iron(ctx, vyosPath, vyosData)
	tools.Trace(ctx, "GenerateVyosOps return", map[string]interface{}{"vyosOps": vyosOps})
	return vyosOps
}

// iron flatens out the input and adds the vyosPath as a prefix to each element
func iron(ctx context.Context, vyosPath []string, values map[string]interface{}) [][]string {

	ret := [][]string{}

	// Make operation deterministic, helps during debugging
	keys := maps.Keys(values)
	slices.Sort(keys)
	slices.Reverse(keys)
	for _, key := range keys {

		// This Clone has proven itself necessary due underlyding mangling happening otherwise
		cVyosPath := append(slices.Clone(vyosPath), key)

		value := values[key]

		switch value := value.(type) {

		// LeafNodes
		case string:

			tools.Trace(ctx, "ironing string value", map[string]interface{}{"current-vyos-path": cVyosPath, "type": fmt.Sprintf("%T", value), "value": fmt.Sprintf("%#v", value)})
			val := append(cVyosPath, value)
			tools.Trace(ctx, "appending to ret", map[string]interface{}{"ret": fmt.Sprintf("%#v", ret), "val": fmt.Sprintf("%#v", val)})
			ret = append(ret, val)
		// LeafNodes
		case bool:

			tools.Trace(ctx, "ironing bool value", map[string]interface{}{"current-vyos-path": cVyosPath, "type": fmt.Sprintf("%T", value), "value": fmt.Sprintf("%#v", value)})
			if value {
				val := append(cVyosPath, "{}")
				tools.Trace(ctx, "appending to ret", map[string]interface{}{"ret": fmt.Sprintf("%#v", ret), "val": fmt.Sprintf("%#v", val)})
				ret = append(ret, val)
			} else {
				tools.Trace(ctx, "NOT appending to ret because bool is false", map[string]interface{}{"ret": fmt.Sprintf("%#v", ret)})
			}
		// LeafNodes
		case float64:

			tools.Trace(ctx, "ironing float value", map[string]interface{}{"current-vyos-path": cVyosPath, "type": fmt.Sprintf("%T", value), "value": fmt.Sprintf("%#v", value)})
			val := append(cVyosPath, strconv.FormatFloat(value, 'f', -1, 64))
			tools.Trace(ctx, "appending to ret", map[string]interface{}{"ret": fmt.Sprintf("%#v", ret), "val": fmt.Sprintf("%#v", val)})
			ret = append(ret, val)
		// LeafNodes multi value
		case []string:

			tools.Trace(ctx, "ironing slice of strings value", map[string]interface{}{"current-vyos-path": cVyosPath, "type": fmt.Sprintf("%T", value), "value": fmt.Sprintf("%#v", value)})
			for _, element := range value {
				val := append(cVyosPath, element)
				tools.Trace(ctx, "appending to ret", map[string]interface{}{"ret": fmt.Sprintf("%#v", ret), "val": fmt.Sprintf("%#v", val)})
				ret = append(ret, val)
			}

		// TagNodes and Nodes
		case map[string]interface{}:

			tools.Trace(ctx, "ironing nested map value", map[string]interface{}{"current-vyos-path": cVyosPath, "type": fmt.Sprintf("%T", value), "value": fmt.Sprintf("%#v", value)})
			tools.Trace(ctx, "recursing for ret", map[string]interface{}{"cVyosPath": fmt.Sprintf("%#v", cVyosPath)})
			val := iron(ctx, cVyosPath, value)
			ret = append(
				ret,
				val...,
			)

		// ERROR
		default:
			tools.Error(ctx, "No handling of value type", map[string]interface{}{"type": fmt.Sprintf("%T", value), "key": key, "cVyosPath": cVyosPath})
			panic("unhandled type see last log entry")
		}

	}

	if len(ret) == 0 {

		tools.Trace(ctx, "ironing empty value", map[string]interface{}{"vyosPath": vyosPath})
		ret = [][]string{vyosPath}
	}

	tools.Trace(ctx, "ironing result", map[string]interface{}{"result": fmt.Sprintf("%#v", ret)})

	return ret
}
