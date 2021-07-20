package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Player struct {
	name     string
	position string
}

func NewPlayer(n string, p string) *Player {
	return &Player{
		name:     n,
		position: p,
	}
}

type Team struct {
	name    string
	players []Player
}

func NewTeam(n string) *Team {
	return &Team{
		name: n,
	}
}

func (t *Team) AddPlayer(p Player) *Team {
	t.players = append(t.players, p)
	return t
}

func readPlayers() ([]string, []string) {

	return []string{
			"Celso",
			"Martina",
			"Brandão",
			"Ian",
			"Yousha",
			"Caíque",
			"Ricardo",
			"Richard",
		},
		[]string{
			"Meio",
			"Atacante",
			"Atacante",
			"Zagueiro",
			"Atacante",
			"Zagueiro",
			"Meio",
			"Zagueiro",
		}
}

func buildPlayers(lstNames []string, lstPositions []string) []Player {

	lstPlayers := []Player{}

	for i := range lstNames {
		lstPlayers = append(lstPlayers, *NewPlayer(lstNames[i], lstPositions[i]))
	}

	return lstPlayers
}

func buildTeams(qtyTeams int) []Team {

	lstTeams := []Team{}

	for i := 0; i < qtyTeams; i++ {
		lstTeams = append(lstTeams, *NewTeam("Team " + strconv.Itoa(i+1)))
	}

	return lstTeams
}

func shufflePlayers(src []Player) []Player {
	final := make([]Player, len(src))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}
	return final
}

func splitPlayersTeams(lstBlock []Player, lstTeams *[]Team) {

	lstSuffle := shufflePlayers(lstBlock)

	for i := range lstSuffle {
		(*lstTeams)[i].AddPlayer(lstSuffle[i])
	}

}

func selectTeams(lstPlayers []Player, lstTeams []Team) {

	lstBlock := []Player{}
	qtyTeams := len(lstTeams)
	qtyPlayers := len(lstPlayers)
	ptrTeams := 0

	for i := range lstPlayers {
		if ptrTeams < qtyTeams {
			lstBlock = append(lstBlock, lstPlayers[i])
			ptrTeams++
		}

		if ptrTeams == qtyTeams || i == qtyPlayers-1 {
			splitPlayersTeams(lstBlock, &lstTeams)
			lstBlock = lstBlock[:0]
			ptrTeams = 0
		}

	}

}

func main() {

	lstNames, lstPositions := readPlayers()
	lstPlayers := buildPlayers(lstNames, lstPositions)
	lstTeams := buildTeams(4)
	selectTeams(lstPlayers, lstTeams)

	fmt.Println("Jogadores:", lstPlayers)
	fmt.Println("Times:", lstTeams)

}
