package json

import (
	"encoding/json"
)

// Read parses []byte in the json format into map.
func Read(b []byte) (map[string]interface{}, error) {
	var ret map[string]interface{}
	err := json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
