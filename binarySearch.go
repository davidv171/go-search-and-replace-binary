package main

type Node struct {
	bit  bool
	next *Node
}
type LinkedList struct {
	start *Node
	end   *Node
	//Not sure if we need
	size int
}

//Add element to end of the linkedlist
//Get last element, add a new node to its end
func (linkedList *LinkedList) append(nextBit bool) {
	node := &Node{nextBit, nil}
	linkedList.end.next = node
	linkedList.end = node
}

//Get elements, replace the sublist on an index
//Take the original, loop to the targeted node, unlink the next node,
//Link the next node to the start node of the replacement lsit
//Put the end of replacement list to the next node that got unlinked before
func (linkedList *LinkedList) replaceSublist(replacementIndex int, originalList *LinkedList, replacementList *LinkedList) *LinkedList {
	//initialize the first value
	originalNode := originalList.start
	replacementNode := replacementList.start
	//Find the index we're looking for
	for i := 0; i < replacementIndex && originalNode.next != nil; i++ {
		originalNode = originalNode.next
		if i == replacementIndex {
			//Unlink the culprit node
			cachedNode := originalNode.next
			originalNode.next = replacementNode
			//Link the end of the replacementnode to the start of the unlinked node
			originalList.end = cachedNode
		}
	}
	return originalList
}

//Returns the indexes, where the bits are matching
//Needs the required search bool array and data array, which is the file we're parsing through in binary form
func binarySearch(search []bool, data []bool, bufferOverflow int) []int {
	matchingIndexes := make([]int, 0, len(data))
	//Loop through search boolean slice and look for matching patterns
	//There is no interleaving so we're only checking X booleans at a time for matches
	//Advance is the variable that decides by how much we're advacing each check
	//So... check X booleans for equality, move on to the next X booleans, ALL have to match
	advance := len(search)

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
			//fmt.Print(since)
			//fmt.Print(" to ")
			//The upper matching border
			//until := since + advance
			//fmt.Print(until)
			//fmt.Print(" ")
			matchingIndexes = matchingIndexes[:matchesDiscovered+1]
			matchingIndexes[matchesDiscovered] = since
			matchesDiscovered++
		}

	}
	return matchingIndexes
}

//Find the bit sequence and replace it with a different one
func binaryReplace(search []bool, data []bool, replace []bool) []bool {

	//If variable sized search and replace
	//search > replace: delete values at the edge until you reach replace length
	//replace > search add n values until you've filled out, then shift the latter ones
	//replace == search find the searched indexes and replace the values
	//Bufferoverflow isn't important, we replace on indexes, no matter what
	matchedData := binarySearch(search, data, 0)
	var diff = len(replace) - len(search)
	firstNode := &Node{data[0], nil}
	list := LinkedList{firstNode, firstNode, 1}
	//Create a replacement linked list
	firstReplaceNode := &Node{replace[0], nil}
	replaceList := LinkedList{firstReplaceNode, firstReplaceNode, 1}
	for i := 0; i < len(replace); i++ {
		replaceList.append(replace[i])
	}
	replacedData := make([]bool, 0, len(data)+(len(matchedData)*diff))
	if diff == 0 {
		replacedData = data
		for _, value := range matchedData {
			/*Run through the indexes, which denote positions of matching sequences, and set the values on matching values to the replace values
			it's going to work if the replace and search are the same size*/
			//Expand slice to extra elements, empty values will be added that way
			//Append extra values
			for j := 0; j < len(replace); j++ {
				//Shifting from starting index to last index of the replacement candidate
				shifting := value + j
				replacedData[shifting] = replace[j]
			}

		}
	} else if diff > 0 {

		var inMatchedData bool
		for i := 0; i < len(data); i++ {
			for j := 0; j < len(matchedData); j++ {
				if i != matchedData[j] {
					inMatchedData = false
				} else {
					inMatchedData = true
					break
				}
			}
			if inMatchedData {
				//for z := 0; z < len(replace); z++ {
				//Big bottleneck here
				list.replaceSublist(i, &list, &replaceList)
				//list.append(replace[z])
				//replacedData = append(replacedData, replace[z])

				//}
				i += diff - 1
			} else {
				list.append(data[i])
				//replacedData = append(replacedData, data[i])
			}
		}
	}
	//TODO: Solve with linked list

	return replacedData
}
