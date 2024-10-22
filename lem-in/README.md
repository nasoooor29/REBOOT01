Lem-in
Objectives
The goal of this project is to create a digital version of an ant farm. The program lem-in will read from a file (describing the ants and the colony) provided as an argument. Upon successfully finding the quickest path, lem-in will display the content of the file passed as an argument and each move the ants make from room to room.

How Does It Work?
Setup the Ant Farm:

Create rooms and tunnels.
Place the ants at the start room and observe how they find the exit.
Find the Quickest Path:

Move n ants across the colony (composed of rooms and tunnels) from the room ##start to the room ##end with as few moves as possible.
Handle various edge cases such as:
No path between ##start and ##end.
Rooms that link to themselves.
Invalid or poorly formatted input (e.g., invalid number of ants, missing ##start or ##end room, duplicated rooms, links to unknown rooms, invalid room coordinates).
Input File Format
The input file should be structured as follows:

Number of Ants: An integer indicating the number of ants.
Rooms: Each room is defined by name coord_x coord_y, for example, Room 1 2, nameoftheroom 1 6, 4 6 7.
Links: Each link is defined by name1-name2, for example, 1-2, 2-5.
Commands: Special commands like ##start and ##end.
Example
shell
Copy code
3
##start
1 23 3
2 16 7
#comment
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
#another comment
4-2
2-1
7-6
7-2
7-4
6-5
Graphical Representation
css
Copy code
        _________________
       /                 \
  ____[5]----[3]--[1]     |
 /            |    /      |
[6]---[0]----[4]  /       |
 \   ________/|  /        |
  \ /        [2]/________/
  [7]_________/
Instructions
Rooms:

Must not start with the letter L or #.
Must not contain spaces.
Can be linked to multiple rooms.
Tunnels:

Join only two rooms.
A room can only contain one ant at a time (except ##start and ##end which can contain multiple ants).
Each tunnel can only be used once per turn.
Ant Movement:

Display only the ants that moved at each turn.
Each ant can move only once per turn through an empty tunnel.
Error Handling:

Handle errors gracefully, without unexpected crashes.
Coordinates of rooms will always be integers.
Ignore unknown commands.
Usage
Run the program with an input file describing the ant colony:

php
Copy code
lem-in <input_file>
Development
Language: Go
Good Practices: Ensure code quality and readability.
Testing: Provide unit tests for various scenarios.
Allowed Packages
Only standard Go packages are allowed.

Example Execution
ruby
Copy code
$ lem-in example.txt
This will display the content of example.txt and each move the ants make from room to room.

Conclusion
This project challenges you to simulate an ant farm and find the quickest path for ants to travel from start to end, considering various edge cases and constraints. Happy coding!