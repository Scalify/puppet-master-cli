package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

// Load fetches the content of a file as a string
func Load(logger *logrus.Entry, fileName string) (string, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to load file %q: %v", fileName, err)
	}

	return string(b), nil
}

// LoadJSON loads the content of a file in the given target
func LoadJSON(logger *logrus.Entry, fileName string, target interface{}) (error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("failed to load file %q: %v", fileName, err)
	}

	if err := json.Unmarshal(b, target); err != nil {
		return fmt.Errorf("failed to decode json file %q content: %v", fileName, err)
	}

	return nil
}
