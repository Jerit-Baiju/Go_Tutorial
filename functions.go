package main

import "fmt"

func add(num1, num2 int)  int{
	var op = num1 + num2
	return op
}

func main()  {
	fmt.Println(add(1,5))
}