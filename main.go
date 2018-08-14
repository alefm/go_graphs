package main

import (
	"github.com/fogleman/gg"
	"math/rand"
	"time"
)

type Node struct {
	position_x float64
	position_y float64
	// node_list  *Node
}

const screen_width int = 1024
const screen_height int = 1024

const screen_size float64 = 1024
const radio float64 = 20

func main() {

	rand.Seed(time.Now().UnixNano())

	dc := gg.NewContext(screen_width, screen_height)

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(0, 0, 0)
	x, y := random_position()
	dc.DrawCircle(x, y, 20)
	dc.Fill()

	dc.SetRGB(1, 1, 1)

	if err := dc.LoadFontFace("/Library/Fonts/Arial.ttf", 28); err != nil {
		panic(err)
	}

	dc.DrawStringAnchored("A", x, y, 0.5, 0.5)
	dc.SavePNG("out.png")
}

func random_position() (float64, float64) {
	x := radio + rand.Float64()*((screen_size-radio)-radio)
	y := radio + rand.Float64()*((screen_size-radio)-radio)
	return x, y
}

func check_distance(x1, y1, x2, y2 int) int {

	distanceX := x1 - x2
	distanceY := y1 - y2
	distance := distanceX*distanceX + distanceY*distanceY

	return distance
}
