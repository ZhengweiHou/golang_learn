package hzwrefl

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	type1 := reflect.TypeOf("123")
	value1 := reflect.ValueOf("123")

	fmt.Println(type1, value1)
	fmt.Println(type1.Kind())

	var c interface{}
	c = "hzw"
	fmt.Println(c)
	// c.(type1.Kind())
}

func Test2(t *testing.T) {
	// 创建一个实现接口的结构体实例
	myStructInstance := MyStruct{
		Id:   1,
		Name: "hzw"}

	// 将结构体实例转换为接口类型
	myInterfaceInstance := interface{}(&myStructInstance)

	// 使用反射获取接口的类型
	interfaceType := reflect.TypeOf(myInterfaceInstance)

	// 使用反射创建结构体实例的指针
	structPtr := reflect.New(interfaceType.Elem())

	// 使用反射将接口实例转换为结构体实例
	structPtr.Elem().Set(reflect.ValueOf(&myStructInstance).Elem())

	// 使用反射调用结构体的方法
	method := structPtr.MethodByName("MyMethod")
	method.Call(nil)

	// 将反射对象还原为原始结构体类型
	// originalStruct := structPtr.Interface().(MyStruct)

	// 将反射对象还原为原始结构体类型（通过 Elem() 获取指针指向的值）
	originalStruct := structPtr.Elem().Interface().(MyStruct)

	// 直接访问结构体的字段
	fmt.Println("Name:", originalStruct.Name)

}

func Test3(t *testing.T) {
	// 创建一个实现接口的结构体实例
	myStructInstance := &MyStruct{Id: 1, Name: "hzw"}

	// 将结构体实例转换为接口类型
	var ms MyInterface = myStructInstance

	// 使用反射获取接口的值
	value := reflect.ValueOf(ms)

	// 确保接口持有的是指针类型
	if value.Kind() == reflect.Ptr {
		// 获取指针指向的结构体值
		structValue := value.Elem()

		// 将反射对象还原为原始结构体类型
		originalStruct := structValue.Interface().(MyStruct)

		// 直接访问结构体的字段
		fmt.Println("Name:", originalStruct.Name)
	} else {
		fmt.Println("Not a pointer to a struct")
	}

}

func Test4(t *testing.T) {
	// 创建一个 MyStruct 类型的实例
	myv := MyStruct{Id: 1, Name: "hzw"}
	// 创建一个指向 MyStruct 类型实例的指针
	myp := &MyStruct{Id: 1, Name: "hzw"}
	// 将实例和指针转换为空接口类型
	amyv := any(myv)
	amyp := any(myp)

	// 使用反射获取值 reflect.ValueOf()返回的是 reflect.Value 类型
	vv := reflect.ValueOf(myv)
	pv := reflect.ValueOf(myp)
	avv := reflect.ValueOf(amyv)
	apv := reflect.ValueOf(amyp)
	// 使用反射获取类型 reflect.TypeOf()返回的是 *reflect.Type 类型
	vt := reflect.TypeOf(myv) // 获取的值为目标的类型
	pt := reflect.TypeOf(myp)
	avt := reflect.TypeOf(amyv)
	apt := reflect.TypeOf(amyp)
	fmt.Printf("myv.T %T = vt.v %v\n", myv, vt) // myv.T hzwrefl.MyStruct = vt.v hzwrefl.MyStruct 值的type 和reflect.TypeOf获取的值是一样的

	fmt.Printf("myv: %v %T \n", myv, myv) // myv: {1 hzw} hzwrefl.MyStruct
	fmt.Printf("myp: %v %T \n", myp, myp) // myp: &{1 hzw} *hzwrefl.MyStruct
	fmt.Printf("vv: %v %T \n", vv, vv)    // vv: {1 hzw} reflect.Value
	fmt.Printf("pv: %v %T \n", pv, pv)    // pv: &{1 hzw} reflect.Value
	fmt.Printf("avv: %v %T \n", avv, avv) // avv: {1 hzw} reflect.Value
	fmt.Printf("apv: %v %T \n", apv, apv) // apv: &{1 hzw} reflect.Value
	fmt.Printf("vt: %v %T \n", vt, vt)    // vt: hzwrefl.MyStruct *reflect.rtype
	fmt.Printf("pt: %v %T \n", pt, pt)    // pt: *hzwrefl.MyStruct *reflect.rtype
	fmt.Printf("avt: %v %T \n", avt, avt) // avt: hzwrefl.MyStruct *reflect.rtype
	fmt.Printf("apt: %v %T \n", apt, apt) // apt: *hzwrefl.MyStruct *reflect.rtype

	/* ============使用=============*/
	/* ------------Kind------------*/
	// 获取反射值的种类并打印 value 和 type 获取的kind一样
	vvkind := vv.Kind()
	fmt.Printf("vvkind: %v %T \n", vvkind, vvkind) // vvkind: struct reflect.Kind
	pvkind := pv.Kind()
	fmt.Printf("pvkind: %v %T \n", pvkind, pvkind) // pvkind: ptr reflect.Kind
	vtkind := vt.Kind()
	fmt.Printf("vtkind: %v %T \n", vtkind, vtkind) // vtkind: struct reflect.Kind
	ptkind := pt.Kind()
	fmt.Printf("ptkind: %v %T \n", ptkind, ptkind) // ptkind: ptr reflect.Kind

	/* -----------value.Interface 获取原始值-----------*/
	//reflect.ValueOf()返回的是 reflect.Value 类型，其中包含了原始值的值信息。reflect.Value 与原始值之间可以互相转换
	vvi := vv.Interface()                 // 因传入的是struct，所以将反射值转换为any
	fmt.Printf("vvi: %v %T \n", vvi, vvi) // vvi: {1 hzw} hzwrefl.MyStruct
	pvi := pv.Interface()
	fmt.Printf("pvi: %v %T \n", pvi, pvi) // pvi: &{1 hzw} *hzwrefl.MyStruct

	vvf0 := vv.Field(0)                           // 通过value获取第一个属性的value
	fmt.Printf("vvf0: %v %T \n", vvf0, vvf0)      // vvf0: 1 reflect.Value
	fmt.Printf("vvf0.canset:%v\n", vvf0.CanSet()) // vvf0.canset:false 属性不可设置
	// vvf0.SetInt(5)                           // vv 是结构体的副本，非原始变量，从副本获取的字段同样不可设置

	/* -----------value.Elem 获取原始值-----------
	解引用指针类型的反射值，获取指针指向的内容。若 pv 非指针，调用会 panic。
	和value.Interface()获取原始值不同，value.Elem()返回的是 reflect.Value 类型，其中包含了原始值的值信息。
	*/
	// vv 和 pv 分别不可直接调用Elem和Field
	// vv.Elem()   // vv不是指针类型，所以不能调用Elem获取指针内容，会panic
	// pv.Field(0) // pv是指向结构体的指针，只有类型为结构体时才可以调用Field方法，会panic
	pve := pv.Elem()
	fmt.Printf("pve: %v %T \n", pve, pve) // pve: {1 hzw} reflect.Value
	pvef0 := pve.Field(0)
	fmt.Printf("pvef0: %v %T \n", pvef0, pvef0)     // pvef0: 1 reflect.Value
	fmt.Printf("pvef0.canset:%v\n", pvef0.CanSet()) // pvef0.canset:true 属性可设置
	fmt.Printf("vvf0.canset:%v\n", vvf0.CanSet())   // vvf0.canset:false 属性不可设置

	pvef0.SetInt(5)                                  // 修改属性值，因该属性是指针类型，所以不会panic
	fmt.Printf("pvef0设置后,原始指针指向的值也被修改了: %v \n", myp) // pvef0设置后,原始指针指向的值也被修改了: &{5 hzw}

	// reflect.TypeOf()返回的是 *reflect.Type 类型, 包含反射对象信息，可以获得方法和属性等信息
	numf := vt.NumField()           // 获得属性数量
	fmt.Printf("numf: %v \n", numf) // numf: 2
	vtf0 := vt.Field(0)             // 获得第一个属性，获取StructField类型
	// >>value 和 type 获取Filed的区别<<
	fmt.Printf("vtf0: %v %T \n", vtf0, vtf0) // vtf0: {Id  int  0 [0] false} reflect.StructField
	fmt.Printf("vvf0: %v %T \n", vvf0, vvf0) // vvf0: 1 reflect.Value

	// vtf0.Tag() // 获取属性的tag

	// f0json, _ := json.Marshal(f0)
	// fmt.Printf("f0: %s \n", f0json) // f0: {"Name":"Id","PkgPath":"","Type":{},"Tag":"","Offset":0,"Index":[0],"Anonymous":false}

	numm := vt.NumMethod()          // 获得方法数量
	fmt.Printf("numm: %v \n", numm) // numm: 2
	m1 := vt.Method(1)              // 获得第二个方法，获取Method类型
	m1json, _ := json.Marshal(m1)
	fmt.Printf("m1: %s \n", m1json) // m1: {"Name":"MyMethod2","PkgPath":"","Type":{},"Func":{},"Index":1}

	/* ----------type.Elem 获取元素类型-----------*/
	// 若type 不是Array, Chan, Map, Pointer, Slice类型，会panic
	pte := pt.Elem()
	fmt.Printf("pt: %v %T kind:%v \n", pt, pt, pt.Kind())     // pt: *hzwrefl.MyStruct *reflect.rtype kind:ptr
	fmt.Printf("pte: %v %T kind:%v \n", pte, pte, pte.Kind()) // pte: hzwrefl.MyStruct *reflect.rtype kind:struct

}

func Test5(t *testing.T) {
	// 创建一个 MyStruct 类型的实例
	mv := MyStruct{Id: 1, Name: "hzw"}
	// 创建一个 MyStruct 类型的指针
	mp := &MyStruct{Id: 1, Name: "hzw"}

	// reflect.ValueOf()
	vv := reflect.ValueOf(mv)
	pv := reflect.ValueOf(mp)
	// reflect.TypeOf()
	vt := reflect.TypeOf(mv)
	pt := reflect.TypeOf(mp)

	vvf := vv.Field(0)
	fmt.Printf("vvf canset:%v %v %T \n", vvf.CanSet(), vvf, vvf)
	// pv.Field(0) // panic: reflect: call of reflect.Value.Field on ptr Value
	pvef := pv.Elem().Field(0)
	fmt.Printf("pvef canset:%v %v %T \n", pvef.CanSet(), pvef, pvef)
	vtf := vt.Field(0)
	fmt.Printf("vtf %v %T \n", vtf, vtf)
	// pt.Field(0) // panic: reflect: Field of non-struct type *hzwrefl.MyStruct
	ptef := pt.Elem().Field(0)
	fmt.Printf("ptef %v %T \n", ptef, ptef)
	/*
		vvf canset:false 1 reflect.Value
		pvef canset:true 1 reflect.Value
		vtf {Id  int  0 [0] false} reflect.StructField
		ptef {Id  int  0 [0] false} reflect.StructField
	*/
}

func TestSlice1(t *testing.T) {
	ms := &MyStruct{Id: 1, Name: "hzw"}
	mt := reflect.TypeOf(ms)
	mte := mt.Elem()

	// 创建元素类型为MyStruct的slice
	sliceType := reflect.SliceOf(mte)
	slice := reflect.MakeSlice(sliceType, 0, 0)

	// 创建并添加第一个元素
	elem1 := reflect.New(mte).Elem()
	elem1.FieldByName("Id").SetInt(1)
	elem1.FieldByName("Name").SetString("hzw")
	slice = reflect.Append(slice, elem1)

	// 创建并添加第二个元素
	elem2 := reflect.New(mte).Elem()
	elem2.FieldByName("Id").SetInt(2)
	elem2.FieldByName("Name").SetString("test")
	slice = reflect.Append(slice, elem2)

	// 验证结果
	if slice.Len() != 2 {
		t.Errorf("Expected slice length 2, got %d", slice.Len())
	}

	firstElem := slice.Index(0).Interface().(MyStruct)
	if firstElem.Id != 1 || firstElem.Name != "hzw" {
		t.Errorf("First element mismatch, got %v", firstElem)
	}

	secondElem := slice.Index(1).Interface().(MyStruct)
	if secondElem.Id != 2 || secondElem.Name != "test" {
		t.Errorf("Second element mismatch, got %v", secondElem)
	}

	fmt.Printf("slice: %v %T \n", slice, slice)
}

func TestSlice2(t *testing.T) {
	// 创建一个初始slice
	slice := make([]MyStruct, 2)

	// 获取slice的反射值
	sliceValue := reflect.ValueOf(&slice).Elem()

	sve := sliceValue.Type().Elem()

	// 创建第一个元素并设置
	elem1 := reflect.New(sve).Elem()
	elem1.FieldByName("Id").SetInt(10)
	elem1.FieldByName("Name").SetString("element1")
	sliceValue.Index(0).Set(elem1)

	// 创建第二个元素并设置
	elem2 := reflect.New(sliceValue.Type().Elem()).Elem()
	elem2.FieldByName("Id").SetInt(20)
	elem2.FieldByName("Name").SetString("element2")
	sliceValue.Index(1).Set(elem2)

	// 验证结果
	if len(slice) != 2 {
		t.Errorf("Expected slice length 2, got %d", len(slice))
	}

	if slice[0].Id != 10 || slice[0].Name != "element1" {
		t.Errorf("First element mismatch, got %v", slice[0])
	}

	if slice[1].Id != 20 || slice[1].Name != "element2" {
		t.Errorf("Second element mismatch, got %v", slice[1])
	}

	fmt.Printf("slice: %v %T \n", slice, slice)
}

func TestSlice3(t *testing.T) {
	// 定义一个空的MyStruct切片
	// var slice []MyStruct
	slice := make([]MyStruct, 2)

	// 获取切片的反射值
	sliceValue := reflect.ValueOf(&slice)
	if sliceValue.Kind() == reflect.Ptr {
		// 如果是指针类型，获取其元素
		sliceValue = sliceValue.Elem()
	}

	// 获取切片的类型
	sliceType := sliceValue.Type()

	// 获取切片元素的类型
	myType := sliceType.Elem()

	// 创建一个新的空切片
	rslice := reflect.MakeSlice(sliceType, 0, 0)

	// 创建一个新的MyStruct实例并设置其字段值
	myValue := reflect.New(myType).Elem()
	myValue.FieldByName("Id").SetInt(1)
	myValue.FieldByName("Name").SetString("hzw1")

	// 将新创建的MyStruct实例追加到切片中
	rslice = reflect.Append(rslice, myValue)

	myValue.FieldByName("Id").SetInt(22)
	rslice = reflect.Append(rslice, myValue)

	// 将新切片设置回原切片
	sliceValue.Set(rslice)
	// 打印切片的值和类型
	fmt.Printf("slice: %v %T \n", slice, slice)

	// 取出rslice
	// slice2 := rslice.Interface().([]MyStruct)
	slice2 := rslice.Interface()
	fmt.Printf("slice2: %v %T \n", slice2, slice2)

}

type MyInterface interface {
	MyMethod()
}

type MyStruct struct {
	Id   int
	Name string
}

func (s MyStruct) MyMethod() {
	fmt.Println("Hello from MyMethod")
}

func (s MyStruct) MyMethod2(msg string) string {
	fmt.Println("Hello from MyMethod2")
	return fmt.Sprintf("name:%s msg:%s", s.Name, msg)
}
