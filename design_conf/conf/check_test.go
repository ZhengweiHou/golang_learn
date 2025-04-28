package conf

import "testing"

// 用于测试类型的接口实现情况

var _PropertyResolver IPropertyResolver
var _Environment IEnvironment
var _ConfigurablePeopertyResolver IConfigurablePeopertyResolver
var _ConfigurableEnvironment IConfigurableEnvironment
var _IAbstractEnvironment IAbstractEnvironment

var _AbstractEnvironment *AbstractEnvironment
var _StandardEnvironment *StandardEnvironment
var _AbstractPropertyResolver *AbstractPropertyResolver

func TestCheck(t *testing.T) {
	_PropertyResolver = _Environment
	_PropertyResolver = _ConfigurablePeopertyResolver
	_PropertyResolver = _ConfigurableEnvironment
	_Environment = _ConfigurableEnvironment
	_ConfigurablePeopertyResolver = _ConfigurableEnvironment

	_ConfigurableEnvironment = _AbstractEnvironment
	// _AbstractEnvironment = _StandardEnvironment // 子实现无法赋值给父struct
	// _IAbstractEnvironment = _AbstractEnvironment // 使用一个接口来约束, AbstractEnvironment 不真正实现customizePropertySources方法
	_IAbstractEnvironment = _StandardEnvironment // 子实现满足同样的接口约束

}

// sources 部分类型测试
var _IPropertySources IPropertySources
var _MutablePropertySources *MutablePropertySources

var _IPropertySource IPropertySource
var _PropertySource *MapPropertySource

func TestSourcesCheck(t *testing.T) {
	_IPropertySource = _PropertySource
	_IPropertySources = _MutablePropertySources
}
