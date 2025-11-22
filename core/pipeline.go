package core

import (
	"fmt"

	"github.com/BramRodenboog/RISCVM/core/opcodes"
	pipelineTypes "github.com/BramRodenboog/RISCVM/core/types"
)

// ToDo cleanup Go-routine 
func (cpu *cpu) fetch(pipelineChannel chan pipelineTypes.ChainPipe) {
	pipelineChannel <- pipelineTypes.ChainPipe{Instr: cpu.bus.busLoad(cpu.pc, 32)}
	close(pipelineChannel)
}

// ToDo cleanup Goroutine
func (cpu *cpu) identify(pipelineChannel chan pipelineTypes.ChainPipe) (interface{}, error) {
	dataChain, ok := <-pipelineChannel
	if !ok {
		return nil, fmt.Errorf("channel clossed")
	}
	Opcode := uint8(dataChain.Instr & 0x7f)
	switch Opcode {
	case opcodes.ITypeCode:
		return opcodes.DecodeIType(dataChain.Instr), nil
	default:
		return opcodes.DecodeIType(dataChain.Instr), fmt.Errorf("unsupported opcode: 0x%x", Opcode)
	}
}

func (cpu *cpu) execute(Type interface{}) (uint32, error) {
	switch isaType := Type.(type) {
	case opcodes.IType:
		switch isaType.Func3 {
		case opcodes.FUNCT3_ADDI:
			Rs1Value := cpu.registers[isaType.Rs1]
			Imm := isaType.Imm
			return cpu.alu.ADDI(Rs1Value, uint32(Imm)), nil
		default:
			return 0, fmt.Errorf("unsuppored func3: 0x%x", isaType.Func3)
		}
	default:
		return 0, fmt.Errorf("unsupported isa type: 0x%x", isaType)
	}
}

func (cpu *cpu) memory(address uint32, size uint, data uint32) {
	cpu.bus.busStore(address, size, data)
}

func (cpu *cpu) writeBack(Type opcodes.WriteBack, value uint32) {
	cpu.registers[Type.GetRd()] = value
}
