package something

import "fmt"

func sayBye() { // No Export (Private)
	fmt.Println("Bye")
}

func SayHello() { // Export Public
	fmt.Println("Hello")
}
