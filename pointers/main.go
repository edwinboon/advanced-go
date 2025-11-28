package main

import (
	"fmt"
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

type ElectricTruck struct {
	id           string
	cargo        int
	batteryLevel float64
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 100
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 100
	e.batteryLevel -= 10.0
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.batteryLevel -= 5.0
	return nil
}

// processTruck handles the loading and unloading of a truck.
func processTruck(truck Truck) error {
	fmt.Printf("processing truck %+v\n", truck)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo for truck: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo for truck: %w", err)
	}

	return nil
}

func fillTruckWithCargo(t *NormalTruck) {
	t.cargo = 100
}

func main() {
	// truckID := "Truck-1"
	// fmt.Println("id:", truckID)
	// fmt.Println("address in memory:", &truckID) // we can use & to get the address of a variable

	t := NormalTruck{id: "Truck-1", cargo: 0}
	fillTruckWithCargo(&t)

	fmt.Printf("Truck after filling: %+v\n", t)
}
