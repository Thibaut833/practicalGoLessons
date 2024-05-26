/*
We will build a hotel application. We will begin with the front desk program used by receptionists. At the front desk, the receptionist will welcome customers. Clients will ask the receptionist for an available room. He needs to know :

    How many rooms are available now

    For how long each room is available

We will write a program that will do that based on fake data. In figure 11 you can see the wireframes of the application.

Note that we have presented the same application in two different states.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const nameHotel = "Gopher Paris Inn"
	const numberRooms = 134

	rand.New(rand.NewSource(time.Now().UnixMicro()))
	availableRooms := rand.Int31n(numberRooms)
	numberRoomsOccupied := numberRooms - availableRooms
	occupancyRate := 100 - uint32((float32(numberRoomsOccupied)/float32(numberRooms))*100)

	fmt.Println("\nHotel : ", nameHotel)
	switch {
	case occupancyRate >= 60:
		fmt.Println("				Occupancy level : high")
	case occupancyRate >= 30 && occupancyRate < 60:
		fmt.Println("				Occupancy level : medium")
	case occupancyRate >= 0 && occupancyRate < 30:
		fmt.Println("				Occupancy level : low")
	}
	fmt.Println("				Occupancy rate : ", occupancyRate)
	fmt.Println("Number of rooms :", numberRooms)
	fmt.Println("Rooms available :", availableRooms)

	if availableRooms > 0 {
		fmt.Println("\nRooms :")
		for i := 0; i < int(availableRooms); i++ {
			randPeople := rand.Int31n(10)
			randNight := rand.Int31n(20)
			room := i + 110
			fmt.Println("	- ", room, " : ", randPeople, " people / ", randNight, " nights")
		}
	} else {
		fmt.Println("\nNo rooms available for tonight")
	}
}
