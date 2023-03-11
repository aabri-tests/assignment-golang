package decoder

import (
	"strings"
	"testing"

	logger2 "github.com/aabri-assignments/assignment-golang/pkg/logger"
)

func TestJSONDecoder_Decode(t *testing.T) {
	logger := logger2.InitLogger()
	decoder := New(logger)

	tests := []struct {
		input   string
		value   interface{}
		wantErr bool
	}{
		{`{"foo": "bar"}`, &struct{ Foo string }{}, false},
		{`{"foo": "bar"}`, nil, true},
		{``, &struct{ Foo string }{}, true},
		{`{"foo": }`, &struct{ Foo string }{}, true},
		{`{"baz": "qux"}`, &struct{ Foo string }{}, true},
	}

	for i, tt := range tests {
		r := strings.NewReader(tt.input)
		err := decoder.Decode(r, tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("test %d: Decode(%q, %T) returned unexpected error status: %v", i, tt.input, tt.value, err)
		}
	}
}
