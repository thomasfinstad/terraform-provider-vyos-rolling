package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MarshalVyos takes a Terraform resource model pointer and returns a vyos config representation
func MarshalVyos(ctx context.Context, data any) ([]byte, error) {
	res := make(map[string]interface{})

	valueOfData := reflect.ValueOf(data).Elem()
	typeOfData := valueOfData.Type()

	for i := 0; i < valueOfData.NumField(); i++ {
		fName := typeOfData.Field(i).Name
		fType := typeOfData.Field(i).Type
		fValue := valueOfData.Field(i)
		fTags := strings.Split(typeOfData.Field(i).Tag.Get("vyos"), ",")
		tflog.Debug(ctx, "processing field", map[string]interface{}{"Field": fName, "Type": fType, "Value-Interface": fValue, "Tags": fTags})
		fmt.Printf("Field: %s\tType: %s\tValue: %v\tTags: %v\n", fName, fType, fValue, fTags)

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
		if flags["child"].(bool) || flags["self-id"].(bool) || flags["parent-id"].(bool) {
			//fmt.Printf("Not configuring field: %s\n", fName)
			tflog.Debug(ctx, "Not configuring field", map[string]interface{}{"field-name": fName})
			continue
		}

		switch v := fValue.Interface().(type) {
		case basetypes.StringValue:
			if !(v.IsNull() || v.IsUnknown()) {
				fmt.Printf("Marshalling String Field: %s\t%s=%s\n", fName, flags["name"].(string), v.ValueString())
				tflog.Debug(ctx, "Marshalling String Field", map[string]interface{}{"field-name": fName, flags["name"].(string): v.ValueString()})
				res[flags["name"].(string)] = v.ValueString()
			} else if !flags["omitempty"].(bool) {
				panic("Missing value")
			}
		case basetypes.BoolValue:
			fmt.Printf("Marshalling Bool Field: %s\n", fName)
			tflog.Debug(ctx, "Marshalling Bool Field", map[string]interface{}{"field-name": fName})
			if v.IsNull() || v.IsUnknown() {
				fmt.Printf("Marshalling Bool Field: %s is nil, skipping\n", fName)
				tflog.Debug(ctx, "Marshalling Bool Field  is nil, skipping", map[string]interface{}{"field-name": fName})
			} else if !v.ValueBool() {
				fmt.Printf("Marshalling Bool Field: %s is false, skipping: %s=%s\n", fName, flags["name"].(string), v)
				tflog.Debug(ctx, "Marshalling Bool Field  is false, skipping", map[string]interface{}{"field-name": fName, flags["name"].(string): v})
			} else {
				fmt.Printf("Marshalling Bool Field: %s is true: %s=%s\n", fName, flags["name"].(string), v)
				tflog.Debug(ctx, "Marshalling Bool Field  is true", map[string]interface{}{"field-name": fName, flags["name"].(string): v})
				res[flags["name"].(string)] = map[string]interface{}{}
			}
		case basetypes.NumberValue:
			if !(v.IsNull() || v.IsUnknown()) {
				fmt.Printf("Marshalling Number Field: %s\t%s=%s\n", fName, flags["name"].(string), v.ValueBigFloat())
				tflog.Debug(ctx, "Marshalling Number Field", map[string]interface{}{"field-name": fName, flags["name"].(string): v.ValueBigFloat()})
				res[flags["name"].(string)], _ = v.ValueBigFloat().Float64()
			} else if !flags["omitempty"].(bool) {
				panic("Missing value")
			}
		case basetypes.ListValue:
			if !(v.IsNull() || v.IsUnknown()) {
				fmt.Printf("Marshalling List Field: %s\t%s=%s\n", fName, flags["name"].(string), v)
				switch v.ElementType(ctx).(type) {
				case basetypes.StringType:
					var lst []string
					diags := v.ElementsAs(ctx, &lst, false)
					if diags != nil {
						return nil, fmt.Errorf("ERROR: %v", diags)
					}
					res[flags["name"].(string)] = lst
				case basetypes.NumberType:
					var lst []float64
					diags := v.ElementsAs(ctx, &lst, false)
					if diags != nil {
						return nil, fmt.Errorf("ERROR: %v", diags)
					}
					res[flags["name"].(string)] = lst
				default:
					panic("Unhandled type in list")
				}

			} else if !flags["omitempty"].(bool) {
				panic("Missing value")
			}
		default:
			// TODO cleanup nested ifs inside switch and flatten out to only use switch statement, this might require not using a type switch
			if fType.Kind() == reflect.Ptr {
				fType = fValue.Type().Elem()
			}
			if fType.Kind() == reflect.Struct {
				fmt.Printf("Marshalling Struct Field: %s\t%s=%s\n", fName, flags["name"].(string), fValue.Interface())
				tflog.Debug(ctx, "Marshalling Struct Field", map[string]interface{}{"field-name": fName, flags["name"].(string): fValue.Interface()})
				if !(fValue.IsNil() || fValue.IsZero()) {
					type subctxkey string
					subctx := context.WithValue(ctx, subctxkey(typeOfData.Name()), flags["name"].(string))
					j, err := MarshalVyos(subctx, fValue.Interface())
					if err != nil {
						return nil, err
					}
					subv := make(map[string]any)
					json.Unmarshal(j, &subv)
					res[flags["name"].(string)] = subv
				} else {
					fmt.Printf("Marshalling Struct Field: %s is nil, skipping\n", fName)
					tflog.Debug(ctx, "Marshalling Struct Field is nil, skipping", map[string]interface{}{"field-name": fName})

				}
			} else {
				fmt.Printf("ERROR!!! Currently unhandled tf type: %s\tkind: %v\tvalue:%#v\n", fType, fType.Kind(), v)
				tflog.Error(ctx, "Currently unhandled tf type", map[string]interface{}{"field-name": fName, "kind": fType.Kind(), "value": v})
			}
		}
	}

	// Return compiled data
	ret, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
