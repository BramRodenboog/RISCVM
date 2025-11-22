package core

import (
	"fmt"

	"github.com/BramRodenboog/RISCVM/core/opcodes"
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
	if wb, ok := instrDecoded.(opcodes.WriteBack); ok {
		cpu.writeBack(wb, value)
	}
	fmt.Printf("registers: %v\n", cpu.dump())
}

func NewCpu() *cpu {
	return &cpu{
		registers: [32]uint32{},
		pc:        uint32(DRAMBASE),
		bus:       *newBus(),
	}
}

func (cpu *cpu) dump() [32]uint32 {
	return cpu.registers
}
