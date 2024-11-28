package helpers

import (
	"context"
	"fmt"
	"reflect"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers/tools"
)

// UnknownToNull
// Takes in a Resource Data Model pointer and
// recursively sets all unknown fields to null
func UnknownToNull(ctx context.Context, data VyosResourceDataModel) {
	dataValueElem := reflect.ValueOf(data).Elem()
	dataTypeElem := reflect.TypeOf(data).Elem()
	tools.Trace(ctx, "checking for fields with unknown values", map[string]any{"value": data})
	for idx := range dataValueElem.NumField() {
		fieldValue := dataValueElem.Field(idx)
		fieldName := dataTypeElem.Field(idx).Name
		tools.Trace(ctx, "checking if field implements attr.Value interface", map[string]any{"field": fieldName})
		if dataTypeElem.Field(idx).Tag.Get("tfsdk") == "-" {
			continue
		} else if (slices.Contains([]reflect.Kind{reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice}, fieldValue.Kind())) && fieldValue.IsNil() {
			continue
		} else if fieldValue.Kind() == reflect.Map {
			iter := fieldValue.MapRange()
			for iter.Next() {
				// k := iter.Key()
				v := iter.Value()
				UnknownToNull(ctx, v.Interface().(VyosResourceDataModel))
			}
		} else if fieldValue.Type().Implements(reflect.TypeOf((*VyosResourceDataModel)(nil)).Elem()) {
			UnknownToNull(ctx, fieldValue.Interface().(VyosResourceDataModel))
		} else if fieldValue.Type().Implements(reflect.TypeOf((*attr.Value)(nil)).Elem()) {
			valueIface := fieldValue.Interface().(attr.Value)
			tools.Trace(ctx, "checking if field value is unknown", map[string]any{"field": fieldName, "value": valueIface})
			if valueIface.IsUnknown() {
				tools.Debug(ctx, "field value is unknown", map[string]any{"field": fieldName})
				switch valueIface.(type) {
				case basetypes.BoolValue:
					fieldValue.Set(reflect.ValueOf(basetypes.NewBoolNull()))
				case basetypes.Float64Value:
					fieldValue.Set(reflect.ValueOf(basetypes.NewFloat64Null()))
				case basetypes.Int64Value:
					fieldValue.Set(reflect.ValueOf(basetypes.NewInt64Null()))
				case basetypes.ListValue:
					attrType := valueIface.Type(ctx)
					attrValue := attrType.ValueType(ctx)
					fieldValue.Set(reflect.ValueOf(attrValue.(basetypes.ListValue)))
				case basetypes.MapValue:
					tools.Error(ctx, "invalid code path, maps should be structs", map[string]interface{}{"value": valueIface})
					panic("invalid code path, maps should be structs")
				case basetypes.NumberValue:
					fieldValue.Set(reflect.ValueOf(basetypes.NewNumberNull()))
				case basetypes.ObjectValue:
					tools.Error(ctx, "invalid code path, objects should be structs", map[string]interface{}{"value": valueIface})
					panic("invalid code path, objects should be structs")
				case basetypes.SetValue:
					tools.Error(ctx, "invalid code path, sets should be structs", map[string]interface{}{"value": valueIface})
					panic("invalid code path, sets should be structs")
				case basetypes.StringValue:
					fieldValue.Set(reflect.ValueOf(basetypes.NewStringNull()))
				default:
					tools.Error(ctx, "invalid code path, unhandled type", map[string]interface{}{"value": valueIface})
					panic(fmt.Sprintf("invalid code path, unhandled type: %t", valueIface))
				}
			}
		} else {
			tools.Error(ctx, "invalid code path, unhandled interface", map[string]interface{}{"interfaceType": dataTypeElem.Field(idx).Name, "interfaceKind": dataTypeElem.Field(idx).Type.Kind()})
			panic(fmt.Sprintf("invalid code path, unhandled interface: %s of kind %s", dataTypeElem.Field(idx).Name, dataTypeElem.Field(idx).Type.Kind()))
		}
	}
}
