package core

type alu struct{}

func (alu *alu) ADDI(x, y uint32) uint32 {
	return x + y
}
