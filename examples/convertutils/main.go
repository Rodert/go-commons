package main

import (
	"fmt"

	"github.com/Rodert/go-commons/convertutils"
)

func main() {
	fmt.Println("=== Go Commons Convert Utils Demo ===")
	fmt.Println()

	// 字符串转数字
	fmt.Println("字符串转数字 / String to Number:")
	str := "123"
	num, err := convertutils.StringToInt(str)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("'%s' -> %d\n", str, num)
	}

	strFloat := "123.45"
	numFloat, err := convertutils.StringToFloat64(strFloat)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("'%s' -> %f\n", strFloat, numFloat)
	}

	strBool := "true"
	b, err := convertutils.StringToBool(strBool)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("'%s' -> %v\n", strBool, b)
	}
	fmt.Println()

	// 数字转字符串
	fmt.Println("数字转字符串 / Number to String:")
	num = 123
	fmt.Printf("%d -> '%s'\n", num, convertutils.IntToString(num))

	numFloat64 := 123.456
	fmt.Printf("%f -> '%s'\n", numFloat64, convertutils.Float64ToString(numFloat64, 2))

	b = true
	fmt.Printf("%v -> '%s'\n", b, convertutils.BoolToString(b))
	fmt.Println()

	// 通用类型转换
	fmt.Println("通用类型转换 / Generic Type Conversion:")
	values := []interface{}{123, int64(456), 789.0, "999"}
	for _, v := range values {
		num, err := convertutils.ToInt(v)
		if err != nil {
			fmt.Printf("无法转换 %v (%T): %v\n", v, v, err)
		} else {
			fmt.Printf("%v (%T) -> %d\n", v, v, num)
		}
	}
	fmt.Println()

	// ToString
	fmt.Println("ToString转换 / ToString Conversion:")
	values = []interface{}{123, int64(456), 789.45, true, "hello", []byte("world")}
	for _, v := range values {
		str := convertutils.ToString(v)
		fmt.Printf("%v (%T) -> '%s'\n", v, v, str)
	}
	fmt.Println()

	// 深拷贝
	fmt.Println("深拷贝 / Deep Copy:")
	type Person struct {
		Name string
		Age  int
	}
	src := &Person{Name: "John", Age: 30}
	var dst Person
	err = convertutils.DeepCopy(src, &dst)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("源对象: %+v\n", src)
		fmt.Printf("目标对象: %+v\n", dst)
		// 修改源对象，目标对象不应受影响
		src.Name = "Jane"
		fmt.Printf("修改源对象后:\n")
		fmt.Printf("源对象: %+v\n", src)
		fmt.Printf("目标对象: %+v (应保持不变)\n", dst)
	}
}
