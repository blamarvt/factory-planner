// Package data contains technology and research data structures.
package data

// Technology represents a Factorio research technology.
type Technology struct {
	Name          string             `json:"name"`
	Prerequisites []string           `json:"prerequisites"`
	Research      map[string]int     `json:"research"` // science pack -> count
	Effects       []TechnologyEffect `json:"effects"`
}

// TechnologyEffect represents what a technology unlocks.
type TechnologyEffect struct {
	Type     string  `json:"type"` // "unlock-recipe", "modifier", etc.
	Recipe   string  `json:"recipe,omitempty"`
	Modifier string  `json:"modifier,omitempty"`
	Change   float64 `json:"change,omitempty"`
}

// ResearchProgress tracks which technologies have been unlocked.
type ResearchProgress struct {
	UnlockedTechnologies map[string]bool `json:"unlocked"`
	AvailableRecipes     map[string]bool `json:"recipes"`
}

// TechnologyData holds all technology information.
type TechnologyData struct {
	Version      string                 `json:"version"`
	Technologies map[string]*Technology `json:"technologies"`
}

// LoadTechnologies loads technology data for research progression.
func LoadTechnologies() (*TechnologyData, error) {
	// TODO: Implement technology data loading from game files

	// Placeholder: Create basic technology tree
	techData := &TechnologyData{
		Version:      "1.1.0",
		Technologies: make(map[string]*Technology),
	}

	// Basic automation technology
	automation := &Technology{
		Name:          "automation",
		Prerequisites: []string{},
		Research: map[string]int{
			"automation-science-pack": 10,
		},
		Effects: []TechnologyEffect{
			{Type: "unlock-recipe", Recipe: "assembling-machine-1"},
			{Type: "unlock-recipe", Recipe: "long-handed-inserter"},
		},
	}
	techData.Technologies["automation"] = automation

	// Electronics technology
	electronics := &Technology{
		Name:          "electronics",
		Prerequisites: []string{},
		Research: map[string]int{
			"automation-science-pack": 30,
		},
		Effects: []TechnologyEffect{
			{Type: "unlock-recipe", Recipe: "electronic-circuit"},
			{Type: "unlock-recipe", Recipe: "inserter"},
		},
	}
	techData.Technologies["electronics"] = electronics

	return techData, nil
}

// CreateResearchProgress creates a research progress tracker for a given level.
func CreateResearchProgress(level string) *ResearchProgress {
	progress := &ResearchProgress{
		UnlockedTechnologies: make(map[string]bool),
		AvailableRecipes:     make(map[string]bool),
	}

	// TODO: Implement proper research level parsing
	// For now, handle basic predefined levels
	switch level {
	case "basic-science":
		progress.UnlockedTechnologies["automation"] = true
		progress.AvailableRecipes["automation-science-pack"] = true
		progress.AvailableRecipes["iron-gear-wheel"] = true
		progress.AvailableRecipes["iron-plate"] = true
		progress.AvailableRecipes["copper-plate"] = true
	case "early-game":
		progress.UnlockedTechnologies["automation"] = true
		progress.UnlockedTechnologies["electronics"] = true
		progress.AvailableRecipes["automation-science-pack"] = true
		progress.AvailableRecipes["iron-gear-wheel"] = true
		progress.AvailableRecipes["iron-plate"] = true
		progress.AvailableRecipes["copper-plate"] = true
		progress.AvailableRecipes["electronic-circuit"] = true
	}

	return progress
}

// IsRecipeUnlocked checks if a recipe is available with current research.
func (rp *ResearchProgress) IsRecipeUnlocked(recipeName string) bool {
	return rp.AvailableRecipes[recipeName]
}

// IsTechnologyUnlocked checks if a technology has been researched.
func (rp *ResearchProgress) IsTechnologyUnlocked(techName string) bool {
	return rp.UnlockedTechnologies[techName]
}
