package core

type program struct {
	code []uint32
}

func (program *program) Run(cpu *cpu) {

	var addr uint32 = DRAMBASE

	for _, instruction := range program.code {
		cpu.bus.busStore(addr, 32, instruction)
		addr += 4
	}

	for {
		cpu.run()
		cpu.pc += 4 // !!!

	}
}

func NewProgram(code []uint32) *program {
	return &program{
		code: code,
	}
}
