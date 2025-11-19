package opcodes

type IType struct {
	Imm    int32 // 12 bits
	Rs1    uint8
	Func3  uint8
	Rd     uint8
	Opcode uint8
	Shamt  uint8
}
