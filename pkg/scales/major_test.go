package scales_test

import (
	"testing"

	"github.com/mimikwang/scales/pkg/scales"
	"github.com/stretchr/testify/assert"
)

func TestMajor(t *testing.T) {
	testCases := []struct {
		name      string
		key       string
		expected  []string
		expectErr bool
	}{
		{
			"Should return an error if key is not valid",
			"P",
			[]string{},
			true,
		},
		{
			"Should return the C major scale",
			"C",
			[]string{"C", "D", "E", "F", "G", "A", "B", "C"},
			false,
		},
		{
			"Should return the F major scale",
			"F",
			[]string{"F", "G", "A", "Bb", "C", "D", "E", "F"},
			false,
		},
		{
			"Should return the C flat major scale",
			"Cb",
			[]string{"Cb", "Db", "Eb", "Fb", "Gb", "Ab", "Bb", "Cb"},
			false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := scales.Major(testCase.key)
			assert.Equal(t, testCase.expected, actual)
			if testCase.expectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
