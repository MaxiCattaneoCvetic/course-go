package vehicles

import "fmt"

type Car struct {
	Time int
}

func (c Car) Distance() float64 {
	return 100 * (float64(c.Time) / 60)
}

type Motorcycle struct {
	Time int
}

func (m Motorcycle) Distance() float64 {
	return 120 * (float64(m.Time) / 60)
}

type Truck struct {
	Time int
}

func (t Truck) Distance() float64 {
	return 60 * (float64(t.Time) / 60)
}

// Ahora debemos crear una entidad que represente todas estas estructuras

type Vehicle interface {
	Distance() float64
}

const (
	CarVehicle        = "CAR"
	MotorcycleVehicle = "MOTORCYCLE"
	TruckVehicle      = "TRUCK"
)

func New(vehicle string, time int) (Vehicle, error) {

	switch vehicle { // retornamos la direccion en memoria
	case CarVehicle:
		return &Car{Time: time}, nil
	case MotorcycleVehicle:
		return &Motorcycle{Time: time}, nil
	case TruckVehicle:
		return &Truck{Time: time}, nil
	default:
		return nil, fmt.Errorf("vehicle not found")

	}
}
