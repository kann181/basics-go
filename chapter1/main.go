package main

import (
	"fmt"
	"time"
)

func getName() string {
	return "world!"
}

func main() {
	name := getName()
	fmt.Println("Hello ", name)

	concurrent()

}

func concurrent() {
	go count()

	time.Sleep(time.Millisecond * 2)
	fmt.Println("Hello World")
	time.Sleep(time.Millisecond * 5)
}
