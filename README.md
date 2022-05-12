# 🎲 Roll

[![Go Reference](https://pkg.go.dev/badge/github.com/brittonhayes/roll.svg)](https://pkg.go.dev/github.com/brittonhayes/roll)
![Latest Release](https://img.shields.io/github/v/release/brittonhayes/roll?label=latest%20release)
[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/roll)](https://goreportcard.com/report/github.com/brittonhayes/roll)

A simple Go package and CLI for dice rolling.

## ⚡ Usage

How to use Roll as a CLI or a library

### CLI

Installation

```bash
# Install
go install github.com/brittonhayes/roll/cmd/roll
```

Using the command line tool

```bash
# Run the CLI
roll --help

Usage: roll <dice>

A simple CLI for dice rolling

Arguments:
  <dice>

Flags:
  -h, --help            Show context-sensitive help.
  -v, --verbose         Display verbose log output ($VERBOSE)
  -s, --skip-spinner    Skip loading spinner ($SKIP_SPINNER)

# Roll a D6
roll 1d6

# Roll with modifiers
roll 1d6+2
```

### Package

Using the package

```go
func main() {
	// Create a new d6
	d := roll.NewD6()

	// Roll the die
	fmt.Println("Rolling", d)
	result := d.Roll()

	// Print the result
	fmt.Println(result)
}
```
