package main

// import bring external packages into the file that we are working
//where you actually need it

import (
	"fmt"
	"math"
)

func main() {

	//packagename.functionname = to call a function from package
	//so for example math.functions
	fmt.Println("Sqrt(25) :", math.Sqrt(25))
	fmt.Println("Sqrt(16): ", math.Sqrt(16))
}
