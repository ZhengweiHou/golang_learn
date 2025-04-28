package prop

import "github.com/magiconair/properties"

// Read parses []byte in the properties format into map.
func Read(b []byte) (map[string]interface{}, error) {

	p := properties.NewProperties()
	p.DisableExpansion = true
	_ = p.Load(b, properties.UTF8) // always no error

	ret := make(map[string]interface{})
	for k, v := range p.Map() {
		ret[k] = v
	}
	return ret, nil
}
