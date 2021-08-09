package main

import (
	"fmt"
	"strings"
)

/*
func lec101() {
	fmt.Println("Hello, World!")
	something.SayHello()
}

func lec102() {
	const name_c string = "nico" // constant
	var name_v string = "nico"   // variable 1
	name_vv := "nico"            // variable 2 (only inside function)

	name_v = "hono"
	name_vv = "omit"

	fmt.Println(name_c)
	fmt.Println(name_v)
	fmt.Println(name_vv)
}

func multiply(a int, b int) int { // 2 Inputs, 1 Output
	return a * b
}
func lenAndUpper(s string) (int, string) { // 1 Input, 2 Outputs
	return len(s), strings.ToUpper(s)
}
func repeatMe(words ...string) {
	fmt.Println(words[2])
}
func lec103() {
	fmt.Println(multiply(2, 2))

	totalLen, upperName := lenAndUpper("Heonho Kim")
	fmt.Println(totalLen)
	fmt.Println(upperName)

	repeatMe("Apple Pie", "Banana Bread", "Cup Cake", "Doughnut")
}
*/

func lenAndUpper(s string) (len_s int, upper_s string) { // Naked Returns
	len_s = len(s) // update variable
	upper_s = strings.ToUpper(s)
	return
}
func lec104() {
	totalLen, upperName := lenAndUpper("Heonho Kim")
	fmt.Println(totalLen)
	fmt.Println(upperName)
}

func main() {
	lec104()
}
