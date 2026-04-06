package config_test

import (
	"dotmgr/internal/config"
	"testing"
)

const testDataDirectory = "testdata/"

type testConfiguration struct {
	name        string
	file        string
	validConfig bool
}

// All separate allTestConfigurations and if they are valid declared here. Tests run against each config declared
var allTestConfigurations = []testConfiguration{
	{
		name:        "Full valid config",
		file:        "config.toml",
		validConfig: true,
	},
	{
		name:        "Invalid Git URL",
		file:        "invalidURL.toml",
		validConfig: false,
	},
}

// For each defined config runs all tests against it. Configurations with the validConfig bool
// set to false have inverse checks.
func TestAllConfigs(t *testing.T) {
	for _, testConf := range allTestConfigurations {
		t.Run(testConf.name, func(t *testing.T) {
			// Load current configuration, working directory is the package directory
			_, err := config.LoadFrom(testDataDirectory + testConf.file)

			// For case when config is invalid and should fail
			if !testConf.validConfig {
				if err == nil {
					t.Fatalf("expected error, got nil for case %v", testConf.name)
				}
				t.Logf("got expected error: %v", err)
				return
			}

			if err != nil {
				t.Error(err)
			}
		})
	}
}

