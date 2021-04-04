package main

import "fmt"

// ExportFunc export function named with Capital
func ExportFunc(arg string) string {
	fmt.Println(arg)
	return arg
}

func multipleArg(args ...string) {
	fmt.Println(args)
}

func multipleReturn() (name string, age int) {
	name = "moon"
	age = 38

	return name, age
}

func deferFunc() (nakedReturnVal string) {
	// defer called when function is returned
	// maybe it used for delete resources
	defer fmt.Println("done deferFunc")

	nakedReturnVal = "return val"
	fmt.Println("deferFunc doing ")
	return
}

func forAndRange(numbers ...int) int {
	totalFor := 0
	// go supports loop logic 'for' syntax only
	for i := 0; i < len(numbers); i++ {
		totalFor += numbers[i]
	}

	// range used with for
	// it returns index & value
	totalRange := 0
	for index, num := range numbers {
		totalRange += num
		fmt.Println(index, num)
	}

	return totalRange
}

func canIDrink(age int) bool {
	// variable expression : variable lives in if context
	if koreanAge := age + 1; koreanAge < 18 {
		return false
	}

	// variable expression also can used in switch
	switch switchAge := age; switchAge {
	case 10:
		return true
	case 18:
		return false
	}

	// avoid to if else if ... else if context
	switch {
	case age > 18:
		return true
	case age > 20:
		return true
	}

	return true
}

// call by ref
func usePointer(ref *int) {
	*ref = 10
}

func useSlice() {
	sclies := []string{"moon", "60"}
	sclies = append(sclies, "catoro")
	fmt.Println(sclies)

	arrays := [3]string{"moon", "60", "catoro"}
	fmt.Println(arrays)

	// why not?
	// testArr := ['a', 'b', 'c'];
}

func useMap() {
	objs := map[string]string{"key": "value", "key2": "value2"}
	for k, v := range objs {
		fmt.Println(k, v)
	}
}

type person struct {
	name string
	age  int
}

func useStruct() {
	dbrud := person{"60", 16}
	fmt.Println(dbrud)
	moon := person{name: "moon", age: 18}
	fmt.Println(moon)
}

func testMain() {
	multipleArg("hello", "go")
	// _ ignore return value
	name, _ := multipleReturn()
	fmt.Println(name)

	total := forAndRange(1, 2, 3, 4, 5)
	fmt.Println(total)

	fmt.Println("can I drink :", canIDrink(16))

	ref := 1
	usePointer(&ref)
	fmt.Println(ref)

	useSlice()
	useMap()
	useStruct()
}
