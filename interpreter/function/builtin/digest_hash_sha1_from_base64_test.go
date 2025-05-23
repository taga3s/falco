// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"testing"

	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of digest.hash_sha1_from_base64
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-sha1-from-base64/
func Test_Digest_hash_sha1_from_base64(t *testing.T) {
	ret, err := Digest_hash_sha1_from_base64(
		&context.Context{},
		&value.String{Value: "SGVsbG8sIHdvcmxkIQo="},
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if ret.Type() != value.StringType {
		t.Errorf("Unexpected return type, expect=STRING, got=%s", ret.Type())
	}
	v := value.Unwrap[*value.String](ret)
	expect := "09fac8dbfd27bd9b4d23a00eb648aa751789536d"
	if v.Value != expect {
		t.Errorf("return value unmatch, expect=%s, got=%s", expect, v.Value)
	}
}
