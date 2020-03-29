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
	name   string
	won    int
	drawn  int
	lost   int
	points int
}

// Tally creates a league table from raw results input
func Tally(reader io.Reader, writer io.Writer) error {
	var records = make(map[string]*Record)

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

		if _, exists := records[t1]; !exists {
			records[t1] = &Record{
				t1,
				0,
				0,
				0,
				0,
			}
		}

		if _, exists := records[t2]; !exists {
			records[t2] = &Record{
				t2,
				0,
				0,
				0,
				0,
			}
		}

		switch result {
		case "win":
			records[t1].won++
			records[t1].points += 3
			records[t2].lost++
		case "loss":
			records[t1].lost++
			records[t2].won++
			records[t2].points += 3
		case "draw":
			records[t1].drawn++
			records[t1].points++
			records[t2].drawn++
			records[t2].points++
		default:
			return errors.New("Invalid match result")
		}
	}

	recordsSlice := make([]*Record, 0, len(records))
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

	writer.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))

	for _, r := range recordsSlice {
		fmt.Fprintf(
			writer,
			"%-30s |%3d |%3d |%3d |%3d |%3d\n",
			r.name,
			r.won+r.drawn+r.lost,
			r.won,
			r.drawn,
			r.lost,
			r.points,
		)
	}

	return nil
}
