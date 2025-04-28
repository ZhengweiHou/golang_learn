package conf

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestStandardEnvironment(t *testing.T) {
	senv := &StandardEnvironment{}
	assert.Equal(t, 0, len(senv.GetSystemProperties()))

	// 子类无法直接作为父类使用
	doSomethingWithAbstractEnvironment(senv.AbstractEnvironment)
	// 可以作为父接口使用
	doSomethingWithConfigurableEnvironment(senv)

	var i interface{} = senv
	e, ok := i.(StandardEnvironment)

	// 使用类型断言判断是否为 AbstractEnvironment 类型
	if ok {
		fmt.Printf("It is an AbstractEnvironment with propertySources address: %p\n", e)
	} else {
		fmt.Println("It is not an AbstractEnvironment")
	}

	// 使用类型断言判断是否为 ConfigurableEnvironment 类型
	if _, ok := i.(IConfigurableEnvironment); ok {
		fmt.Println("It is a ConfigurableEnvironment")
	}

}

func doSomethingWithAbstractEnvironment(env AbstractEnvironment) {
	return
}

func doSomethingWithConfigurableEnvironment(env IConfigurableEnvironment) {
	return
}
