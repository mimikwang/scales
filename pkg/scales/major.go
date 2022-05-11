package scales

import (
	"fmt"

	"github.com/mimikwang/scales/pkg/types"
)

// Major generates a slice which contains the notes of a major scale
func Major(key string) ([]string, error) {
	intervals := []types.Interval{
		types.Whole,
		types.Whole,
		types.Half,
		types.Whole,
		types.Whole,
		types.Whole,
		types.Half,
	}
	lookup := map[string]types.Notes{
		"C":  types.Sharp,
		"G":  types.Sharp,
		"D":  types.Sharp,
		"A":  types.Sharp,
		"E":  types.Sharp,
		"B":  types.Sharp,
		"F#": types.Sharp,
		"C#": types.Sharp,

		"F":  types.Flat,
		"Bb": types.Flat,
		"Eb": types.Flat,
		"Ab": types.Flat,

		"Cb": types.Flat2,
		"Gb": types.Flat2,
		"Db": types.Flat,
	}

	if _, ok := lookup[key]; !ok {
		return []string{}, ErrKeyNotFound
	}

	scale, err := types.NewScale(lookup[key], key)
	if err != nil {
		return []string{}, err
	}
	fmt.Println(scale)

	output := make([]string, len(intervals)+1)
	output[0] = scale.GetCurrentNote()
	for i, interval := range intervals {
		scale.Move(interval)
		output[i+1] = scale.GetCurrentNote()
	}

	return output, nil
}
