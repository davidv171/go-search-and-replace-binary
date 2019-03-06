package main

import "fmt"

//Returns the indexes, where the bits are matching
//Needs the required search bool array and bits array, which is the file we're parsing through in binary form
func binarySearch(search *[]bool, bits *[]bool, bufferOverflow int64) []int {
	matchingIndexes := make([]int, 0, len(*bits))
	//Loop through search boolean slice and look for matching patterns
	//There is no interleaving so we're only checking X booleans at a time for matches
	//Advance is the variable that decides by how much we're advacing each check
	//So... check X booleans for equality, move on to the next X booleans, ALL have to match
	advance := len(*search)
	//A count of matches, needs to reach "advance" for there to be a full match between two arrays
	matchCount := 0
	/*A count of matching sequences, while matchCount counts amount of matches of a single element and resets on mismatches and full matches
	matches discovered only triggers once per full match, so if we have 1 matching 40character match, matchesDiscovered is 1
	this value is then used for the return array, so its well sized*/
	i := 0
	if bufferOverflow == 0 {
		i = 112
	}
	matchesDiscovered := 0
	//Probably wont be matching at 0, TODO: fix for last and first elements
	for ; i < len(*bits) && i+advance < len(*bits); i++ {
		for j := 0; j < advance; j++ {
			if (*bits)[i+j] == (*search)[j] {
				//Look for identical matches, every element needs to match
				matchCount++
			} else {
				//If one doesn't match, stop checking and advance
				matchCount = 0
				break
			}
		}
		if matchCount == advance {
			matchCount = 0
			//Whenever we fill up multiple buffers we need to take that into account when printing out indexes
			since := int64(i)
			fmt.Print(since + bufferOverflow*8)
			fmt.Print(" ")
			matchingIndexes = matchingIndexes[:matchesDiscovered+1]
			matchingIndexes[matchesDiscovered] = int(since)
			matchesDiscovered++
		}

	}
	return matchingIndexes
}

//Find the bit sequence and replace it with a different one
func binaryReplace(search *[]bool, bits *[]bool, replace *[]bool, bufferOverflow int64) []byte {

	//If variable sized search and replace
	//search > replace: delete values at the edge until you reach replace length
	//replace > search add n values until you've filled out, then shift the latter ones
	//replace == search find the searched indexes and replace the values
	//Bufferoverflow isn't important, we replace on indexes, no matter what, the search needs it to print out the names
	matchedData := binarySearch(search, bits, bufferOverflow)
	var diff = len(*replace) - len(*search)
	//Create an array/slice that is large enough to fit every replacement diff times
	replacedData := make([]bool, len(*bits)+(len(matchedData)*diff))
	if diff == 0 {
		replacedData = *bits
		for _, value := range matchedData {
			/*Run through the indexes, which denote positions of matching sequences, and set the values on matching values to the replace values
			it's going to work if the replace and search are the same size*/
			//Expand slice to extra elements, empty values will be added that way
			//Append extra values
			for j := 0; j < len(*replace); j++ {
				//Shifting from starting index to last index of the replacement candidate
				shifting := value + j
				replacedData[shifting] = (*replace)[j]
			}

		}
	} else if diff > 0 {
		replacementCounter := 0
		var inMatchedData bool
		for i := 0; i < len(*bits); i++ {
			for j := 0; j < len(matchedData); j++ {
				if i != matchedData[j] {
					inMatchedData = false
				} else {
					inMatchedData = true
					break
				}
			}
			if inMatchedData {
				for z := 0; z < len(*replace); z++ {
					index := i + (replacementCounter * diff) + z
					replacedData[index] = (*replace)[z]
					fmt.Print("i ", i, " ")
					fmt.Print("Replacing on: ")
					fmt.Println(index)
					//replacedData = append(replacedData, (*replace)[z])
				}
				i += len(*search) - 1
				replacementCounter++
			} else {
				index := i + (replacementCounter * diff)
				fmt.Print("i ", i, " ")
				fmt.Print("Adding stuff normally: ")
				fmt.Println(index)
				replacedData[index] = (*bits)[i]
			}
		}
	}
	bytesToWrite := make([]byte, len(replacedData)/8)
	var sentSlice []bool
	//TODO: Merge the with the replacing, replacing can be merged with the reading...AKA replace as you read
	for i := 0; i < len(replacedData); i += 8 {
		for j := 0; j < 8; j++ {
			sentSlice = replacedData[i : i+j+1]
		}
		bytesToWrite[i/8] = bitSliceToByte(&sentSlice)
	}
	return bytesToWrite
}
