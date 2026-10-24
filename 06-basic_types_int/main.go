package main

import "fmt"

func main() {

	views1 := 1000
	views2 := 2000

	totalviews := views1 + views2

	likes := 10 //increment using ++
	likes++
	likes++

	avgViews := totalviews / 2

	fmt.Println(likes, totalviews, avgViews)

}
