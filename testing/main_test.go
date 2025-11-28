package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {
		t.Run("should load and unload a trucks cargo without errors", func(t *testing.T) {
			nt := &NormalTruck{id: "Truck-1", cargo: 0}
			et := &ElectricTruck{id: "ETruck-1", cargo: 0, batteryLevel: 100.0}

			err := processTruck(nt)
			if err != nil {
				t.Fatalf("Error processing normal truck: %s", err)
			}

			err = processTruck(et)
			if err != nil {
				t.Fatalf("Error processing normal ETruck: %s", err)
			}

			// assertions
			if nt.cargo != 0 {
				t.Fatalf("NormalTruck cargo should be 0 after unloading, got %d", nt.cargo)
			}

			if et.batteryLevel != 85.0 {
				t.Fatalf("ElectricTruck battery level should be 85.0 after loading and unloading, got %f", et.batteryLevel)
			}
		})
	})
}
