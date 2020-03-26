package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

const (
	WIDTH  = 640
	HEIGHT = 360
	SCALE  = 2
)

var grid [WIDTH][HEIGHT]bool

func step(screen *ebiten.Image) {
	var gridNext [WIDTH][HEIGHT]bool

	for x := 1; x < WIDTH-1; x++ {
		for y := 1; y < HEIGHT-1; y++ {
			c := grid[x][y]

			t := grid[x][y-1]
			r := grid[x+1][y]
			b := grid[x][y+1]
			l := grid[x-1][y]

			tr := grid[x+1][y-1]
			br := grid[x+1][y+1]
			bl := grid[x-1][y+1]
			tl := grid[x-1][y-1]

			cells := [8]bool{t, r, b, l, tr, br, bl, tl}

			sum := 0
			for _, cell := range cells {
				if cell {
					sum++
				}
			}

			isCellNew := !c && sum == 3
			isCellSurvive := c && (sum == 2 || sum == 3)

			if isCellNew || isCellSurvive {
				gridNext[x][y] = true
				screen.Set(x, y, color.RGBA{0xff, 0xff, 0xff, 0xff})
			}
		}
	}

	grid = gridNext
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	step(screen)

	return nil
}

func setup() {
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			// Random state
			grid[x][y] = rand.Intn(2) == 0
		}
	}
}

func main() {
	setup()

	err := ebiten.Run(update, WIDTH, HEIGHT, SCALE, "Gogol")
	if err != nil {
		panic(err)
	}
}
