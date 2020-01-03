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

//Team represents all team's name + stats
type Team struct {
	name string // team name
	mp   int    //matches played
	mw   int    //matches won
	md   int    //matches tied
	ml   int    //matches lost
	ts   int    //team score
}

//map cointaing all teams
var teamStats = make(map[string]*Team)

//parseLine pasres line and updates stats in teamStats
func parseLine(line string) error {
	//ignore comments and empty line
	if len(line) == 0 || line[0] == '#' {
		return nil
	}
	//split line and check if input is correct
	words := strings.Split(line, ";")
	if len(words) != 3 {
		return errors.New("wrong line")
	}
	team1, team2, result := words[0], words[1], words[2]
	//if such team1 or team2 don't exist in teamStat creat new Team
	if teamStats[team1] == nil {
		teamStats[team1] = new(Team)
		teamStats[team1].name = team1

	}
	//update stats according to result of a match
	if teamStats[team2] == nil {
		teamStats[team2] = new(Team)
		teamStats[team2].name = team2
	}
	//check if result is correct
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
	//table's header
	header := "Team                           | MP |  W |  D |  L |  P\n"
	var mainTable string
	//create new slice of pointers and put in there Teams in order of points
	var sortedTable = make([]*Team, 0)
	for _, team := range teamStats {
		sortedTable = append(sortedTable, team)
	}
	sort.Slice(sortedTable, func(i, j int) bool {
		//if points the same order aplhabetically
		if sortedTable[i].ts == sortedTable[j].ts {
			return sortedTable[i].name < sortedTable[j].name
		}
		return sortedTable[i].ts > sortedTable[j].ts
	})
	//generate rest of the tavle line by line
	for _, team := range sortedTable {

		mainTable += fmt.Sprintf("%-31s|  %d |  %d |  %d |  %d |  %d\n",
			team.name, team.mp, team.mw, team.md, team.ml, team.ts)
	}
	return header + mainTable
}

//Tally given input in the form team1;team2;resulst returns results table
func Tally(in io.Reader, out io.Writer) error {
	//clean the map containg teams
	for name := range teamStats {
		delete(teamStats, name)
	}
	//read input line by line
	input := bufio.NewScanner(in)

	input.Split(bufio.ScanLines)
	for input.Scan() {
		//parse each line
		err := parseLine(input.Text())
		if err != nil {
			return err
		}
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
	//generate table from map and return result to io.Write
	_, err := io.WriteString(out, generateTable())
	if err != nil {
		return errors.New("couldn't write table")
	}
	return nil
}
