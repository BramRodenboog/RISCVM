package core

import (
	"fmt"

	"github.com/BramRodenboog/RISCVM/core/opcodes"
)

// ToDo Goroutine
func (cpu *cpu) fetch() uint32 {
	return cpu.bus.busLoad(cpu.pc, 32)
}

// ToDo Goroutine
func (cpu *cpu) identify(instr uint32) (interface{}, error) {
	Opcode := uint8(instr & 0x7f)
	switch Opcode {
	case opcodes.ITypeCode:
		return opcodes.DecodeIType(instr), nil
	default:
		return nil, fmt.Errorf("unsupported opcode: %d", Opcode)
	}
}

func execute() {

}

func memory() {

}

func writeBack() {

}
