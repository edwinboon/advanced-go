package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	// "time"
)

type ContextKey string

var UserIDKey ContextKey = "userID"

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
	// fmt.Printf("processing truck %+v\n", truck)
	//
	// // simulate some processing time
	// time.Sleep(time.Second * 1)
	//
	// if err := truck.LoadCargo(); err != nil {
	// 	return fmt.Errorf("error loading cargo for truck: %w", err)
	// }
	//
	// if err := truck.UnloadCargo(); err != nil {
	// 	return fmt.Errorf("error unloading cargo for truck: %w", err)
	// }
	//
	// fmt.Printf("finished processing truck %+v\n", truck)
	return errors.New("truck not found")
}

/**
 * only add the keyword 'go' will not work as expected because the main function may exit before goroutines complete.
 * there fore, we need to add a waitGroup mechanism to ensure all goroutines finish before main exits.
 */
func processFleet(trucks []Truck) error {
	var wg sync.WaitGroup

	errorsChan := make(chan error, len(trucks)) // buffered channel to hold errors

	for _, t := range trucks {
		wg.Add(1) // wait for 1 goroutine each iteration

		go func(t Truck) {
			if err := processTruck(t); err != nil {
				errorsChan <- err // sending error to channel
			}
			wg.Done()
		}(t) // launch goroutine to process each truck and we do this because we want to call wg.Done() after processing
	}

	wg.Wait()         // wait for all goroutines to finish
	close(errorsChan) // close the errors channel after all goroutines are done

	var errs []error

	for err := range errorsChan {
		log.Printf("Error processing truck: %s\n", err)
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("fleet processing had %d errors", len(errs))
	}

	return nil
}

func main() {
	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, batteryLevel: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, batteryLevel: 100},
	}

	// Process all trucks concurrently
	if err := processFleet(fleet); err != nil {
		fmt.Printf("Error processing fleed: %v\n", err)
		return
	}

	fmt.Println("All trucks processed succesfully")
}
