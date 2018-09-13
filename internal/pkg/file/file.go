package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

)

// Load fetches the content of a file as a string
// nolint: gosec
func Load(fileName string) (string, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to load file %q: %v", fileName, err)
	}

	return string(b), nil
}

// LoadJSON loads the content of a file in the given target
// nolint: gosec
func LoadJSON(fileName string, target interface{}) (error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("failed to load file %q: %v", fileName, err)
	}

	if err := json.Unmarshal(b, target); err != nil {
		return fmt.Errorf("failed to decode json file %q content: %v", fileName, err)
	}

	return nil
}
