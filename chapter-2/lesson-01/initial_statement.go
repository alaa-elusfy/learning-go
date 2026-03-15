package main

import "fmt"

func main() {
	email := "alaa"

	if length := getLength(email); length < 10 {
    		fmt.Printf("Email must be at least 10 characters, is %d\n", length)
	}
}

func getLength(email string) int{
	return len(email)
}
