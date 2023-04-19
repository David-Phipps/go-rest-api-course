package main

import "fmt"

// Instantiates and starts our GO app
func Run() error {
	fmt.Println("Starting our application")
	return nil
}

func main() {
	fmt.Println("GO rest API course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
