package branch

import (
	"fmt"
	"github.com/andredurao/prs/pkg/github"
	"github.com/andredurao/prs/pkg/tree"
)

var rootNode tree.Tree
var branchesMap BranchesMap

var pullRequests github.TQuery

func MountMap() {
	initializeMap()
	populateBranches()
	populateTree()
}

func initializeMap() {
	rootNode = tree.New("master", nil)
	branchesMap = make(BranchesMap)
	addBranch("master", nil)
	pullRequests = github.PullRequests().(github.TQuery)
	fmt.Printf("%+v\n", pullRequests)
}

func addBranch(name string, pull interface{}) {
	branchesMap[name] = &Branch{name, tree.New(name, pull)}
}

func populateBranches() {
	for _, pr := range pullRequests.Search.Nodes {
		pullRequest := pr.PullRequest
		addBranch(pullRequest.HeadRefName, pullRequest)
		if branchesMap[pullRequest.BaseRefName] == nil {
			addBranch(pullRequest.BaseRefName, pullRequest)
		}
	}
}

func populateTree() {
	for _, branch := range branchesMap {
		parentName := FindParentBranch(branch.Name)
		if parentName != "master" {
			parent := branchesMap[parentName]
			parent.Tree.AddTree(branch.Tree)
		} else if branch.Tree.Text() != "master" {
			rootNode.AddTree(branch.Tree)
		}
	}
}

func FindParentBranch(name string) string {
	for _, pr := range pullRequests.Search.Nodes {
		pullRequest := pr.PullRequest
		if pullRequest.HeadRefName == name {
			return pullRequest.BaseRefName
		}
	}
	return "master"
}
