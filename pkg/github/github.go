package github

import (
	"context"
	"fmt"
	"github.com/andredurao/prs/pkg/git"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"log"
	"os"
)

type PullRequest struct {
	URL         string
	HeadRefName string
	BaseRefName string
}

var args string = "is:open is:pr repo:my_repo"

type TQuery struct {
	Search struct {
		Nodes []struct {
			PullRequest PullRequest `graphql:"... on PullRequest"`
		}
	} `graphql:"search(query: $args, type: ISSUE, first: 10)"`
}

var query TQuery

func githubToken() string {
	return os.Getenv("GITHUB_AUTH_TOKEN")
}

func variables() map[string]interface{} {
	query_args := fmt.Sprintf("is:open is:pr repo:%s", git.RepositoryPath())
	return map[string]interface{}{
		"args": githubv4.String(query_args),
	}
}

func PullRequests() interface{} {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken()},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)
	err := client.Query(context.Background(), &query, variables())
	if err != nil {
		log.Panicln(err)
	}
	return query
}
