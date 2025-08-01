package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		research  = flag.String("research", "", "Research progress level (e.g., 'basic-science')")
		target    = flag.String("target", "", "Production target (e.g., 'science-pack-1:60/min')")
		output    = flag.String("output", "", "Output file path for PNG image")
		blueprint = flag.Bool("blueprint", false, "Generate Factorio blueprint string")
	)
	flag.Parse()

	fmt.Println("Factorio Factory Planner")
	fmt.Println("========================")

	if *research == "" || *target == "" || *output == "" {
		fmt.Println("Error: Missing required parameters")
		fmt.Println("Usage: factory-planner --research <level> --target <item:rate> --output <file.png>")
		fmt.Println("Example: factory-planner --research basic-science --target \"science-pack-1:60/min\" --output factory.png")
		os.Exit(1)
	}

	fmt.Printf("Research level: %s\n", *research)
	fmt.Printf("Production target: %s\n", *target)
	fmt.Printf("Output file: %s\n", *output)

	if *blueprint {
		fmt.Println("Blueprint generation: enabled")
	}

	fmt.Println("\nTODO: Implement factory planning logic")
	fmt.Println("This is the basic program layout - core functionality coming soon!")
}
