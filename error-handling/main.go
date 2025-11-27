package main

import (
	"fmt"
	"log"
)

type Truck struct {
	id string
}

func (t *Truck) LoadCargo() error {
	return nil
}

func (t *Truck) UnloadCargo() error {
	return nil
}

// processTruck handles the loading and unloading of a truck.
func processTruck(truck Truck) error {
	fmt.Printf("Processing truck: %s\n", truck.id)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo for truck %s: %w", truck.id, err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo for truck %s: %w", truck.id, err)
	}

	return nil
}

func main() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.id)

		err := processTruck(truck)
		if err != nil {
			log.Fatalf("Error processing truck: %s", err)
		}
	}
}
