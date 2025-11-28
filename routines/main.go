package main

import (
	"fmt"
	"log"
	"sync"
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

func processFleet(fleet []Truck) error {
	var wg sync.WaitGroup

	for _, t := range fleet {

		wg.Add(1)

		go func(t Truck) {
			if err := processTruck(t); err != nil {
				log.Println(err)
			}
			wg.Done()
		}(t)

	}
	wg.Wait()

	return nil
}

func main() {
	Fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, batteryLevel: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, batteryLevel: 100},
	}

	if err := processFleet(Fleet); err != nil {
		log.Fatalf("Error processing fleet %v", err)
	}

	fmt.Println("All trucks processed succesfully")
}
