package tree

import (
	"errors"
	"sort"
)

//Record represnts forum post
type Record struct {
	ID     int
	Parent int
}

//Node represents single node a tree
type Node struct {
	ID       int
	Children []*Node
}

//Build parses slice of recods from database into tree
func Build(records []Record) (*Node, error) {
	sort.SliceStable(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	//Check if recods are empty
	if len(records) == 0 {
		return nil, nil
	}
	//Check if we have root
	if records[0].ID != 0 {

		return nil, errors.New("no root, no go")
	}
	//We don't need to do nothing more if we have a single record
	if len(records) == 1 {
		return &Node{ID: 0}, nil
	}
	//Root node can't have a parent
	if records[0].Parent != 0 {
		return nil, errors.New("root papa is a no no")
	}

	var whereAreThey = make([]*Node, len(records))
	result := Node{ID: records[0].ID, Children: make([]*Node, 0)}
	whereAreThey[0] = &result

	for i, rec := range records[1:] {
		//Check if tree is valid
		if rec.ID <= rec.Parent || whereAreThey[i].ID-rec.ID != -1 || whereAreThey[i].ID == rec.ID {
			return nil, errors.New("Invalid tree")
		}
		whereAreThey[rec.ID] = &Node{ID: rec.ID}
		whereAreThey[rec.Parent].Children = append(whereAreThey[rec.Parent].Children, whereAreThey[rec.ID])

	}
	return &result, nil
}
