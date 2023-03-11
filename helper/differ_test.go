package helper

import "testing"

func TestFindFirstDifference(t *testing.T) {
	tests := []struct {
		a, b []string
		want string
	}{
		{[]string{"foo", "bar", "baz"}, []string{"foo", "baz"}, "bar"},
		{[]string{"foo", "bar"}, []string{"foo", "bar"}, ""},
		{[]string{"foo", "bar", "baz"}, []string{"baz", "bar", "foo"}, ""},
	}

	for _, tt := range tests {
		got := FindFirstDifference(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("FindFirstDifference(%v, %v) = %q; want %q", tt.a, tt.b, got, tt.want)
		}
	}
}
