package yaml_parser

import (
	Configurator "Universal-Rate-Limiter/src/BaseConfig"
	RateConfigurator "Universal-Rate-Limiter/src/ConfigReader"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_init(t *testing.T) {
	testfileconfig := Configurator.FileConfigurationData{
		FilePathType: "Local",
		FilePath:     "config/Test.yaml",
	}
	// parser := YamlParser{}
	// Test success case
	isAnyError, err, obj := _init(&testfileconfig)
	assert.NoError(t, err, "valid input should not cause an error")
	assert.NotEqual(t, isAnyError, true, "valid input should not isAnyError as true")

	switch reflect.TypeOf(obj).Name() {
	case "YamlParser":
	default:
		assert.Error(t, errors.New("Invalid Type Recieved"))
	}

}

func TestReadFile(t *testing.T) {
	testfileconfig := Configurator.FileConfigurationData{
		FilePathType: "Local",
		FilePath:     "D:\\Coding\\projects\\Go\\Universal-Rate-Limiter\\config\\Test.yaml",
	}

	isAnyError, err, obj := _init(&testfileconfig)
	assert.NoError(t, err, "valid input should not cause an error")
	assert.NotEqual(t, isAnyError, true, "valid input should not isAnyError as true")

	switch reflect.TypeOf(obj).Name() {
	case "YamlParser":
	default:
		assert.Error(t, errors.New("Invalid Type Recieved"))
	}
	store := RateConfigurator.RateLimiterConfiguration{}
	isAnyError, err = obj.ReadFile(&store)

	assert.NoError(t, err, "valid input should not cause an error")
	assert.NotEqual(t, isAnyError, true, "valid input should not isAnyError as true")

}

func BenchmarkMyFunc(b *testing.B) {
	testfileconfig := Configurator.FileConfigurationData{
		FilePathType: "Local",
		FilePath:     "D:\\Coding\\projects\\Go\\Universal-Rate-Limiter\\config\\Test.yaml",
	}

	isAnyError, err, obj := _init(&testfileconfig)
	if err != nil {

		fmt.Print(isAnyError)
		fmt.Print(err.Error())
	}
	store := RateConfigurator.RateLimiterConfiguration{}
	for i := 0; i < b.N; i++ {
		isAnyError, err := obj.ReadFile(&store)
		if err != nil {

			fmt.Print(isAnyError)
			fmt.Print(err.Error())
		}
	}
}

/*
func Testunmarshal(t *testing.T) {
	// Test success case
	result, err := Parse("valid input")
	assert.NoError(t, err, "valid input should not cause an error")
	assert.Equal(t, 42, result, "valid input should return 42")

	// Test error case
	result, err = Parse("")
	assert.Error(t, err, "empty input should cause an error")
	assert.Equal(t, "empty input", err.Error(), "error message should match")
	assert.Equal(t, 0, result, "error case should return zero value")
}

func TestrateLimiterConfigurationParser(t *testing.T) {
	// Test success case
	result, err := Parse("valid input")
	assert.NoError(t, err, "valid input should not cause an error")
	assert.Equal(t, 42, result, "valid input should return 42")

	// Test error case
	result, err = Parse("")
	assert.Error(t, err, "empty input should cause an error")
	assert.Equal(t, "empty input", err.Error(), "error message should match")
	assert.Equal(t, 0, result, "error case should return zero value")
}

func TestdescriptorParser(t *testing.T) {
	// Test success case
	result, err := Parse("valid input")
	assert.NoError(t, err, "valid input should not cause an error")
	assert.Equal(t, 42, result, "valid input should return 42")

	// Test error case
	result, err = Parse("")
	assert.Error(t, err, "empty input should cause an error")
	assert.Equal(t, "empty input", err.Error(), "error message should match")
	assert.Equal(t, 0, result, "error case should return zero value")
}

func TestrateLimitParser(t *testing.T) {
	// Test success case
	result, err := Parse("valid input")
	assert.NoError(t, err, "valid input should not cause an error")
	assert.Equal(t, 42, result, "valid input should return 42")

	// Test error case
	result, err = Parse("")
	assert.Error(t, err, "empty input should cause an error")
	assert.Equal(t, "empty input", err.Error(), "error message should match")
	assert.Equal(t, 0, result, "error case should return zero value")
}

func TeststringParser(t *testing.T) {
	// Test success case
	result, err := Parse("valid input")
	assert.NoError(t, err, "valid input should not cause an error")
	assert.Equal(t, 42, result, "valid input should return 42")

	// Test error case
	result, err = Parse("")
	assert.Error(t, err, "empty input should cause an error")
	assert.Equal(t, "empty input", err.Error(), "error message should match")
	assert.Equal(t, 0, result, "error case should return zero value")
}

func Testint64Parser(t *testing.T) {
	// Test success case
	result, err := Parse("valid input")
	assert.NoError(t, err, "valid input should not cause an error")
	assert.Equal(t, 42, result, "valid input should return 42")

	// Test error case
	result, err = Parse("")
	assert.Error(t, err, "empty input should cause an error")
	assert.Equal(t, "empty input", err.Error(), "error message should match")
	assert.Equal(t, 0, result, "error case should return zero value")
}
*/
