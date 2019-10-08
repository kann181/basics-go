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
}

