package BaseConfig

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)
	envPath := filepath.Join(currentDir, "..", "..", ".env")

	if err := godotenv.Load(envPath); err != nil {
		// FALLBACK: If .env doesn't exist (like in a CI/CD pipeline),
		// set default test environment variables manually so tests still pass.
		_ = os.Setenv("FilePathType", "yaml")
		_ = os.Setenv("FilePath", "config.yaml")
	}
}

func Test_initFileConfigurationData(t *testing.T) {
	config := ConfigurationData{}
	isAnyError, err := config._initFileConfigurationData()
	assert.NoError(t, err, "valid input should not cause an error")
	assert.NotEqual(t, isAnyError, true, "valid input should not isAnyError as true")

	assert.NotEqual(t, len(config.File.FilePath) > 0, true, "valid input should not isAnyError as true")
	assert.NotEqual(t, len(config.File.FilePathType) > 0, true, "valid input should not isAnyError as true")

}

func Test_init(t *testing.T) {
	isAnyError, err, config := _init()
	assert.NoError(t, err, "valid input should not cause an error")
	assert.NotEqual(t, isAnyError, true, "valid input should not isAnyError as true")

	assert.NotEqual(t, len(config.File.FilePath) > 0, true, "valid input should not isAnyError as true")
	assert.NotEqual(t, len(config.File.FilePathType) > 0, true, "valid input should not isAnyError as true")

}
