package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func main() {
	filepath := os.Args[1]
	operation := os.Args[2]
	search := os.Args[3]
	fmt.Println("Searching for " + search + " in " + filepath)
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
	if err != nil {
		fmt.Println("No file ", filepath,  " exists ")
		os.Exit(1)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	errCheck(err)
	fileSize := fileInfo.Size()
	fmt.Print("File size is ")
	fmt.Print(fileSize)
	//Make sure the search slice fits, this will be our buffer
	var bufferSize int64
	bufferSize = 8192
	if fileSize < bufferSize {
		bufferSize = fileSize
	}
	var bufferOverflow int64 = 0
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
			//TODO: remove append this is not performant
			//Unnecessary complexity, we could be doing this per byte array of buffer size instead of per each byte
			bits = append(bits, byteToBitSlice(&aByte)...)
		}
		if operation == "f" {
			binarySearch(&searchSlice, &bits, bufferOverflow)
		}
		if operation == "fr" {
			bytesToWrite := binaryReplace(&searchSlice, &bits, &replaceSlice, bufferOverflow)
			writeBinaryFile("out", &bytesToWrite, bufferOverflow)

		}
		//So we're aware of indexes if the file is larger than buffer size
		bufferOverflow += bufferSize

		bits = nil
	}
	os.Exit(0)
}

//Writes into the binary file at bufferOverflow offset
func writeBinaryFile(fileName string, bytesToWrite *[]byte, bufferOverflow int64) {
	if bufferOverflow == 0 {
		_, err := os.Create(fileName)
		errCheck(err)
	}
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	defer file.Close()
	errCheck(err)
	_, err = file.WriteAt(*bytesToWrite, bufferOverflow)
	errCheck(err)
}

//- Binarni zapis ali branje podatkovnih tipov 32-bitni int, 32-bitni float in 8-bitni char.
//Read amount of data, and based on that conver to type
func bytesToInt32(bytesToConvert []byte) uint32 {
	var data uint32
	data = binary.BigEndian.Uint32(bytesToConvert)
	return data
}
func bytesToFloat32(bytesToConvert []byte) float32 {
	var data float32
	buff := bytes.NewReader(bytesToConvert)
	err := binary.Read(buff, binary.LittleEndian, &data)
	errCheck(err)
	return data
}
func bytesTo8Char(bytesToConvert []byte) uint8 {
	var data uint8
	buff := bytes.NewReader(bytesToConvert)
	err := binary.Read(buff, binary.LittleEndian, &data)
	errCheck(err)
	return data
}

func charToBytes(f rune) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}

func float64ToBytes(f float64) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}
func float32ToBytes(f float32) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}
func int32ToBytes(f int32) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}
