package lecture

import (
	"fmt"
	"strings"

	"github.com/kimhono97/learngo/something"
)

func Lec101() {
	fmt.Println("Hello, World!")
	something.SayHello()
}

func Lec102() {
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
func lenAndUpper1(s string) (int, string) { // 1 Input, 2 Outputs
	return len(s), strings.ToUpper(s)
}
func repeatMe(words ...string) {
	fmt.Println(words[2])
}
func Lec103() {
	fmt.Println(multiply(2, 2))

	totalLen, upperName := lenAndUpper1("Heonho Kim")
	fmt.Println(totalLen)
	fmt.Println(upperName)

	repeatMe("Apple Pie", "Banana Bread", "Cup Cake", "Doughnut")
}

func lenAndUpper2(s string) (len_s int, upper_s string) { // Naked Returns
	defer fmt.Println("lenAndUpper() is done.") // defer : excuted after this func is done

	len_s = len(s) // update variable
	upper_s = strings.ToUpper(s)

	return
}
func Lec104() {
	totalLen, upperName := lenAndUpper2("Heonho Kim")
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
func Lec105() {
	total := superAdd(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(total)
}

func canIDrink1(age int) bool {
	if koreanAge := age - 2; koreanAge < 18 { // Declare Variables only for if-else
		return false
	} else if koreanAge == 18 {
		return true
	} else {
		return true
	}
}
func Lec106() {
	fmt.Println(canIDrink1(16))
}

func canIDrink2(age int) bool {
	switch koreanAge := age - 2; { // Declare Variables only for if-else
	case koreanAge < 18:
		return false
	default:
		return true
	}
}
func Lec107() {
	fmt.Println(canIDrink2(16))
}

func Lec108() {
	a := 2
	b := &a

	fmt.Println(&a, b)
	fmt.Println(a, *b)
}

func Lec109() {
	names_a := [10]string{"ApplePie", "BananaBread"} // Array
	names_a[8] = "CupCake"
	names_a[9] = "Doughnut"
	fmt.Println(names_a)

	names_s := []string{"ApplePie", "BananaBread", "CupCake"} // Slice
	names_s = append(names_s, "Doughnut")
	fmt.Println(names_s)
}

func Lec110() {
	map_obj := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(map_obj)
	for key, val := range map_obj {
		fmt.Println(key, val)
	}
}

type person struct {
	name    string
	age     int
	favFood []string
}

func Lec111() {
	st_obj := person{name: "nico", age: 12, favFood: []string{"kimchi", "ramen"}}
	//st_obj := person{"nico", 12, []string{"kimchi", "ramen"}}

	fmt.Println(st_obj)
}
