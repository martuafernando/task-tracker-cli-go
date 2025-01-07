package models

type Status = int

const (
	Todo Status = iota
	Inprogress
	Done
)

type Task struct {
	Id     int
	Name   string
	Status Status
}
