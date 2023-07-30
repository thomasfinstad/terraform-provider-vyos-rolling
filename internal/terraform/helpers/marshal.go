package helpers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// MarshalVyos takes a Terraform resource model pointer and returns a vyos config representation
func MarshalVyos(data any) ([]byte, error) {
	res := make(map[string]interface{})

	valueOfData := reflect.ValueOf(data).Elem()
	typeOfData := valueOfData.Type()

	for i := 0; i < valueOfData.NumField(); i++ {
		fName := typeOfData.Field(i).Name
		fType := typeOfData.Field(i).Type
		fValue := valueOfData.Field(i)
		fTags := strings.Split(typeOfData.Field(i).Tag.Get("vyos"), ",")
		fmt.Printf("Field: %s\tType: %s\tValue: %v\tTags: %v\n", fName, fType, fValue.Interface(), fTags)

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
			fmt.Printf("Not configuring field: %s\n", fName)
			continue
		}

		switch v := fValue.Interface().(type) {
		case basetypes.StringValue:
			if !(v.IsNull() || v.IsUnknown()) {
				fmt.Printf("Marshalling String Field: %s\t%s=%s\n", fName, flags["name"].(string), v.ValueString())
				res[flags["name"].(string)] = v.ValueString()
			} else if !flags["omitempty"].(bool) {
				panic("Missing value")
			}
		default:
			// TODO cleanup nested ifs inside switch and flatten out to only use switch statement
			if fType.Kind() == reflect.Ptr {
				fType = fValue.Type().Elem()
			}
			if fType.Kind() == reflect.Struct {
				fmt.Printf("Marshalling Struct Field: %s\t%s=%s\n", fName, flags["name"].(string), fValue.Interface())
				if !fValue.IsNil() {

					j, err := MarshalVyos(fValue.Interface())
					if err != nil {
						return nil, err
					}
					subv := make(map[string]any)
					json.Unmarshal(j, &subv)
					res[flags["name"].(string)] = subv
				} else {
					fmt.Printf("Marshalling Struct Field: %s is nil, skipping\n", fName)
				}
			} else {
				fmt.Printf("Currently unhandled tf type: %s\tkind: %v\tvalue:%#v\n", fType, fType.Kind(), v)
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
