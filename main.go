package main

import "fmt"

func main() {
	 num1,num2 := 10,20
	c:= func (no1,no2 int)int{

		return no1+no2
}(num1,num2)
fmt.Printf("The ADDITION of two numbers %d + %d is  %d",num1,num2,c)



	d:= func(n1,n2 int)int{
		return n2-n1
	}(num1,num2)
	fmt.Printf("\n \n The SUBTRACTION of two numbers %d - %d is  %d",num2,num1,d)
}