package parse

import (
	"testing"

	"github.com/brittonhayes/roll"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
}

func TestMatch_valid(t *testing.T) {
	tests := []struct {
		name  string
		got   string
		valid bool
	}{
		{
			name: "pattern is valid",
			got:  "1d12",
		},
		{
			name: "pattern with addition modifier is valid",
			got:  "1d6+3",
		},
		{
			name: "pattern with subtraction modifier is valid",
			got:  "1d6-3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Match(tt.got)
			assert.NoError(t, err)
		})
	}
}

func TestMatch_invalid(t *testing.T) {
	tests := []struct {
		name  string
		got   string
		valid bool
	}{
		{
			name: "pattern with no quantity is invalid",
			got:  "d6",
		},
		{
			name: "pattern with no quantity and modifier is invalid",
			got:  "d6+1",
		},
		{
			name: "pattern with zero quantity is invalid",
			got:  "0d6",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Match(tt.got)
			assert.Error(t, err)
		})
	}
}

func TestMatch(t *testing.T) {
	type fields struct {
		input string
		want  []string
	}

	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{"can parse 1d6", fields{input: "1d6"}, []string{"1", "d6", "", ""}},
		{"can parse 1d6 with modifier", fields{input: "1d6+3"}, []string{"1", "d6", "+", "3"}},
		{"can parse 1d6 with subtraction modifier", fields{input: "1d6-3"}, []string{"1", "d6", "-", "3"}},
		{"can parse 1d13 with subtraction modifier", fields{input: "1d13-3"}, []string{"1", "d13", "-", "3"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Match(tt.fields.input)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestParser(t *testing.T) {
	t.Run("create parser and roll", func(t *testing.T) {
		p, err := NewParser("1d6+2")

		if assert.NoError(t, err) {
			assert.EqualValues(t, Parser{
				Quantity: 1,
				Dice:     []*roll.Die{{Min: 1, Max: 6}},
				Operator: "+",
				Modifier: 2},
				*p)

			assert.Positive(t, p.Roll())
		}
	})
}
