package conf

import (
	"log/slog"
)

// PropertySourcesPropertyResolver 属性解析器
type PropertySourcesPropertyResolver struct {
	// 嵌入抽象实现
	AbstractPropertyResolver

	// 属性源集合
	PropertySources IPropertySources
}

// NewPropertySourcesPropertyResolver 构造函数
func NewPropertySourcesPropertyResolver(propertySources IPropertySources) *PropertySourcesPropertyResolver {
	apr := &PropertySourcesPropertyResolver{
		AbstractPropertyResolver: AbstractPropertyResolver{
			PlaceholderPrefix:                    ConstPlaceholderPrefix,
			PlaceholderSuffix:                    ConstPlaceholderSuffix,
			ValueSeparator:                       ConstValueSeparator,
			IgnoreUnresolvableNestedPlaceholders: false,
		},
	}
	// 初始化获取属性的方法
	apr.AbstractPropertyResolver.GetProperty = apr.GetProperty

	apr.PropertySources = propertySources

	return apr
}

// GetProperty impl IPropertyResolver.GetProperty 具体实现，实例化时要重写赋值给AbstractPropertyResolver.GetProperty
func (pspr *PropertySourcesPropertyResolver) GetProperty(key string) string {
	value := ""
	if pspr.PropertySources != nil {
		err := pspr.PropertySources.RangePropertySourceHandler(func(ps IPropertySource) (end bool, err error) {
			if slog.Default().Enabled(nil, slog.LevelDebug) {
				slog.Debug("Searching for key '" + key + "' in PropertySource '" + ps.GetName() + "'")
			}

			// 遍历TODO
			v := ps.GetProperty(key)
			if v != nil {
				// TODO:
				// 1. 是否需要解析嵌套占位符
				// 2. 类型转换
				// 3. 记录日志
				value = v.(string) // TODO 需要类型推断，value是否有用

				return true, nil // 已经找到，结束遍历
			}
			return false, nil // 继续遍历
		})
		if err != nil {

		}
	}
	if slog.Default().Enabled(nil, slog.LevelDebug) {
		if value == "" {
			slog.Debug("Could not find key '" + key + "' in any property source")
		}
	}

	return value
}

/*
	虽然AbstractPropertyResolver已经实现了ContainsProperty，这里仍然可以重新实现，类似于java的重写
*/
// ContainsProperty impl IPropertyResolver.ContainsProperty
// func (pspr *PropertySourcesPropertyResolver) ContainsProperty(key string) bool {
// 	v := pspr.GetProperty(key)
// 	return v != ""
// }
