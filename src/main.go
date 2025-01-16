package main

import "fmt"

func Run() error {
	fmt.Println("Hello, World!")
	return nil
}

func main() {
	if err := Run(); err != nil {
		panic(err)
	}
}
