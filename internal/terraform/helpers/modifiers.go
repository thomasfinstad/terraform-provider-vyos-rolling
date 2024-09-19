package helpers

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// UnknownToNull
// Takes in a Resource Data Model pointer and
// recursively sets all unknown fields to null
func UnknownToNull(ctx context.Context, m VyosResourceDataModel) {
	t := reflect.ValueOf(m).Elem()
	for idx := range t.NumField() {
		f := t.Field(idx)
		if f.Type().Implements(reflect.TypeOf((*attr.Value)(nil)).Elem()) {
			v := f.Interface().(attr.Value)
			if v.IsUnknown() {
				switch v.(type) {
				case basetypes.BoolValue:
					f.Set(reflect.ValueOf(basetypes.NewBoolNull()))
				case basetypes.Float64Value:
					f.Set(reflect.ValueOf(basetypes.NewFloat64Null()))
				case basetypes.Int64Value:
					f.Set(reflect.ValueOf(basetypes.NewInt64Null()))
				case basetypes.ListValue:
					// f.Set(reflect.ValueOf(basetypes.NewListNull()))
					panic("invalid code path, lists should be structs")
				case basetypes.MapValue:
					// f.Set(reflect.ValueOf(basetypes.NewMapNull()))
					panic("invalid code path, maps should be structs")
				case basetypes.NumberValue:
					f.Set(reflect.ValueOf(basetypes.NewNumberNull()))
				case basetypes.ObjectValue:
					panic("invalid code path, objects should be structs")
					// f.Set(reflect.ValueOf(basetypes.NewObjectNull()))
				case basetypes.SetValue:
					// f.Set(reflect.ValueOf(basetypes.NewSetNull()))
					panic("invalid code path, sets should be structs")
				case basetypes.StringValue:
					f.Set(reflect.ValueOf(basetypes.NewStringNull()))
				default:
					panic(fmt.Sprintf("invalid code path, unhandled type: %t", v))
				}
			}
		} else if f.Type().Implements(reflect.TypeOf((*VyosResourceDataModel)(nil)).Elem()) && !f.IsNil() {
			UnknownToNull(ctx, f.Interface().(VyosResourceDataModel))
		}
	}
}
