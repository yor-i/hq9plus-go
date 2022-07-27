package main

import (
	"fmt"
	"os"
)

func main() {
	// Open and load file into array
	if len(os.Args) < 2 {
		fmt.Println("One argument (file path) must be supplied.")
		return
	}

	f, e := os.Open(os.Args[1])
	if e != nil {
		fmt.Println("Couldn't open file.")
		return
	}

	stat, e := f.Stat()
	if e != nil {
		fmt.Println("Couldn't get file info.")
		return
	}

	bytes := make([]byte, stat.Size())

	_, e = f.Read(bytes)
	if e != nil {
		fmt.Println("Couldn't read from file.")
		return
	}

	// Accumulator variable (very important)
	accumulator := 0

	// Looping through the file
	for i := 0; i < len(bytes); i++ {
		// Switch to execute the current instruction
		switch bytes[i] {
		// Prints "Hello, World!"
		case 'h':
			fmt.Println("Hello, World!")
			break
		// Prints the program's source code
		case 'q':
			fmt.Println(string(bytes))
			break
		// Prints the lyrics of "99 Bottles of Beer"
		case '9':
			for i := 99; i > 1; i-- {
				fmt.Println(i, "bottles of beer on the wall,")
				fmt.Println(i, "bottles of beer.")
				fmt.Println("Take one down, pass it around.")
			}

			fmt.Println("1 bottle of beer on the wall,")
			fmt.Println("1 bottle of beer.")
			fmt.Println("Take one down, pass it around.")
			fmt.Println("No bottles of beer on the wall.")
			break
		// Incremens accumulator
		case '+':
			accumulator++
			break
		}
	}
}
