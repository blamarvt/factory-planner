// Package core contains optimization algorithms for efficient factory planning.
package core

import "math"

// ProductionTarget represents a desired production rate for an item.
type ProductionTarget struct {
	Item string  // item name
	Rate float64 // items per minute
}

// ProductionPlan represents the calculated production requirements.
type ProductionPlan struct {
	Targets          []ProductionTarget
	RequiredMachines map[string]int     // recipe name -> number of machines needed
	ResourceFlow     map[string]float64 // item name -> items per minute
	TotalPowerUsage  float64            // estimated power consumption in MW
}

// Optimizer handles production optimization calculations.
type Optimizer struct {
	RecipeGraph *RecipeGraph
	Research    map[string]bool // available technologies
}

// NewOptimizer creates a new optimizer with the given recipe graph and research.
func NewOptimizer(graph *RecipeGraph, research map[string]bool) *Optimizer {
	return &Optimizer{
		RecipeGraph: graph,
		Research:    research,
	}
}

// OptimizeProduction calculates the optimal production plan for given targets.
func (opt *Optimizer) OptimizeProduction(targets []ProductionTarget) (*ProductionPlan, error) {
	plan := &ProductionPlan{
		Targets:          targets,
		RequiredMachines: make(map[string]int),
		ResourceFlow:     make(map[string]float64),
		TotalPowerUsage:  0.0,
	}

	// TODO: Implement actual optimization algorithm
	// This would involve:
	// 1. Building dependency tree for all target items
	// 2. Calculating required production rates for intermediate items
	// 3. Determining optimal recipe choices (when multiple options exist)
	// 4. Computing machine counts based on crafting times
	// 5. Minimizing resource waste and bottlenecks

	// Placeholder calculation
	for _, target := range targets {
		recipes := opt.RecipeGraph.GetRecipesForItem(target.Item)
		if len(recipes) > 0 {
			recipe := recipes[0] // Use first available recipe for now

			// Calculate machines needed (simplified)
			itemsPerSecond := target.Rate / 60.0
			outputPerSecond := recipe.Outputs[target.Item] / recipe.CraftingTime
			machinesNeeded := math.Ceil(itemsPerSecond / outputPerSecond)

			plan.RequiredMachines[recipe.Name] = int(machinesNeeded)
			plan.ResourceFlow[target.Item] = target.Rate
		}
	}

	return plan, nil
}

// IsRecipeAvailable checks if a recipe can be used with current research.
func (opt *Optimizer) IsRecipeAvailable(recipeName string) bool {
	// TODO: Implement technology dependency checking
	// For now, assume all recipes are available
	return true
}
