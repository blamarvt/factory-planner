# Factorio Factory Planner

A sophisticated tool for automatically generating efficient Factorio factory layouts based on your unlocked research and production goals.

## Overview

The Factorio Factory Planner takes your current research progress and production requirements as input, then generates optimized factory designs that maximize efficiency while minimizing resource waste and space usage. The tool outputs visual factory layouts that can be directly implemented in-game.

## Goals and Features

### Core Features
- **Research-Aware Planning**: Analyzes your unlocked technologies to determine available recipes and production methods
- **Efficient Layout Generation**: Creates optimized factory designs with minimal belt congestion and resource waste
- **Visual Output**: Generates clear, implementable factory blueprints and images
- **Production Goal Optimization**: Calculates optimal production ratios for target items per minute/second
- **Resource Flow Analysis**: Traces material flow from raw resources to final products

### Advanced Features (Planned)
- **Modular Factory Design**: Creates reusable factory modules that can be expanded
- **Multi-Product Optimization**: Handles complex production chains with multiple end products
- **Space Efficiency Analysis**: Optimizes for compact designs when space is limited
- **Power and Infrastructure Planning**: Includes power pole and logistics infrastructure in designs
- **Blueprint Export**: Direct integration with Factorio's blueprint system

## Technology Stack Options

### Backend Options
1. **Python-based Solution**
   - **Pros**: Rich data science ecosystem, excellent for optimization algorithms, mature libraries
   - **Libraries**: NetworkX (graph algorithms), SciPy (optimization), PIL/Pillow (image generation)
   - **Frameworks**: Flask/FastAPI for web interface, Click for CLI

2. **JavaScript/Node.js Solution**
   - **Pros**: Single language for full-stack, excellent web integration, active community
   - **Libraries**: D3.js (visualization), Canvas API (image generation), Express.js (web server)
   - **Frameworks**: React/Vue.js for frontend, Electron for desktop app

3. **Rust Solution**
   - **Pros**: High performance, memory safety, excellent for complex algorithms
   - **Libraries**: Petgraph (graph algorithms), Image crate (image processing), Serde (data serialization)
   - **Frameworks**: Actix-web for API, Tauri for desktop interface

### Frontend Options
1. **Web Application**
   - **React/TypeScript**: Component-based UI with strong typing
   - **Vue.js**: Progressive framework with excellent developer experience
   - **Svelte**: Compile-time optimizations for fast rendering

2. **Desktop Application**
   - **Electron**: Cross-platform with web technologies
   - **Tauri**: Lightweight Rust-based alternative to Electron
   - **Native**: Platform-specific apps (Qt, GTK, WinUI)

3. **Command-Line Interface**
   - **JSON/YAML input**: Configuration files for batch processing
   - **Interactive CLI**: Step-by-step factory planning wizard

### Data Storage Options
1. **File-based**: JSON/YAML for recipes, technologies, and user configurations
2. **SQLite**: Lightweight database for complex queries and relationships
3. **Graph Database**: Neo4j or similar for recipe dependency graphs

## Project Structure (Planned)

```
factory-planner/
├── src/
│   ├── core/              # Core planning algorithms
│   │   ├── recipe_parser.py
│   │   ├── optimization.py
│   │   └── layout_generator.py
│   ├── data/              # Factorio game data
│   │   ├── recipes.json
│   │   ├── technologies.json
│   │   └── items.json
│   ├── visualization/     # Image and blueprint generation
│   │   ├── factory_renderer.py
│   │   └── blueprint_exporter.py
│   └── api/               # Web API or CLI interface
│       ├── routes.py
│       └── models.py
├── tests/                 # Unit and integration tests
├── docs/                  # Documentation and examples
├── examples/              # Sample configurations and outputs
└── data/                  # Game data and user configurations
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

### 4. Visualization
- Render factory layouts as clear, labeled images
- Export as Factorio-compatible blueprints
- Provide material lists and construction guides

## Getting Started (Coming Soon)

The tool is currently in early development. Once ready, it will support:

```bash
# Install the factory planner
pip install factorio-factory-planner

# Generate a factory for basic science production
factory-planner --research basic-science --target "science-pack-1:60/min"

# Web interface
factory-planner serve --port 8080
```

## Contributing

This project is in the initial planning phase. Contributions and feedback are welcome! 

### Development Setup (Future)
1. Clone the repository
2. Install dependencies
3. Run tests
4. Start the development server

### Areas for Contribution
- Algorithm optimization and efficiency improvements
- UI/UX design for the planning interface
- Factorio game data parsing and updates
- Performance benchmarking and testing
- Documentation and examples

## License

This project will be open source under the MIT License. The tool is not affiliated with Factorio or Wube Software Ltd.

## Acknowledgments

- Factorio by Wube Software Ltd.
- The Factorio community for game data and optimization insights
- Open source optimization and visualization libraries