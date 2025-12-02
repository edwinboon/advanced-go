package main

import (
	"errors"
	"fmt"
	"sync"
)

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks       map[string]*Truck
	sync.RWMutex // compose a read-write mutex for concurrent access
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (tm *truckManager) AddTruck(id string, cargo int) error {
	tm.Lock() // lock for writing
	defer tm.Unlock()

	if _, exists := tm.trucks[id]; exists {
		return fmt.Errorf("truck with id %s already exists", id)
	}
	tm.trucks[id] = &Truck{ID: id, Cargo: cargo}
	// tm.Unlock() // other option is to use defer
	return nil
}

func (tm *truckManager) GetTruck(id string) (Truck, error) {
	tm.RLock() // lock for reading
	defer tm.RUnlock()

	truck, exists := tm.trucks[id]

	if !(exists) {
		return Truck{}, ErrTruckNotFound
	}
	return *truck, nil
}

func (tm *truckManager) RemoveTruck(id string) error {
	tm.Lock()
	defer tm.Unlock()
	if _, exists := tm.trucks[id]; !exists {
		return ErrTruckNotFound
	}
	delete(tm.trucks, id)
	return nil
}

func (tm *truckManager) UpdateTruckCargo(id string, cargo int) error {
	tm.Lock()
	defer tm.Unlock()
	truck, exists := tm.trucks[id]
	if !exists {
		return ErrTruckNotFound
	}
	truck.Cargo = cargo
	return nil
}
