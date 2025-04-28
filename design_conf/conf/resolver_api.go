package conf

// IPropertyResolver 属性解析器接口
type IPropertyResolver interface {
	// ContainsProperty 判断属性源中是否包含指定键 key 的属性。
	ContainsProperty(key string) bool
	// GetProperty 属性源中获取指定键 key 对应的属性值，找不到返回nil。
	GetProperty(key string) string
	// GetPropertyAsType[T any] 属性源中获取指定键 key 对应的属性值并转换为类型T，找不到返回nil。
	// GetPropertyAsType[T any](key string) // TODO 接口方法不支持泛型，寻找其他解决方案
	// GetPropertyDefault 属性源中获取指定键 key 对应的属性值，找不到返回默认值。
	GetPropertyDefault(key string, defaultValue string) string
	// GetRequiredProperty 获取指定键 key 对应的属性值，找不到返回错误。
	GetRequiredProperty(key string) (string, error)
	//ResolvePlaceholders 解析文本中的占位符，并将其替换为实际的属性值，没有默认值的无法解析占位符将被忽略并传递。
	ResolvePlaceholders(text string) string
	// ResolveRequiredPlaceholders 解析文本中的占位符，并将其替换为实际的属性值，没有默认值的无法解析占位符将返回错误。
	ResolveRequiredPlaceholders(text string) (string, error)
}

// IEnvironment 环境接口
type IEnvironment interface {
	IPropertyResolver
	// GetActiveProfiles() []string
	GetDefaultProfiles() []string // TODO 暂无使用 暂时不实现
	// AcceptsProfiles(profiles ...string) bool
}

// IConfigurablePeopertyResolver 可配置的属性解析器接口
type IConfigurablePeopertyResolver interface {
	IPropertyResolver
	// SetConversionService(service ConversionService)
	SetPlaceholdPrefix(placeholderPrefic string)
	SetPlaceholdSuffix(placeholderSuffix string)
	SetValueSeparator(valueSeparator string)
	SetIgnoreUnresolvableNestedPlaceholders(ignoreUnresolvableNestedPlaceholders bool)
	SetRequiredProperties(requiredProperties ...string)
	ValidateRequiredProperties() error
}

// IConfigurableEnvironment 可配置的环境属性接口
type IConfigurableEnvironment interface {
	// 继承Enviroment
	IEnvironment
	// 继承ConfigurablePeopertyResolver
	IConfigurablePeopertyResolver

	// 返回Environment的属性源
	GetPropertySources() *MutablePropertySources

	// Merge 将指定的父环境合并到当前环境中
	Merge(parent IConfigurableEnvironment)

	// 获取命令行参数
	GetSystemProperties() map[string]any

	// 获取系统环境变量
	GetSystemEnvironment() map[string]any
}

// IAbstractEnvironment 抽象环境接口
type IAbstractEnvironment interface {
	IConfigurableEnvironment
	customizePropertySources(ps *MutablePropertySources)
}
