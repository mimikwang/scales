package types_test

import (
	"testing"

	"github.com/mimikwang/scales/internal/testutils"
	"github.com/mimikwang/scales/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestNewScale(t *testing.T) {
	testCases := []struct {
		name      string
		notes     types.Notes
		key       string
		expectErr bool
	}{
		{
			"Should not return an error on valid inputs",
			types.Flat,
			"Gb",
			false,
		},
		{
			"Should return an error on invalid inputs",
			types.Sharp,
			"Gb",
			true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := types.NewScale(testCase.notes, testCase.key)
			if testCase.expectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestScaleGetNote(t *testing.T) {
	testCases := []struct {
		name     string
		scale    types.Scale
		index    int
		expected string
	}{
		{
			"Should return the right note",
			testutils.NewScale(types.Sharp, "C"),
			0,
			"C",
		},
		{
			"Should return the right note",
			testutils.NewScale(types.Sharp, "C"),
			2,
			"D",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := testCase.scale.GetNote(testCase.index)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
