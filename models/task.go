package models

type Status = int

const (
	Todo Status = iota
	Inprogress
	Done
)

var StatusNames = map[Status]string{
	Todo:       "Todo",
	Inprogress: "In progress",
	Done:       "Done",
}

type Task struct {
	Id     int
	Name   string
	Status Status
}
