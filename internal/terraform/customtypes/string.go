package customtypes

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.StringTypable  = CustomStringType{}
	_ basetypes.StringValuable = CustomStringValue{}
)

// CustomStringType for use in terraform schemas, returns as CustomStringValue
type CustomStringType struct {
}

// ApplyTerraform5AttributePathStep applies the given AttributePathStep to the
// type.
func (t CustomStringType) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (interface{}, error) {
	return nil, fmt.Errorf("cannot apply AttributePathStep %T to %s", step, t.String())
}

// Equal returns true if the given type is equivalent.
func (t CustomStringType) Equal(o attr.Type) bool {
	_, ok := o.(CustomStringType)

	return ok
}

// String returns a human readable string of the type name.
func (t CustomStringType) String() string {
	return "customtypes.CustomStringType"
}

// TerraformType returns the tftypes.Type that should be used to represent this
// framework type.
func (t CustomStringType) TerraformType(_ context.Context) tftypes.Type {
	return tftypes.String
}

// ValueFromString returns a StringValuable type given a StringValue.
func (t CustomStringType) ValueFromString(_ context.Context, v basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return v, nil
}

// ValueFromTerraform returns a Value given a tftypes.Value.  This is meant to
// convert the tftypes.Value into a more convenient Go type for the provider to
// consume the data with.
func (t CustomStringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if !in.IsKnown() {
		return CustomStringValue{
			state: attr.ValueStateNull,
		}, nil
	}

	if in.IsNull() {
		return CustomStringValue{
			state: attr.ValueStateNull,
		}, nil
	}

	var s string

	err := in.As(&s)

	if err != nil {
		return nil, err
	}

	return CustomStringValue{
		state: attr.ValueStateKnown,
		value: s,
	}, nil
}

// ValueType returns the Value type.
func (t CustomStringType) ValueType(_ context.Context) attr.Value {
	// This Value does not need to be valid.
	return CustomStringValue{}
}

// CustomStringValue represents a UTF-8 string value. And implements custom json (un)marshaling
type CustomStringValue struct {

	// state represents whether the value is null, unknown, or known. The
	// zero-value is null.
	state attr.ValueState

	// value contains the known value, if not null or unknown.
	value string
}

// Type returns a StringType.
func (o CustomStringValue) Type(_ context.Context) attr.Type {
	return CustomStringType{}
}

// ToTerraformValue returns the data contained in the *String as a tftypes.Value.
func (o CustomStringValue) ToTerraformValue(_ context.Context) (tftypes.Value, error) {
	switch o.state {
	case attr.ValueStateKnown:
		if err := tftypes.ValidateValue(tftypes.String, o.value); err != nil {
			return tftypes.NewValue(tftypes.String, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(tftypes.String, o.value), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(tftypes.String, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(tftypes.String, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled String state in ToTerraformValue: %s", o.state))
	}
}

// Equal returns true if `other` is a String and has the same value as `s`.
func (o CustomStringValue) Equal(otherValue attr.Value) bool {
	other, ok := otherValue.(CustomStringValue)

	if !ok {
		return false
	}

	if o.state != other.state {
		return false
	}

	if o.state != attr.ValueStateKnown {
		return true
	}

	return o.value == other.value
}

// IsNull returns true if the String represents a null value.
func (o CustomStringValue) IsNull() bool {
	return o.state == attr.ValueStateNull
}

// IsUnknown returns true if the String represents a currently unknown value.
func (o CustomStringValue) IsUnknown() bool {
	return o.state == attr.ValueStateUnknown
}

// String returns a human-readable representation of the String value. Use
// the ValueString method for Terraform data handling instead.
//
// The string returned here is not protected by any compatibility guarantees,
// and is intended for logging and error reporting.
func (o CustomStringValue) String() string {
	if o.IsUnknown() {
		return attr.UnknownValueString
	}

	if o.IsNull() {
		return attr.NullValueString
	}

	return fmt.Sprintf("%q", o.value)
}

// ValueString returns the known string value. If String is null or unknown, returns
// "".
// func (s CustomStringValue) ValueString() string {
// 	return s.value
// }

// ValueStringPointer returns a pointer to the known string value, nil for a
// null value, or a pointer to "" for an unknown value.
// func (s CustomStringValue) ValueStringPointer() *string {
// 	if s.IsNull() {
// 		return nil
// 	}

// 	return &s.value
// }

// ToStringValue returns String.
func (o CustomStringValue) ToStringValue(context.Context) (basetypes.StringValue, diag.Diagnostics) {
	if o.IsUnknown() {
		return basetypes.NewStringUnknown(), nil
	}

	if o.IsNull() {
		return basetypes.NewStringNull(), nil
	}

	return basetypes.NewStringValue(o.value), nil
}

// MarshalJSON for native go types instead of tftypes
func (o *CustomStringValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.String())
}
