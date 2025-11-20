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
	offset := address - DRAMBASE
	return uint32(dram.mem[offset])
}

func (dram *dram) dramLoad16(address uint32) uint32 {
	offset := address - DRAMBASE
	return uint32(dram.mem[offset]) |
		(uint32(dram.mem[offset+1]) << 8)
}

func (dram *dram) dramLoad32(address uint32) uint32 {
	offset := address - DRAMBASE
	return uint32(dram.mem[offset]) |
		(uint32(dram.mem[offset+1]) << 8) |
		(uint32(dram.mem[offset+2]) << 16) |
		(uint32(dram.mem[offset+3]) << 24)
}

func (dram *dram) dramStore(address uint32, size uint, data uint32) (uint8, error) {
	switch size {
	case 8:
		dram.dramStore8(address, data)
		return 0, nil
	case 16:
		dram.dramStore16(address, data)
		return 0, nil
	case 32:
		dram.dramStore32(address, data)
		return 0, nil
	default:
		return 1, fmt.Errorf("unsupported size: %d", size)
	}
}

func (dram *dram) dramStore8(address uint32, data uint32) {
	dram.mem[address-DRAMBASE] = uint8(data)
}

func (dram *dram) dramStore16(address uint32, data uint32) {
	dram.mem[address-DRAMBASE] = uint8(data & 0xff)
	dram.mem[address-DRAMBASE+1] = uint8((data >> 8) & 0xff)
}

func (dram *dram) dramStore32(address uint32, data uint32) {
	dram.mem[address-DRAMBASE] = uint8(data & 0xff)
	dram.mem[address-DRAMBASE+1] = uint8((data >> 8) & 0xff)
	dram.mem[address-DRAMBASE+2] = uint8((data >> 16) & 0xff)
	dram.mem[address-DRAMBASE+3] = uint8((data >> 24) & 0xff)
}
func newDram() *dram {
	return &dram{
		mem: [DRAMSIZE]uint8{},
	}
}
