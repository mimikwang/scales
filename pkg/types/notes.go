package types

// Notes is an enum that defines the notes in a scale
type Notes int

const (
	Sharp Notes = iota
	Flat
	Flat2
)

// Define the notes here
var (
	sharpNotes []string = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	flatNotes  []string = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	flat2Notes []string = []string{"C", "Db", "D", "Eb", "Fb", "F", "Gb", "G", "Ab", "A", "Bb", "Cb"}

	notesLookUp map[Notes][]string = map[Notes][]string{
		Sharp: sharpNotes,
		Flat:  flatNotes,
		Flat2: flat2Notes,
	}
)

// FindKeyPosition finds the position of the key in a scale
func (n Notes) FindKeyPosition(key string) (int, error) {
	position := findPosition(key, notesLookUp[n])
	if position < 0 {
		return position, ErrInvalidNote
	}
	return position, nil
}

// GetNote returns the note in a scale at index
func (n Notes) GetNote(index int) string {
	return traverse(index, notesLookUp[n])
}

// findPosition finds the position of element in a slice.  -1 is returned if the element does not
// exist in the slice.
func findPosition(element string, slice []string) int {
	for index, s := range slice {
		if s == element {
			return index
		}
	}
	return -1
}

// traverse takes in an index and a slice and returns the element as if the slice is circular.
func traverse(index int, slice []string) string {
	// Return an empty string if the slice is empty
	if len(slice) == 0 {
		return ""
	}
	for {
		if index >= 0 {
			break
		}
		index += len(slice)
	}
	index = index % len(slice)
	return slice[index]
}
