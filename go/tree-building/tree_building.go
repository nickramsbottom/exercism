package tree

import (
	"errors"
	"sort"
)

// Record of an individual post
type Record struct {
	ID     int
	Parent int
}

// Node a node in the tree structure
type Node struct {
	ID       int
	Children []*Node
}

// Build creates a tree data structure
func Build(records []Record) (*Node, error) {
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make(map[int]*Node)

	for i, record := range records {
		if record.ID < record.Parent {
			return nil, errors.New("Parent can't have a higher ID than Node")
		}

		if record.ID != 0 && record.ID == record.Parent {
			return nil, errors.New("Node can't be its own parent")
		}

		if record.ID != i {
			return nil, errors.New("Records must be continuous")
		}

		nodes[record.ID] = &Node{
			record.ID,
			nil,
		}

		if record.ID > 0 {
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, nodes[record.ID])
		}
	}

	return nodes[0], nil
}
