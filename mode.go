package main

type mode int

const (
	modeWork mode = iota
	modeRelax
)

func (m mode) String() string {
	switch m {
	case modeWork:
		return tModeWork
	case modeRelax:
		return tModeRelax
	}

	return tWrongMode
}

// getSwap returns toggled mode (maybe rename)
func (m mode) getSwap() mode {
	// func (m *mode) swap() mode {
	switch m {
	case modeWork:
		return modeRelax
	case modeRelax:
		return modeWork
	}

	return m
}
