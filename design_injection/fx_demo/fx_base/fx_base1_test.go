package fxbase

import (
	"fmt"
	"testing"

	"go.uber.org/fx"
)

func TestFxBase(t *testing.T) {
	fx.New()

	/* ### Container ### */
	/* *** providing values *** */
	// fx.Provide(constructors ...interface{})
	// fx.Supply(values ...interface{})
	// fx.Decorate(decorators ...interface{}) // 装饰器？

	/* *** using values *** */
	// fx.Invoke(funcs ...interface{})

	/* ### Lifecycle ### */
	// fx.In
	// fx.Out
	// fx.Annotate(t interface{}, anns ...fx.Annotation)

	// fx.Private //??
	// fx.Invoke(funcs ...interface{})
	// fx.Populate(targets ...interface{}) ??

}

type S1 string
type S2 string
type M1S1 string
type M1PriS1 string
type M1S2 string
type DS1 string

func TestBase1modules(t *testing.T) {

	var m1 = fx.Module("M1",
		fx.Supply(M1S1("m1s1")),
		fx.Supply(
			fx.Private, // 标记s1是模块私有的
			M1PriS1("m1pris1"),
		),

		fx.Provide(
			func(s M1S1, ps M1PriS1) M1S2 { //ps是模块私有的能被module内部使用
				fmt.Println(s)
				return M1S2(fmt.Sprintf("%s using:%s,%s", "m1s2", s, ps))
			},
		),

		fx.Decorate(func(s M1S1) M1S1 { // 装饰器 修改原来的对象
			return M1S1(fmt.Sprintf("d1[%s]", s))
		}),

		// fx.Decorate(func(s M1S1) M1S1 {
		// 	return M1S1(fmt.Sprintf("d2[%s]", s))
		// }),

		// fx.Decorate(func(s M1S1) DS1 { // 装饰器 使用已有的对象，返回新的对象
		// 	return DS1(fmt.Sprintf("ds1[%s]", s))
		// }),

		// fx.Invoke(func(s DS1) {
		// 	fmt.Println(s)
		// }),
	)

	fx.New(

		m1,

		fx.Invoke(func(s M1S2) {
			fmt.Println(s)
		}),

		// fx.Invoke(func(s M1PriS1) { // 错误 PriS1是模块私有的，不能被外部使用
		// 	fmt.Println(s)
		// }),
	)
}

type MyString string

// 多装饰器问题
func TestMultipleDecorations(t *testing.T) {
	app := fx.New(
		fx.Supply(MyString("hello")),

		// 装饰器1：先执行（因为它在后面）
		fx.Decorate(
			func(s MyString) MyString {
				return MyString(fmt.Sprintf("decorated1(%s)", s))
			},
		),

		// 装饰器2：后执行（因为它在前）
		fx.Decorate(
			func(s MyString) MyString {
				return MyString(fmt.Sprintf("decorated2(%s)", s))
			},
		),

		fx.Invoke(func(s MyString) { // TODO 实际错误：fxbase.MyString already decorated
			fmt.Println(s) // 输出: decorated2(decorated1(hello))
		}),
	)

	if err := app.Err(); err != nil {
		t.Fatal(err)
	}
}
