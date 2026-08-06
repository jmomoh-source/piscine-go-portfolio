package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}

	for _, arg := range os.Args[1:] {
		if len(arg) >= 2 && arg[0] == '-' && arg[1] == 'h' {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		}
	}

	var options uint32

	for _, arg := range os.Args[1:] {
		if len(arg) < 2 || arg[0] != '-' {
			fmt.Println("Invalid Option")
			return
		}
		for i := 1; i < len(arg); i++ {
			ch := arg[i]
			if ch < 'a' || ch > 'z' {
				fmt.Println("Invalid Option")
				return
			}
			options |= (1 << (ch - 'a'))
		}
	}

	b3 := (options >> 24) & 0xFF
	b2 := (options >> 16) & 0xFF
	b1 := (options >> 8) & 0xFF
	b0 := options & 0xFF

	fmt.Printf("%08b %08b %08b %08b\n", b3, b2, b1, b0)
}
