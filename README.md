# Factorio Factory Planner

A sophisticated tool for automatically generating efficient Factorio factory layouts based on your unlocked research and production goals.

## Overview

The Factorio Factory Planner is a command-line tool that takes your current research progress and production requirements as input, then generates optimized factory designs. The tool outputs both visual PNG images and Factorio blueprint strings that can be directly imported into the game.

## Goals and Features

### Core Features
- **Research-Aware Planning**: Analyzes your unlocked technologies to determine available recipes and production methods
- **Efficient Layout Generation**: Creates optimized factory designs with minimal belt congestion and resource waste
- **PNG Output**: Generates clear, visual factory layouts for easy implementation
- **Blueprint Export**: Outputs Factorio blueprint strings for direct import into the game
- **Production Goal Optimization**: Calculates optimal production ratios for target items per minute/second

## Technology Stack

This is a command-line program written in **Go** that generates:
- **PNG images**: Visual factory layouts for easy implementation
- **Factorio blueprint strings**: Direct import into the game's blueprint system

## Project Structure

```
factory-planner/
├── cmd/                     # CLI application entry point
│   └── main.go
├── internal/                # Private application code
│   ├── core/                # Core planning algorithms
│   │   ├── recipe.go
│   │   ├── optimization.go
│   │   └── layout.go
│   ├── data/                # Game data structures
│   │   ├── recipes.go
│   │   ├── technologies.go
│   │   └── items.go
│   ├── render/              # PNG generation
│   │   └── factory_image.go
│   └── blueprint/           # Blueprint string generation
│       └── exporter.go
├── pkg/                     # Public API packages
├── testdata/                # Test fixtures and sample data
├── examples/                # Sample configurations and outputs
├── go.mod                   # Go module definition
└── go.sum                   # Go module checksums
```

## Algorithm Approach

### 1. Recipe Graph Analysis
- Parse Factorio recipe data into a directed acyclic graph (DAG)
- Identify production chains and dependencies
- Calculate material flow requirements

### 2. Constraint Satisfaction
- Apply research constraints (available recipes)
- Optimize for production targets (items/minute)
- Minimize resource waste and production bottlenecks

### 3. Layout Generation
- Translate optimized production ratios into physical layouts
- Consider belt throughput, inserter speeds, and assembler positioning
- Generate collision-free factory designs

### 4. Output Generation
- Render factory layouts as PNG images with clear labels
- Generate Factorio-compatible blueprint strings
- Provide material lists and construction guides

## Getting Started (Coming Soon)

The tool is currently in early development. Once ready, it will support:

```bash
# Build the factory planner
go build -o factory-planner cmd/main.go

# Generate a factory for basic science production
./factory-planner --research basic-science --target "science-pack-1:60/min" --output factory.png

# Generate both PNG and blueprint string
./factory-planner --research basic-science --target "science-pack-1:60/min" --output factory.png --blueprint
```

## Contributing

This project is in the initial planning phase. Contributions and feedback are welcome! 

### Development Setup (Future)
1. Clone the repository
2. Install Go 1.21 or later
3. Run `go mod tidy` to install dependencies
4. Run `go test ./...` to run tests
5. Build with `go build -o factory-planner cmd/main.go`

### Areas for Contribution
- Algorithm optimization and efficiency improvements
- Command-line interface design and usability
- Factorio game data parsing and updates
- Performance benchmarking and testing
- Documentation and examples

## License

This project will be open source under the MIT License. The tool is not affiliated with Factorio or Wube Software Ltd.

## Acknowledgments

- Factorio by Wube Software Ltd.
- The Factorio community for game data and optimization insights
- Open source optimization and visualization libraries