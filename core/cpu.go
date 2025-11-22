package core

import (
	"fmt"

	pkg "github.com/BramRodenboog/RISCVM/pkg"
)

type cpu struct {
	registers [32]uint32
	pc        uint32
	bus       bus
	alu       alu
}

func (cpu *cpu) run() {
	instr := cpu.fetch()

	instrDecoded := pkg.ErrorCheck(cpu.identify(instr))
	value := pkg.ErrorCheck(cpu.execute(instrDecoded))

	fmt.Printf("value: %v\n", value)
}

func NewCpu() *cpu {
	return &cpu{
		registers: [32]uint32{},
		pc:        uint32(DRAMBASE),
		bus:       *newBus(),
	}
}
