package main

import (
	"fmt"
	"myapp/packageone"
)

var one ="one"

func main(){
	var somethingElse = "this is a block variable"
	fmt.Println(somethingElse)

	myFunc()

	newString:=packageone.PublicVar
	fmt.Println("From packageone:",newString)

	//package.Exported is available
}

func myFunc(){
	fmt.Println(one)
}