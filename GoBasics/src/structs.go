package main

import (
	"fmt"
)

type bikeRide struct { // structs are like types in typescript
	dist     float32
	hours    float32
	avgSpeed float32
}

// type swim struct {
// 	hundredMeters float32
// 	hours         float32
// 	avgSpeed      float32
// }

func (r *bikeRide) getAvgSpeed() float32 {
	r.avgSpeed = r.dist / r.hours
	return r.dist / r.hours
}

// func (r *swim) getAvgSpeed() float32 {
// 	r.avgSpeed = r.hundredMeters / r.hours
// 	return r.hundredMeters / r.hours
// }

type trip interface {
	getAvgSpeed() float32
}

func acceptableSpeed(t trip) {
	fmt.Println(t.getAvgSpeed())
	fmt.Println(t)
	// if r > 20 {
	// 	fmt.Println("Your speed is good!")
	// } else {
	// 	fmt.Println("Go FASTER!!")
	// }
}

func structs() {
	var myBikeRide *bikeRide = &bikeRide{75, 2.6, 0}
	acceptableSpeed(myBikeRide)
}
