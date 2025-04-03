package config

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	cases := []struct {
		name       string
		env        map[string]string
		expectFail bool
	}{
		{
			name: "Testing correct config",
			env: map[string]string{
				"ENV":              "local",
				"UUID":             "6f018ae3-2f4a-415b-a13c-adce12af1nm7",
				"REGEXP":           `^[a-zA-Z0-9.]+@[a-z0-9.-]+\.ru$`,
				"GRPC_PORT":        "44044",
				"GRPC_TIMEOUT":     "10h",
				"REFRESHTOKEN_TTL": "12h",
				"ACCESSTOKEN_TTL":  "30m",
				"POSTGRES_URL":     "postgres://default:default@localhost:5432/auth",
				"REDIS_URL":        "redis://default:default@sso_redis:6379/",
			},
			expectFail: false,
		},
		{
			name: "Missing required variable",
			env: map[string]string{
				"ENV":              "local",
				"UUID":             "6f018ae3-2f4a-415b-a13c-adce12af1nm7",
				"REGEXP":           `^[a-zA-Z0-9.]+@[a-z0-9.-]+\.ru$`,
				"GRPC_PORT":        "44044",
				"GRPC_TIMEOUT":     "10h",
				"REFRESHTOKEN_TTL": "12h",
				"ACCESSTOKEN_TTL":  "30m",
			},
			expectFail: true,
		},
		{
			name: "Invalid duration format",
			env: map[string]string{
				"ENV":              "local",
				"UUID":             "6f018ae3-2f4a-415b-a13c-adce12af1nm7",
				"REGEXP":           `^[a-zA-Z0-9.]+@[a-z0-9.-]+\.ru$`,
				"GRPC_PORT":        "44044",
				"GRPC_TIMEOUT":     "10h",
				"REFRESHTOKEN_TTL": "invalid_duration",
				"ACCESSTOKEN_TTL":  "30m",
				"POSTGRES_URL":     "postgres://default:default@localhost:5432/auth",
				"REDIS_URL":        "redis://default:default@sso_redis:6379/",
			},
			expectFail: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for key, value := range tc.env {
				if err := os.Setenv(key, value); err != nil {
					t.Fatalf("Ошибка установки переменной окружения %s: %v", key, err)
				}
			}
			defer func() {
				for key := range tc.env {
					os.Unsetenv(key)
				}
			}()
			_, err := MustLoad()

			if (err != nil) != tc.expectFail {
				t.Errorf("Expected error: %v, got: %v", tc.expectFail, err)
			}

		})

	}
}
