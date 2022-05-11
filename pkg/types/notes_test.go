package types_test

import (
	"testing"

	"github.com/mimikwang/scales/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestNotesFindKeyPosition(t *testing.T) {
	testCases := []struct {
		name        string
		notes       types.Notes
		key         string
		expected    int
		expectedErr error
	}{
		{
			"Should return the key position for a sharp scale",
			types.Sharp,
			"E",
			4,
			nil,
		},
		{
			"Should return the key position for a flat scale",
			types.Flat,
			"Gb",
			6,
			nil,
		},
		{
			"Should return the key position for a flat2 scale",
			types.Flat2,
			"Cb",
			11,
			nil,
		},
		{
			"Should return an error for an invalid key",
			types.Flat,
			"Z",
			-1,
			types.ErrInvalidNote,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := testCase.notes.FindKeyPosition(testCase.key)
			assert.Equal(t, testCase.expected, actual)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestNotesGetNote(t *testing.T) {
	testCases := []struct {
		name     string
		notes    types.Notes
		index    int
		expected string
	}{
		{
			"Should return the right note for a sharp scale",
			types.Sharp,
			8,
			"G#",
		},
		{
			"Should return the right note for a flat scale",
			types.Flat,
			8,
			"Ab",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := testCase.notes.GetNote(testCase.index)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func TestTraverse(t *testing.T) {
	testCases := []struct {
		name     string
		index    int
		slice    []string
		expected string
	}{
		{
			"Should return an empty string",
			10,
			[]string{},
			"",
		},
		{
			"Should return the element at index",
			1,
			[]string{"a", "b"},
			"b",
		},
		{
			"Should handle a negative index",
			-1,
			[]string{"a", "b"},
			"b",
		},
		{
			"Should circle back and return the element at index",
			9,
			[]string{"a", "b"},
			"b",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := types.Traverse(testCase.index, testCase.slice)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func TestFindPosition(t *testing.T) {
	testCases := []struct {
		name     string
		element  string
		slice    []string
		expected int
	}{
		{
			"Should return the index of the element",
			"A",
			[]string{"R", "I", "T", "A"},
			3,
		},
		{
			"Should return -1 if element not found",
			"A",
			[]string{},
			-1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := types.FindPosition(testCase.element, testCase.slice)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
