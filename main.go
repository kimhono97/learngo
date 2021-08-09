package main

import "fmt"

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

func lenAndUpper(s string) (len_s int, upper_s string) { // Naked Returns
	defer fmt.Println("lenAndUpper() is done.") // defer : excuted after this func is done

	len_s = len(s) // update variable
	upper_s = strings.ToUpper(s)

	return
}
func lec104() {
	totalLen, upperName := lenAndUpper("Heonho Kim")
	fmt.Println(totalLen)
	fmt.Println(upperName)
}

func superAdd(numbers ...int) (sum int) {
	sum = 0

	//for i:=0; i<len(numbers); i++ { sum += numbers[i] }

	//for index, number := range numbers
	for number := range numbers {
		sum += number
	}

	return
}
func lec105() {
	total := superAdd(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(total)
}

func canIDrink(age int) bool {
	if koreanAge := age - 2; koreanAge < 18 { // Declare Variables only for if-else
		return false
	} else if koreanAge == 18 {
		return true
	} else {
		return true
	}
}
func lec106() {
	fmt.Println(canIDrink(16))
}

func canIDrink(age int) bool {
	switch koreanAge := age - 2; { // Declare Variables only for if-else
	case koreanAge < 18:
		return false
	default:
		return true
	}
}
func lec107() {
	fmt.Println(canIDrink(16))
}

func lec108() {
	a := 2
	b := &a

	fmt.Println(&a, b)
	fmt.Println(a, *b)
}
*/

func lec109() {
	names_a := [10]string{"ApplePie", "BananaBread"} // Array
	names_a[8] = "CupCake"
	names_a[9] = "Doughnut"
	fmt.Println(names_a)

	names_s := []string{"ApplePie", "BananaBread", "CupCake"} // Slice
	names_s = append(names_s, "Doughnut")
	fmt.Println(names_s)
}

func main() {
	lec109()
}
