package main

import (
	"fmt"

	"github.com/gosagawa/go-error-handling-sample/account"
	"github.com/gosagawa/go-error-handling-sample/errors"
)

func main() {

	user, err := account.GetUser(1)

	if err != nil {
		fmt.Println("fail!")
		if appError, ok := err.(*errors.Error); ok {
			fmt.Println(errors.Ops(appError))
		} else {
			fmt.Println(err)
		}
		return
	}
	fmt.Println("success!")
	fmt.Println(user)
}
