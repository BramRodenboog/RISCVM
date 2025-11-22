package opcodes

func DecodeIType(instr uint32) IType {
	return IType{
		Imm:    int32((instr & 0xfff00000) >> 20),
		Rs1:    uint8((instr >> 15) & 0x1f),
		Func3:  uint8((instr >> 12) & 0x7),
		Opcode: uint8(instr & 0x7f),
		Shamt:  uint8(((instr >> 20) & 0xfff) & 0x1f),

		CanWriteBack: CanWriteBack{
			Rd: uint8((instr >> 7) & 0x1f),
		},
	}
}
