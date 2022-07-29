package main

import (
	"fmt"
	"poc-sonar/services"
)

func main() {
	message := services.GetMessage()
	fmt.Println(message)
}
