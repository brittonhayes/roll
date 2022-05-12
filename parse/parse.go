package parse

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/brittonhayes/roll"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var validate = validator.New()

const (
	indexQuantity int = iota
	indexDie
	indexOperator
	indexModifier
)

type Parser struct {
	Quantity int
	Dice     []*roll.Die
	Operator string
	Modifier int
}

const pattern = `^(?P<quantity>\d+)(?P<die>d\d+)(?P<operator>[-+]?)(?P<modifier>\d+?)?$`

func Match(s string) ([]string, error) {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	if regex.MatchString(s) {
		return regex.FindStringSubmatch(s)[1:], nil
	}

	return nil, errors.New("no matches found for dice roll input")
}

func NewParser(s string) (*Parser, error) {
	matches, err := Match(s)
	if err != nil {
		return nil, err
	}

	p := &Parser{}

	// Apply quantity to formula
	p.Quantity, err = strconv.Atoi(matches[indexQuantity])
	if err != nil {
		return nil, err
	}

	// Find max die integer
	dieInt := strings.ReplaceAll(matches[indexDie], "d", "")
	max, err := strconv.Atoi(dieInt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert die signature to int")
	}

	for i := 0; i <= p.Quantity; i++ {
		// Apply die to formula
		die, err := roll.NewDie(1, max)
		if err != nil {
			return nil, err
		}

		p.Dice = append(p.Dice, die)
	}

	// Apply modifer to formula
	if matches[indexOperator] != "" && matches[indexModifier] != "" {
		// Apply operator to formula
		p.Operator = matches[indexOperator]

		p.Modifier, err = strconv.Atoi(matches[indexModifier])
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert modifier to int")
		}
	}

	log.Info().Msgf("found %q", p.String())

	return p, nil
}

func (p *Parser) String() string {
	return fmt.Sprintf("%d %s %s %d", p.Quantity, p.Dice[0], p.Operator, p.Modifier)
}

// Roll rolls the dice with modifiers
// found in the parsed string
func (p *Parser) Roll() int {
	results := 0
	for i, die := range p.Dice {
		roll := die.Roll()
		log.Info().Msgf("rolled %d with %s", roll, p.Dice[i])

		results += roll
	}

	switch p.Operator {
	case "+":
		log.Info().Msgf("adding %d", p.Modifier)
		return results + p.Modifier
	case "-":
		log.Info().Msgf("subtracting %d", p.Modifier)
		return results - p.Modifier
	default:
		return results
	}
}
