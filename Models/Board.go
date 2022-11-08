package models

type Board struct {
	ID     uint
	RowNum uint
	ColNum uint
	Boats  []Boat
	Pins   []Pin
}
