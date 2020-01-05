package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

//Team represents all team's name + stats
type Team struct {
	name    string // team name
	matches int    //matches played
	won     int    //matches won
	tied    int    //matches tied
	lost    int    //matches lost
	score   int    //team score
}

//parseLine pasres line and updates stats in teamStats
func parseLine(teamStats map[string]Team, line string) error {
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
	team1Stats := teamStats[team1]
	team2Stats := teamStats[team2]
	team1Stats.name, team2Stats.name = team1, team2

	//check if result is correct
	if !(result == "win" || result == "loss" || result == "draw") {
		return errors.New("can't parse that")
	}
	switch result {
	case "win":
		team1Stats.won++
		team1Stats.score += 3
		team2Stats.lost++
	case "loss":
		team1Stats.lost++
		team2Stats.won++
		team2Stats.score += 3
	case "draw":
		team1Stats.tied++
		team1Stats.score++
		team2Stats.tied++
		team2Stats.score++
	}

	team1Stats.matches++
	team2Stats.matches++
	teamStats[team1] = team1Stats
	teamStats[team2] = team2Stats
	return nil
}

func generateTable(out io.Writer, teamStats map[string]Team) {
	//table's header
	header := "Team                           | MP |  W |  D |  L |  P\n"
	//create slice to put sorted Teams
	var sortedTable = make([]Team, 0, len(teamStats))
	for _, team := range teamStats {
		sortedTable = append(sortedTable, team)
	}
	sort.Slice(sortedTable, func(i, j int) bool {
		//if points the same order aplhabetically
		if sortedTable[i].score == sortedTable[j].score {
			return sortedTable[i].name < sortedTable[j].name
		}
		return sortedTable[i].score > sortedTable[j].score
	})
	//generate rest of the tavle line by line
	fmt.Fprintf(out, header)
	for _, team := range sortedTable {
		fmt.Fprintf(out, "%-31s|  %d |  %d |  %d |  %d |  %d\n",
			team.name, team.matches, team.won, team.tied, team.lost, team.score)
	}
}

//Tally given input in the form team1;team2;resulst returns results table
func Tally(in io.Reader, out io.Writer) error {
	//map cointaing all teams
	var teamStats = make(map[string]Team)
	//read input line by line
	input := bufio.NewScanner(in)

	input.Split(bufio.ScanLines)
	for input.Scan() {
		//parse each line
		err := parseLine(teamStats, input.Text())
		if err != nil {
			return err
		}
	}
	if err := input.Err(); err != nil {
		return err
	}
	//generate table from map and return result to io.Write
	generateTable(out, teamStats)

	return nil
}
