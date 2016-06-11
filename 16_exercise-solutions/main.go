package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

//-------------------------------------------
//	ZOMBIELANDMANIA 2: RISE OF ZED
//-------------------------------------------
/*
	This is a text-based turn-based shooter played in a rectangular grid against zombie AI opponents.
	On the player's turn, the game prompts the user for a line of input
	The player can move, shoot, and stand still
	After the player acts, the zombies get their turn to move
	If a zombie moves onto the same tile as the player, it's game over!

	If the player manages to kill all of the zombies, he wins!
*/

//-------------------------------------------
//               Constants
//-------------------------------------------
//zombie attributes
//none yet!
const z1 = 1
const z2 = 2
const z3 = 3
const z4 = 4
const z5 = 5

const zInitialPopulation = 5 //defines the number of zombies that spawn when the game loads

//player attributes
const playerAim = 0.9       //determines how likely a player is to hit a zombie when shooting from 1 space away
const playerAimDecay = 0.01 //OPTIONAL: determines how much worse the player's aim gets for each space between him and his target

//weapon attributes
//none yet!

//terrain attributes
//none yet!

//graphics
//these are the symbols printed to the console each frame/turn of the game
const openSymbol = '.'   //indicates an empty space
const playerSymbol = 'P' //player location
const zSymbol = 'z'      //zombie locations
//NOTE: if multiple zombies are on the same space, their number is displayed instead
//e.g.  3 zombies on one space results in a '3' being shown

const gameXSize = 20 //horizontal size of the game. x = 0 is the left-most column, x = gameXSize-1 is the right-most column
const gameYSize = 20 //vertical size of the game. y = 0 is the bottom row, y = gameYSize-1 is the top row

//AI modes
const zRandomAI = 0
const zSimpleAI = 1
const zSmartRange = 10
const playerRandomAI = 0
const playerSimpleAI = 1

//Inputs
const inputUp = "w"
const inputLeft = "a"
const inputDown = "s"
const inputRight = "d"
const inputShootUp = " w"
const inputShootLeft = " a"
const inputShootDown = " s"
const inputShootRight = " d"
const inputPass = "p"
const inputHelp = "h"
const inputQuit = "q"

//Directions
const up = 0
const left = 1
const down = 2
const right = 3

//-------------------------------------------

//-------------------------------------------
//      Global Variables (package scope)
//-------------------------------------------
//5 zombies
var x1, x2, x3, x4, x5 int  //x coordinates
var y1, y2, y3, y4, y5 int  //y coordinates
var l1, l2, l3, l4, l5 bool //alive?

//1 player
var playerX, playerY int
var playerAlive bool

//-------------------------------------------

//-------------------------------------------
//           Game Logic Functions
//-------------------------------------------

func move(x, y, dir int) (int, int) {
	switch dir {
	default:
		return x, y
	case up:
		if y >= gameYSize-1 {
			return x, y
		}
		return x, y + 1
	case down:
		if y <= 0 {
			return x, y
		}
		return x, y - 1
	case right:
		if x >= gameXSize-1 {
			return x, y
		}
		return x + 1, y
	case left:
		if x <= 0 {
			return x, y
		}
		return x - 1, y
	}

	//return the new coordinates (x_new, y_new)
	//which result from starting at (x,y) and moving one space in direction dir
	//consult the constants to learn what different values of dir mean
}

func movePlayer(dir int) {
	playerX, playerY = move(playerX, playerY, dir)
	//TODO: move the player from his current position
	//one space in direction dir
}

func moveZombie(n, dir int) {
	if n == z1 {
		x1, y1 = move(x1, y1, dir)
	}
	if n == z2 {
		x2, y2 = move(x2, y2, dir)
	}
	if n == z3 {
		x3, y3 = move(x3, y3, dir)
	}
	if n == z4 {
		x4, y4 = move(x4, y4, dir)
	}
	if n == z5 {
		x5, y5 = move(x5, y5, dir)
	}

	//move the nth zombie from its current position
	//one space in direction dir
}

func killPlayer() {
	playerAlive = false
	// change whatever program state (variables) necessary
	//to indicate that the player is dead
}

func killZombie(n int) {
	switch n {
	default:

	case z1:
		l1 = false
	case z2:
		l2 = false
	case z3:
		l3 = false
	case z4:
		l4 = false
	case z5:
		l5 = false
	}
	//change whatever program state (variables) necessary
	//to indicate that the nth zombie is dead
}

func shoot(dir int) {
	switch dir {
	default:

	case up:
		closeZ := gameYSize
		closeI := 0

		for i := 1; i <= zInitialPopulation; i++ {
			zx, zy := zombieLocation(i)
			if zy >= playerY && zx == playerX {
				if zy < closeZ {
					closeZ = zy
					closeI = i
				}

			}
		}
		killZombie(closeI)
	case down:
		closeZ := 0
		closeI := 0

		for i := 1; i <= zInitialPopulation; i++ {
			zx, zy := zombieLocation(i)
			if zy <= playerY && zx == playerX {
				if zy > closeZ {
					closeZ = zy
					closeI = i
				}

			}
		}
		killZombie(closeI)
	case left:
		closeZ := 0
		closeI := 0

		for i := 1; i <= zInitialPopulation; i++ {
			zx, zy := zombieLocation(i)
			if zy == playerY && zx <= playerX {
				if zx > closeZ {
					closeZ = zx
					closeI = i
				}

			}
		}
		killZombie(closeI)
	case right:
		closeZ := gameXSize
		closeI := 0

		for i := 1; i <= zInitialPopulation; i++ {
			zx, zy := zombieLocation(i)
			if zy == playerY && zx >= playerX {
				if zx < closeZ {
					closeZ = zx
					closeI = i
				}

			}
		}
		killZombie(closeI)
	}
	/* aim = aim - decay * zombieD
	if (aim >= 0.45) {
		killZombie(i)
	} else {
		return
	}
	*/
	//TODO: calculate the result of the player shooting in direction dir
	//	This involves a couple things:
	//		figure out where the nearest zombie, if any, in that direction is
	//		calculate whether or not the shot was a hit  (see playerAim and 			playerAimDecay constants)
	//		apply the result
}

func zombieAI(mode int) {
	// Automatically move each of the zombies
	// zRandomAI can move each zombie in a random direction (staying still is a possibility)
	// zSimpleAI can behave a couple ways:
	//			if the player is within zSmartRange (see constants) of the player
	//				move towards the player
	//			else
	//				move in a random direction
	switch mode {
	default:
	case zSimpleAI:

	case zRandomAI:
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		for i := 1; i <= zInitialPopulation; i++ {
			moveZombie(i, r.Intn(5))
		}

	}
}

//func playerAI(mode int) {
//TODO: optional!
//Similar to the zombie AI, implement an AI for the player
//Which is activated by a certain input, which you can define
//playerRandomAI performs a random action for the player, including moving or shooting
//
//There are several possibilities for what to do with playerSimpleAI. Here are some ideas:
//	Detect nearby zombies, and move away from them
//	Always shoot at zombies in range
//	If not in range, aggressively move to get in range
//	determine the BEST zombie to shoot at / move towards
//	This will likely require defining some extra functions for doing things like
//	detecting when zombies are in range, and in which direction
//switch mode {
//case playerSimpleAI:

//case playerRandomAI:

//}
//}

func zombieLocation(n int) (int, int) {
	switch n {
	default:
		return -1, -1
	case z1:
		return x1, y1
	case z2:
		return x2, y2
	case z3:
		return x3, y3
	case z4:
		return x4, y4
	case z5:
		return x5, y5
	}
	//return the (x,y) coordinates of the nth zombie
}

func getZombieCount(x, y int) int {
	howMany := 0
	for i := 1; i <= zInitialPopulation; i++ {
		zx, zy := zombieLocation(i)
		if zx == x && zy == y && isZombieAlive(i) {
			howMany++
		}
	}
	//return the number of zombies currently at position (x,y)
	return howMany
}

func isZombieAlive(n int) bool {
	switch n {
	default:
		return false
	case z1:
		return l1
	case z2:
		return l2
	case z3:
		return l3
	case z4:
		return l4
	case z5:
		return l5
	}
}

//-------------------------------------------
//-------------------------------------------

//-------------------------------------------
//           Boilerplate Functions
//-------------------------------------------
func initGame() {
	//Initialize a random number generator  (a 'rand', see package "math/rand")
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	//Initialize zombie stats
	l1, l2, l3, l4, l5 = true, true, true, true, true
	x1, x2, x3, x4, x5 = r.Intn(gameXSize), r.Intn(gameXSize), r.Intn(gameXSize), r.Intn(gameXSize), r.Intn(gameXSize)
	y1, y2, y3, y4, y5 = r.Intn(gameYSize), r.Intn(gameYSize), r.Intn(gameYSize), r.Intn(gameYSize), r.Intn(gameYSize)

	//Initialize player stats
	playerAlive = true
	playerX = gameXSize / 2
	playerY = gameYSize / 2

	//Make sure zombies don't start on top of the player
	if gameXSize > 1 && gameYSize > 1 {
		for getZombieCount(playerX, playerY) > 0 {
			if x1 == playerX && y1 == playerY {
				moveZombie(1, r.Intn(4))
			}
			if x2 == playerX && y2 == playerY {
				moveZombie(2, r.Intn(4))
			}
			if x3 == playerX && y3 == playerY {
				moveZombie(3, r.Intn(4))
			}
			if x4 == playerX && y4 == playerY {
				moveZombie(4, r.Intn(4))
			}
			if x5 == playerX && y5 == playerY {
				moveZombie(5, r.Intn(4))
			}
		}
	}
}

func getPixel(x, y int) rune {

	if x1 == x && y1 == y && isZombieAlive(z1) {
		return zSymbol
	}
	if x2 == x && y2 == y && isZombieAlive(z2) {
		return zSymbol
	}
	if x3 == x && y3 == y && isZombieAlive(z3) {
		return zSymbol
	}
	if x4 == x && y4 == y && isZombieAlive(z4) {
		return zSymbol
	}
	if x5 == x && y5 == y && isZombieAlive(z5) {
		return zSymbol
	}
	if playerX == x && playerY == y && playerAlive == true {
		return playerSymbol
	}
	return openSymbol

	//based on the current positions of zombies and the player
	//output the appropriate symbol for position (x,y)
	//Use the symbols defined in the Constants section at the top of this file (by their 	identifiers, don't re-type the literal value)
	// The reason for having those constants is that if you ever want to change which symbol is 	used for something,
	//you can just change it at the top. You don't have to go hunting through the whole file for 	every place it gets used
}

func drawFrame() {
	fmt.Println()
	for i := gameYSize - 1; i >= 0; i-- {
		line := ""
		for j := 0; j < gameXSize; j++ {
			line = line + string(getPixel(j, i))
		}
		fmt.Println(line)
	}
}

func getInput() string {
	fmt.Println("It's your turn! What will you do?  (w/a/s/d/ w/ a/ s/ d/q/p) then hit Enter.")
	bio := bufio.NewReader(os.Stdin)
	line, _ := bio.ReadString('\n')
	return strings.Trim(line, "\r\n")
}

//-------------------------------------------

func main() {
	initGame()

	for { //main loop starts here
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		drawFrame() //draw a frame

		//TODO: how should the game respond to movement and shooting inputs?
		//hint: check the game logic functions for something that seems useful
		//these should only take one or two lines of code each
		if playerAlive == false {
			fmt.Println("You are ded, Mang!")
			os.Exit(0)
		}
		switch getInput() { //NOTE: this is a CALL to the getInput() function. It returns the 		command typed in by the user as a string
		default:
			fmt.Println("Invalid input!")
		case inputPass: //do nothing
		case inputHelp: //print a helpful message to the player, perhaps explaining what some 		of the commands do
		//movement starts here
		case inputUp:
			movePlayer(up)
		case inputLeft:
			movePlayer(left)
		case inputDown:
			movePlayer(down)
		case inputRight:
			movePlayer(right)
		//shooting starts here
		case inputShootUp:
			shoot(up)
		case inputShootLeft:
			shoot(left)
		case inputShootDown:
			shoot(down)
		case inputShootRight:
			shoot(right)
		case inputQuit:
			fmt.Println("Quitting the game.")
			os.Exit(0) //this is how we get out of our infinite loop: tell the os to exit 			the program.
			//Status 0 means success, non-zero exit status means some kind of error caused 			the exit
		}
		zombieAI(zRandomAI)
		for i := 1; i <= zInitialPopulation; i++ {
			zx, zy := zombieLocation(i)
			if zx == playerX && zy == playerY && isZombieAlive(i) {
				killPlayer()
			}
		}
		if l1 == false && l2 == false && l3 == false && l4 == false && l5 == false {
			fmt.Println("You survived Zombiemania Mang!")
			os.Exit(0)
			//Player Wins!!!
		}
	}
}
