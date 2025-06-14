// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Std_count_Name = "std.count"

var Std_count_ArgumentTypes = []value.Type{value.IdentType}

func Std_count_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough(Std_count_Name, 1, args)
	}
	for i := range args {
		if args[i].Type() != Std_count_ArgumentTypes[i] {
			return errors.TypeMismatch(Std_count_Name, i+1, Std_count_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of std.count
// Arguments may be:
// - ID
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/std-count/
func Std_count(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Std_count_Validate(args); err != nil {
		return value.Null, err
	}

	// TODO: On Fastly example, uncodumented "req.headers" variable is used...what?
	// But currently we could support xxx.headers ident only...
	ident := value.Unwrap[*value.Ident](args[0])

	switch ident.Value {
	case "req.headers":
		return &value.Integer{Value: int64(len(ctx.Request.Header))}, nil
	case "bereq.headers":
		return &value.Integer{Value: int64(len(ctx.BackendRequest.Header))}, nil
	case "beresp.headers":
		return &value.Integer{Value: int64(len(ctx.BackendResponse.Header))}, nil
	case "obj.headers":
		return &value.Integer{Value: int64(len(ctx.Object.Header))}, nil
	case "resp.headers":
		return &value.Integer{Value: int64(len(ctx.Response.Header))}, nil
	default:
		return value.Null, errors.New(
			Std_count_Name,
			"Unexpected or Unsupported collection ident %s is provided", ident.Value,
		)
	}
}
