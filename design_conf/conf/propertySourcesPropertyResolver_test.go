package conf

import "testing"

func TestPropertySourcesPropertyResolver(t *testing.T) {
	pspr := NewPropertySourcesPropertyResolver(nil)

	pspr.GetProperty("aaa")
	pspr.GetProperty("bbb")
	pspr.GetPropertyDefault("ccc", "CCC")

	pspr.ContainsProperty("key")

}
