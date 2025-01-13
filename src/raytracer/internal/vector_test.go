package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVectorLength(t *testing.T) {
	v := Vector{3, 4, 0}
	expected := 5.0
	result := v.Length()
	assert.Equal(t, expected, result, "they should be equal")
}

func TestVectorDot(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{4, 5, 6}
	expected := 32.0
	result := v1.Dot(v2)
	assert.Equal(t, expected, result, "they should be equal")
}

func TestVectorNormalize(t *testing.T) {
	v := Vector{3, 4, 0}
	expected := Vector{0.6, 0.8, 0}
	result := v.Normalize()
	assert.Equal(t, expected, result, "they should be equal")
}

func TestVectorMultiplyScalar(t *testing.T) {
	v := Vector{1, 2, 3}
	scalar := 2.0
	expected := Vector{2, 4, 6}
	result := v.MultiplyScalar(scalar)
	assert.Equal(t, expected, result, "they should be equal")
}

func TestVectorAddScalar(t *testing.T) {
	v := Vector{1, 2, 3}
	scalar := 2.0
	expected := Vector{3, 4, 5}
	result := v.AddScalar(scalar)
	assert.Equal(t, expected, result, "they should be equal")
}

func TestVectorSubtract(t *testing.T) {
	v1 := Vector{5, 6, 7}
	v2 := Vector{1, 2, 3}
	expected := Vector{4, 4, 4}
	result := v1.Subtract(v2)
	assert.Equal(t, expected, result)
}
