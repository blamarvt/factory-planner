// Package core contains layout generation algorithms for factory design.
package core

import (
	"fmt"
	"image/color"
)

// ItemColorProvider provides color information for items.
type ItemColorProvider interface {
	GetItemColor(itemName string) (color.Color, bool)
}

// Position represents a 2D coordinate in the factory layout.
type Position struct {
	X, Y int
}

// Building represents a factory building (assembler, furnace, etc.).
type Building struct {
	ID       string
	Type     string // "assembler", "furnace", "belt", "inserter", etc.
	Position Position
	Recipe   string     // recipe being crafted (for machines)
	Rotation int        // 0, 90, 180, 270 degrees
	Color    color.Color // color for rendering
}

// FactoryLayout represents the complete physical layout of a factory.
type FactoryLayout struct {
	Buildings []Building
	Width     int
	Height    int
	Title     string
}

// LayoutGenerator creates physical factory layouts from production plans.
type LayoutGenerator struct {
	MinSpacing    int                // minimum space between buildings
	ColorProvider ItemColorProvider  // provider for building colors
}

// NewLayoutGenerator creates a new layout generator.
func NewLayoutGenerator() *LayoutGenerator {
	return &LayoutGenerator{
		MinSpacing: 2, // default spacing
	}
}

// NewLayoutGeneratorWithColorProvider creates a new layout generator with a color provider.
func NewLayoutGeneratorWithColorProvider(colorProvider ItemColorProvider) *LayoutGenerator {
	return &LayoutGenerator{
		MinSpacing:    2,
		ColorProvider: colorProvider,
	}
}

// GenerateLayout creates a factory layout from a production plan.
func (lg *LayoutGenerator) GenerateLayout(plan *ProductionPlan) (*FactoryLayout, error) {
	if plan == nil {
		return nil, fmt.Errorf("production plan cannot be nil")
	}

	layout := &FactoryLayout{
		Buildings: make([]Building, 0),
		Width:     10, // placeholder dimensions
		Height:    10,
		Title:     "Factory Layout",
	}

	// TODO: Implement actual layout generation algorithm
	// This would involve:
	// 1. Placing production buildings based on the plan
	// 2. Routing belts and inserters for material flow
	// 3. Optimizing for minimal belt length and congestion
	// 4. Ensuring proper spacing for inserter reach
	// 5. Adding power poles and infrastructure

	// Placeholder: Add a basic building for each recipe in the plan
	x, y := 0, 0
	buildingID := 0

	for recipeName, count := range plan.RequiredMachines {
		for i := 0; i < count; i++ {
			buildingType := "assembler" // simplified for now
			building := Building{
				ID:       fmt.Sprintf("building_%d", buildingID),
				Type:     buildingType,
				Position: Position{X: x, Y: y},
				Recipe:   recipeName,
				Rotation: 0,
			}

			// Set building color from color provider
			if lg.ColorProvider != nil {
				// Map building type to item name
				itemName := lg.getBuildingItemName(buildingType)
				if buildingColor, hasColor := lg.ColorProvider.GetItemColor(itemName); hasColor {
					building.Color = buildingColor
				} else {
					// Default color if not found
					building.Color = color.RGBA{150, 150, 150, 255} // gray
				}
			} else {
				// Default color if no provider
				building.Color = color.RGBA{150, 150, 150, 255} // gray
			}

			layout.Buildings = append(layout.Buildings, building)

			buildingID++
			x += lg.MinSpacing + 1
			if x > 20 { // wrap to next row
				x = 0
				y += lg.MinSpacing + 1
			}
		}
	}

	// Update layout dimensions based on placed buildings
	if len(layout.Buildings) > 0 {
		maxX, maxY := 0, 0
		for _, building := range layout.Buildings {
			if building.Position.X > maxX {
				maxX = building.Position.X
			}
			if building.Position.Y > maxY {
				maxY = building.Position.Y
			}
		}
		layout.Width = maxX + 5 // add some padding
		layout.Height = maxY + 5
	}

	return layout, nil
}

// ValidateLayout checks if a layout is valid and collision-free.
func (lg *LayoutGenerator) ValidateLayout(layout *FactoryLayout) error {
	// TODO: Implement collision detection and validation
	// Check for overlapping buildings, invalid connections, etc.

	if len(layout.Buildings) == 0 {
		return fmt.Errorf("layout contains no buildings")
	}

	return nil
}

// getBuildingItemName maps building types to their corresponding item names.
func (lg *LayoutGenerator) getBuildingItemName(buildingType string) string {
	switch buildingType {
	case "assembler":
		return "assembling-machine-1"
	case "furnace":
		return "stone-furnace"
	case "belt":
		return "transport-belt"
	case "inserter":
		return "inserter"
	case "power":
		return "small-electric-pole"
	default:
		return "assembling-machine-1" // default fallback
	}
}
