package main

import "fmt"

type Boat struct{}

type Transportation interface {
	Travel()
}

type BoatAdaptor struct {
	boat *Boat
}

type Client struct{}

type Car struct{}

func (m *Car) Travel() {
	fmt.Println("Car is navigating to destination")
}

func (b *Boat) Travel() {
	fmt.Println("Boat is navigating to destination")
}

func (w *BoatAdaptor) Travel() {
	fmt.Println("Adaptor to move boats on roads")
	w.boat.travelToDestination()
}

func (w *Boat) travelToDestination() {
	fmt.Println("Boat is navigating to destination")
}

func (c *Client) StartingTheJourney(com Transportation) {
	fmt.Println("Starting navigation process")
	com.Travel()
}

func main() {
	client := &Client{}
	// car := &Car{}

	// fmt.Println("Car started")
	// client.StartingTheJourney(car)

	fmt.Println("Boat started")
	boat := &Boat{}
	boatAdaptor := &BoatAdaptor{
		boat: boat,
	}
	client.StartingTheJourney(boatAdaptor)
}
