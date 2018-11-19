package main

import "fmt"

func main() {
	 num1,num2 := 10,20
	c:= func (no1,no2 int)int{

		return no1+no2
}(num1,num2)
fmt.Printf("Thw adtition of two numbers %d + %d is  %d",num1,num2,c)
}