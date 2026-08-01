package BaseConfig

type ConfigurationData struct {
	File FileConfigurationData
}

type FileConfigurationData struct {
	FilePathType string `yaml: "filepathtype"`
	FilePath     string `yaml: "filepath"`
}
