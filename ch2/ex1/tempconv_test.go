package tempconv

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	errorFmt    string = "Expected %g, received %g"
	tolerance          = 0.00001
	invalidTemp        = -500
)

// Comparer ignores float rounding errors within a given tolerance
var comparer = cmp.Comparer(func(x, y float64) bool {
	diff := math.Abs(x - y)
	mean := math.Abs(x+y) / 2.0
	return (diff / mean) < tolerance
})

func TestValidCToF(t *testing.T) {
	rec := CToF(BoilingC)
	if !cmp.Equal(float64(BoilingF), float64(rec), comparer) {
		t.Errorf(errorFmt, BoilingF, rec)
	}
}

func TestInvalidCToF(t *testing.T) {
	if rec := CToF(invalidTemp); rec != AbsoluteZeroF {
		t.Errorf(errorFmt, AbsoluteZeroF, rec)
	}
}

func TestValidCToK(t *testing.T) {
	rec := CToK(BoilingC)
	if !cmp.Equal(float64(BoilingK), float64(rec), comparer) {
		t.Errorf(errorFmt, BoilingK, rec)
	}
}

func TestInvalidCToK(t *testing.T) {
	if rec := CToK(invalidTemp); rec != AbsoluteZeroK {
		t.Errorf(errorFmt, AbsoluteZeroK, rec)
	}
}

func TestValidFToC(t *testing.T) {
	rec := FToC(BoilingF)
	if !cmp.Equal(float64(BoilingC), float64(rec), comparer) {
		t.Errorf(errorFmt, BoilingC, rec)
	}
}

func TestInvalidFToC(t *testing.T) {
	if rec := FToC(invalidTemp); rec != AbsoluteZeroC {
		t.Errorf(errorFmt, AbsoluteZeroC, rec)
	}
}

func TestValidFToK(t *testing.T) {
	rec := FToK(BoilingF)
	if !cmp.Equal(float64(BoilingK), float64(rec), comparer) {
		t.Errorf(errorFmt, BoilingK, rec)
	}
}

func TestInvalidFToK(t *testing.T) {
	if rec := FToK(invalidTemp); rec != AbsoluteZeroK {
		t.Errorf(errorFmt, AbsoluteZeroK, rec)
	}
}

func TestValidKToC(t *testing.T) {
	rec := KToC(BoilingK)
	if !cmp.Equal(float64(BoilingC), float64(rec), comparer) {
		t.Errorf(errorFmt, BoilingC, rec)
	}
}

func TestInvalidKToC(t *testing.T) {
	if rec := KToC(invalidTemp); rec != AbsoluteZeroC {
		t.Errorf(errorFmt, AbsoluteZeroC, rec)
	}
}

func TestValidKToF(t *testing.T) {
	rec := KToF(BoilingK)
	if !cmp.Equal(float64(BoilingF), float64(rec), comparer) {
		t.Errorf(errorFmt, BoilingF, rec)
	}
}

func TestInvalidKToF(t *testing.T) {
	if rec := KToF(invalidTemp); rec != AbsoluteZeroF {
		t.Errorf(errorFmt, AbsoluteZeroF, rec)
	}
}
