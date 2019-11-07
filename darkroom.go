package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func generateNest(x, cmd, diff int) []int {
	xindex := 0
	nest := make([]int, 0)
	for xindex = 0; xindex < x; xindex++ {
		nest = append(nest, 0)
	}
	switch cmd {
	case 1:
		for xindex = 0; xindex < x-1; xindex++ {
			nest[randomNumber(0, x)] = randomNumber(0, 12)
		}
	}
	return nest
}

func generateSlice(x, y, cmd, diff int) [][]int {
	yindex := 0
	slice := make([][]int, 0, 0)
	for yindex = 0; yindex < y; yindex++ {
		slice = append(slice, generateNest(x, cmd, diff))
	}
	switch cmd {
	case 0:
		slice[0][0] = 2
	case 1:
		slice[0][0] = 2
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

func resetSlice(a int, b int, diff int, zeros, ones *[][]int) {
	i := *zeros
	j := *ones
	i = generateSlice(a, b, 0, diff)
	j = generateSlice(a, b, 1, diff)
	*zeros = i
	*ones = j
}

func checknumber(y, x int, game [][]int, he *Hero) {
	h := *he
	number := game[y][x]
	switch number {
	case 0:
		fmt.Println("You found nothing here.")
	case 1:
		fmt.Println("You have found a Dagger!")
		i := *h.attack
		i = i + 1
		*h.attack = i
	case 3:
		fmt.Println("You have found a Axe!")
		i := *h.attack
		i = i + 3
		*h.attack = i
	case 4:
		fmt.Println("You have found a Sword!")
		i := *h.attack
		i = i + 5
		*h.attack = i
	case 5:
		fmt.Println("You have found Excalibur!")
		i := *h.attack
		i = i + 8
		*h.attack = i
	case 6:
		fmt.Println("You have encountered a small toad.")
		initiateFight(he, 6)
	case 7:
		fmt.Println("You have encountered a goblin.")
		initiateFight(he, 7)
	case 8:
		fmt.Println("A giant moth descends from the ceiling.")
		initiateFight(he, 8)
	case 9:
		fmt.Println("The ghost of bad luck spooks you.")
		initiateFight(he, 9)
	case 10:
		fmt.Println("A shadow beast approaches.")
		initiateFight(he, 10)
	case 11:
		fmt.Println("A hooded rogue has been waiting...")
		initiateFight(he, 11)
	case 12:
		fmt.Println("A demon appears!")
		initiateFight(he, 11)
	}
	game[y][x] = 0
	*he = h
}

func initiateFight(h *Hero, i int) {
	toad := Monster{"toad", 20, 1}
	goblin := Monster{"goblin", 15, 2}
	moth := Monster{"giant moth", 30, 3}
	shadowbeast := Monster{"shadow beast", 30, 5}
	rogue := Monster{"hooded rogue", 35, 6}
	demon := Monster{"demon", 40, 6}
	switch i {
	case 6:
		Fight(h, toad)
	case 7:
		Fight(h, goblin)
	case 8:
		Fight(h, moth)
	case 10:
		Fight(h, shadowbeast)
	case 11:
		Fight(h, rogue)
	case 12:
		Fight(h, demon)
	}
}

func Fight(h *Hero, m Monster) {
	herohealth := *h.health
	heroattack := *h.attack
	win := 0
	for win == 0 {
		herohealth = herohealth - m.attack
		m.health = m.health - heroattack
		if m.health <= 0 {
			fmt.Println("You win the fight against" + m.name)
			fmt.Println("You have" + strconv.Itoa(herohealth) + "remaining.")
			if herohealth <= 0 {
				fmt.Println("You Died.")
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

func Move(c *[][]int, d [][]int, s string, h *Hero) { //Moves the number 2 in the slice around, up,down,left,right
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
						checknumber(a, b, d, h)
						*c = i
						return
					}
					if i[a][b] == 2 {
						i[a][b] = 0
						i[len(i)-1][b] = 2
						checknumber(len(i)-1, b, d, h)
						*c = i
						return
					}
				}
			}
			if a != 0 {
				for b := range i[a] {
					if i[a][b] == 2 {
						i[a-1][b] = 2
						checknumber(a-1, b, d, h)
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
						checknumber(a+1, b, d, h)
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
						checknumber(0, b, d, h)
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
						checknumber(a, len(i[b])-1, d, h)
						*c = i
						return
					}
				}
				if b != 0 {
					if i[a][b] == 2 {
						i[a][b] = 0
						i[a][b-1] = 2
						checknumber(a, b-1, d, h)
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
						checknumber(a, 0, d, h)
						*c = i
						return
					}
				}
				if b != len(i[0])-1 {
					if i[a][b] == 2 {
						i[a][b] = 0
						i[a][b+1] = 2
						checknumber(a, b+1, d, h)
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
	fmt.Println("Darkrooms---")
	fmt.Println("Type in your name:")
	Scanner.Scan()
	you.name = Scanner.Text()
	fmt.Println("Good luck" + you.name + "!")
	hp := 100
	attk := 5
	you.health = &hp
	you.attack = &attk
	gameover := 0
	difficulty := 0
	zeroslice := generateSlice(2, 2, 0, 1)
	gameslice := generateSlice(2, 2, 1, 1)
	for gameover != 1 {
		difficulty = difficulty + 1
		mapx, mapy := randomNumber(2, 5), randomNumber(2, 5)
		fmt.Println("Difficulty:" + strconv.Itoa(difficulty))
		printSlice(zeroslice)
		fmt.Println("")
		printSlice(gameslice)
		fmt.Println("Type here:")
		Scanner.Scan()
		result := Scanner.Text()
		switch result {
		case "q":
			gameover = 1
			os.Exit(3)
		case "n":
			resetSlice(mapx, mapy, difficulty, &zeroslice, &gameslice)
		case "w":
			Move(&zeroslice, gameslice, "w", &you)
		case "s":
			Move(&zeroslice, gameslice, "s", &you)
		case "a":
			Move(&zeroslice, gameslice, "a", &you)
		case "d":
			Move(&zeroslice, gameslice, "d", &you)
		case "p":
			fmt.Println("You have " + strconv.Itoa(*you.health) + " health remaining.")
			fmt.Println("You have " + strconv.Itoa(*you.attack) + " attack.")
		}
	}
}
