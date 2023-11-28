package functional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2, 3)
	assert.Equal(t, 6, result, "Expected sum of 1, 2, 3 to be 6")
}

func TestAddWithZero(t *testing.T) {
	result := Add(0, 0, 0)
	assert.Equal(t, 0, result, "Expected sum of 0, 0, 0 to be 0")
}

func TestAddWithNegativeNumbers(t *testing.T) {
	result := Add(-1, -2, -3)
	assert.Equal(t, -6, result, "Expected sum of -1, -2, -3 to be -6")
}

func TestAddCurry(t *testing.T) {
	curriedAdd := AddCurry(1)
	result := curriedAdd(2)(3)
	assert.Equal(t, 6, result, "Expected curried sum of 1, 2, 3 to be 6")
}

func TestAddCurryWithZero(t *testing.T) {
	curriedAdd := AddCurry(0)
	result := curriedAdd(0)(0)
	assert.Equal(t, 0, result, "Expected curried sum of 0, 0, 0 to be 0")
}

func TestAddCurryWithNegativeNumbers(t *testing.T) {
	curriedAdd := AddCurry(-1)
	result := curriedAdd(-2)(-3)
	assert.Equal(t, -6, result, "Expected curried sum of -1, -2, -3 to be -6")
}

func TestAddCurryBehaviour(t *testing.T) {
	assert.Equal(t, 6, AddCurry(1)(2)(3), "Expected sum of 1, 2, 3 to be 6")
}
