package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func newNode(startInp, endInp, visitedInp bool, roomNameInp string) *node {
	tempNode := &node{start: startInp, end: endInp, visited: visitedInp, roomName: roomNameInp}
	return tempNode
}

func newAnt(num int, ptr *path) *ant {
	tempAnt := &ant{number: num, pathPtr: ptr}
	return tempAnt
}

func processFile(inputFileName string) (error, int) {
	in, err := os.ReadFile(inputFileName)
	if err != nil {
		return fmt.Errorf("file not found"), -1
	}
	inputFile := string(in)
	inputFile = strings.ReplaceAll(inputFile, "\r\n", "\n")
	splittedInput := strings.Split(inputFile, "\n")
	str := ""

	commentRegex, _ := regexp.Compile(`\n?^#.*`)
	endRegex, _ := regexp.Compile(`##end\n(?P<end_room>\w+) \d+ \d+`)
	startRegex, _ := regexp.Compile(`##start\n(?P<start_room>\w+) \d+ \d+`)
	startEmpty, _ := regexp.Compile(`##start`)
	endEmpty, _ := regexp.Compile(`##end`)
	roomRegex, _ := regexp.Compile(`\w+ \d+ \d+`)
	linkRegex, _ := regexp.Compile(`(\w+-\w+)`)

	antNumber := 0
	endRoom := ""
	startRoom := ""
	rooms := []string{}
	links := []string{}

	for _, v := range splittedInput {
		matched := commentRegex.MatchString(v)
		if matched && (v != "##end" && v != "##start") {
			continue
		}
		str += v + "\n"
	}

	if len(endRegex.FindAllStringIndex(str, -1)) != 1 || len(endEmpty.FindAllStringIndex(str, -1)) != 1 {
		return fmt.Errorf("end room not found, or multiple endrooms are there"), 0
	}
	if len(startRegex.FindAllStringIndex(str, -1)) != 1 || len(startEmpty.FindAllStringIndex(str, -1)) != 1 {
		return fmt.Errorf("start room not found, or multiple startrooms are there"), 0
	}

	endRoom = endRegex.FindString(str)
	startRoom = startRegex.FindString(str)

	str = strings.Replace(str, endRoom, "", 1)
	str = strings.Replace(str, startRoom, "", 1)
	splittedInput = strings.Split(str, "\n")

	antNumber, err = strconv.Atoi(splittedInput[0])
	if err != nil {
		return fmt.Errorf("ERROR: invalid data format"), 0
	}

	if len(splittedInput) > 0 {
		splittedInput = splittedInput[1:]
	}

	str = ""
	for _, line := range splittedInput {
		matched := roomRegex.MatchString(line)
		if matched {
			rooms = append(rooms, line)
		} else {
			str += line + "\n"
		}
	}

	splittedInput = strings.Split(str, "\n")

	for _, line := range splittedInput {
		matched := linkRegex.MatchString(line)
		if matched {
			links = append(links, line)
		} else {
			str += line + "\n"
		}
	}

	startingNode := newNode(true, false, false, startRoom[8:strings.Index(startRoom[8:], " ")+8])
	endingNode := newNode(false, true, false, endRoom[6:strings.Index(endRoom[6:], " ")+6])
	nodes = append(nodes, startingNode)
	nodes = append(nodes, endingNode)

	cords := []string{}

	for _, room := range rooms {
		nodes = append(nodes, newNode(false, false, false, room[0:strings.Index(room, " ")]))
		cords = append(cords, room[strings.Index(room, " ")+1:])
		str = strings.Replace(str, room, "", 1)
	}

	if isDupedCords(cords) {
		return fmt.Errorf("duped coords"), -1
	}
	if isDuped() {
		return fmt.Errorf("duped room"), -1
	}

	for _, link := range links {
		var node1 *node
		var node2 *node
		first := false
		second := false
		split := strings.Split(link, "-")
		for _, node := range nodes {
			if node.roomName == split[0] {
				node1 = node
				first = true
			} else if node.roomName == split[1] {
				node2 = node
				second = true
			}
		}
		if first && second {
			str = strings.Replace(str, link, "", 1)
			if linkContains(node1.link, node2) {
				return fmt.Errorf("repeat Link"), -1
			}
			node1.link = append(node1.link, node2)
			node2.link = append(node2.link, node1)
		} else {
			fmt.Println(link)
			fmt.Println(first, second)
			return fmt.Errorf("incorrect links"), -1
		}
	}

	str = strings.ReplaceAll(str, "\n", "")

	if len(str) != 0 {
		fmt.Println(str)
		return fmt.Errorf("ERROR: invalid data format"), 0
	}

	return nil, antNumber
}

func isDuped() bool {
	for i, node := range nodes {
		for j, node1 := range nodes {
			if i == j {
				continue
			}
			if node.roomName == node1.roomName {
				return true
			}
		}
	}
	return false
}

func isDupedCords(cords []string) bool {
	for i, cord := range cords {
		for j, cord1 := range cords {
			if i == j {
				continue
			}
			if cord == cord1 {
				return true
			}
		}
	}
	return false
}

func findValidPath(home, start, prev, end *node) node {
	if start.end {
		possiblePath = append(possiblePath, start)

		validPaths = append(validPaths, possiblePath)
		return *start
	}

	possiblePath = append(possiblePath, start)
	for _, nodeLink := range start.link {
		if nodeLink.roomName == prev.roomName { 
			continue
		}
		if !nodeLink.visited {
			if !start.visited || start == home {
				start.visited = true

				track := []*node{}
				for _, thing := range possiblePath {
					track = append(track, thing)
				}

				findValidPath(home, nodeLink, start, end)

				if len(track) < len(possiblePath) {
					possiblePath = track
					start.visited = false
				} else if len(track) > len(possiblePath) {
					start.visited = false
				}

				if areAllVisited() {
					return *start
				}

			}
		} else if nodeLink.end {
			findValidPath(home, nodeLink, start, end)
		} else {
		}
	}

	return *end
}

func removeInside(paths []*path) []*path {
	res := []*path{}
	for _, path := range paths {
		duped := false
		comp1 := path.rooms[1 : len(path.rooms)-1]
		for _, path1 := range paths {
			comp2 := path1.rooms[1 : len(path1.rooms)-1]
			if path == path1 {
				continue
			}
			track := len(comp2)
			for i := 0; i < len(comp1); i++ {
				if track >= len(comp1) || track == 0 {
					break
				}
				if reflect.DeepEqual(comp1[i:track], comp2) {
					duped = true
					// fmt.Println("DEEZ")
				}
				track++
			}

		}
		if !duped {
			res = append(res, path)
		}
	}
	return res
}

func countLinks(paths []*path) []*path {

	startingRooms := []*node{}
	pathLinks := map[*node][]*path{}
	minLink := map[*path]int{}
	bestPathsUpd := []*path{}
	bestPaths := []*path{}

	startingRooms = append(startingRooms, startingNode.link...)
	for _, way := range paths {
		check, idx := isInNodes(way.rooms[1], startingRooms)
		if check {
			pathLinks[startingRooms[idx]] = append(pathLinks[startingRooms[idx]], way)
		}
	}

	for _, room := range startingRooms {
		for _, way := range pathLinks[room] {
			count := 0
			for _, room1 := range startingRooms {
				if room == room1 {
					continue
				}
				for _, way1 := range pathLinks[room1] {
					for _, node := range way.rooms[1 : len(way.rooms)-1] {
						for _, node1 := range way1.rooms[1 : len(way1.rooms)-1] {
							if node == node1 {
								// if node == startingNode || node == endingNode {
								// 	continue
								// }
								count++
							}
						}
					}
				}
			}
			minLink[way] = count
		}
		min, track := 0, &path{}
		if len(pathLinks[room]) != 0 {
			min, track = minLink[pathLinks[room][0]], pathLinks[room][0]
		}
		for i, way := range pathLinks[room] {
			if i == 0 {
				continue
			}
			if minLink[way] < min {
				min, track = minLink[way], way
			}
		}

		pathLinks[room] = []*path{track}

		bestPaths = append(bestPaths, track)

	}

	if len(validPaths) == 0 {
		return nil
	}

	for _, way := range bestPaths {
		flag := false
		for _, way1 := range bestPaths {
			if way == way1 {
				continue
			}
			for _, node := range way.rooms[1 : len(way.rooms)-1] {
				for _, node1 := range way1.rooms[1 : len(way1.rooms)-1] {
					if node == node1 {
						if !flag {
							way.otherRoomLink++
							flag = true
						}
					}
				}
			}

		}
	}
	for _, way := range bestPaths {
		if way.otherRoomLink == 0 {
			bestPathsUpd = append(bestPathsUpd, way)
			continue
		} else {
			for _, way1 := range bestPaths {
				if way == way1 {
					continue
				}
				if way.otherRoomLink == way1.otherRoomLink {
					if len(way.rooms) < len(way1.rooms) {
						bestPathsUpd = append(bestPathsUpd, way)
					} else if len(way.rooms) > len(way1.rooms) {
						bestPathsUpd = append(bestPathsUpd, way1)
					} else {
						bestPathsUpd = append(bestPathsUpd, way)
						bestPathsUpd = append(bestPathsUpd, way1)
					}
				}

			}
		}
	}
	bestPathsUpd = removeDupe(bestPathsUpd)

	return bestPathsUpd
}

func removeDupe(ways []*path) []*path {
	filter := []*path{}
	for _, way := range ways {
		flag := false
		for _, check := range filter {
			if way == check {
				flag = true
			}
		}

		if !flag {
			filter = append(filter, way)
		}
	}
	return filter
}

func linkContains(nodes []*node, checkNode *node) bool {
	for _, room := range nodes {
		if room == checkNode {
			return true
		}
	}
	return false
}

func sort(paths []*path) []*path {
	res := []*path{}
	track := paths
	for i := 0; i < len(paths); i++ {
		min := track[i]
		for j := i; j < len(paths); j++ {
			check := track[j]
			if len(check.rooms) < len(min.rooms) {
				min = check
				temp := track[i]
				track[i] = track[j]
				track[j] = temp

			}
		}
		res = append(res, min)

	}
	return res
}

func areAllVisited() bool {
	check := true
	for _, node := range nodes {
		if !node.visited {
			check = false
		}
	}
	return check
}

func assignPath(ant *ant, paths []*path) {
	min, idx := (paths[0].noAssigned + (len(paths[0].rooms)) + minLink[paths[0]]), 0
	for i := 1; i < len(paths); i++ {
		if (paths[i].noAssigned + (len(paths[i].rooms)) + minLink[paths[i]]) < min {
			min, idx = (paths[i].noAssigned + (len(paths[i].rooms)) + minLink[paths[i]]), i
		}
	}
	ant.pathPtr = paths[idx]
	paths[idx].noAssigned++

}

func solve(numOfAnts int) int {
	counter := 0
	antsTracker := ants
	order := [][]*ant{}
	moved := []*ant{}

	for ants[0].pathPtr.rooms[len(ants[0].pathPtr.rooms)-1].noOccupied < numOfAnts {
		if counter == 0 {
			track := false
			for _, ant := range antsTracker {
				if ant.pathPtr.rooms[1].occupier == nil {
					ant.pathPtr.rooms[1].occupier = ant
					if ant.pathPtr.rooms[1].end {
						ant.idx = -1
						ant.pathPtr.rooms[1].noOccupied++
						track = true
					} else {
						ant.idx = 1
					}
					fmt.Print("L", ant.number, "-", ant.pathPtr.rooms[1].roomName, " ")
					moved = append(moved, ant)
					antsTracker = remove(antsTracker, ant)
				}
			}
			if track {
				ants[0].pathPtr.rooms[len(ants[0].pathPtr.rooms)-1].occupier = nil
			}

		} else {
			for _, group := range order {
				for _, ant := range group {
					if ant.pathPtr.rooms[1].end && ant.idx != -1 {
						ant.pathPtr.rooms[1].noOccupied++
						fmt.Print("L", ant.number, "-", ant.pathPtr.rooms[1].roomName, " ")
						ant.idx = -1
					}
					if ant.idx != -1 {
						toGo := ant.pathPtr.rooms[ant.idx+1]
						if toGo.end {
							toGo.noOccupied++
							ant.pathPtr.rooms[ant.idx].occupier = nil
							fmt.Print("L", ant.number, "-", ant.pathPtr.rooms[ant.idx+1].roomName, " ")
							ant.idx = -1
						} else if toGo.occupier == nil {
							ant.pathPtr.rooms[ant.idx].occupier = nil
							toGo.occupier = ant
							fmt.Print("L", ant.number, "-", toGo.roomName, " ")
							ant.moved = true
						}
						if ant.idx != -1 && ant.moved {
							ant.moved = false
							ant.idx++
						}
					}

				}
			}
			flag := true

			for _, ant := range antsTracker {
				if ant.pathPtr.rooms[1].occupier == nil {
					if ant.pathPtr.rooms[1].end && flag {
						ant.pathPtr.rooms[1].noOccupied++
						flag = false
						fmt.Print("L", ant.number, "-", ant.pathPtr.rooms[1].roomName, " ")
						ant.idx = -1
						antsTracker = remove(antsTracker, ant)
					} else if !ant.pathPtr.rooms[1].end {
						ant.pathPtr.rooms[1].occupier = ant
						ant.idx = 1
						fmt.Print("L", ant.number, "-", ant.pathPtr.rooms[1].roomName, " ")

						moved = append(moved, ant)
						antsTracker = remove(antsTracker, ant)
					}
				}
			}
		}

		order = append(order, moved)
		moved = []*ant{}
		fmt.Println("")
		counter++
	}

	return counter
}

func isInNodes(node *node, arr []*node) (bool, int) {
	for i, thing := range arr {
		if node == thing {
			return true, i
		}
	}
	return false, -1
}

func indexOf(ant *ant, arr []*ant) int {
	for i, thing := range arr {
		if ant == thing {
			return i
		}
	}
	return -1
}

func remove(s []*ant, i *ant) []*ant {
	s[indexOf(i, s)] = s[len(s)-1]
	return s[:len(s)-1]
}
