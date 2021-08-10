package lecture

import (
	"fmt"

	"github.com/kimhono97/learngo/accounts"
)

func Lec201() {
	account := accounts.NewAccount("nico")
	account.Memo = "Hello"

	account.Deposit(10)
	fmt.Println(account.Balance())

	err := account.Withdraw(20)
	if err != nil {
		//log.Fatalln(err) // Print Error Msg and Kill Process
		fmt.Println(err)
	}
	fmt.Println(account.Balance())

	fmt.Println(account.Owner())
	account.ChangeOwner("hono")
	fmt.Println(account.Owner())

	fmt.Println(account)
}
