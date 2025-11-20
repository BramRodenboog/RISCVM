package main

import (
	core "github.com/BramRodenboog/RISCVM/core"
)

func main() {
	cpu := core.NewCpu()
	program := core.NewProgram([]uint32{0x00100093})
	program.Run(cpu)
}
