package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var PRINT_MEMORY_AT_THE_END = true

const (
	OP_MV_DATAP_LEFT  = '<' //Move data pointer to the left
	OP_MV_DATAP_RIGHT = '>' //Move data pointer to the right
	OP_INC_MEM        = '+' //Increment memory
	OP_DEC_MEM        = '-' //Decrement memory
	OP_INPUT          = ',' //Read key input and store it's ascii number as a byte in the current memory location
	OP_OUTPUT         = '.' //Output current byte as ASCII character
	OP_LEFT_BRACE     = '[' //If memory byte != 0; jump forward to the command after the matching ]
	OP_RIGHT_BRACE    = ']' //If memory byte != 0; jump backward to the command after the matching []
)

func Eval(instructions []byte) ([]byte, error) {
	stdinReader := bufio.NewReader(os.Stdin)

	newMemory := make([]byte, 10)
	memory := make([]byte, 10)
	pointer := 0

	for idx := 0; idx < len(instructions); idx++ {
		// Increase memory size if the pointer needs to move beyond the original bounds
		if pointer >= len(memory) {
			memory = append(memory, newMemory...)
		}
		switch instructions[idx] {
		case OP_MV_DATAP_RIGHT:
			pointer++

		case OP_MV_DATAP_LEFT:
			pointer--

		case OP_INC_MEM:
			memory[pointer]++

		case OP_DEC_MEM:

			memory[pointer]--
		case OP_OUTPUT:
			fmt.Printf("%c", memory[pointer])

		case OP_INPUT:
			input, err := stdinReader.ReadByte()
			if err != nil {
				return memory, err
			}
			memory[pointer] = input

		case OP_LEFT_BRACE:
			if memory[pointer] == 0 {
				i := 1
				for i > 0 {
					idx++
					if instructions[idx] == OP_LEFT_BRACE {
						i++
					} else if instructions[idx] == OP_RIGHT_BRACE {
						i--
					}
				}
			}
		case OP_RIGHT_BRACE:

			if memory[pointer] != 0 {
				i := 1
				for i > 0 {
					idx--
					if instructions[idx] == OP_RIGHT_BRACE {
						i++
					}
					if instructions[idx] == OP_LEFT_BRACE {
						i--
					}
				}
			}
		}
	}

	return memory, nil
}

func main() {
	fileName := flag.String("file", "", "The name of the source file")
	flag.Parse()
	fileContent, err := os.ReadFile(*fileName)
	if err != nil {
		log.Fatalf("ERROR: Couldn't read %s", *fileName)
	}

	memory, err := Eval(fileContent)
	if err != nil {
		log.Panic("ERROR: Couldn't interpret the program")
	}

	fmt.Println()
	if PRINT_MEMORY_AT_THE_END == true {
		fmt.Println(memory)
	}
}
