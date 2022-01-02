package main

import (
	"myapp/packageone"
)
var myVar ="This is a package level var"

func main(){
	//variables
	//declare a package level var for the main
	//package named myVar

	//declare a block variable for the main function
	//called blockVar
	var blockVar = "This is a block level var"

	//declare a package level variable in packageone
	//package named PackageVar

	//create an exported functionn in packageone named PrintMe

	//in the main function print out the values of myVar,
	//blockVar, and PackageVar in one line
	//using the PrintMe func of packageone

	
	packageone.PrintMe(myVar,blockVar)

}

