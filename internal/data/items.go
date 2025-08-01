// Package data contains item definitions and properties.
package data

import "image/color"

// ItemType represents the category of an item.
type ItemType string

const (
	ItemTypeRaw          ItemType = "raw"          // Raw materials (ores, crude oil)
	ItemTypeIntermediate ItemType = "intermediate" // Processed materials (plates, circuits)
	ItemTypeConsumable   ItemType = "consumable"   // Science packs, fuel
	ItemTypeBuilding     ItemType = "building"     // Machines, structures
	ItemTypeTool         ItemType = "tool"         // Equipment, weapons
)

// Item represents a Factorio item with its properties.
type Item struct {
	Name       string                 `json:"name"`
	Type       ItemType               `json:"type"`
	StackSize  int                    `json:"stack_size"`
	FuelValue  float64                `json:"fuel_value,omitempty"` // in MJ
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// ItemDatabase holds all item definitions.
type ItemDatabase struct {
	Version string           `json:"version"`
	Items   map[string]*Item `json:"items"`
}

// LoadItems loads item data from game files or embedded data.
func LoadItems() (*ItemDatabase, error) {
	// TODO: Implement item data loading from Factorio data files

	// Placeholder: Create basic item definitions
	itemDB := &ItemDatabase{
		Version: "1.1.0",
		Items:   make(map[string]*Item),
	}

	// Raw materials
	items := []*Item{
		{Name: "iron-ore", Type: ItemTypeRaw, StackSize: 50},
		{Name: "copper-ore", Type: ItemTypeRaw, StackSize: 50},
		{Name: "coal", Type: ItemTypeRaw, StackSize: 50, FuelValue: 4.0},
		{Name: "stone", Type: ItemTypeRaw, StackSize: 50},
		{Name: "wood", Type: ItemTypeRaw, StackSize: 50, FuelValue: 2.0},

		// Intermediate products
		{Name: "iron-plate", Type: ItemTypeIntermediate, StackSize: 100},
		{Name: "copper-plate", Type: ItemTypeIntermediate, StackSize: 100},
		{Name: "steel-plate", Type: ItemTypeIntermediate, StackSize: 100},
		{Name: "iron-gear-wheel", Type: ItemTypeIntermediate, StackSize: 100},
		{Name: "copper-cable", Type: ItemTypeIntermediate, StackSize: 200},
		{Name: "electronic-circuit", Type: ItemTypeIntermediate, StackSize: 200},

		// Science packs
		{Name: "automation-science-pack", Type: ItemTypeConsumable, StackSize: 200},
		{Name: "logistic-science-pack", Type: ItemTypeConsumable, StackSize: 200},
		{Name: "chemical-science-pack", Type: ItemTypeConsumable, StackSize: 200},
		{Name: "production-science-pack", Type: ItemTypeConsumable, StackSize: 200},
		{Name: "utility-science-pack", Type: ItemTypeConsumable, StackSize: 200},

		// Buildings
		{Name: "assembling-machine-1", Type: ItemTypeBuilding, StackSize: 50, Properties: map[string]interface{}{"color": color.RGBA{100, 150, 255, 255}}},
		{Name: "assembling-machine-2", Type: ItemTypeBuilding, StackSize: 50, Properties: map[string]interface{}{"color": color.RGBA{100, 150, 255, 255}}},
		{Name: "assembling-machine-3", Type: ItemTypeBuilding, StackSize: 50, Properties: map[string]interface{}{"color": color.RGBA{100, 150, 255, 255}}},
		{Name: "stone-furnace", Type: ItemTypeBuilding, StackSize: 50, Properties: map[string]interface{}{"color": color.RGBA{255, 100, 100, 255}}},
		{Name: "steel-furnace", Type: ItemTypeBuilding, StackSize: 50, Properties: map[string]interface{}{"color": color.RGBA{255, 100, 100, 255}}},
		{Name: "electric-furnace", Type: ItemTypeBuilding, StackSize: 50, Properties: map[string]interface{}{"color": color.RGBA{255, 100, 100, 255}}},
		{Name: "transport-belt", Type: ItemTypeBuilding, StackSize: 100, Properties: map[string]interface{}{"color": color.RGBA{255, 255, 100, 255}}},
		{Name: "inserter", Type: ItemTypeBuilding, StackSize: 50, Properties: map[string]interface{}{"color": color.RGBA{100, 255, 100, 255}}},
		{Name: "small-electric-pole", Type: ItemTypeBuilding, StackSize: 50, Properties: map[string]interface{}{"color": color.RGBA{255, 150, 100, 255}}},
	}

	for _, item := range items {
		itemDB.Items[item.Name] = item
	}

	return itemDB, nil
}

// GetItem retrieves an item by name.
func (db *ItemDatabase) GetItem(name string) (*Item, bool) {
	item, exists := db.Items[name]
	return item, exists
}

// GetItemsByType returns all items of a specific type.
func (db *ItemDatabase) GetItemsByType(itemType ItemType) []*Item {
	var result []*Item
	for _, item := range db.Items {
		if item.Type == itemType {
			result = append(result, item)
		}
	}
	return result
}

// IsRawMaterial checks if an item is a raw material.
func (db *ItemDatabase) IsRawMaterial(itemName string) bool {
	if item, exists := db.GetItem(itemName); exists {
		return item.Type == ItemTypeRaw
	}
	return false
}

// IsFuel checks if an item can be used as fuel.
func (db *ItemDatabase) IsFuel(itemName string) bool {
	if item, exists := db.GetItem(itemName); exists {
		return item.FuelValue > 0
	}
	return false
}

// GetItemColor retrieves the color property of an item.
func (db *ItemDatabase) GetItemColor(itemName string) (color.Color, bool) {
	if item, exists := db.GetItem(itemName); exists {
		if item.Properties != nil {
			if colorValue, hasColor := item.Properties["color"]; hasColor {
				if c, ok := colorValue.(color.Color); ok {
					return c, true
				}
			}
		}
	}
	return nil, false
}
