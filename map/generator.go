package _map

import (
	"golang.org/x/image/colornames"
	"image/color"
	"math/rand"
)

func generateWall(m *MapTilemap, x1, y1 uint, x2, y2 uint) {
	if x1 > x2 || y1 > y2 {
		panic("Invalid coordinates")
	}

	dy := float64(y2 - y1) / float64(x2-x1+1)
	iy := y1

	for ix := x1; ix <= x2; ix++ {
		for di := iy; float64(di) <= float64(iy)+dy; di++ {
            if (*m)[ix][di] == nil {
            	(*m)[ix][di] = &MapTile{'#', color.White}
			}
		}

		iy = uint(float64(iy) + dy)
	}
}

func generateFloor(m *MapTilemap, x1, y1, x2, y2 uint) bool {
	if x1 > x2 || y1 > y2 {
		panic("Invalid coordinates")
	}

	connected := false
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			if (*m)[x][y] != nil {
				connected = true
			}

			(*m)[x][y] = &MapTile{'.', colornames.Darkgray}
		}
	}

	return connected
}

func generateRoom(m *MapTilemap, x, y uint, size uint) {
    generateWall(m, x, y, x+size, y)
	generateWall(m, x, y, x, y+size)
	generateWall(m, x+size, y, x+size, y+size)
	generateWall(m, x, y+size, x+size, y+size)
	generateFloor(m, x+1, y+1, x+size-1, y+size-1)
}

func Gen() MapTilemap {
	capacity := uint(250)
    m := NewTilemap(capacity)

	for i := 0; i < 150; i++ {
		x := uint(rand.Intn(int(capacity)-5))
		y := uint(rand.Intn(int(capacity)-5))

		size := uint(0)
		if x > y {
			size = uint(rand.Intn(int(capacity / 7)))
		} else {
			size = uint(rand.Intn(int(capacity / 7)))
		}

		if size + x >= capacity {
			size = capacity - x - 1
		}

		if size + y >= capacity {
			size = capacity - y - 1
		}

		if size < 3 {
			size = 3
		}

		generateRoom(&m, x, y, size)
	}

	return m
}