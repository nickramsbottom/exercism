package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Record an individual element in a Tally
type Record struct {
	name    string
	matches int
	won     int
	drawn   int
	lost    int
	points  int
}

// Tally creates a league table from raw results input
func Tally(reader io.Reader, writer io.Writer) error {
	var records = make(map[string]Record)

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

		t1, t2, result := separated[0], separated[1], separated[2]
		r1, r2 := records[t1], records[t2]

		r1.name = t1
		r2.name = t2

		r1.matches++
		r2.matches++

		switch result {
		case "win":
			r1.won++
			r1.points += 3
			r2.lost++
		case "loss":
			r1.lost++
			r2.won++
			r2.points += 3
		case "draw":
			r1.drawn++
			r1.points++
			r2.drawn++
			r2.points++
		default:
			return errors.New("Invalid match result")
		}

		records[t1], records[t2] = r1, r2
	}

	recordsSlice := make([]Record, 0, len(records))
	for _, record := range records {
		recordsSlice = append(recordsSlice, record)
	}

	sort.Slice(recordsSlice, func(i, j int) bool {
		record1 := recordsSlice[i]
		record2 := recordsSlice[j]

		if record1.points == record2.points {
			return record1.name < record2.name
		}

		return record1.points > record2.points
	})

	fmt.Fprintln(writer, "Team                           | MP |  W |  D |  L |  P")

	for _, r := range recordsSlice {
		fmt.Fprintf(
			writer,
			"%-30s |%3d |%3d |%3d |%3d |%3d\n",
			r.name,
			r.matches,
			r.won,
			r.drawn,
			r.lost,
			r.points,
		)
	}

	return nil
}
