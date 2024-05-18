package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	if t.Hour() >= 0 && t.Hour() < 12 {
		fmt.Printf("Nous sommes le %v %v %v et il est %v:%v du matin\n", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
	} else if t.Hour() >= 12 && t.Hour() <= 18 {
		fmt.Printf("Nous sommes le %v %v %v et il est %v:%v de l'après-midi\n", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
	} else {
		fmt.Printf("Nous sommes le %v %v %v et il est %v:%v du soir\n", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
	}
}
