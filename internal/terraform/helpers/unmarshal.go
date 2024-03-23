package helpers

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers/tools"
)

// UnmarshalVyos takes unmarshalled json data and a Terraform resource model pointer and populates it with the data
func UnmarshalVyos(ctx context.Context, data map[string]any, value VyosResourceDataModel) error {
	valueReflection := reflect.ValueOf(value).Elem()
	typeReflection := valueReflection.Type()

	for i := 0; i < valueReflection.NumField(); i++ {
		fName := typeReflection.Field(i).Name
		fType := typeReflection.Field(i).Type
		fValue := valueReflection.Field(i)
		fTags := strings.Split(typeReflection.Field(i).Tag.Get("vyos"), ",")
		tools.Debug(ctx, "processing field", map[string]interface{}{"Field": fName, "Type": fType, "Tags": fTags})

		// Set flags based on tags, first tag must be the vyos field name, the rest are bools with default of false
		// TODO create struct of valid options
		flags := map[string]any{
			"name":      fTags[0],
			"self-id":   false,
			"parent-id": false,
			"omitempty": false,
			"child":     false,
			"timeout":   false,
		}
		for _, tag := range fTags[1:] {
			flags[tag] = true
		}

		if flags["self-id"].(bool) || flags["parent-id"].(bool) || flags["timeout"].(bool) {

			tools.Debug(ctx, "Not configuring field", map[string]interface{}{"field-name": fName})
			continue
		}

		if flags["child"].(bool) {

			if !tools.KeyInMap(flags["name"].(string), data) {

				tfValueRefection := reflect.ValueOf(false)
				fValue.Set(tfValueRefection)
				continue
			}

			if data == nil {

				tfValueRefection := reflect.ValueOf(false)
				fValue.Set(tfValueRefection)
				continue
			}

			tfValueRefection := reflect.ValueOf(true)
			fValue.Set(tfValueRefection)
			continue
		}

		switch fValue.Interface().(type) {
		case basetypes.StringValue:
			var tfval basetypes.StringValuable

			if !tools.KeyInMap(flags["name"].(string), data) {

				tools.Debug(ctx, "No data for field, setting to empty string", map[string]interface{}{"field-name": fName})

				tfval = basetypes.NewStringNull()
			} else {
				dataValue := data[flags["name"].(string)]

				tools.Debug(ctx, "Unmarshalling String Field", map[string]interface{}{"field-name": fName, flags["name"].(string): dataValue})

				tfval = basetypes.NewStringValue(dataValue.(string))
			}
			tfValueRefection := reflect.ValueOf(tfval)
			fValue.Set(tfValueRefection)

		case basetypes.BoolValue:

			tools.Debug(ctx, "Unmarshalling Bool Field", map[string]interface{}{"field-name": fName})
			tfValue := basetypes.NewBoolValue(tools.KeyInMap(flags["name"].(string), data))
			tfValueRefection := reflect.ValueOf(tfValue)
			fValue.Set(tfValueRefection)
		case basetypes.NumberValue:
			var tfval basetypes.NumberValuable

			if !tools.KeyInMap(flags["name"].(string), data) {

				tools.Debug(ctx, "No data for field, skipping", map[string]interface{}{"field-name": fName})

				tfval = basetypes.NewNumberNull()
			} else {
				dataValue := data[flags["name"].(string)]

				tools.Debug(ctx, "Unmarshalling Number Field", map[string]interface{}{"field-name": fName, flags["name"].(string): dataValue})

				var dataValueF float64
				switch dataValue := dataValue.(type) {
				case string:
					var err error
					dataValueF, err = strconv.ParseFloat(dataValue, 64)
					if err != nil {
						tools.Error(ctx, "cannot convert string to float64", map[string]interface{}{"value": dataValue})
						return fmt.Errorf("cannot convert string '%#v' to float64", dataValue)
					}
				case int:
					dataValueF = float64(dataValue)
				case int8:
					dataValueF = float64(dataValue)
				case int16:
					dataValueF = float64(dataValue)
				case int32:
					dataValueF = float64(dataValue)
				case int64:
					dataValueF = float64(dataValue)
				case uint:
					dataValueF = float64(dataValue)
				case uint8:
					dataValueF = float64(dataValue)
				case uint16:
					dataValueF = float64(dataValue)
				case uint32:
					dataValueF = float64(dataValue)
				case uint64:
					dataValueF = float64(dataValue)
				case float32:
					dataValueF = float64(dataValue)
				case float64:
					dataValueF = float64(dataValue)
				default:
					tools.Error(ctx, "cannot convert UNHANDLED to float64", map[string]interface{}{"value": dataValue})
					return fmt.Errorf("cannot convert UNHANDLED '%#v' to float64", dataValue)
				}

				bf := big.NewFloat(dataValueF)
				tfval = basetypes.NewNumberValue(bf)
			}
			tfValueRefection := reflect.ValueOf(tfval)
			fValue.Set(tfValueRefection)
			tools.Trace(ctx, "Setting Value", map[string]interface{}{"tfval": tfval, "tfValueRefection": tfValueRefection, "fValue": fValue})

		case basetypes.ListValue:
			var tfval basetypes.ListValuable
			tfSchemaName := strings.Split(typeReflection.Field(i).Tag.Get("tfsdk"), ",")[0]
			schemaAttr := value.ResourceSchemaAttributes(ctx)[tfSchemaName]
			listAttr := schemaAttr.(schema.ListAttribute)
			elemType := listAttr.ElementType

			if !tools.KeyInMap(flags["name"].(string), data) {

				tools.Debug(ctx, "No data for field, skipping", map[string]interface{}{"field-name": fName})

				tfval = basetypes.NewListNull(elemType)
			} else {
				dataValue := data[flags["name"].(string)]

				tools.Debug(ctx, "Unmarshalling List Field", map[string]interface{}{"field-name": fName, flags["name"].(string): dataValue})

				if reflect.TypeOf(dataValue).Kind() != reflect.Slice {

					tools.Debug(ctx, "dataValue is not a slice, wrapping the value in list", map[string]interface{}{flags["name"].(string): dataValue})
					dataValue = []interface{}{dataValue}
				}

				var diags diag.Diagnostics
				tfval, diags = basetypes.NewListValueFrom(ctx, elemType, dataValue)
				if diags != nil {
					return fmt.Errorf("ERROR: %v", diags)
				}
			}
			tools.Trace(ctx, "setting list value", map[string]interface{}{"tfval": tfval})
			tfValueRefection := reflect.ValueOf(tfval)
			fValue.Set(tfValueRefection)

		default:
			if !tools.KeyInMap(flags["name"].(string), data) {

				tools.Debug(ctx, "No data for field, skipping", map[string]interface{}{"field-name": fName})
				continue
			}
			dataValue := data[flags["name"].(string)]

			tools.Debug(ctx, "Unmarshalling Struct Field", map[string]interface{}{"field-name": fName, flags["name"].(string): dataValue})

			subSctructReflection := reflect.New(fType.Elem())
			subStructIf := subSctructReflection.Interface()
			subStruct := subStructIf.(VyosResourceDataModel)
			err := UnmarshalVyos(ctx, dataValue.(map[string]any), subStruct)
			if err != nil {
				return err
			}

			fValue.Set(subSctructReflection)
		}
	}

	// Return no error
	return nil
}
