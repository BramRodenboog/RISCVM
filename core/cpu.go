package core

import (
	"fmt"

	pkg "github.com/BramRodenboog/RISCVM/pkg"
)

type cpu struct {
	registers [32]uint32
	pc        uint32
	bus       bus
}

func (cpu *cpu) run() {
	instr := cpu.fetch()
	instrDecoded := pkg.ErrorCheck(cpu.identify(instr))
	fmt.Printf("InstrDecoded: %v\n", instrDecoded)
}
