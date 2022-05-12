# 🎲 Roll

[![Go Reference](https://pkg.go.dev/badge/github.com/brittonhayes/roll.svg)](https://pkg.go.dev/github.com/brittonhayes/roll)
![Latest Release](https://img.shields.io/github/v/release/brittonhayes/roll?label=latest%20release)
[![Lint](https://github.com/brittonhayes/roll/actions/workflows/lint.yml/badge.svg)](https://github.com/brittonhayes/roll/actions/workflows/lint.yml)
[![Test](https://github.com/brittonhayes/roll/actions/workflows/test.yml/badge.svg)](https://github.com/brittonhayes/roll/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/roll)](https://goreportcard.com/report/github.com/brittonhayes/roll)

A simple Go package and CLI for dice rolling.

## ⚡ Usage

How to use Roll as a CLI or a library

### CLI

Installation via Go or Docker

```bash
# Go
go install github.com/brittonhayes/roll/cmd/roll
```

```bash
# Docker
docker run --rm -it ghcr.io/brittonhayes/roll:latest 1d6+2
```

Using the command line tool

```bash
# Run the CLI
Usage: roll <dice>

A simple CLI for dice rolling

Arguments:
  <dice>    Dice to roll +/- modifiers e.g. 'roll 1d6', 'roll 2d12+20', or 'roll 1d20-5'

Flags:
  -h, --help            Show context-sensitive help.
  -v, --verbose         Display verbose log output ($VERBOSE)
  -s, --skip-spinner    Skip loading spinner ($SKIP_SPINNER)

# Roll a D6
roll 1d6

# Roll with modifiers
roll 1d6+2
```

## 📺 Preview

[![asciicast](https://asciinema.org/a/ylBbsg1NjKb8WVciMUnoJSEgG.svg)](https://asciinema.org/a/ylBbsg1NjKb8WVciMUnoJSEgG)

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
