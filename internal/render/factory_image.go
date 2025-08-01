// Package render provides PNG image generation for factory layouts.
package render

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/blamarvt/factory-planner/internal/core"
)

// ImageRenderer handles the generation of factory layout images.
type ImageRenderer struct {
	TileSize       int                    // pixels per tile
	GridColor      color.Color            // grid line color
	BgColor        color.Color            // background color
	BuildingColors map[string]color.Color // building type -> color
}

// NewImageRenderer creates a new image renderer with default settings.
func NewImageRenderer() *ImageRenderer {
	return &ImageRenderer{
		TileSize:  32,                             // 32x32 pixels per tile
		GridColor: color.RGBA{200, 200, 200, 255}, // light gray
		BgColor:   color.RGBA{50, 50, 50, 255},    // dark gray
		BuildingColors: map[string]color.Color{
			"assembler": color.RGBA{100, 150, 255, 255}, // blue
			"furnace":   color.RGBA{255, 100, 100, 255}, // red
			"belt":      color.RGBA{255, 255, 100, 255}, // yellow
			"inserter":  color.RGBA{100, 255, 100, 255}, // green
			"power":     color.RGBA{255, 150, 100, 255}, // orange
		},
	}
}

// RenderLayout generates a PNG image of the factory layout.
func (ir *ImageRenderer) RenderLayout(layout *core.FactoryLayout, outputPath string) error {
	if layout == nil {
		return fmt.Errorf("layout cannot be nil")
	}

	// Calculate image dimensions
	width := layout.Width * ir.TileSize
	height := layout.Height * ir.TileSize

	// Create image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill background
	draw.Draw(img, img.Bounds(), &image.Uniform{ir.BgColor}, image.Point{}, draw.Src)

	// Draw grid
	ir.drawGrid(img, layout.Width, layout.Height)

	// Draw buildings
	for _, building := range layout.Buildings {
		ir.drawBuilding(img, &building)
	}

	// Draw title if provided
	if layout.Title != "" {
		// TODO: Add text rendering for title
		// For now, we'll skip text rendering to avoid external dependencies
	}

	// Save to file
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return fmt.Errorf("failed to encode PNG: %w", err)
	}

	return nil
}

// drawGrid draws a grid overlay on the image.
func (ir *ImageRenderer) drawGrid(img *image.RGBA, gridWidth, gridHeight int) {
	bounds := img.Bounds()

	// Vertical lines
	for x := 0; x <= gridWidth; x++ {
		xPos := x * ir.TileSize
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			if xPos < bounds.Max.X {
				img.Set(xPos, y, ir.GridColor)
			}
		}
	}

	// Horizontal lines
	for y := 0; y <= gridHeight; y++ {
		yPos := y * ir.TileSize
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if yPos < bounds.Max.Y {
				img.Set(x, yPos, ir.GridColor)
			}
		}
	}
}

// drawBuilding draws a single building on the image.
func (ir *ImageRenderer) drawBuilding(img *image.RGBA, building *core.Building) {
	// Get building color
	buildingColor, exists := ir.BuildingColors[building.Type]
	if !exists {
		buildingColor = color.RGBA{150, 150, 150, 255} // default gray
	}

	// Calculate pixel coordinates
	x1 := building.Position.X*ir.TileSize + 2 // small padding
	y1 := building.Position.Y*ir.TileSize + 2
	x2 := x1 + ir.TileSize - 4 // leave border
	y2 := y1 + ir.TileSize - 4

	// Draw building rectangle
	buildingRect := image.Rect(x1, y1, x2, y2)
	draw.Draw(img, buildingRect, &image.Uniform{buildingColor}, image.Point{}, draw.Src)

	// Draw border
	borderColor := color.RGBA{0, 0, 0, 255} // black border

	// Top and bottom borders
	for x := x1; x < x2; x++ {
		img.Set(x, y1, borderColor)
		img.Set(x, y2-1, borderColor)
	}

	// Left and right borders
	for y := y1; y < y2; y++ {
		img.Set(x1, y, borderColor)
		img.Set(x2-1, y, borderColor)
	}
}

// SetBuildingColor allows customization of building colors.
func (ir *ImageRenderer) SetBuildingColor(buildingType string, color color.Color) {
	ir.BuildingColors[buildingType] = color
}
