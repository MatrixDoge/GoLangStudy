/* Rock-paper-scissors-lizard-spock
 * @Author: Parthas-Menethil
 * @Last Modified: Parthas-Menethil on Oct 13th, 2013
 * @Github: https://github.com/Parthas-Menethil/GoLangStudy
 * */ 

package main

import ("fmt"
		"math/rand"
		"time"
)
var patterns = []string{"NULL", "Scissors","Paper", "Rock", "Lizard", "Spock", "BuffeOops!"}

func main(){
	fmt.Println("Welcome to Rock-paper-scissors-lizard-spock")
	fmt.Println("Modes Available:")
	fmt.Println("1.\tPlayer Versus AI")
	fmt.Println("2.\tTwo Players")
	fmt.Println("0.\tExit Game")
	fmt.Print("Please select a game mode:")
	var mode int;
	fmt.Scanf("%d", &mode);
	switch mode {
	default: fmt.Println("Error->Invalid Selection")
	case 0: break
	case 1: modeAI()
	case 2: modeTwoPlayers()
	}
}
func printSelections(){
	fmt.Println("===========================")
	fmt.Println("\tSelections")
	fmt.Println("===========================")
	fmt.Println("1. Scissors")
	fmt.Println("2. Paper")
	fmt.Println("3. Rock")
	fmt.Println("4. Lizard")
	fmt.Println("5. Spock")
}
func modeTwoPlayers(){
	var p1, p2, result int
	// Ask for input
	printSelections()
	fmt.Print("Please enter player 1's choice:")
	fmt.Scanf("%d", &p1)
	printSelections()
	fmt.Print("Please enter player 2's choice:")
	fmt.Scanf("%d", &p2)
	// Judge result
	result = judgeResult(p1, p2, true)
	// Print result
	switch result{
	case -1: fmt.Println("Player two wins")
	case 0 : fmt.Println("Tie")
	case 1 : fmt.Println("Player one wins")
	}
}
func modeAI(){
	var p1, p2, result int
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))	// Instantiate an Rand instance and randomize
	// Ask for input
	printSelections()
	fmt.Print("Please enter your choice:")
	fmt.Scanf("%d", &p1)
	// Generate AI's choice
	p2 = r.Intn(5) + 1;
	fmt.Println("AI's choice is", patterns[p2])
	// Judge result
	result = judgeResult(p1, p2, true)
	// Print result
	switch result{
	case -1: fmt.Println("AI wins")
	case 0 : fmt.Println("Tie")
	case 1 : fmt.Println("You win!")
	}
}
func judgeResult(p1 int, p2 int, autoDisplay bool) int{
	
	var ret	int;
	// Determine the result
	switch (p1 - p2){
	default: 		ret = -1	// p1 lose
	case 0:			ret = 0		// Tie
	case -1, -3, 2, 4:	ret = 1	// p1 win
	}
	// Auto display result
	if autoDisplay == true {
		if ret == -1 {
			fmt.Println(patterns[p1], "is beaten by", patterns[p2])
		}
		if ret == 0 {
			fmt.Println(patterns[p1], "ties with", patterns[p2])
		}
		if ret == 1 {
			fmt.Println(patterns[p1], "beats", patterns[p2])
		}
	}
	return ret
}
