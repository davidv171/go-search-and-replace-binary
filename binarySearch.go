package main

import "fmt"

//Returns the indexes, where the bits are matching
//Needs the required search bool array and data array, which is the file we're parsing through in binary form
func binarySearch(search []bool, data []bool, bufferOverflow int) []int {
	matchingIndexes := make([]int, len(data))
	//Loop through search boolean slice and look for matching patterns
	//There is no interleaving so we're only checking X booleans at a time for matches
	//Advance is the variable that decides by how much we're advacing each check
	//So... check X booleans for equality, move on to the next X booleans, ALL have to match
	advance := len(search)
	fmt.Println("")
	fmt.Println(data[0:])
	fmt.Println("")

	//A count of matches, needs to reach "advance" for there to be a full match between two arrays
	matchCount := 0
	/*A count of matching sequences, while matchCount counts amount of matches of a single element and resets on mismatches and full matches
	matches discovered only triggers once per full match, so if we have 1 matching 40character match, matchesDiscovered is 1
	this value is then used for the return array, so its well sized*/
	matchesDiscovered := 0
	for i := 0; i < len(data) && i+advance < len(data); i += advance {
		for j := 0; j < advance; j++ {
			if data[i+j] == search[j] {
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
			since := i + bufferOverflow
			fmt.Print(since)
			fmt.Print(" to ")
			//The upper matching border
			until := since + advance
			fmt.Print(until)
			fmt.Print(" ")
			matchingIndexes[matchesDiscovered] = since
			matchesDiscovered++
		}

	}
	return matchingIndexes
}

//Find the bit sequence and replace it with a different one
func binaryReplace(search []bool, data []bool, replace []bool) []bool {
	replacedData := make([]bool, len(data)+len(replace))
	replacedData = data
	//If variable sized search and replace
	//search > replace: delete values at the edge until you reach replace length
	//replace > search add n values until you've filled out, then shift the latter ones
	//replace == search find the searched indexes and replace the values
	//Bufferoverflow isn't important, we replace on indexes, no matter what
	matchedData := binarySearch(search, data, 0)
	for _, value := range matchedData {
		/*Run through the indexes, which denote positions of matching sequences, and set the values on matching values to the replace values
		it's going to work if the replace and search are the same size*/
		for j := 0; j < len(replace); j++ {
			replacedData[value+j] = replace[j]
		}
	}
	fmt.Println("")
	fmt.Println(data[0:])
	fmt.Print(replacedData[0:])
	return replacedData
}
