package _map

import (
	"image/color"
)

type (
	mapEntity interface {
		update(*Map)
	}

	MapTile struct {
		Rune rune
		Color color.Color
	}

	MapTilemap [][]*MapTile

	Map struct {
		Tiles MapTilemap
		entities []mapEntity
	}
)

func NewMap(tiles MapTilemap) Map {
	var entities []mapEntity

	return Map {
		tiles,
		entities,
	}
}

func NewTilemap(capacity uint) MapTilemap {
    result := make(MapTilemap, capacity)
    for i, _ := range result {
    	result[i] = make([]*MapTile, capacity)
	}

    return result
}

func (self *Map) Update() {
	for _, entity := range self.entities {
		entity.update(self)
	}
}