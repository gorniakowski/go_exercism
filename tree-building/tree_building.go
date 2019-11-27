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
	var tree = make(map[int]*Node)

	for i, rec := range records {
		//Check if tree is valid
		if rec.ID != i || rec.Parent > rec.ID || rec.ID > 0 && rec.Parent == rec.ID {

			return nil, errors.New("wrong tree")
		}
		tree[rec.ID] = &Node{ID: rec.ID}
		if i > 0 {
			tree[rec.Parent].Children = append(tree[rec.Parent].Children, tree[rec.ID])
		}
	}

	return tree[0], nil
}
