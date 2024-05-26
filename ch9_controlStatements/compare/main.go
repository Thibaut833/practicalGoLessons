/*
Write a program that defines two int variables ageJohn, agePaul.
Those are randomly initialized. Those two variables may vary from 0 to 110 (excluded, consequently the maximum number is 109).
ageJohn represents the age of John, agePaul the age of Paul. For example, the program outputs the following lines:
John is 42 years old
Paul is 1 years old
It is true that John is older than Paul
It is false that John and Paul have the same age
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ageJohn, agePaul := rand.Intn(110), rand.Intn(110)
	fmt.Println("John is", ageJohn, "years old")
	fmt.Println("Paul is", agePaul, "years old")
	fmt.Println("It is", ageJohn > agePaul, "that John is older than Paul")
	fmt.Println("It is", ageJohn == agePaul, "that John and Paul have the same age")
}
