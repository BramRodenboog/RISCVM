package core

import "fmt"

const DRAMSIZE = 1024 * 1024 * 1
const DRAMBASE = 0x80000000

type dram struct {
	mem [DRAMSIZE]uint8
}

func (dram *dram) dramLoad(address uint32, size uint) (uint32, error) {

	switch size {
	case 8:
		return dram.dramLoad8(address), nil
	case 16:
		return dram.dramLoad16(address), nil
	case 32:
		return dram.dramLoad32(address), nil
	default:
		return 0, fmt.Errorf("unsupported size: %d", size)
	}
}

func (dram *dram) dramLoad8(address uint32) uint32 {
	return uint32(dram.mem[address-DRAMBASE])
}

func (dram *dram) dramLoad16(address uint32) uint32 {
	return uint32(dram.mem[address-DRAMBASE] | (dram.mem[address-DRAMBASE+1] << 8))
}

func (dram *dram) dramLoad32(address uint32) uint32 {
	return uint32(dram.mem[address-DRAMBASE] | (dram.mem[address-DRAMBASE+1] << 8) | (dram.mem[address-DRAMBASE+2] << 16) | (dram.mem[address-DRAMBASE+3] << 24))
}
