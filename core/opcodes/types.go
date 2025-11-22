package opcodes

type WriteBack interface {
	GetRd() uint8
}

type CanWriteBack struct {
	Rd uint8
}

func (c CanWriteBack) GetRd() uint8 {
	return c.Rd
}

type IType struct {
	Imm    int32 // 12 bits
	Rs1    uint8
	Func3  uint8
	Opcode uint8
	Shamt  uint8
	CanWriteBack
}
