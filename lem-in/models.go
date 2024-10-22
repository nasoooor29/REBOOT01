package main

type node struct {
	start      bool
	end        bool
	roomName   string
	visited    bool
	link       []*node
	occupier   *ant
	noOccupied int
}

type path struct {
	rooms         []*node
	noAssigned    int
	otherRoomLink int
}

type ant struct {
	number  int
	idx     int
	pathPtr *path
	moved   bool
}
