package tools

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

/*
Trace logs `msg` at the trace level to the logger in `ctx`, with optional `additionalFields` structured key-value fields in the log output. Fields are shallow merged with any defined on the logger, e.g. by the `SetField()` function, and across multiple maps.
*/
func Trace(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	for i, v := range additionalFields {
		for fk, fv := range v {
			additionalFields[i][fk] = stringify(fv)
		}
	}

	tflog.Trace(ctx, msg, additionalFields...)
}

/*
Debug logs `msg` at the debug level to the logger in `ctx`, with optional `additionalFields` structured key-value fields in the log output. Fields are shallow merged with any defined on the logger, e.g. by the `SetField()` function, and across multiple maps.
*/
func Debug(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	for i, v := range additionalFields {
		for fk, fv := range v {
			additionalFields[i][fk] = stringify(fv)
		}
	}

	tflog.Debug(ctx, msg, additionalFields...)
}

/*
Info logs `msg` at the info level to the logger in `ctx`, with optional `additionalFields` structured key-value fields in the log output. Fields are shallow merged with any defined on the logger, e.g. by the `SetField()` function, and across multiple maps.
*/
func Info(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	for i, v := range additionalFields {
		for fk, fv := range v {
			additionalFields[i][fk] = stringify(fv)
		}
	}

	tflog.Info(ctx, msg, additionalFields...)
}

/*
Warn logs `msg` at the warn level to the logger in `ctx`, with optional `additionalFields` structured key-value fields in the log output. Fields are shallow merged with any defined on the logger, e.g. by the `SetField()` function, and across multiple maps.
*/
func Warn(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	for i, v := range additionalFields {
		for fk, fv := range v {
			additionalFields[i][fk] = stringify(fv)
		}
	}

	tflog.Warn(ctx, msg, additionalFields...)
}

/*
Error logs `msg` at the error level to the logger in `ctx`, with optional `additionalFields` structured key-value fields in the log output. Fields are shallow merged with any defined on the logger, e.g. by the `SetField()` function, and across multiple maps.
*/
func Error(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	for i, v := range additionalFields {
		for fk, fv := range v {
			additionalFields[i][fk] = stringify(fv)
		}
	}

	tflog.Error(ctx, msg, additionalFields...)
}
