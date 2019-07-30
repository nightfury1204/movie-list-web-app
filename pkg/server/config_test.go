package server

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	os.Setenv("OMDB_API_TOKEN", "test")
	cfg := NewConfig()
	assert.Equal(t, cfg.Port, DefaultPort)
	assert.Equal(t, cfg.OMDBAPIUrl, OMDBAPIUrl)
	assert.Equal(t, cfg.OMDBAPIToken, "test")
}

func TestConfig_Validation(t *testing.T) {
	testData := []struct {
		testName  string
		cfg       *Config
		expectErr bool
	}{
		{
			testName: "All required value is provided, validation passed",
			cfg: &Config{
				Port:         "443",
				OMDBAPIToken: "test",
				OMDBAPIUrl:   "api.com",
			},
			expectErr: false,
		},
		{
			testName: "Port is empty, validation failed",
			cfg: &Config{
				Port:         "",
				OMDBAPIToken: "test",
				OMDBAPIUrl:   "api.com",
			},
			expectErr: true,
		},
		{
			testName: "OMDB api token is empty, validation failed",
			cfg: &Config{
				Port:         "443",
				OMDBAPIToken: "",
				OMDBAPIUrl:   "api.com",
			},
			expectErr: true,
		},
		{
			testName: "OMDB api url is empty, validation failed",
			cfg: &Config{
				Port:         "443",
				OMDBAPIToken: "test",
				OMDBAPIUrl:   "",
			},
			expectErr: true,
		},
	}

	for _, tc := range testData {
		t.Run(tc.testName, func(t *testing.T) {
			err := tc.cfg.Validate()
			if tc.expectErr {
				assert.NotNil(t, err, "error expected")
			} else {
				assert.Nil(t, err, "error is not expected")
			}
		})
	}
}
