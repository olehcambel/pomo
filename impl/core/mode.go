package core

type mode int

// Mode is type of the activity
const (
	ModeWork mode = iota
	ModeRelax
)

func (m mode) String() string {
	switch m {
	case ModeWork:
		return tModeWork
	case ModeRelax:
		return tModeRelax
	}

	return tWrongMode
}

// getSwap returns toggled mode (maybe rename)
func (m mode) getSwap() mode {
	// func (m *mode) swap() mode {
	switch m {
	case ModeWork:
		return ModeRelax
	case ModeRelax:
		return ModeWork
	}

	return m
}
