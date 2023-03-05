package main

import "fmt"

type mobil struct {
	Type   string
	fuelIn float64
}

func (mobil mobil) JarakTempuh() float64 {

	return mobil.fuelIn / 1.5
}

func main() {
	mobil := mobil{
		Type:   "Pajero Sport",
		fuelIn: 100}

	jarakTempuh := mobil.JarakTempuh()

	fmt.Println("jarak yang sanggup ditempuh:", mobil.Type, jarakTempuh)
}
