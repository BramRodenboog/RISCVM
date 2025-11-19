package core

type cpu struct {
	registers [32]uint32
	pc        uint32
	bus       bus
}
