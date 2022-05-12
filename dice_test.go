package roll_test

import (
	"testing"

	"github.com/brittonhayes/roll"
	"github.com/stretchr/testify/assert"
)

func TestDice_String(t *testing.T) {
	type fields struct {
		Min int
		Max int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "D6", fields: fields{Min: 1, Max: 6}, want: "D6"},
		{name: "D12", fields: fields{Min: 1, Max: 12}, want: "D12"},
		{name: "D20", fields: fields{Min: 1, Max: 20}, want: "D20"},
		{name: "D100", fields: fields{Min: 1, Max: 100}, want: "D100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := roll.Die{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			assert.Equal(t, tt.want, d.String())
		})
	}
}

func TestDice_Roll(t *testing.T) {

	t.Run("rolled greater than or equal to dice minimum", func(t *testing.T) {
		d := &roll.Die{
			Min: 1,
			Max: 6,
		}

		got := d.Roll()
		assert.GreaterOrEqual(t, got, d.Min)
	})

	t.Run("rolled less than or equal to dice maximum", func(t *testing.T) {
		d := &roll.Die{
			Min: 1,
			Max: 6,
		}

		got := d.Roll()
		assert.LessOrEqual(t, got, d.Max)
	})
}

func TestDice_Validate(t *testing.T) {
	dice := []*roll.Die{
		{Min: 1, Max: 6},
		{Min: 1, Max: 20},
		{Min: 1, Max: 100},
		{Min: 1, Max: 12},
		{Min: 1, Max: 8},
	}

	t.Run("Dice are valid", func(t *testing.T) {
		for _, d := range dice {
			err := d.Validate()
			assert.NoError(t, err)
		}
	})
}

func TestNewD6(t *testing.T) {
	t.Run("New D6 has min 1 and max 6", func(t *testing.T) {
		want := &roll.Die{Min: 1, Max: 6}
		got := roll.NewD6()

		assert.Equal(t, want, got)
	})
}
