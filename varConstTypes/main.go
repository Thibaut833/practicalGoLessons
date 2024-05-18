/*
Write a program that will:
- create a string constant name hotelName and with value "Gopher Hotel"
- create two untyped constants that will contain respectively 24.806078 and - 78.243027. The names of those two constants are longitude and latitude. (somewhere in the Bahamas)
- Create a variable of type int named occupancy which is initialized with the value 12.
- Print the hotelName the longitude, the latitude and the occupancy on a new line
*/
package main

import "fmt"

func main() {
	const hotelName = "Gopher Hotel"
	const longitude, latitude = 24.806078, -78.243027
	var occupancy int = 12

	fmt.Println(hotelName, longitude, latitude)
	fmt.Println(occupancy)
}
