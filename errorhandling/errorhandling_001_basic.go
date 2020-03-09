package main

import (
	"errors"
	"fmt"
)

func status(status int) (ret bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	if status > 5 {
		panic(errors.New("WHOA Status > 5"))
	}
	return true, nil
}

func main() {
	for i := 0; i < 10; i++ {
		_, err := status(i)
		if err != nil {
			fmt.Println("Shit", i, err)
		} else {
			fmt.Println("Alles super", i, err)
		}
	}
}
