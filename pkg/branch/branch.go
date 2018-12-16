package branch

import (
	"github.com/andredurao/prs/pkg/tree"
)

type Branch struct {
	Name string
	Tree tree.Tree
}

type BranchesMap map[string]*Branch
