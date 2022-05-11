package types

// Scale represents a musical scale
type Scale struct {
	notes           Notes
	key             string
	startPosition   int
	currentPosition int
}

// NewScale constructs a new scale
func NewScale(notes Notes, key string) (Scale, error) {
	startPosition, err := notes.FindKeyPosition(key)
	if err != nil {
		return Scale{}, err
	}
	return Scale{
		notes:           notes,
		key:             key,
		startPosition:   startPosition,
		currentPosition: 0,
	}, nil
}

// GetNote returns the note at index in a scale
func (s Scale) GetNote(index int) string {
	return s.notes.GetNote(s.startPosition + index)
}

// GetCurrentNote returns the note at the current position
func (s Scale) GetCurrentNote() string {
	return s.GetNote(s.currentPosition)
}

// Move moves the scale based on the interval
func (s *Scale) Move(interval Interval) {
	s.currentPosition += int(interval)
}
