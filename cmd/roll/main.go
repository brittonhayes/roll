package main

import (
	"fmt"
	"os"
	"time"

	"github.com/alecthomas/kong"
	"github.com/briandowns/spinner"
	"github.com/brittonhayes/roll/parse"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var CLI struct {
	Verbose bool `help:"Display verbose log output" short:"v" env:"VERBOSE" default:"false"`

	SkipSpinner bool   `help:"Skip loading spinner" short:"s" env:"SKIP_SPINNER" default:"false"`
	Dice        string `arg:"" help:"Dice to roll +/- modifiers e.g. 'roll 1d6', 'roll 2d12+20', or 'roll 1d20-5'" required:""`
}

func main() {
	ctx := kong.Parse(&CLI,
		kong.Name("roll"),
		kong.Description("A simple CLI for dice rolling"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}))

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.Disabled)
	if *&CLI.Verbose {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	p, err := parse.NewParser(CLI.Dice)
	ctx.FatalIfErrorf(err)

	s := spinner.New(spinner.CharSets[2], 100*time.Millisecond, spinner.WithSuffix(" Rolling..."))
	s.Start()
	defer s.Stop()
	result := p.Roll()
	s.FinalMSG = fmt.Sprintf("🎲 %d\n", result)

	if !CLI.SkipSpinner {
		time.Sleep(400 * time.Millisecond)
	}
}
