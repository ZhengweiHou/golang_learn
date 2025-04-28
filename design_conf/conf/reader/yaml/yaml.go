package yaml

import (
	"gopkg.in/yaml.v2"
)

// Read parses []byte in the yaml format into map.
func Read(b []byte) (map[string]interface{}, error) {
	ret := make(map[string]interface{})
	err := yaml.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
