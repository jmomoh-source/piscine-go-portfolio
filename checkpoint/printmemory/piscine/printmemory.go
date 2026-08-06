package piscine

import "fmt"

func PrintMemory(arr [10]byte) {
	fmt.Printf("%02x %02x %02x %02x\n", arr[0], arr[1], arr[2], arr[3])
	fmt.Printf("%02x %02x %02x %02x\n", arr[4], arr[5], arr[6], arr[7])
	fmt.Printf("%02x %02x\n", arr[8], arr[9])

	for _, b := range arr {
		if b >= 32 && b <= 126 {
			fmt.Printf("%c", b)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}
