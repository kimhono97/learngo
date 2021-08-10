package lecture

import (
	"fmt"

	"github.com/kimhono97/learngo/mydict"
)

func Lec301() {
	d := mydict.Dictionary{"name": "nico", "age": "12"}
	d["hello"] = "Hello!"
	fmt.Println(d)

	data, err := d.Search("phone")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}

	err = d.Add("phone", "010-1111-1111")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("d :", d)
	}

	dd := d // dd Refers d (Not Copied)

	err = dd.Update("phone", "010-2222-2222")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("d :", d)
	}

	err = dd.Delete("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("d :", d)
	}
}
