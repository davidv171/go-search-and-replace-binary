package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	filepath := os.Args[1]
	operation := os.Args[2]
	search := os.Args[3]
	fmt.Println("Searching for " + search + " in " + filepath)
	benchmark, err := os.Create("benchmark.txt") // can use open and O_CREAT instead
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(benchmark)
	defer pprof.StopCPUProfile()
	//Turn the searched string, usually in the format of "0101010001010000000101010" into a slice of bools
	searchSlice := argumentToBinary(search)
	if operation == "f" {
		//Reads binary file and based on the inputted operation performs an action of find or replace
		readBinaryFile(filepath, searchSlice, "f", nil)
	} else if operation == "fr" {
		replace := os.Args[4]
		replaceSlice := argumentToBinary(replace)
		readBinaryFile(filepath, searchSlice, "fr", replaceSlice)
	}
}
func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}
func readBinaryFile(filepath string, searchSlice []bool, operation string, replaceSlice []bool) {
	file, err := os.Open(filepath)
	defer file.Close()
	bufferOverflow := 0
	fileInfo, err := file.Stat()
	errCheck(err)
	fileSize := fileInfo.Size()
	fmt.Print("File size is ")
	fmt.Print(fileSize)
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
	bits = nil
	fmt.Println("\nMatching on indexes: ")
	for {
		//Loop through the file, retrieve the bytes as integers
		//:cap(data) = capacity of array, how many elements it can take before it has to resize
		//Init slice
		data = data[:cap(data)]
		//Byte in the file
		readByte, err := file.Read(data)
		if err != nil {
			if err == io.EOF {
				fmt.Print("\n")
				fmt.Println("Done reading file")
				break
			}
			fmt.Println(err)
			return
		}
		data = data[:readByte]
		for _, aByte := range data {
			bits = append(bits, byteToBitSlice(&aByte)...)
		}
		if operation == "f" {
			binarySearch(searchSlice, bits, bufferOverflow)
		}
		if operation == "fr" {
			binaryReplace(searchSlice, bits, replaceSlice)
		}
		bits = nil
		//So we're aware of indexes if the file is larger
		bufferOverflow += 4096
	}
}
