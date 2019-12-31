package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

type Team struct {
	name string //team name
	mp   int    //matches played
	mw   int    //matches won
	md   int    //matches tied
	ml   int    //matches lost
}

var teamStats = make(map[string]*Team)

func parseLine(line string) error {
	words := strings.Split(line, ";")
	if len(words) != 3 {
		return nil
	}
	team1, team2, result := words[0], words[1], words[2]

	if teamStats[team1] == nil {
		teamStats[team1] = new(Team)
	}
	if teamStats[team2] == nil {
		teamStats[team2] = new(Team)
	}
	if !(result == "win" || result == "loss" || result == "draw") {
		return errors.New("can't parse that")
	}
	if result == "win" {
		teamStats[team1].mp++
		teamStats[team1].mw++
		teamStats[team2].mp++
		teamStats[team2].ml++

	}
	if result == "loss" {
		teamStats[team1].mp++
		teamStats[team1].ml++
		teamStats[team2].mp++
		teamStats[team2].mw++

	}
	if result == "draw" {
		teamStats[team1].mp++
		teamStats[team1].md++
		teamStats[team2].mp++
		teamStats[team2].md++

	}
	return nil
}

func Tally(in io.Reader, out io.Writer) error {

	input := bufio.NewScanner(in)

	input.Split(bufio.ScanLines)
	for input.Scan() {
		err := parseLine(input.Text())
		if err != nil {
			return err
		}
	}
	for team := range teamStats {
		fmt.Println(teamStats[team].mp)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}
