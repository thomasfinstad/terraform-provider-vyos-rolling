package helpers

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers/tools"
)

// MarshalVyos takes a Terraform resource model pointer and returns a vyos config representation
func MarshalVyos(ctx context.Context, data any) (map[string]any, error) {
	tools.Trace(ctx, "Marshal for VyOS", map[string]interface{}{"data-model": data})
	res := make(map[string]interface{})

	valueReflection := reflect.ValueOf(data)
	valueReflection = valueReflection.Elem()
	typeReflection := valueReflection.Type()

	for i := 0; i < valueReflection.NumField(); i++ {
		fName := typeReflection.Field(i).Name
		fType := typeReflection.Field(i).Type
		fValue := valueReflection.Field(i)
		fTags := strings.Split(typeReflection.Field(i).Tag.Get("vyos"), ",")
		for i, v := range fTags {
			if v == "" {
				fTags = append(fTags[:i], fTags[i+1:]...)
			}
		}
		tools.Debug(ctx, "processing field", map[string]interface{}{"Field": fName, "Type": fType, "Value-Interface": fValue, "Tags": fTags})
		tools.Trace(ctx, "processing field", map[string]interface{}{"Field": fName, "Type": fType, "Value-Interface": fValue, "Tags": fTags, "data": data})

		if len(fTags) < 1 {
			tools.Error(ctx, "no vyos tags found on field, at least the name field must be filled. an underscore can be used if no real value is propriate")
			panic("no vyos tags found on field, at least the name field must be filled. an underscore can be used if no real value is propriate")
		}

		// Set flags based on tags, first tag must be the vyos field name, the rest are bools with default of false
		// TODO create struct of valid struct field flag options
		//  milestone:6
		flags := map[string]any{
			"name":      fTags[0],
			"self-id":   false,
			"parent-id": false,
			"omitempty": false,
			"child":     false,
			"timeout":   false,
			"tfsdk-id":  false,
		}
		for _, tag := range fTags[1:] {

			tools.Trace(ctx, "Enabling flag", map[string]interface{}{"flag": tag})
			flags[tag] = true
		}
		if flags["child"].(bool) || flags["self-id"].(bool) || flags["parent-id"].(bool) || flags["tfsdk-id"].(bool) || flags["timeout"].(bool) {

			tools.Debug(ctx, "Not configuring field", map[string]interface{}{"field-name": fName, "flags": flags})
			continue
		}

		switch v := fValue.Interface().(type) {
		case basetypes.StringValue:
			if !(v.IsNull() || v.IsUnknown()) {

				tools.Debug(ctx, "Marshalling String Field", map[string]interface{}{"field-name": fName, "flag:name": flags["name"], "value-string": v.ValueString()})
				res[flags["name"].(string)] = v.ValueString()
			} else if !flags["omitempty"].(bool) {
				panic(fmt.Sprintf("Missing value: %s", fName))
			}
		case basetypes.BoolValue:

			tools.Debug(ctx, "Marshalling Bool Field", map[string]interface{}{"field-name": fName})
			if v.IsNull() || v.IsUnknown() {

				tools.Debug(ctx, "Marshalling Bool Field  is nil, skipping", map[string]interface{}{"field-name": fName})
			} else if !v.ValueBool() {

				tools.Debug(ctx, "Marshalling Bool Field  is false, skipping", map[string]interface{}{"field-name": fName, flags["name"].(string): v})
			} else {

				tools.Debug(ctx, "Marshalling Bool Field  is true", map[string]interface{}{"field-name": fName, flags["name"].(string): v})
				res[flags["name"].(string)] = map[string]interface{}{}
			}
		case basetypes.NumberValue:
			if !(v.IsNull() || v.IsUnknown()) {

				tools.Debug(ctx, "Marshalling Number Field", map[string]interface{}{"field-name": fName, flags["name"].(string): v.ValueBigFloat()})
				res[flags["name"].(string)], _ = v.ValueBigFloat().Float64()
			} else if !flags["omitempty"].(bool) {
				panic(fmt.Sprintf("Missing value: %s", fName))
			}
		case basetypes.ListValue:
			if !(v.IsNull() || v.IsUnknown()) {

				tools.Debug(ctx, "Marshalling List Field", map[string]interface{}{"field-name": flags["name"].(string), "field-value": fmt.Sprintf("%#v", v)})
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
				tools.Debug(ctx, "Add list elements", map[string]interface{}{"list": res[flags["name"].(string)]})

			} else if !flags["omitempty"].(bool) {
				panic(fmt.Sprintf("Missing value: %s", fName))
			}
		default:
			// TODO cleanup nested ifs inside switch and flatten out to only use switch statement, this might require not using a type switch
			//  milestone:6
			if fType.Kind() == reflect.Ptr {
				fType = fValue.Type().Elem()
			}
			if fType.Kind() == reflect.Struct {

				tools.Debug(ctx, "Marshalling Struct Field", map[string]interface{}{"field-name": fName, flags["name"].(string): fValue.Interface()})
				if !(fValue.IsNil() || fValue.IsZero()) {
					type subctxkey string
					subctx := context.WithValue(ctx, subctxkey(typeReflection.Name()), flags["name"].(string))
					subv, err := MarshalVyos(subctx, fValue.Interface().(VyosResourceDataModel))
					if err != nil {
						return nil, err
					}
					res[flags["name"].(string)] = subv
				} else {

					tools.Debug(ctx, "Marshalling Struct Field is nil, skipping", map[string]interface{}{"field-name": fName})

				}
			} else {

				tools.Error(ctx, "Currently unhandled tf type", map[string]interface{}{"field-name": fName, "kind": fType.Kind(), "value": v})
			}
		}
	}

	// Return data

	tools.Trace(ctx, "Marshal return", map[string]interface{}{"res": res})
	return res, nil
}
