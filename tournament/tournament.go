package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

type Team struct {
	name string // team name
	mp   int    //matches played
	mw   int    //matches won
	md   int    //matches tied
	ml   int    //matches lost
	ts   int    //team score
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
		teamStats[team1].name = team1

	}
	if teamStats[team2] == nil {
		teamStats[team2] = new(Team)
		teamStats[team2].name = team2
	}
	if !(result == "win" || result == "loss" || result == "draw") {
		return errors.New("can't parse that")
	}
	if result == "win" {
		teamStats[team1].mp++
		teamStats[team1].mw++
		teamStats[team1].ts += 3
		teamStats[team2].mp++
		teamStats[team2].ml++

	}
	if result == "loss" {
		teamStats[team1].mp++
		teamStats[team1].ml++
		teamStats[team2].mp++
		teamStats[team2].mw++
		teamStats[team2].ts += 3

	}
	if result == "draw" {
		teamStats[team1].mp++
		teamStats[team1].md++
		teamStats[team1].ts++
		teamStats[team2].mp++
		teamStats[team2].md++
		teamStats[team2].ts++

	}
	return nil
}

func generateTable() string {
	header := "Team                           | MP |  W |  D |  L |  P"
	var sortedTable = make([]*Team, 0)
	for _, team := range teamStats {
		sortedTable = append(sortedTable, team)
	}
	sort.Slice(sortedTable, func(i, j int) bool {
		return sortedTable[i].ts > sortedTable[j].ts
	})

	for _, team := range sortedTable {
		fmt.Println(team)
	}
	return header
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
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(generateTable())
	return nil
}
