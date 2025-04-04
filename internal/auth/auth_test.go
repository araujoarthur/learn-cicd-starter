package auth

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAPIKey(t *testing.T) {
	type test struct {
		key         string
		val         string
		expected    string
		expectedErr string
	}

	tests := []test{
		{
			key:         "",
			val:         "",
			expectedErr: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			key:      "Authorization",
			val:      "ApiKey ABCD-123sXyz",
			expected: "ABCD-123sXyz",
		},
		{
			key:         "Authorization",
			val:         "ABCD-123sXyz",
			expectedErr: "malformed authorization header",
		},
		{
			key:         "Authorization",
			val:         "ApiKey",
			expectedErr: "malformed authorization header",
		},
		{
			key:         "Authorization",
			val:         "",
			expectedErr: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			key:         "Authoization",
			val:         "ABCD-123sXyz",
			expected:    "ABCD-123sXyz",
			expectedErr: ErrNoAuthHeaderIncluded.Error(),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(tt.key, tt.val)
			outp, err := GetAPIKey(header)
			if err != nil {
				if reflect.DeepEqual(err.Error(), tt.expectedErr) {
					return
				} else {
					t.Errorf("Unexpected error: %q\n", err)
					return
				}
			}

			if outp != tt.expected {
				t.Errorf("Unexpected return value: %q. Wanted %q", outp, tt.expected)
				return
			}
		})

	}
}
