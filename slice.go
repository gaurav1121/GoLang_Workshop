package main

import (
	"fmt"
)

func main() {
	games := []string{"kokemon", "sims"}
	newGames := []string{"pacman", "doom", "pong"}

	newGames = games
	games = nil

	games = []string{}
	var ok string
	for i, game := range games {
		if game != newGames[i] {
			ok = "not "
			break
		}
	}

	// if games == nil {
	// 	ok = "not"
	// }

	if len(games) != len(newGames) {
		ok = "not"
	}

	fmt.Printf("games and newGames are %sequal \n\n", ok)
	fmt.Printf("games : %#v\n", games)
	fmt.Printf("new games :%#v\n", newGames)
}
