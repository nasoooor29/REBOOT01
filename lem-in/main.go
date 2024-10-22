package main


import (
	"fmt"
	"os"
)

// var validPaths []path
var validPaths [][]*node
var nodes []*node
var possiblePath []*node
var ants []*ant
var minLink map[*path]int
var startingNode *node
var endingNode *node

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Not Enough Arguments")
		return
	}
	inputFileName := os.Args[1]
	if inputFileName[len(inputFileName)-4:] != ".txt" {
		fmt.Println("Incorrect input file extension")
		return
	}
	err, numberOfAnts := processFile(inputFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	if numberOfAnts <= 0 || numberOfAnts > 2000 {
		fmt.Println("Invalid Number of Ants")
		return
	}


	for _, node := range nodes {
		if node.start {
			startingNode = node
		} else if node.end {
			endingNode = node
		}
	}

	findValidPath(startingNode, startingNode, startingNode, endingNode)

	paths := []*path{}
	for _, nodes := range validPaths {
		temp := &path{}
		temp.rooms = nodes
		paths = append(paths, temp)
	}

	paths = removeInside(paths)

	if numberOfAnts != 1 {
		paths = countLinks(paths)
		if paths == nil {
			fmt.Println("ERROR: invalid data format")
			return
		}
	}

	paths = sort(paths)

	for i := 1; i <= numberOfAnts; i++ {
		ants = append(ants, newAnt(i, nil))
	}

	for _, ant := range ants {
		assignPath(ant, paths)
	}

	temp := solve(numberOfAnts)
	fmt.Println("\nStep Count", temp)
}
