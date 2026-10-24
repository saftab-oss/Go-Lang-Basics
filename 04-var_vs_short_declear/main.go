package main

import "fmt"

func main() {

	var city string
	city = "karachi"

	fmt.Println("My city is : ", city)

	var country = "pakistan" //inferred to string

	// :=
	ranking := 60
	ranking = ranking + 40 //you can add
	safety_rank, air_quality := 30, 40
	fmt.Println("The name of my city, country and ranking are : ", city, country, ranking)
	fmt.Println("Safety ranking and air quality : ", safety_rank, air_quality)
}
