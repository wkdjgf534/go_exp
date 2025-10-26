package tilemap

import (
	"encoding/json"
	"os"
	"path"

	"rpg-tutorial/tileset"
)

// Data we want for one layer in our list of layers
type TilemapLayerJSON struct {
	Data   []int  `json:"data"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Name   string `json:"name"`
}

// All layers in a tilemap
type TilemapJSON struct {
	Layers []TilemapLayerJSON `json:"layers"`
	// Raw data for each tileset (path, gid)
	Tilesets []map[string]any `json:"tilesets"`
}

// Temp function to generate all of our tilesets and return a slice of them
func (t *TilemapJSON) GenTilesets() ([]tileset.Tileset, error) {
	tilesets := make([]tileset.Tileset, 0)

	for _, tilesetData := range t.Tilesets {
		// Convert map relative path to project relative path
		tilesetPath := path.Join("assets/maps/", tilesetData["source"].(string))
		tileset, err := tileset.NewTileset(tilesetPath, int(tilesetData["firstgid"].(float64)))
		if err != nil {
			return nil, err
		}

		tilesets = append(tilesets, tileset)
	}

	return tilesets, nil
}

// Opens the file, parses it, and returns the json object + potential error
func NewTilemapJSON(filepath string) (*TilemapJSON, error) {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var tilemapJSON TilemapJSON
	err = json.Unmarshal(contents, &tilemapJSON)
	if err != nil {
		return nil, err
	}

	return &tilemapJSON, nil
}
