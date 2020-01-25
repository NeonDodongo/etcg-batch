package env

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
)

// LoadConfig will load settings from the environment into the application
func LoadConfig(t interface{}, path string) error {
	if v := reflect.ValueOf(t); v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("Target for config loading must be a pointer")
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Failed to read config file: %s [ %w ]", path, err)
	}

	if err := json.Unmarshal(content, t); err != nil {
		return fmt.Errorf("Failed to unmarshal config file contents [ %w ]", err)
	}

	return nil
}
