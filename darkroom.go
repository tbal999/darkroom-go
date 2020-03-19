package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func generateNest(x, cmd int, diff *int) []int {
	//diffy := *diff
	xindex := 0
	nest := make([]int, 0)
	for xindex = 0; xindex < x; xindex++ {
		nest = append(nest, 0)
	}
	switch cmd {
	case 1:
		for xindex = 0; xindex < x-1; xindex++ {
			nest[randomNumber(0, x)] = randomNumber(3, 9)
		}
	}
	return nest
}

func generateSlice(x, y, cmd int, diffy *int) [][]int {
	//thedifficulty := *diffy
	yindex := 0
	slice := make([][]int, 0, 0)
	for yindex = 0; yindex < y; yindex++ {
		slice = append(slice, generateNest(x, cmd, diffy))
	}
	switch cmd {
	case 0:
		slice[0][0] = 2
	case 1:
		for yindex = 0; yindex < y; yindex++ {
			slice[randomNumber(0, len(slice)-1)][randomNumber(0, len(slice[0])-1)] = randomNumber(3, 9)
		}
		for yindex = 0; yindex < y; yindex++ {
			slice[randomNumber(0, len(slice)-1)][randomNumber(0, len(slice[0])-1)] = randomNumber(10, 16)
		}
		slice[0][0] = 0
	}
	return slice
}

func printSlice(x [][]int) {
	for i := range x {
		fmt.Println(x[i])
	}
}

func randomNumber(min, max int) int {
	z := rand.Intn(max)
	if z < min {
		z = min
	}
	return z
}

func resetSlice(a, b int, diff *int, zeros, ones *[][]int) {
	diffy := *diff
	diffy = diffy + 1
	*diff = diffy
	i := *zeros
	j := *ones
	i = generateSlice(a, b, 0, diff)
	j = generateSlice(a, b, 1, diff)
	*zeros = i
	*ones = j
}

func checkMapClear(a, b int, diff *int, zeros, ones *[][]int, he *Hero) {
	o := *ones
	for iy := range o {
		for ix := range o[iy] {
			if o[iy][ix] >= 6 {
				return
			}
		}
	}
	//difficult := *diff
	//h := *he
	fmt.Println("The dungeon is clear of monsters!")
	fmt.Println(`		
  |                    _,.-----.,_         o    |          
           +    *    .-'.         .'-.          -O-         
     *            .'.-'   .---.   '.'.         |     *    
.                /_.-'   /     \   .'-.\                   
        ' -=*<  |-._.-  |       |   '-._|  >*=-    .     + 
-- )--           \-.    \     /    .-'/                   
      *     +     .'.    '---'    .'.'    +       o       
                 .  '-._         _.-'  .                   
         |               ~~~~~~~       - --===D       @   
`)
	fmt.Println("Moving to new Dungeon!")
	resetSlice(a, b, diff, zeros, ones)
}

func checknumber(y, x int, game [][]int, he *Hero, diff int) {
	h := *he
	number := game[y][x]
	switch number {
	case 0:
		fmt.Println("You found nothing here.")
	case 3:
		printSlice(game)
		fmt.Println("You found a map of the dungeon! What does it all mean?")
		fmt.Println("The map states numbers below 10 are good, and above 10 are bad...")
	case 4:
		fmt.Println("You have found a Dagger!")
		i := *h.attack
		i = i + 1
		*h.attack = i
	case 5:
		fmt.Println("You have found a Axe!")
		i := *h.attack
		i = i + 2
		*h.attack = i
	case 6:
		fmt.Println("You have found a Sword!")
		i := *h.attack
		i = i + 3
		*h.attack = i
	case 7:
		fmt.Println("You have found Excalibur!")
		i := *h.attack
		i = i + 4
		*h.attack = i
	case 8:
		fmt.Println("You have found some rejuvinating water!")
		fmt.Println("You have gained " + strconv.Itoa(diff) + " health.")
		i2 := *h.health
		i2 = i2 + diff
		*h.health = i2
	case 10:
		fmt.Println("You have encountered a small toad.")
		initiateFight(he, 6, diff)
	case 11:
		fmt.Println("You have encountered a goblin.")
		initiateFight(he, 7, diff)
	case 12:
		fmt.Println("A giant moth descends from the ceiling.")
		initiateFight(he, 8, diff)
	case 13:
		fmt.Println("The ghost of bad luck spooks you.")
		fmt.Println("You've dropped all your weapons and now have 1 attack")
		i3 := *h.attack
		i3 = 1
		*h.attack = i3
		initiateFight(he, 9, diff)
	case 14:
		fmt.Println("A shadow beast approaches.")
		initiateFight(he, 10, diff)
	case 15:
		fmt.Println("A hooded rogue has been waiting...")
		initiateFight(he, 11, diff)
	case 16:
		fmt.Println("A demon appears!")
		initiateFight(he, 11, diff)
	}
	game[y][x] = 0
	*he = h
}

func initiateFight(h *Hero, i, difficulty int) {
	toad := Monster{"toad", 20, 3}
	goblin := Monster{"goblin", 15, 5}
	moth := Monster{"giant moth", 30, 7}
	shadowbeast := Monster{"shadow beast", 30, 9}
	rogue := Monster{"hooded rogue", 35, 11}
	demon := Monster{"demon", 40, 13}
	switch i {
	case 6:
		Fight(h, toad, difficulty)
	case 7:
		Fight(h, goblin, difficulty)
	case 8:
		Fight(h, moth, difficulty)
	case 10:
		Fight(h, shadowbeast, difficulty)
	case 11:
		Fight(h, rogue, difficulty)
	case 12:
		Fight(h, demon, difficulty)
	}
}

func Fight(h *Hero, m Monster, d int) {
	herohealth := *h.health
	heroattack := *h.attack
	win := 0
	for win == 0 {
		herohealth = herohealth - m.attack
		herohealth = herohealth - d
		m.health = m.health - heroattack
		if m.health <= 0 {
			fmt.Println("You win the fight against the " + m.name)
			fmt.Println("You have " + strconv.Itoa(herohealth) + " health remaining.")
			if herohealth <= 0 {
				fmt.Println("You Died.")
				fmt.Println("GAME OVER")
				main()
			}
			*h.health = herohealth
			*h.attack = heroattack
			win = 1
		}
		if herohealth <= 0 {
			fmt.Println("You Died.")
			main()
		}
	}
}

func Move(c *[][]int, d [][]int, s string, h *Hero, diff int) { //Moves the number 2 in the slice around, up,down,left,right
	i := *c
	switch s {
	case "w":
		// MOVE UP
		fmt.Println("Moving Up")
		for a := range i {
			if a == 0 {
				for b := range i[a] {
					if i[a+1][b] == 2 {
						i[a+1][b] = 0
						i[a][b] = 2
						checknumber(a, b, d, h, diff)
						*c = i
						return
					}
					if i[a][b] == 2 {
						i[a][b] = 0
						i[len(i)-1][b] = 2
						checknumber(len(i)-1, b, d, h, diff)
						*c = i
						return
					}
				}
			}
			if a != 0 {
				for b := range i[a] {
					if i[a][b] == 2 {
						i[a-1][b] = 2
						checknumber(a-1, b, d, h, diff)
						i[a][b] = 0
						*c = i
						return
					}
				}
			}

		} // END MOVE UP
	case "s":
		// MOVE DOWN
		fmt.Println("Moving Down")
		for a := range i {
			if a != len(i)-1 {
				for b := range i[a] {
					if i[a][b] == 2 {
						i[a][b] = 0
						i[a+1][b] = 2
						checknumber(a+1, b, d, h, diff)
						*c = i
						return
					}
				}
			}
			if a == len(i)-1 {
				for b := range i[a] {
					if i[a][b] == 2 {
						i[a][b] = 0
						i[0][b] = 2
						checknumber(0, b, d, h, diff)
						*c = i
						return
					}
				}
			}

		} // END MOVE DOWN
	case "a":
		// MOVE LEFT
		fmt.Println("Moving Left")
		for a := range i {
			for b := range i[a] {
				if b == 0 {
					if i[a][b] == 2 {
						i[a][b] = 0
						i[a][len(i[b])-1] = 2
						checknumber(a, len(i[b])-1, d, h, diff)
						*c = i
						return
					}
				}
				if b != 0 {
					if i[a][b] == 2 {
						i[a][b] = 0
						i[a][b-1] = 2
						checknumber(a, b-1, d, h, diff)
						*c = i
						return
					}
				}
			}
		} //END MOVE LEFT
	case "d":
		// MOVE RIGHT
		fmt.Println("Moving Right")
		for a := range i {
			for b := range i[a] {
				if b == len(i[0])-1 {
					if i[a][b] == 2 {
						i[a][b] = 0
						i[a][0] = 2
						checknumber(a, 0, d, h, diff)
						*c = i
						return
					}
				}
				if b != len(i[0])-1 {
					if i[a][b] == 2 {
						i[a][b] = 0
						i[a][b+1] = 2
						checknumber(a, b+1, d, h, diff)
						*c = i
						return
					}
				}
			}
		} //END MOVE RIGHT
	} //END CASES
	return
} // END FUNCTION

type Hero struct {
	name   string
	health *int
	attack *int
}

type Monster struct {
	name   string
	health int
	attack int
}

func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	you := Hero{}
	fmt.Println(`
 ______   _______  ______    ___   _          
|      | |   _   ||    _ |  |   | | |         
|  _    ||  |_|  ||   | ||  |   |_| |         
| | |   ||       ||   |_||_ |      _|         
| |_|   ||       ||    __  ||     |_          
|       ||   _   ||   |  | ||    _  |         
|______| |__| |__||___|  |_||___| |_|         
 ______    _______  _______  __   __  _______ 
|    _ |  |       ||       ||  |_|  ||       |
|   | ||  |   _   ||   _   ||       ||  _____|
|   |_||_ |  | |  ||  | |  ||       || |_____ 
|    __  ||  |_|  ||  |_|  ||       ||_____  |
|   |  | ||       ||       || ||_|| | _____| |
|___|  |_||_______||_______||_|   |_||_______|
	`)

	fmt.Println("Type in your name:")
	Scanner.Scan()
	you.name = Scanner.Text()
	fmt.Println("Good luck " + you.name + "!")
	hp := 200
	attk := 5
	you.health = &hp
	you.attack = &attk
	gameover := 0
	difficulty := 0
	initx := 6
	inity := 6
	zeroslice := generateSlice(4, 2, 0, &initx)
	gameslice := generateSlice(4, 2, 1, &inity)
	fmt.Println("Type 'w' 's' 'a' 'd' to move, 'p' to view player stats, 'q' to quit the game:")
	for gameover != 1 {
		mapx, mapy := randomNumber(2, 5), randomNumber(2, 5)
		checkMapClear(mapx, mapy, &difficulty, &zeroslice, &gameslice, &you)
		fmt.Println("Difficulty is currently:" + strconv.Itoa(difficulty))
		printSlice(zeroslice)
		fmt.Println("")
		fmt.Println("Type here:")
		Scanner.Scan()
		result := Scanner.Text()
		switch result {
		case "q":
			gameover = 1
			os.Exit(3)
		case "n":
			resetSlice(mapx, mapy, &difficulty, &zeroslice, &gameslice)
		case "w":
			Move(&zeroslice, gameslice, "w", &you, difficulty)
		case "s":
			Move(&zeroslice, gameslice, "s", &you, difficulty)
		case "a":
			Move(&zeroslice, gameslice, "a", &you, difficulty)
		case "d":
			Move(&zeroslice, gameslice, "d", &you, difficulty)
		case "p":
			fmt.Println("You have " + strconv.Itoa(*you.health) + " health remaining.")
			fmt.Println("You have " + strconv.Itoa(*you.attack) + " attack.")
		case "ok":
			printSlice(gameslice)
		}
	}
}
