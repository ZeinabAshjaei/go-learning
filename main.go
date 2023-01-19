package main

import (
	"fmt"
	"github.com/ZeinabAshjaei/helper/nameProducer"
)

func main() {
	fmt.Println("Hello", nameProducer.GetGreeting("Farshad"))
	nameProducer.GetHelloWorld()
}
