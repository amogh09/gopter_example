package base64

import (
	"encoding/base64"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestEncodeDecode(t *testing.T) {
	params := gopter.DefaultTestParameters()
	params.MaxSize = 10000
	params.MinSuccessfulTests = 1000
	properties := gopter.NewProperties(params)

	properties.Property("encode-decode works", prop.ForAll(
		func(s string) bool {
			return Decode((Encode(s))) == s
		},
		gen.AnyString(),
	))

	properties.TestingRun(t)
}

func TestEncode(t *testing.T) {
	params := gopter.DefaultTestParameters()
	params.MaxSize = 10000
	params.MinSuccessfulTests = 1000
	properties := gopter.NewProperties(params)

	properties.Property("encode matches standard lib", prop.ForAll(
		func(s string) bool {
			return Encode(s) == base64.StdEncoding.EncodeToString([]byte(s))
		},
		gen.AnyString(),
	))

	properties.TestingRun(t)
}
