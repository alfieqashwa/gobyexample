package main

import (
	"fmt"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte // use Color as alias of byte.

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box // a slice of boxes

func (b Box) Volume() float64 { // Volume() uses Box as its receiver and returns the volume of Box
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) { // SetColor(c Color) changes Box's color.
	b.color = c
}

func (bl BoxList) BiggestsColor() Color { // BiggestsColor() returns the color which has the biggest volume.
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() { // PaintItBlack() sets color for all Box in BoxList to black.
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string { // String() use Color as its receiver, returns the string format of color name.
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm3")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestsColor().String())
	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestsColor().String())
}

/*
=== OUTPUT ===

We have 6 boxes in our set
The volume of the first one is 64 cm3
The color of the last one is YELLOW
The biggest one is YELLOW
Let's paint them all black
The color of the second one is BLACK
Obviously, now, the biggest one is BLACK

=== SOME EXPLANATION ===

Take a look at SetColor method. Its receiver is a pointer of Bos. You can use *Box as a receiver.
Why do we use a pointer here is because we want to change Box's color in this method.
Thus, if we don't use a pointer, it will only change the value inside a copy of Box.
*/
