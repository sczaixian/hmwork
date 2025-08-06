package util

import "fmt"

type T any

func Print(user T, err error) {
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(user)
}
