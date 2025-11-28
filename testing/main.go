package main

import (
	"fmt"
	"log"
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

func main() {
	nt := &NormalTruck{id: "Truck-1", cargo: 0}
	et := &ElectricTruck{id: "ETruck-1", cargo: 0, batteryLevel: 100.0}

	err := processTruck(nt)
	if err != nil {
		log.Fatalf("Error processing normal truck: %s", err)
	}

	err = processTruck(et)
	if err != nil {
		log.Fatalf("Error processing normal ETruck: %s", err)
	}

	log.Printf("Normal Truck after processing: %+v\n", nt)
	log.Printf("Electric Truck after processing: %+v\n", et)
	log.Printf("All trucks processed successfully.")
}
