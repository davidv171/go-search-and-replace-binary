package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filepath := os.Args[1]
	operation := os.Args[2]
	search := os.Args[3]
	fmt.Println("Searching for " + search + " in " + filepath)
	file, err := os.Open(filepath)
	errCheck(err)
	defer file.Close()
	//Turn the searched string, usually in the format of "0101010001010000000101010" into a slice of bools
	searchSlice := argumentToBinary(search)
	if operation == "f" {
		fileInfo, err := file.Stat()
		errCheck(err)
		fileSize := fileInfo.Size()
		fmt.Print("File size is ")
		fmt.Print(fileSize)
		fmt.Print(searchSlice)
		//Make sure the search slice fits, this will be our buffer
		var bufferSize int64
		if fileSize > 4096 {
			bufferSize = 4096
		} else {
			bufferSize = fileSize
		}
		//Data where we put the read bytes into
		data := make([]byte, bufferSize)
		//Bits where we put arrays of bools, signifying bits
		bits := make([]bool, bufferSize)
		//Pattern trigger, counter of how many times we've found the matching pattern, defined in the search OS argument
		for {
			//Loop through the file, retrieve the bytes as integers
			//:cap(data) = capacity of array, how many elements it can take before it has to resize
			//Init slice
			data = data[:cap(data)]
			//Byte in the file
			readByte, err := file.Read(data)
			if err != nil {
				if err == io.EOF {
					fmt.Println("Done reading file")
					fmt.Println(len(data))
					break
				}
				fmt.Println(err)
				return
			}
			data = data[:readByte]
			for _, aByte := range data {
				fmt.Println("")
				fmt.Print(aByte)
				fmt.Print(" ")
				bits = append(bits, byteToBitSlice(&aByte)...)
			}

		}
	} else if operation == "fr" {
		//replace:=os.Args[4];
		fmt.Print("Replacing ...")
	}
}
func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}
func byteToBitSlice(byteSlice *uint8) []bool {
	bits := make([]bool, 8)
	var i uint8
	//Loop through the byte and turn it into bit sequence using XOR and masking
	//Using an unsigned integer, so
	//7 -> 0
	fmt.Println("")
	for i = 0; i < 8; i++ {
		mask := byte(1 << i)
		if (*byteSlice & mask) > 0 {
			bits[7-i] = true
		} else {
			bits[7-i] = false
		}

	}
	return bits
}
