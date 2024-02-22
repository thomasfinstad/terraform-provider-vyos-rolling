package helpers

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
		tflog.Debug(ctx, "processing field", map[string]interface{}{"Field": fName, "Type": fType, "Tags": fTags})
		fmt.Printf("Field: %s\tType: %s\tTags: %v\n", fName, fType, fTags)

		// Set flags based on tags, first tag must be the vyos field name, the rest are bools with default of false
		// TODO create struct of valid options
		flags := map[string]any{
			"name":      fTags[0],
			"self-id":   false,
			"parent-id": false,
			"omitempty": false,
			"child":     false,
		}
		for _, tag := range fTags[1:] {
			flags[tag] = true
		}
		if flags["self-id"].(bool) || flags["parent-id"].(bool) {
			fmt.Printf("\tNot configuring field: %s\n", fName)
			tflog.Debug(ctx, "Not configuring field", map[string]interface{}{"field-name": fName})
			continue
		}

		if flags["child"].(bool) {
			// TODO if field is a child field/resource, check if it is populated and set bool to true if it is so delete function can choose correct action to take
			continue
		}

		switch fValue.Interface().(type) {
		case basetypes.StringValue:
			var tfval basetypes.StringValuable

			if !KeyInMap(flags["name"].(string), data) {
				fmt.Printf("\tNo data for field: %s, setting to empty string\n", fName)
				tflog.Debug(ctx, "No data for field, setting to empty string", map[string]interface{}{"field-name": fName})

				tfval = basetypes.NewStringNull()
			} else {
				dataValue := data[flags["name"].(string)]

				fmt.Printf("\tUnmarshalling String Field: %s\t%s=%s\n", fName, flags["name"].(string), dataValue)
				tflog.Debug(ctx, "Unmarshalling String Field", map[string]interface{}{"field-name": fName, flags["name"].(string): dataValue})

				tfval = basetypes.NewStringValue(dataValue.(string))
			}
			tfValueRefection := reflect.ValueOf(tfval)
			fValue.Set(tfValueRefection)

		case basetypes.BoolValue:
			fmt.Printf("\tUnmarshalling Bool Field: %s\n", fName)
			tflog.Debug(ctx, "Unmarshalling Bool Field", map[string]interface{}{"field-name": fName})
			tfValue := basetypes.NewBoolValue(KeyInMap(flags["name"].(string), data))
			tfValueRefection := reflect.ValueOf(tfValue)
			fValue.Set(tfValueRefection)
		case basetypes.NumberValue:
			var tfval basetypes.NumberValuable

			if !KeyInMap(flags["name"].(string), data) {
				fmt.Printf("\tNo data for field: %s, skipping\n", fName)
				tflog.Debug(ctx, "No data for field, skipping", map[string]interface{}{"field-name": fName})

				tfval = basetypes.NewNumberNull()
			} else {
				dataValue := data[flags["name"].(string)]

				fmt.Printf("\tUnmarshalling Number Field: %s\t%s=%s\n", fName, flags["name"].(string), dataValue)
				tflog.Debug(ctx, "Unmarshalling Number Field", map[string]interface{}{"field-name": fName, flags["name"].(string): dataValue})

				// Use reflection to convert anything (possible) to float, powered by https://stackoverflow.com/a/50008579
				var floatTypeReflection = reflect.TypeOf(float64(0))
				dataValueReflection := reflect.ValueOf(dataValue)
				dataValueReflection = reflect.Indirect(dataValueReflection)
				if !dataValueReflection.Type().ConvertibleTo(floatTypeReflection) {
					return fmt.Errorf("cannot convert %v to float64", dataValueReflection.Type())
				}
				dataValueFloat := dataValueReflection.Convert(floatTypeReflection).Float()

				bf := big.NewFloat(dataValueFloat)
				tfval = basetypes.NewNumberValue(bf)
			}
			tfValueRefection := reflect.ValueOf(tfval)
			fValue.Set(tfValueRefection)

		case basetypes.ListValue:
			var tfval basetypes.ListValuable
			tfSchemaName := strings.Split(typeReflection.Field(i).Tag.Get("tfsdk"), ",")[0]
			schemaAttr := value.ResourceSchemaAttributes()[tfSchemaName]
			listAttr := schemaAttr.(schema.ListAttribute)
			elemType := listAttr.ElementType

			if !KeyInMap(flags["name"].(string), data) {
				fmt.Printf("\tNo data for field: %s, skipping\n", fName)
				tflog.Debug(ctx, "No data for field, skipping", map[string]interface{}{"field-name": fName})

				tfval = basetypes.NewListNull(elemType)
			} else {
				dataValue := data[flags["name"].(string)]

				fmt.Printf("\tUnmarshalling List Field: %s\t%s=%s\n", fName, flags["name"].(string), dataValue)
				tflog.Debug(ctx, "Unmarshalling List Field", map[string]interface{}{"field-name": fName, flags["name"].(string): dataValue})

				if reflect.TypeOf(dataValue).Kind() != reflect.Slice {
					fmt.Printf("\tdataValue is not a slice, wrapping the value in list: %#v\n", dataValue)
					tflog.Debug(ctx, "dataValue is not a slice, wrapping the value in list", map[string]interface{}{flags["name"].(string): dataValue})
					dataValue = []interface{}{dataValue}
				}

				var diags diag.Diagnostics
				tfval, diags = basetypes.NewListValueFrom(ctx, elemType, dataValue)
				if diags != nil {
					return fmt.Errorf("ERROR: %v", diags)
				}
			}
			tflog.Trace(ctx, "setting list value", map[string]interface{}{"tfval": tfval})
			tfValueRefection := reflect.ValueOf(tfval)
			fValue.Set(tfValueRefection)

		default:
			//var tfval VyosResourceDataModel

			if !KeyInMap(flags["name"].(string), data) {
				fmt.Printf("\tNo data for field: %s, skipping\n", fName)
				tflog.Debug(ctx, "No data for field, skipping", map[string]interface{}{"field-name": fName})

				// TODO set nil / null for terraform state on sub-objects
				//tfval = basetypes.NewObjectNull()
				continue
			}
			dataValue := data[flags["name"].(string)]

			fmt.Printf("\tUnmarshalling Struct Field: %s\t%s=%s\n", fName, flags["name"].(string), dataValue)
			tflog.Debug(ctx, "Unmarshalling Struct Field", map[string]interface{}{"field-name": fName, flags["name"].(string): dataValue})

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
