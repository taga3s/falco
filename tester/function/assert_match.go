package function

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
	regexp "go.elara.ws/pcre"
)

const Assert_match_Name = "assert.match"

var Assert_match_ArgumentTypes = []value.Type{value.StringType, value.StringType}

func Assert_match_Validate(args []value.Value) error {
	if len(args) < 2 || len(args) > 3 {
		return errors.ArgumentNotInRange(Assert_match_Name, 2, 3, args)
	}

	for i := range Assert_match_ArgumentTypes {
		if args[i].Type() != Assert_match_ArgumentTypes[i] {
			return errors.TypeMismatch(Assert_match_Name, i+1, Assert_match_ArgumentTypes[i], args[i].Type())
		}
	}

	if len(args) == 3 {
		if args[2].Type() != value.StringType {
			return errors.TypeMismatch(Assert_match_Name, 3, value.StringType, args[2].Type())
		}
	}
	return nil
}

func Assert_match(ctx *context.Context, args ...value.Value) (value.Value, error) {
	if err := Assert_match_Validate(args); err != nil {
		return nil, errors.NewTestingError("%s", err.Error())
	}

	// Check custom message
	var message string
	if len(args) == 3 {
		message = value.Unwrap[*value.String](args[2]).Value
	}

	actual := value.Unwrap[*value.String](args[0])
	expect := value.Unwrap[*value.String](args[1])

	re, err := regexp.Compile(expect.Value)
	if err != nil {
		return nil, errors.NewTestingError(
			"Invalid regexp string provided: %s",
			expect.Value,
		)
	}
	ret := &value.Boolean{Value: re.MatchString(actual.Value)}
	if !ret.Value {
		if message != "" {
			return ret, errors.NewAssertionError(actual, "%s", message)
		}
		return ret, errors.NewAssertionError(
			actual,
			`"%s" should match against %s`,
			actual.Value,
			expect.Value,
		)
	}
	return ret, nil
}
