// Package blueprint provides Factorio blueprint string generation.
package blueprint

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/blamarvt/factory-planner/internal/core"
)

// BlueprintEntity represents a single entity in a Factorio blueprint.
type BlueprintEntity struct {
	EntityNumber int      `json:"entity_number"`
	Name         string   `json:"name"`
	Position     Position `json:"position"`
	Direction    *int     `json:"direction,omitempty"`
	Recipe       *string  `json:"recipe,omitempty"`
}

// Position represents coordinates in blueprint format.
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Blueprint represents the main blueprint structure.
type Blueprint struct {
	Item     string            `json:"item"`
	Label    string            `json:"label"`
	Entities []BlueprintEntity `json:"entities"`
	Icons    []BlueprintIcon   `json:"icons,omitempty"`
	Version  int64             `json:"version"`
}

// BlueprintIcon represents an icon in the blueprint.
type BlueprintIcon struct {
	Signal BlueprintSignal `json:"signal"`
	Index  int             `json:"index"`
}

// BlueprintSignal represents a signal/item icon.
type BlueprintSignal struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// BlueprintWrapper wraps the blueprint for encoding.
type BlueprintWrapper struct {
	Blueprint Blueprint `json:"blueprint"`
}

// Exporter handles blueprint string generation.
type Exporter struct {
	Version int64 // Factorio version
}

// NewExporter creates a new blueprint exporter.
func NewExporter() *Exporter {
	return &Exporter{
		Version: 281479275249664, // Factorio 1.1.x version
	}
}

// ExportBlueprint converts a factory layout to a Factorio blueprint string.
func (e *Exporter) ExportBlueprint(layout *core.FactoryLayout) (string, error) {
	if layout == nil {
		return "", fmt.Errorf("layout cannot be nil")
	}

	// Create blueprint structure
	blueprint := Blueprint{
		Item:     "blueprint",
		Label:    layout.Title,
		Entities: make([]BlueprintEntity, 0, len(layout.Buildings)),
		Version:  e.Version,
	}

	// Convert buildings to blueprint entities
	for i, building := range layout.Buildings {
		entity := BlueprintEntity{
			EntityNumber: i + 1,
			Name:         e.getBlueprintEntityName(building.Type),
			Position: Position{
				X: float64(building.Position.X),
				Y: float64(building.Position.Y),
			},
		}

		// Add recipe if it's a crafting machine
		if building.Recipe != "" {
			entity.Recipe = &building.Recipe
		}

		// Add direction/rotation if specified
		if building.Rotation != 0 {
			direction := building.Rotation / 90 * 2 // Convert to Factorio direction format
			entity.Direction = &direction
		}

		blueprint.Entities = append(blueprint.Entities, entity)
	}

	// Add default icon
	if len(blueprint.Entities) > 0 {
		blueprint.Icons = []BlueprintIcon{
			{
				Signal: BlueprintSignal{
					Type: "item",
					Name: "assembling-machine-1",
				},
				Index: 1,
			},
		}
	}

	// Wrap blueprint
	wrapper := BlueprintWrapper{Blueprint: blueprint}

	// Convert to blueprint string
	return e.encodeBlueprint(wrapper)
}

// getBlueprintEntityName maps our building types to Factorio entity names.
func (e *Exporter) getBlueprintEntityName(buildingType string) string {
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

// encodeBlueprint converts a blueprint struct to a Factorio blueprint string.
func (e *Exporter) encodeBlueprint(wrapper BlueprintWrapper) (string, error) {
	// Convert to JSON
	jsonData, err := json.Marshal(wrapper)
	if err != nil {
		return "", fmt.Errorf("failed to marshal blueprint JSON: %w", err)
	}

	// Compress with zlib
	var compressed bytes.Buffer
	writer := zlib.NewWriter(&compressed)
	_, err = writer.Write(jsonData)
	if err != nil {
		writer.Close()
		return "", fmt.Errorf("failed to compress blueprint: %w", err)
	}
	writer.Close()

	// Encode to base64
	encoded := base64.StdEncoding.EncodeToString(compressed.Bytes())

	// Add version prefix (0 for blueprint)
	return "0" + encoded, nil
}

// ValidateBlueprint checks if the blueprint string is valid.
func (e *Exporter) ValidateBlueprint(blueprintString string) error {
	if len(blueprintString) < 2 {
		return fmt.Errorf("blueprint string too short")
	}

	if blueprintString[0] != '0' {
		return fmt.Errorf("invalid blueprint version prefix")
	}

	// Try to decode to verify format
	encoded := blueprintString[1:]

	// Decode base64
	compressed, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return fmt.Errorf("failed to decode base64: %w", err)
	}

	// Decompress zlib
	reader, err := zlib.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return fmt.Errorf("failed to create zlib reader: %w", err)
	}
	defer reader.Close()

	var decompressed bytes.Buffer
	_, err = decompressed.ReadFrom(reader)
	if err != nil {
		return fmt.Errorf("failed to decompress blueprint: %w", err)
	}

	// Validate JSON structure
	var wrapper BlueprintWrapper
	err = json.Unmarshal(decompressed.Bytes(), &wrapper)
	if err != nil {
		return fmt.Errorf("failed to parse blueprint JSON: %w", err)
	}

	return nil
}
