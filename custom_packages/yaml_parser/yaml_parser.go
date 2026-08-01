package yaml_parser

import (
	Configurator "Universal-Rate-Limiter/src/BaseConfig"
	RateConfigurator "Universal-Rate-Limiter/src/ConfigReader"
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type YamlParser struct {
	filePathType string
	filePath     string
}

/*
_init() :	This function will initialise the YamlParser struct.

	This function takes pointer to FileConfigurationData as an argument.
	performs validation on filepathtype and filepath variables.
*/
func _init(config *Configurator.FileConfigurationData) (bool, error, YamlParser) {
	parser := YamlParser{}

	cleanFilePathType := strings.ReplaceAll(config.FilePathType, " ", "")
	cleanFilePath := strings.ReplaceAll(config.FilePath, " ", "")

	if len(cleanFilePathType) == 0 {
		return true, errors.New("File Path type is empty."), parser
	}
	if len(cleanFilePath) == 0 {
		return true, errors.New("File Path is empty."), parser
	}

	parser.filePath = cleanFilePath
	parser.filePathType = cleanFilePathType

	return false, nil, parser
}

/*
ReadFile() :	This function will based on the data in the YamlParser object

	will parse the file and populate the struptr object.
*/
func (Par *YamlParser) ReadFile(StruPtr *RateConfigurator.RateLimiterConfiguration) (bool, error) {
	structPointer := reflect.ValueOf(StruPtr)
	if structPointer.Kind() != reflect.Pointer || structPointer.IsNil() {
		return true, errors.New("must pass a non-nil pointer to a struct")
	}

	switch Par.filePathType {
	case "Local":
		yamlData, err := os.ReadFile(Par.filePath)
		if err != nil {
			return true, errors.New("Error occured while opening the file. Error: " + err.Error())
		}

		err = yaml.Unmarshal(yamlData, StruPtr)
		isAnyError, err := Par.unmarshal(yamlData, StruPtr)
		if isAnyError == true || err != nil {
			return true, errors.New("Error occured while reading the file. Error: " + err.Error())
		}

		return false, nil
	}
	return true, errors.New("Something has gone wrong while opening/reading the file.")
}

/*
	Unmarshal() :	This is custom unmarshaler which will unmarshal the yaml files.
*/

func (Par *YamlParser) unmarshal(data []byte, store *RateConfigurator.RateLimiterConfiguration) (bool, error) {
	lineBytes := [][]byte{}
	temp := []byte{}
	var newlineByte byte = '\n'
	var curr byte
	for i := 0; i < len(data); i++ {
		curr = data[i]
		if curr == newlineByte {
			lineBytes = append(lineBytes, temp)
			temp = []byte{}
		} else {
			temp = append(temp, curr)
		}
	}
	isAnyError, err := Par.rateLimiterConfigurationParser(lineBytes, store)
	if isAnyError == true || err != nil {
		return isAnyError, err
	}
	return false, err
}

/*
	rateLimiterConfigurationParser() :	This function will parse the [][]byte and return RateLimiterConfiguration
*/

func (Par *YamlParser) rateLimiterConfigurationParser(data [][]byte, store *RateConfigurator.RateLimiterConfiguration) (bool, error) {

	isAnyError, err, uniVal := Par.stringParser(data[0])
	if isAnyError == true || err != nil {
		return isAnyError, err
	}
	store.Domain = uniVal

	descriptors := []RateConfigurator.Descriptor{}
	for i := 2; i < len(data)-5; i = i + 7 {
		descriptor := RateConfigurator.Descriptor{}
		descriptorData := data[(i-1):(i + 6)]
		isAnyError, err := Par.descriptorParser(descriptorData, &descriptor)
		if isAnyError == true || err != nil {
			store.Descriptors = descriptors
			return isAnyError, err
		}
		descriptors = append(descriptors, descriptor)
	}
	store.Descriptors = descriptors
	return false, nil
}

/*
	descriptorParser() :	This function will parse the [][]byte and return Descriptor
*/
func (Par *YamlParser) descriptorParser(data [][]byte, store *RateConfigurator.Descriptor) (bool, error) {
	// Key
	isAnyError, err, uniVal := Par.stringParser(data[0])
	if isAnyError == true || err != nil {
		return isAnyError, err
	}
	store.Key = uniVal

	// algorithm
	isAnyError, err, uniVal = Par.stringParser(data[1])
	if isAnyError == true || err != nil {
		return isAnyError, err
	}
	store.Algorithm = uniVal

	// Value
	isAnyError, err, uniVal = Par.stringParser(data[2])
	if isAnyError == true || err != nil {
		return isAnyError, err
	}
	store.Value = uniVal

	// Rate Limit
	rateLimit := RateConfigurator.RateLimit{}
	isAnyError, err = Par.rateLimitParser(data[3:5], &rateLimit)
	if isAnyError == true || err != nil {
		return isAnyError, err
	}
	store.Rate_Limit = rateLimit
	return false, nil
}

/*
	rateLimitParser() :	This function will parse the [][]byte and return RateLimit
*/
func (Par *YamlParser) rateLimitParser(data [][]byte, store *RateConfigurator.RateLimit) (bool, error) {
	isAnyError, err, uniVal := Par.stringParser(data[0])
	if isAnyError == true || err != nil {
		return isAnyError, err
	}
	store.Unit = uniVal

	isAnyError, err, Val := Par.int64Parser(data[1])
	if isAnyError == true || err != nil {
		return isAnyError, err
	}
	store.RequestsPerUnit = Val
	return false, nil
}


/*
	stringParser() :	This function will parse the [][]byte and return string
*/
func (Par *YamlParser) stringParser(data []byte) (bool, error, string) {
	var spaceByte byte = ' '    // Standard space (ASCII 32)
	var tabByte byte = '\t'     // Horizontal tab (ASCII 9)
	var newlineByte byte = '\n' // Newline / Line feed (ASCII 10)
	var carriageReturn byte = '\r'
	var colonByte byte = ':'
	temp := []byte{}
	isColonFound := false
	var currUni byte
	for i := 0; i < len(data); i++ {
		currUni = data[i]
		if isColonFound == true {
			if currUni == spaceByte ||
				currUni == tabByte ||
				currUni == newlineByte ||
				currUni == carriageReturn {
			} else {
				temp = append(temp, currUni)
			}
		} else {
			if currUni == colonByte {
				isColonFound = true
			}
		}
	}
	return false, nil, string(temp)
}

/*
	stringParser() :	This function will parse the [][]byte and return int64
*/
func (Par *YamlParser) int64Parser(data []byte) (bool, error, int64) {
	var spaceByte byte = ' '    // Standard space (ASCII 32)
	var tabByte byte = '\t'     // Horizontal tab (ASCII 9)
	var newlineByte byte = '\n' // Newline / Line feed (ASCII 10)
	var carriageReturn byte = '\r'
	var colonByte byte = ':'
	temp := []byte{}
	isColonFound := false
	var currUni byte
	for i := 0; i < len(data); i++ {
		currUni = data[i]
		if isColonFound == true {
			if currUni == spaceByte ||
				currUni == tabByte ||
				currUni == newlineByte ||
				currUni == carriageReturn {
			} else {
				temp = append(temp, currUni)
			}
		} else {
			if currUni == colonByte {
				isColonFound = true
			}
		}
	}
	str := string(temp)
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return true, err, 0
	}
	return false, nil, num
}
