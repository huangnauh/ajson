package tokens

import (
	"fmt"
	"math"
	"strconv"
)

// Number is a temporary token
type Number struct {
	parent Token
	Alias  string
	Value  float64
	IsInt  bool
}

var _ Token = (*Number)(nil)

func NewNumber(value string) (*Number, error) {
	float, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, fmt.Errorf("value %q can't be converted to float64: %w", value, err)
	}
	return &Number{
		Alias: value,
		Value: float,
		IsInt: math.Mod(float, 1) == 0,
	}, nil
}

func (t *Number) Type() string {
	return "Number"
}

func (t *Number) String() string {
	if t == nil {
		return "<nil>"
	}
	return t.Alias
}

func (t *Number) Token() string {
	if t == nil {
		return "Number(<nil>)"
	}
	return fmt.Sprintf("Number(%s)", t.Alias)
}

func (t *Number) Parent() Token {
	if t == nil {
		return nil
	}
	return t.parent
}

func (t *Number) SetParent(parent Token) {
	if t == nil {
		return
	}
	t.parent = parent
}
