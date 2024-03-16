package helpers

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Verify compliance
var _ validator.String = StringNotValidator{}

// StringNot returns a validator negating the sub validators results
func StringNot(v validator.String) validator.String {
	return StringNotValidator{
		validator: v,
	}
}

// StringNotValidator implements the negating validator.
type StringNotValidator struct {
	validator validator.String
}

// Description explains the validation.
func (v StringNotValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("illigal value: %s", v.validator.Description(ctx))
}

// MarkdownDescription explains the validation in md.
func (v StringNotValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// ValidateString executes negative validation.
func (v StringNotValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateResp := &validator.StringResponse{}
	v.validator.ValidateString(ctx, req, validateResp)
	if !validateResp.Diagnostics.HasError() {
		validateResp.Diagnostics.AddError("negating validation error", v.Description(ctx))
	}
}
