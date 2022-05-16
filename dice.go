package roll

import (
	crypto "crypto/rand"
	"encoding/binary"
	"fmt"
	rand "math/rand"

	"github.com/go-playground/validator/v10"
)

func init() {
	var b [8]byte
	_, err := crypto.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}

	// Generate cryptographically random seed
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

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
	return fmt.Sprintf("d%d", d.Max)
}

// Validate ensures that the Dice are in a valid
// configuration
func (d *Die) Validate() error {
	return validate.Struct(d)
}

// Roll performs a roll of a single die
func (d *Die) Roll() int {
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
