// Package core contains the core planning algorithms for factory optimization.
package core

// Recipe represents a Factorio recipe with inputs, outputs, and production time.
type Recipe struct {
	Name         string
	Inputs       map[string]float64 // item name -> quantity required
	Outputs      map[string]float64 // item name -> quantity produced
	CraftingTime float64            // time in seconds
	Category     string             // crafting category (e.g., "crafting", "smelting")
}

// RecipeGraph represents the dependency graph of all recipes.
type RecipeGraph struct {
	Recipes      map[string]*Recipe
	Dependencies map[string][]string // item -> list of recipes that produce it
}

// NewRecipeGraph creates a new empty recipe graph.
func NewRecipeGraph() *RecipeGraph {
	return &RecipeGraph{
		Recipes:      make(map[string]*Recipe),
		Dependencies: make(map[string][]string),
	}
}

// AddRecipe adds a recipe to the graph and updates dependencies.
func (rg *RecipeGraph) AddRecipe(recipe *Recipe) {
	rg.Recipes[recipe.Name] = recipe

	// Update dependencies for each output item
	for outputItem := range recipe.Outputs {
		rg.Dependencies[outputItem] = append(rg.Dependencies[outputItem], recipe.Name)
	}
}

// GetRecipesForItem returns all recipes that can produce the given item.
func (rg *RecipeGraph) GetRecipesForItem(item string) []*Recipe {
	var recipes []*Recipe
	for _, recipeName := range rg.Dependencies[item] {
		if recipe, exists := rg.Recipes[recipeName]; exists {
			recipes = append(recipes, recipe)
		}
	}
	return recipes
}
