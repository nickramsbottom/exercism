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
	if len(records) == 0 {
		return nil, nil
	}

	recordsSlice := records[:]

	sort.Slice(recordsSlice, func(i, j int) bool {
		return recordsSlice[i].ID < recordsSlice[j].ID
	})

	if recordsSlice[0].ID != 0 {
		return nil, errors.New("No root node")
	}

	nodes := make(map[int]*Node)

	for i, record := range recordsSlice {
		if record.ID == 0 && record.Parent != 0 {
			return nil, errors.New("Root node should have no parent")
		}

		if record.ID < record.Parent {
			return nil, errors.New("Parent can't have a higher ID than Node")
		}

		if record.ID != 0 && record.ID == record.Parent {
			return nil, errors.New("Node can't be its own parent")
		}

		if nodes[record.ID] != nil {
			return nil, errors.New("Duplicate Node passed")
		}

		if record.ID != i {
			return nil, errors.New("Records must be continuous")
		}

		node := Node{
			record.ID,
			nil,
		}

		nodes[record.ID] = &node

		if record.ID == 0 {
			continue
		}

		nodes[record.Parent].Children = append(nodes[record.Parent].Children, &node)
	}

	return nodes[0], nil
}
