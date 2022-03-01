package core

import "time"

type Event struct{}

type SlotParameters struct {
	month time.Time
	day   *time.Time
}

func (e Event) GetAvailableSlots(params SlotParameters) (AvailableDay,error){} {
	return nil, fmt.Errorf("Not implemented")
}

type Slot struct {
	start int
}

type AvailableDay struct {
	Slots []Slot
}
