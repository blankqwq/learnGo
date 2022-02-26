package main

import (
	"errors"
	"fmt"
)

type MyError struct {
}

func (e *MyError) Error() string {
	return "my Error"
}

func main() {
	e := &MyError{}
	wrapE := fmt.Errorf("err is %w", e)
	fmt.Printf("%+v\n", errors.Unwrap(wrapE))
	if e == errors.Unwrap(wrapE) {
		println("wrap success")
	}
	if errors.Is(wrapE, e) {
		println("is success")
	}
	copyE := &MyError{}
	if errors.As(errors.New("dsa"), &copyE) {
		fmt.Println("as ok")
	}
	//errors.Is()
	//errors.As()

	panic(copyE)

}
