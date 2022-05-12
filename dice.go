package roll

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Die is the struct for all
// rollable dice types
type Die struct {
	Min int `validate:"ltfield=Max"`
	Max int `validate:"gtfield=Min"`
}

func NewDie(min, max int) (*Die, error) {
	d := &Die{Min: min, Max: max}

	err := validate.Struct(d)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// String returns the string representation of
// the die
func (d Die) String() string {
	return fmt.Sprintf("D%d", d.Max)
}

// Validate ensures that the Dice are in a valid
// configuration
func (d *Die) Validate() error {
	return validate.Struct(d)
}

// Roll performs a roll of a single die
func (d *Die) Roll() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(d.Max-d.Min+1) + d.Min
}

// NewD6 creates a new instance
// of the D6 dice type
func NewD6() *Die {
	return &Die{
		Min: 1,
		Max: 6,
	}
}
