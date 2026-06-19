package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		authHeader  string
		expectedKey string
		expectError bool
	}{
		{
			name:        "valid api key",
			authHeader:  "ApiKey secret123",
			expectedKey: "secret123",
			expectError: false,
		},
		{
			name:        "missing authorization header",
			authHeader:  "",
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "wrong auth scheme",
			authHeader:  "Bearer secret123",
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "missing api key",
			authHeader:  "ApiKey",
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "too many spaces",
			authHeader:  "ApiKey secret123 extra",
			expectedKey: "secret123",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}

			if tt.authHeader != "" {
				headers.Set("Authorization", tt.authHeader)
			}

			key, err := GetAPIKey(headers)

			if tt.expectError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if key != tt.expectedKey {
				t.Errorf("expected key %q, got %q", tt.expectedKey, key)
			}
		})
	}
}
