package core

import (
	pkg "github.com/BramRodenboog/RISCVM/pkg"
)

type bus struct {
	dram dram
}

func (bus *bus) busLoad(address uint32, size uint) uint32 {

	return pkg.ErrorCheck(bus.dram.dramLoad(address, size))
}
