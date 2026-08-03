package BaseConfig

import (
	"errors"
	"os"
)

type ConfigurationData struct {
	File FileConfigurationData
}

type FileConfigurationData struct {
	FilePathType string `yaml: "filepathtype"`
	FilePath     string `yaml: "filepath"`
}

func _init() (bool, error, ConfigurationData) {
	config := ConfigurationData{}
	isAnyError, err := config._initFileConfigurationData()
	if isAnyError == true || err != nil {
		return isAnyError, err, config
	}
	return false, nil, config
}

func (Config *ConfigurationData) _initFileConfigurationData() (bool, error) {

	filePathType := os.Getenv("FilePathType")
	if filePathType == "" {
		return true, errors.New("filePathType is not set.")
	}

	filePath := os.Getenv("FilePath")
	if filePath == "" {
		return true, errors.New("filePath is not set.")
	}

	fileData := FileConfigurationData{}
	fileData.FilePath = filePathType
	fileData.FilePath = filePath
	return false, nil
}
