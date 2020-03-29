package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Record an indicidual element in a Tally
type Record struct {
	won   int
	drawn int
	lost  int
}

// Tally creates a league table from raw results input
func Tally(reader io.Reader, writer io.Writer) error {
	var teams = make(map[string]*Record)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || line[0] == '#' {
			continue
		}

		separated := strings.Split(line, ";")

		if len(separated) != 3 {
			return errors.New("Invalid record")
		}

		t1 := separated[0]
		t2 := separated[1]
		result := separated[2]

		if _, exists := teams[t1]; !exists {
			teams[t1] = &Record{
				0,
				0,
				0,
			}
		}

		if _, exists := teams[t2]; !exists {
			teams[t2] = &Record{
				0,
				0,
				0,
			}
		}

		if result == "win" {
			teams[t1].won++
			teams[t2].lost++
		} else if result == "loss" {
			teams[t1].lost++
			teams[t2].won++
		} else if result == "draw" {
			teams[t1].drawn++
			teams[t2].drawn++
		} else {
			return errors.New("Invalid match result")
		}
	}

	keys := make([]string, 0, len(teams))
	for k := range teams {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		score1 := teams[keys[i]].drawn + 3*teams[keys[i]].won
		score2 := teams[keys[j]].drawn + 3*teams[keys[j]].won

		if score1 == score2 {
			return keys[i] < keys[j]
		}

		return score1 > score2
	})

	writer.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))

	for _, name := range keys {
		team := teams[name]

		writer.Write([]byte(
			fmt.Sprintf("%-30s |%3d |%3d |%3d |%3d |%3d\n",
				name,
				team.won+team.drawn+team.lost,
				team.won,
				team.drawn,
				team.lost,
				3*team.won+team.drawn,
			)))
	}

	return nil
}
