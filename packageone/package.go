package packageone

import "fmt"

var PackageVar = "This is a package var in packageone"

func PrintMe(myVar, blockVar string) {
	fmt.Println(myVar, blockVar, PackageVar)
}
