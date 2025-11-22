package main

import (
	core "github.com/BramRodenboog/RISCVM/core"
)

func main() {
	cpu := core.NewCpu()
	program := core.NewProgram([]uint32{0x3e800093})
	program.Run(cpu)
}
