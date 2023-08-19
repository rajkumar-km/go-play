package main

import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// DemoDecoder is an example to parse a key-value config data into a struct
func DemoDecoder() {
	// Assume we have a key-value pairs in
	configFileData := `
	ip=172.0.1.10
	port=1234
	enabled=true
	`
	type ServerConfig struct {
		ServerIp   string `json:"ip,omitempty"`
		ServerPort int    `json:"port,omitempty"`
		IsEnabled  bool   `json:"enabled,omitempty"`
	}

	// First, convert the raw data into key-value map
	var configMap map[string]string = parseKeyValues([]byte(configFileData))

	// Next, use reflection to populate any struct
	var server ServerConfig
	parseServerConfig(configMap, &server)
	fmt.Println("Parsed serverConfig =")
	inspect(reflect.ValueOf(server))
}

// parseKeyValues parse the key-value pairs from a raw data to map
func parseKeyValues(data []byte) map[string]string {
	out := make(map[string]string)
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "=")
		if len(tokens) == 2 {
			key := strings.TrimSpace(tokens[0])
			val := strings.TrimSpace(tokens[1])
			out[key] = val
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("scanner error:", err)
	}

	return out
}

// parseServerConfig parses the key-value pair data into a struct without knowing
// the representation.
func parseServerConfig(configMap map[string]string, config interface{}) error {
	// Ensure that a pointer to the struct is provided so that it is addressable
	v := reflect.ValueOf(config)
	if v.Kind() != reflect.Ptr && v.Kind() != reflect.Interface {
		return fmt.Errorf("%s may be not addressable?", v.Kind())
	}

	// Dereference the pointer and obtain the addressable value
	pv := v.Elem()

	// Ensure that the value is addressable and is a struct type
	if !pv.CanSet() {
		return fmt.Errorf("%s is not settable\n", pv.Kind())
	}
	if pv.Kind() != reflect.Struct {
		return fmt.Errorf("parsing not supported for %s type", pv.Kind())
	}

	// Iterate the fields in the struct and set the values from configMap
	for i := 0; i < pv.NumField(); i++ {
		// Get the field name from json tag if available, use struct field name otherwise.
		fieldInfo := pv.Type().Field(i)
		jsonTag := fieldInfo.Tag.Get("json")
		fieldName := strings.Split(jsonTag, ",")[0]
		if len(fieldName) == 0 {
			fieldName = strings.ToLower(fieldInfo.Name)
		}
		fieldVal := pv.Field(i)

		switch fieldVal.Kind() {
		case reflect.Bool:
			b, err := strconv.ParseBool(configMap[fieldName])
			if err != nil {
				return err
			}
			fieldVal.SetBool(b)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			i, err := strconv.ParseInt(configMap[fieldName], 10, 0)
			if err != nil {
				return err
			}
			// But, the below code works for all signed int types such as int8, int16, and etc.,
			fieldVal.SetInt(i)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			ui, err := strconv.ParseUint(configMap[fieldName], 10, 0)
			if err != nil {
				return err
			}
			fieldVal.SetUint(ui)
		case reflect.Float32, reflect.Float64:
			f, err := strconv.ParseFloat(configMap[fieldName], 64)
			if err != nil {
				return err
			}
			fieldVal.SetFloat(f)
		case reflect.String:
			fieldVal.SetString(configMap[fieldName])
		default:
			return fmt.Errorf("%s is unsupported", fieldVal.Kind())
		}
	}
	return nil
}
