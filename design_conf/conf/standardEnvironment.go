package conf

const (
	constSystemPropertiesPropertySourceName  = "systemProperties"
	constSystemEnvironmentPropertySourceName = "systemEnvironment"
)

// StandardEnvironment is a standard implementation of Environment.
type StandardEnvironment struct {
	AbstractEnvironment
}

// NewStandardEnvironment creates a new StandardEnvironment instance.
func NewStandardEnvironment() *StandardEnvironment {
	e := &StandardEnvironment{}
	e.PropertySources = NewMutablePropertySources()
	e.PropertyResolver = NewPropertySourcesPropertyResolver(e.PropertySources)
	e.customizePropertySources(e.PropertySources)
	return e
}

func (e *StandardEnvironment) customizePropertySources(ps *MutablePropertySources) {
	ps.AddLast(NewPropertiesPropertySource(constSystemPropertiesPropertySourceName, e.GetSystemProperties()))
	ps.AddLast(NewSystemEnvironmentPropertySource(constSystemEnvironmentPropertySourceName, e.GetSystemEnvironment()))
}
