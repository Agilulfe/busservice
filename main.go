package main

import (
	"fmt"

	"github.com/Agilulfe/busservice/service"
)

func main() {
	fmt.Println("Starting simulation")
	expressLine := service.NewBus("Express Line")

	s1 := service.BusStop{Name: "Downtown"}
	s2 := service.BusStop{Name: "The University"}
	s3 := service.BusStop{Name: "The Village"}

	expressLine.AddStop(&s1)
	expressLine.AddStop(&s2)
	expressLine.AddStop(&s3)

	john := service.Prospect{
		SSN:         "12345612-22",
		Destination: &s2,
	}
	betty := service.Prospect{
		SSN:         "11223322-67",
		Destination: &s3,
	}
	s1.NotifyProspectArrival(john)
	s1.NotifyProspectArrival(betty)

	for expressLine.Go() {
		expressLine.VisitPassengers(func(p service.Passenger) {
			fmt.Printf("    Passenger with SSN %q is heading to %q\n", p.SSN, p.Destination.Name)
		})
	}
	fmt.Println("Simulation done")
}
