package base

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/shopspring/decimal"
)

func TestFloat641(t *testing.T) {
	maxFloat64 := math.MaxFloat64
	minFloat64 := math.SmallestNonzeroFloat64

	fmt.Printf("Max Float64:%g\n", maxFloat64)
	fmt.Println("Min Float64:", minFloat64)

	a, _ := decimal.NewFromString("999999999999999.999999")
	fmt.Printf("%v", a)
}

func TestFloat642(t *testing.T) {
	strToFloatTest("1234123412341234")  // ok
	strToFloatTest("12341234123412341") // no

	strToFloatTest("12341234.12341234")  // ok
	strToFloatTest("12341234.123412341") // no
	strToFloatTest("123412341234123.41") // no

	strToFloatTest("0.1234123412341234") // ok
	strToFloatTest("1.1234123412341234") // no

	fstr := "987654321123456.9877000000000000"
	strToFloatTest(fstr)

	fstr = "123123123123123123123123123"
	strToFloatTest(fstr)

	fstr = "98765432112345.8877000000000000"
	strToFloatTest(fstr)

	fstr = "999999999999955.5559999"
	strToFloatTest(fstr)

	fstr = "999999999999999.999"
	strToFloatTest(fstr)

	fstr = "99999999999999.9"
	strToFloatTest(fstr)

	fstr = "99999999999991.9"
	strToFloatTest(fstr)

	fstr = "9999999999999999"
	strToFloatTest(fstr)

	fstr = "999999999999999"
	strToFloatTest(fstr)

	fstr = "9999999.9"
	strToFloatTest(fstr)
}

func strToFloatTest(str string) {
	floatdata, _ := strconv.ParseFloat(str, 64)
	fstr2 := fmt.Sprintf("%v", floatdata)
	fmt.Printf("o:%s\nn:%s\n", str, fstr2)

	a := decimal.NewFromFloat(floatdata)
	b, _ := decimal.NewFromString(str)

	fmt.Printf("isEqual:%t\n\n", a.Equal(b))

}

func TestIntxx(t *testing.T) {
	i := 0
	i += 1
	fmt.Printf("%d\n", i)
	i++
	fmt.Printf("%d\n", i)
}
