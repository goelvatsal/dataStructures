package main

import "fmt"

type Card struct {
	face  int
	shape int
	color int
}

func structsPractice() Card {
	card := Card{1, 0, 0}

	fmt.Println("Face Value: ", card.face)
	fmt.Println("Shape: ", card.shape)
	fmt.Println("Color: ", card.color)

	return card
}
