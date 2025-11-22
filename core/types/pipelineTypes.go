package types

type MemData struct {
	Address uint32
	Size    uint
	Data    uint32
}

type ChainPipe struct {
	Instr   uint32
	IsaType interface{}
	MemData
}
