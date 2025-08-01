// Package data contains Factorio game data structures and loading functionality.
package data

import "github.com/blamarvt/factory-planner/internal/core"

// RecipeData holds all recipe information from Factorio.
type RecipeData struct {
	Version string                  `json:"version"`
	Recipes map[string]*core.Recipe `json:"recipes"`
}

// LoadRecipes loads recipe data from various sources.
func LoadRecipes() (*RecipeData, error) {
	// TODO: Implement recipe data loading
	// This could load from:
	// - Embedded JSON files
	// - Factorio mod data
	// - External recipe databases

	// Placeholder: Create some basic recipes for testing
	recipes := &RecipeData{
		Version: "1.1.0",
		Recipes: make(map[string]*core.Recipe),
	}

	// Basic science pack recipe
	sciencePack1 := &core.Recipe{
		Name: "automation-science-pack",
		Inputs: map[string]float64{
			"copper-plate":    1.0,
			"iron-gear-wheel": 1.0,
		},
		Outputs: map[string]float64{
			"automation-science-pack": 1.0,
		},
		CraftingTime: 5.0,
		Category:     "crafting",
	}
	recipes.Recipes["automation-science-pack"] = sciencePack1

	// Iron gear wheel recipe
	ironGear := &core.Recipe{
		Name: "iron-gear-wheel",
		Inputs: map[string]float64{
			"iron-plate": 2.0,
		},
		Outputs: map[string]float64{
			"iron-gear-wheel": 1.0,
		},
		CraftingTime: 0.5,
		Category:     "crafting",
	}
	recipes.Recipes["iron-gear-wheel"] = ironGear

	// Iron plate smelting
	ironPlate := &core.Recipe{
		Name: "iron-plate",
		Inputs: map[string]float64{
			"iron-ore": 1.0,
		},
		Outputs: map[string]float64{
			"iron-plate": 1.0,
		},
		CraftingTime: 3.2,
		Category:     "smelting",
	}
	recipes.Recipes["iron-plate"] = ironPlate

	// Copper plate smelting
	copperPlate := &core.Recipe{
		Name: "copper-plate",
		Inputs: map[string]float64{
			"copper-ore": 1.0,
		},
		Outputs: map[string]float64{
			"copper-plate": 1.0,
		},
		CraftingTime: 3.2,
		Category:     "smelting",
	}
	recipes.Recipes["copper-plate"] = copperPlate

	return recipes, nil
}

// GetRecipeGraph creates a recipe graph from the loaded recipe data.
func (rd *RecipeData) GetRecipeGraph() *core.RecipeGraph {
	graph := core.NewRecipeGraph()

	for _, recipe := range rd.Recipes {
		graph.AddRecipe(recipe)
	}

	return graph
}
