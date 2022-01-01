package main

import "fmt"

var one ="one"

func main(){
	var somethingElse = "this is a block variable"
	fmt.Println(somethingElse)

	myFunc()
}

func myFunc(){
	fmt.Println(one)
}