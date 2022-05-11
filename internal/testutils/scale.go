package testutils

import "github.com/mimikwang/scales/pkg/types"

// NewScale is a helper constructor function for Scale
func NewScale(notes types.Notes, key string) types.Scale {
	scale, _ := types.NewScale(notes, key)
	return scale
}
