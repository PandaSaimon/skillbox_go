package main

import (
	"fmt"
	"math"
)

func main() {
	var uint8Count uint32 = 1
	var uint16Count uint32

	for ; uint32(math.MaxUint8*uint8Count) < math.MaxUint32; uint8Count++ {
		if uint32(uint8Count*math.MaxUint8) == uint32(math.MaxUint16*(uint16Count+1)) {
			uint16Count++
		}
	}
	uint16Count++
	fmt.Println("Переполнений uint8:", uint8Count, "(", math.MaxUint32/uint32(math.MaxUint8), ")")
	fmt.Println("Переполнений uint16:", uint16Count, "(", math.MaxUint32/uint32(math.MaxUint16), ")")
}

//более разумная реализация

/*{
	var uint8Count uint32 = math.MaxUint32 / uint32(math.MaxUint8)
	var uint16Count uint32 = math.MaxUint32 / uint32(math.MaxUint16)

	fmt.Println(math.MaxUint8, math.MaxUint16, math.MaxUint32)

	fmt.Println("Переполнений uint8", uint8Count)
	fmt.Println("Переполнений uint16", uint16Count)
}*/
